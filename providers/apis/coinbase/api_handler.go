package coinbase

import (
	"encoding/json"
	"fmt"
	"math/big"
	"net/http"
	"time"

	"github.com/skip-mev/slinky/oracle/config"
	"github.com/skip-mev/slinky/pkg/math"
	"github.com/skip-mev/slinky/providers/base/api/handlers"
	"github.com/skip-mev/slinky/providers/constants"
	providertypes "github.com/skip-mev/slinky/providers/types"
	mmtypes "github.com/skip-mev/slinky/x/marketmap/types"
)

var _ constants.PriceAPIDataHandler = (*APIHandler)(nil)

// APIHandler implements the APIDataHandler interface for Coinbase, which can be used
// by a base provider. The DataHandler fetches data from the spot price Coinbase API. It is
// atomic in that it must request data from the Coinbase API sequentially for each ticker.
type APIHandler struct {
	// market is the config for the Coinbase API.
	market mmtypes.MarketConfig
	// api is the config for the Coinbase API.
	api config.APIConfig
}

// NewAPIHandler returns a new Coinbase APIDataHandler.
func NewAPIHandler(
	market mmtypes.MarketConfig,
	api config.APIConfig,
) (handlers.APIDataHandler[mmtypes.Ticker, *big.Int], error) {
	if err := market.ValidateBasic(); err != nil {
		return nil, fmt.Errorf("invalid market config for %s: %w", Name, err)
	}

	if market.Name != Name {
		return nil, fmt.Errorf("expected market config name %s, got %s", Name, market.Name)
	}

	if api.Name != Name {
		return nil, fmt.Errorf("expected api config name %s, got %s", Name, api.Name)
	}

	if !api.Enabled {
		return nil, fmt.Errorf("api config for %s is not enabled", Name)

	}

	if err := api.ValidateBasic(); err != nil {
		return nil, fmt.Errorf("invalid api config for %s: %w", Name, err)
	}

	return &APIHandler{
		market: market,
		api:    api,
	}, nil
}

// CreateURL returns the URL that is used to fetch data from the Coinbase API for the
// given tickers. Since the Coinbase API only supports fetching spot prices for a single
// ticker at a time, this function will return an error if the ticker slice contains more
// than one ticker.
func (h *APIHandler) CreateURL(
	tickers []mmtypes.Ticker,
) (string, error) {
	if len(tickers) != 1 {
		return "", fmt.Errorf("expected 1 ticker, got %d", len(tickers))
	}

	// Ensure that the base and quote currencies are supported by the Coinbase API and
	// are configured for the handler.
	ticker := tickers[0]
	market, ok := h.market.TickerConfigs[tickers[0].String()]
	if !ok {
		return "", fmt.Errorf("unknown ticker %s", ticker.String())
	}

	return fmt.Sprintf(h.api.URL, market.OffChainTicker), nil
}

// ParseResponse parses the spot price HTTP response from the Coinbase API and returns
// the resulting price. Note that this can only parse a single ticker at a time.
func (h *APIHandler) ParseResponse(
	tickers []mmtypes.Ticker,
	resp *http.Response,
) providertypes.GetResponse[mmtypes.Ticker, *big.Int] {
	if len(tickers) != 1 {
		return providertypes.NewGetResponseWithErr[mmtypes.Ticker, *big.Int](
			tickers,
			fmt.Errorf("expected 1 ticker, got %d", len(tickers)),
		)
	}

	// Check if this ticker is supported by the Coinbase API market config.
	ticker := tickers[0]
	_, ok := h.market.TickerConfigs[ticker.String()]
	if !ok {
		return providertypes.NewGetResponseWithErr[mmtypes.Ticker, *big.Int](
			tickers,
			fmt.Errorf("unknown ticker %s", ticker.String()),
		)
	}

	// Parse the response into a CoinBaseResponse.
	var result CoinBaseResponse
	if err := json.NewDecoder(resp.Body).Decode(&result); err != nil {
		return providertypes.NewGetResponseWithErr[mmtypes.Ticker, *big.Int](tickers, err)
	}

	// Convert the float64 price into a big.Int.
	price, err := math.Float64StringToBigInt(result.Data.Amount, ticker.Decimals)
	if err != nil {
		return providertypes.NewGetResponseWithErr[mmtypes.Ticker, *big.Int](tickers, err)
	}

	return providertypes.NewGetResponse[mmtypes.Ticker, *big.Int](
		map[mmtypes.Ticker]providertypes.Result[*big.Int]{
			ticker: providertypes.NewResult[*big.Int](price, time.Now()),
		},
		nil,
	)
}
