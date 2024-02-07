package keeper

import (
	"errors"

	"cosmossdk.io/collections"
	"cosmossdk.io/collections/indexes"
	"cosmossdk.io/core/store"

	"github.com/cosmos/cosmos-sdk/codec"
	sdk "github.com/cosmos/cosmos-sdk/types"

	"github.com/skip-mev/slinky/x/oracle/types"
)

type oracleIndices struct {
	// idUnique is a uniqueness constraint on the IDs of CurrencyPairs. i.e id -> CurrencyPair.String() -> CurrencyPairState
	idUnique *indexes.Unique[uint64, string, types.CurrencyPairState]

	// idMulti is a multi-index on the IDs of CurrencyPairs, i.e. id -> CurrencyPair.String() -> CurrencyPairState
	idMulti *indexes.Multi[uint64, string, types.CurrencyPairState]
}

func (o *oracleIndices) IndexesList() []collections.Index[string, types.CurrencyPairState] {
	return []collections.Index[string, types.CurrencyPairState]{
		o.idUnique,
		o.idMulti,
	}
}

func newOracleIndices(sb *collections.SchemaBuilder) *oracleIndices {
	return &oracleIndices{
		idUnique: indexes.NewUnique[uint64, string, types.CurrencyPairState](
			sb, types.UniqueIndexCurrencyPairKeyPrefix, "currency_pair_id_unique_idx", collections.Uint64Key, collections.StringKey,
			func(_ string, cps types.CurrencyPairState) (uint64, error) {
				return cps.Id, nil
			},
		),
		idMulti: indexes.NewMulti[uint64, string, types.CurrencyPairState](
			sb, types.IDIndexCurrencyPairKeyPrefix, "currency_pair_id_idx", collections.Uint64Key, collections.StringKey,
			func(_ string, cps types.CurrencyPairState) (uint64, error) {
				return cps.Id, nil
			},
		),
	}
}

// Keeper is the base keeper for the x/oracle module.
type Keeper struct {
	storeService store.KVStoreService
	cdc          codec.BinaryCodec

	// schema
	nextCurrencyPairID collections.Sequence
	currencyPairs      *collections.IndexedMap[string, types.CurrencyPairState, *oracleIndices]
	schema             collections.Schema

	// indexes
	idIndex *indexes.Multi[uint64, string, types.CurrencyPairState]

	// module authority
	authority sdk.AccAddress
}

// NewKeeper constructs a new keeper from a store-key + authority account address.
func NewKeeper(
	ss store.KVStoreService,
	cdc codec.BinaryCodec,
	authority sdk.AccAddress,
) Keeper {
	// create a new schema builder
	sb := collections.NewSchemaBuilder(ss)

	indices := newOracleIndices(sb)

	idMulti, ok := indices.IndexesList()[1].(*indexes.Multi[uint64, string, types.CurrencyPairState])
	if !ok {
		panic("expected idMulti to be a *indexes.Multi[uint64, string, types.CurrencyPairState]")
	}

	k := Keeper{
		storeService:       ss,
		cdc:                cdc,
		authority:          authority,
		nextCurrencyPairID: collections.NewSequence(sb, types.CurrencyPairIDKeyPrefix, "currency_pair_id"),
		currencyPairs:      collections.NewIndexedMap(sb, types.CurrencyPairKeyPrefix, "currency_pair", collections.StringKey, codec.CollValue[types.CurrencyPairState](cdc), indices),
		idIndex:            idMulti,
	}

	// create the schema
	schema, err := sb.Build()
	if err != nil {
		panic(err)
	}

	k.schema = schema
	return k
}

// RemoveCurrencyPair removes a given CurrencyPair from state, i.e. removes its nonce + QuotePrice from the module's store.
func (k Keeper) RemoveCurrencyPair(ctx sdk.Context, cpID string) {
	k.currencyPairs.Remove(ctx, cpID)
}

// HasCurrencyPair returns true if a given CurrencyPair is stored in state, false otherwise.
func (k Keeper) HasCurrencyPair(ctx sdk.Context, cpID string) bool {
	ok, err := k.currencyPairs.Has(ctx, cpID)
	if err != nil || !ok {
		return false
	}

	return true
}

// GetPriceWithNonceForCurrencyPair returns a QuotePriceWithNonce for a given CurrencyPair. The nonce for the QuotePrice represents
// the number of times that a given QuotePrice has been updated. Notice: prefer GetPriceWithNonceForCurrencyPair over GetPriceForCurrencyPair.
func (k Keeper) GetPriceWithNonceForCurrencyPair(ctx sdk.Context, cpID string) (types.QuotePriceWithNonce, error) {
	// get the QuotePrice for the currency pair
	qp, err := k.GetPriceForCurrencyPair(ctx, cpID)
	if err != nil {
		// only fail if the Price Query failed for a reason other than there being no QuotePrice for cp
		var quotePriceNotExistError types.QuotePriceNotExistError
		if !errors.As(err, &quotePriceNotExistError) {
			return types.QuotePriceWithNonce{}, err
		}
	}

	// get the nonce
	nonce, err := k.GetNonceForCurrencyPair(ctx, cpID)
	if err != nil {
		return types.QuotePriceWithNonce{}, err
	}

	return types.NewQuotePriceWithNonce(qp, nonce), nil
}

// NextCurrencyPairID returns the next String to be assigned to a currency-pair.
func (k Keeper) NextCurrencyPairID(ctx sdk.Context) (uint64, error) {
	return k.nextCurrencyPairID.Peek(ctx)
}

// GetNonceForCurrencyPair returns the nonce for a given CurrencyPair. If one has not been stored, return an error.
func (k Keeper) GetNonceForCurrencyPair(ctx sdk.Context, cpID string) (uint64, error) {
	cps, err := k.currencyPairs.Get(ctx, cpID)
	if err != nil {
		return 0, err
	}

	return cps.Nonce, nil
}

// GetPriceForCurrencyPair retrieves the QuotePrice for a given CurrencyPair. if a QuotePrice does not
// exist for the given CurrencyPair, this function errors and returns an empty QuotePrice.
func (k Keeper) GetPriceForCurrencyPair(ctx sdk.Context, cpID string) (types.QuotePrice, error) {
	cps, err := k.currencyPairs.Get(ctx, cpID)
	if err != nil {
		return types.QuotePrice{}, err
	}

	// nil check
	if cps.Price == nil {
		return types.QuotePrice{}, types.NewQuotePriceNotExistError(cpID)
	}

	return *cps.Price, nil
}

// SetPriceForCurrencyPair sets the given QuotePrice for a given CurrencyPair, and updates the CurrencyPair's nonce. Note, no validation is performed on
// either the CurrencyPair or the QuotePrice (it is expected the caller performs this validation). If the CurrencyPair does not exist, create the currency-pair
// and set its nonce to 0.
func (k Keeper) SetPriceForCurrencyPair(ctx sdk.Context, cp types.CurrencyPair, qp types.QuotePrice) error {
	// get the current state for the currency-pair, fail if it does not exist
	cps, err := k.currencyPairs.Get(ctx, cp.String())
	if err != nil {
		// get the next currency-pair id
		id, err := k.nextCurrencyPairID.Next(ctx)
		if err != nil {
			return err
		}

		cps = types.NewCurrencyPairState(id, 0, &qp, cp.Decimals)
	} else {
		// update the nonce
		cps.Nonce++
		cps.Price = &qp
	}

	// set the updated state
	return k.currencyPairs.Set(ctx, cp.String(), cps)
}

// CreateCurrencyPair creates a CurrencyPair in state, and sets its ID to the next available ID. If the CurrencyPair already exists, return an error.
// the nonce for the CurrencyPair is set to 0.
func (k Keeper) CreateCurrencyPair(ctx sdk.Context, cp types.CurrencyPair) error {
	// check if the currency pair already exists
	if k.HasCurrencyPair(ctx, cp.String()) {
		return types.NewCurrencyPairAlreadyExistsError(cp)
	}

	id, err := k.nextCurrencyPairID.Next(ctx)
	if err != nil {
		return err
	}

	state := types.NewCurrencyPairState(id, 0, nil, cp.Decimals)

	return k.currencyPairs.Set(ctx, cp.String(), state)
}

// GetIDForCurrencyPair returns the ID for a given CurrencyPair. If the CurrencyPair does not exist, return 0, false, if
// it does, return true and the ID.
func (k Keeper) GetIDForCurrencyPair(ctx sdk.Context, cpID string) (uint64, bool) {
	cps, err := k.currencyPairs.Get(ctx, cpID)
	if err != nil {
		return 0, false
	}

	return cps.Id, true
}

// GetCurrencyPairFromID returns the CurrencyPair for a given ID. If the ID does not exist, return an error and an empty CurrencyPair.
// Otherwise, return the currency pair and no error.
func (k Keeper) GetCurrencyPairFromID(ctx sdk.Context, id uint64) (types.CurrencyPair, bool) {
	// use the ID index to match the given ID
	ids, err := k.idIndex.MatchExact(ctx, id)
	if err != nil {
		return types.CurrencyPair{}, false
	}
	// close the iterator
	defer ids.Close()
	if !ids.Valid() {
		return types.CurrencyPair{}, false
	}

	cpStr, err := ids.PrimaryKey()
	if err != nil {
		return types.CurrencyPair{}, false
	}

	cps, err := k.currencyPairs.Get(ctx, cpStr)
	if err != nil {
		return types.CurrencyPair{}, false
	}

	cp, err := types.CurrencyPairFromTicker(cpStr, cps.Decimals)
	return cp, err == nil
}

// GetAllCurrencyPairs returns all CurrencyPairs that have currently been stored to state.
func (k Keeper) GetAllCurrencyPairs(ctx sdk.Context) []types.CurrencyPair {
	cps := make([]types.CurrencyPair, 0)

	// aggregate CurrencyPairs stored under KeyPrefixNonce
	k.IterateCurrencyPairs(ctx, func(cp types.CurrencyPair, _ types.CurrencyPairState) {
		cps = append(cps, cp)
	})

	return cps
}

// IterateCurrencyPairs iterates over all CurrencyPairs in the store, and executes a callback for each CurrencyPair.
func (k Keeper) IterateCurrencyPairs(ctx sdk.Context, cb func(cp types.CurrencyPair, cps types.CurrencyPairState)) error {
	it, err := k.currencyPairs.Iterate(ctx, nil)
	if err != nil {
		return err
	}
	defer it.Close()

	for ; it.Valid(); it.Next() {
		primaryKey, err := it.Key()
		if err != nil {
			return err
		}

		cps, err := it.Value()
		if err != nil {
			return err
		}

		cp, err := types.CurrencyPairFromTicker(primaryKey, cps.Decimals)
		if err != nil {
			return err
		}

		cb(cp, cps)
	}

	return nil
}
