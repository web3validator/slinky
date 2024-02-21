package oracle

import (
	"fmt"
	"net/http"

	"go.uber.org/zap"

	"github.com/skip-mev/slinky/oracle/config"
	"github.com/skip-mev/slinky/oracle/types"
	"github.com/skip-mev/slinky/pkg/math"
	apihandlers "github.com/skip-mev/slinky/providers/base/api/handlers"
	wshandlers "github.com/skip-mev/slinky/providers/base/websocket/handlers"
	wsmetrics "github.com/skip-mev/slinky/providers/base/websocket/metrics"
	"github.com/skip-mev/slinky/providers/websockets/bitfinex"
	"github.com/skip-mev/slinky/providers/websockets/bitstamp"
	"github.com/skip-mev/slinky/providers/websockets/bybit"
	coinbasews "github.com/skip-mev/slinky/providers/websockets/coinbase"
	"github.com/skip-mev/slinky/providers/websockets/cryptodotcom"
	"github.com/skip-mev/slinky/providers/websockets/gate"
	"github.com/skip-mev/slinky/providers/websockets/huobi"
	"github.com/skip-mev/slinky/providers/websockets/kraken"
	"github.com/skip-mev/slinky/providers/websockets/kucoin"
	"github.com/skip-mev/slinky/providers/websockets/mexc"
	"github.com/skip-mev/slinky/providers/websockets/okx"
)

// WebSocketQueryHandlerFactory returns a sample implementation of the websocket query handler
// factory. Specifically, this factory function returns websocket query handlers that are used to
// fetch data from the price providers.
func WebSocketQueryHandlerFactory(marketMap types.ProviderMarketMap) types.PriceWebSocketQueryHandlerFactory {
	return func(
		logger *zap.Logger,
		cfg config.ProviderConfig,
		wsMetrics wsmetrics.WebSocketMetrics,
	) (types.PriceWebSocketQueryHandler, error) {
		// Validate the provider config.
		err := cfg.ValidateBasic()
		if err != nil {
			return nil, err
		}

		// Create the underlying client that can be utilized by websocket providers that need to
		// interact with an API.
		tickers := marketMap.GetTickers()
		maxCons := math.Min(len(tickers), cfg.API.MaxQueries)
		client := &http.Client{
			Transport: &http.Transport{MaxConnsPerHost: maxCons},
			Timeout:   cfg.API.Timeout,
		}

		var (
			requestHandler apihandlers.RequestHandler
			wsDataHandler  types.PriceWebSocketDataHandler
			connHandler    wshandlers.WebSocketConnHandler
		)

		switch cfg.Name {
		case bitfinex.Name:
			wsDataHandler, err = bitfinex.NewWebSocketDataHandler(logger, marketMap, cfg.WebSocket)
		case bitstamp.Name:
			wsDataHandler, err = bitstamp.NewWebSocketDataHandler(logger, marketMap, cfg.WebSocket)
		case bybit.Name:
			wsDataHandler, err = bybit.NewWebSocketDataHandler(logger, marketMap, cfg.WebSocket)
		case coinbasews.Name:
			wsDataHandler, err = coinbasews.NewWebSocketDataHandler(logger, marketMap, cfg.WebSocket)
		case cryptodotcom.Name:
			wsDataHandler, err = cryptodotcom.NewWebSocketDataHandler(logger, marketMap, cfg.WebSocket)
		case gate.Name:
			wsDataHandler, err = gate.NewWebSocketDataHandler(logger, marketMap, cfg.WebSocket)
		case huobi.Name:
			wsDataHandler, err = huobi.NewWebSocketDataHandler(logger, marketMap, cfg.WebSocket)
		case kraken.Name:
			wsDataHandler, err = kraken.NewWebSocketDataHandler(logger, marketMap, cfg.WebSocket)
		case kucoin.Name:
			// Create the KuCoin websocket data handler.
			wsDataHandler, err = kucoin.NewWebSocketDataHandler(logger, marketMap, cfg.WebSocket)
			if err != nil {
				return nil, err
			}

			// The request handler requires POST requests when first establishing the connection.
			requestHandler, err = apihandlers.NewRequestHandlerImpl(
				client,
				apihandlers.WithHTTPMethod(http.MethodPost),
			)
			if err != nil {
				return nil, err
			}

			// Create the KuCoin websocket connection handler.
			connHandler, err = wshandlers.NewWebSocketHandlerImpl(
				cfg.WebSocket,
				wshandlers.WithPreDialHook(kucoin.PreDialHook(cfg.API, requestHandler)),
			)
		case mexc.Name:
			wsDataHandler, err = mexc.NewWebSocketDataHandler(logger, marketMap, cfg.WebSocket)
		case okx.Name:
			wsDataHandler, err = okx.NewWebSocketDataHandler(logger, marketMap, cfg.WebSocket)
		default:
			return nil, fmt.Errorf("unknown provider: %s", cfg.Name)
		}
		if err != nil {
			return nil, err
		}

		// If a custom request handler is not provided, create a new default one.
		if connHandler == nil {
			connHandler, err = wshandlers.NewWebSocketHandlerImpl(cfg.WebSocket)
			if err != nil {
				return nil, err
			}
		}

		// Create the websocket query handler which encapsulates all fetching and parsing logic.
		return types.NewPriceWebSocketQueryHandler(
			logger,
			cfg.WebSocket,
			wsDataHandler,
			connHandler,
			wsMetrics,
		)
	}
}
