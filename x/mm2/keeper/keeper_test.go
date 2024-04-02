package keeper_test

import (
	"testing"

	oraclekeeper "github.com/skip-mev/slinky/x/oracle/keeper"
	oracletypes "github.com/skip-mev/slinky/x/oracle/types"

	storetypes "cosmossdk.io/store/types"
	"github.com/cosmos/cosmos-sdk/runtime"
	"github.com/cosmos/cosmos-sdk/testutil"
	sdk "github.com/cosmos/cosmos-sdk/types"
	moduletestutil "github.com/cosmos/cosmos-sdk/types/module/testutil"
	"github.com/stretchr/testify/suite"

	slinkytypes "github.com/skip-mev/slinky/pkg/types"
	"github.com/skip-mev/slinky/x/mm2/keeper"
	"github.com/skip-mev/slinky/x/mm2/types"
)

type KeeperTestSuite struct {
	suite.Suite

	ctx sdk.Context

	// Keeper variables
	authority    sdk.AccAddress
	keeper       *keeper.Keeper
	oracleKeeper oraclekeeper.Keeper
}

func (s *KeeperTestSuite) initKeeper() *keeper.Keeper {
	mmKey := storetypes.NewKVStoreKey(types.StoreKey)
	oracleKey := storetypes.NewKVStoreKey(oracletypes.StoreKey)
	mmSS := runtime.NewKVStoreService(mmKey)
	oracleSS := runtime.NewKVStoreService(oracleKey)
	encCfg := moduletestutil.MakeTestEncodingConfig()

	keys := map[string]*storetypes.KVStoreKey{
		types.StoreKey:       mmKey,
		oracletypes.StoreKey: oracleKey,
	}

	transientKeys := map[string]*storetypes.TransientStoreKey{
		types.StoreKey:       storetypes.NewTransientStoreKey("transient_mm"),
		oracletypes.StoreKey: storetypes.NewTransientStoreKey("transient_oracle"),
	}

	s.authority = sdk.AccAddress("authority")
	s.ctx = testutil.DefaultContextWithKeys(keys, transientKeys, nil).WithBlockHeight(10)

	k := keeper.NewKeeper(mmSS, encCfg.Codec, s.authority)
	s.Require().NoError(k.SetLastUpdated(s.ctx, uint64(s.ctx.BlockHeight())))

	params := types.Params{
		MarketAuthorities: []string{s.authority.String()},
	}
	s.Require().NoError(k.SetParams(s.ctx, params))

	s.oracleKeeper = oraclekeeper.NewKeeper(oracleSS, encCfg.Codec, k, s.authority)
	hooks := types.MultiMarketMapHooks{
		s.oracleKeeper.Hooks(),
	}
	k.SetHooks(hooks)

	s.Require().NotPanics(func() {
		s.oracleKeeper.InitGenesis(s.ctx, *oracletypes.DefaultGenesisState())
	})

	return k
}

func (s *KeeperTestSuite) SetupTest() {
	s.keeper = s.initKeeper()
}

func TestKeeperTestSuite(t *testing.T) {
	suite.Run(t, new(KeeperTestSuite))
}

var (
	btcusdt = types.Market{
		Ticker: types.Ticker{
			CurrencyPair: slinkytypes.CurrencyPair{
				Base:  "BITCOIN",
				Quote: "USDT",
			},
			Decimals:         8,
			MinProviderCount: 1,
		},
		ProviderConfigs: []types.ProviderConfig{
			{
				Name:           "kucoin",
				OffChainTicker: "btc-usdt",
			},
		},
	}

	usdtusd = types.Market{
		Ticker: types.Ticker{
			CurrencyPair: slinkytypes.CurrencyPair{
				Base:  "USDT",
				Quote: "USD",
			},
			Decimals:         8,
			MinProviderCount: 1,
		},
		ProviderConfigs: []types.ProviderConfig{
			{
				Name:           "kucoin",
				OffChainTicker: "usdt-usd",
			},
		},
	}

	usdcusd = types.Market{
		Ticker: types.Ticker{
			CurrencyPair: slinkytypes.CurrencyPair{
				Base:  "USDC",
				Quote: "USD",
			},
			Decimals:         8,
			MinProviderCount: 1,
		},
		ProviderConfigs: []types.ProviderConfig{
			{
				Name:           "kucoin",
				OffChainTicker: "usdc-usd",
			},
		},
	}

	ethusdt = types.Market{
		Ticker: types.Ticker{
			CurrencyPair: slinkytypes.CurrencyPair{
				Base:  "ETHEREUM",
				Quote: "USDT",
			},
			Decimals:         8,
			MinProviderCount: 1,
		},
		ProviderConfigs: []types.ProviderConfig{
			{
				Name:           "kucoin",
				OffChainTicker: "eth-usdt",
			},
		},
	}

	markets = []types.Market{
		btcusdt,
		usdcusd,
		usdtusd,
		ethusdt,
	}

	marketsMap = map[string]types.Market{
		btcusdt.Ticker.String(): btcusdt,
		usdcusd.Ticker.String(): usdcusd,
		usdtusd.Ticker.String(): usdtusd,
		ethusdt.Ticker.String(): ethusdt,
	}
)

func (s *KeeperTestSuite) TestGets() {
	s.Run("get no tickers", func() {
		got, err := s.keeper.GetAllMarkets(s.ctx)
		s.Require().NoError(err)
		s.Require().Equal([]types.Market(nil), got)
	})

	s.Run("setup initial markets", func() {
		s.Require().NoError(s.keeper.CreateMarkets(s.ctx, markets))

		s.Run("unable to set markets again", func() {
			s.Require().ErrorIs(s.keeper.CreateMarkets(s.ctx, markets), types.NewMarketAlreadyExistsError(types.TickerString(btcusdt.Ticker.String())))
		})
	})

	s.Run("get all tickers", func() {
		got, err := s.keeper.GetAllMarketsMap(s.ctx)
		s.Require().NoError(err)

		s.Require().Equal(len(markets), len(got))
		s.Require().Equal(marketsMap, got)
	})
}

func (s *KeeperTestSuite) TestSetParams() {
	params := types.DefaultParams()

	s.Run("can set and get params", func() {
		err := s.keeper.SetParams(s.ctx, params)
		s.Require().NoError(err)

		params2, err := s.keeper.GetParams(s.ctx)
		s.Require().NoError(err)
		s.Require().Equal(params, params2)
	})
}

func (s *KeeperTestSuite) TestInvalidCreate() {
	// invalid market with a normalize pair not in state
	invalidMarket := types.Market{
		Ticker: types.Ticker{
			CurrencyPair: slinkytypes.CurrencyPair{
				Base:  "BITCOIN",
				Quote: "USDT",
			},
			Decimals:         8,
			MinProviderCount: 1,
		},
		ProviderConfigs: []types.ProviderConfig{
			{
				Name:            "kucoin",
				OffChainTicker:  "btc-usdt",
				NormalizeByPair: &slinkytypes.CurrencyPair{Base: "invalid", Quote: "pair"},
			},
		},
	}

	s.Require().Error(s.keeper.CreateMarkets(s.ctx, []types.Market{invalidMarket}))
}

func (s *KeeperTestSuite) TestInvalidUpdate() {
	// create a valid market
	s.Require().NoError(s.keeper.CreateMarkets(s.ctx, []types.Market{btcusdt}))

	// invalid market with a normalize pair not in state
	invalidMarket := btcusdt
	invalidMarket.ProviderConfigs = append(invalidMarket.ProviderConfigs, types.ProviderConfig{
		Name:            "huobi",
		OffChainTicker:  "btc-usdt",
		NormalizeByPair: &slinkytypes.CurrencyPair{Base: "invalid", Quote: "pair"},
	})

	s.Require().Error(s.keeper.UpdateMarkets(s.ctx, []types.Market{invalidMarket}))
}

func (s *KeeperTestSuite) TestValidUpdate() {
	// create a valid markets
	s.Require().NoError(s.keeper.CreateMarkets(s.ctx, []types.Market{btcusdt, ethusdt}))

	// valid market with a normalize pair that is in state
	invalidMarket := btcusdt
	invalidMarket.ProviderConfigs = append(invalidMarket.ProviderConfigs, types.ProviderConfig{
		Name:            "huobi",
		OffChainTicker:  "btc-usdt",
		NormalizeByPair: &ethusdt.Ticker.CurrencyPair,
	})

	s.Require().NoError(s.keeper.UpdateMarkets(s.ctx, []types.Market{invalidMarket}))
}