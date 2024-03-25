package keeper_test

import (
	"testing"

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
	authority sdk.AccAddress
	keeper    *keeper.Keeper
}

func (s *KeeperTestSuite) initKeeper() *keeper.Keeper {
	mmKey := storetypes.NewKVStoreKey(types.StoreKey)
	mmSS := runtime.NewKVStoreService(mmKey)
	encCfg := moduletestutil.MakeTestEncodingConfig()

	keys := map[string]*storetypes.KVStoreKey{
		types.StoreKey: mmKey,
	}

	transientKeys := map[string]*storetypes.TransientStoreKey{
		types.StoreKey: storetypes.NewTransientStoreKey("transient_mm"),
	}

	s.authority = sdk.AccAddress("authority")
	s.ctx = testutil.DefaultContextWithKeys(keys, transientKeys, nil).WithBlockHeight(10)

	k := keeper.NewKeeper(mmSS, encCfg.Codec, s.authority)
	s.Require().NoError(k.SetLastUpdated(s.ctx, uint64(s.ctx.BlockHeight())))

	params := types.Params{
		MarketAuthorities: []string{s.authority.String()},
		Version:           10,
	}
	s.Require().NoError(k.SetParams(s.ctx, params))

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

	markets = map[string]types.Market{
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
		for _, market := range markets {
			s.Require().NoError(s.keeper.CreateMarket(s.ctx, market))
		}

		s.Run("unable to set markets again", func() {
			for _, market := range markets {
				s.Require().ErrorIs(s.keeper.CreateMarket(s.ctx, market), types.NewMarketAlreadyExistsError(types.TickerString(market.Ticker.String())))
			}
		})
	})

	s.Run("get all tickers", func() {
		got, err := s.keeper.GetAllMarketsMap(s.ctx)
		s.Require().NoError(err)

		s.Require().Equal(len(markets), len(got))
		s.Require().Equal(markets, got)
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