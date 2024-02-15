package coinbase_test

import (
	"fmt"
	"math/big"
	"net/http"
	"testing"
	"time"

	"github.com/stretchr/testify/require"

	"github.com/skip-mev/slinky/providers/apis/coinbase"
	"github.com/skip-mev/slinky/providers/base/testutils"
	"github.com/skip-mev/slinky/providers/constants"
	providertypes "github.com/skip-mev/slinky/providers/types"
	mmtypes "github.com/skip-mev/slinky/x/marketmap/types"
)

var (
	mogusd = mmtypes.NewTicker("MOG", "USD", 8, 1)
)

func TestCreateURL(t *testing.T) {
	testCases := []struct {
		name        string
		cps         []mmtypes.Ticker
		url         string
		expectedErr bool
	}{
		{
			name: "valid",
			cps: []mmtypes.Ticker{
				constants.BITCOIN_USD,
			},
			url:         "https://api.coinbase.com/v2/prices/BTC-USD/spot",
			expectedErr: false,
		},
		{
			name: "multiple currency pairs",
			cps: []mmtypes.Ticker{
				constants.BITCOIN_USD,
				constants.ETHEREUM_USD,
			},
			url:         "",
			expectedErr: true,
		},
		{
			name: "unknown currency",
			cps: []mmtypes.Ticker{
				mogusd,
			},
			url:         "",
			expectedErr: true,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			h, err := coinbase.NewAPIHandler(coinbase.DefaultMarketConfig, coinbase.DefaultAPIConfig)
			require.NoError(t, err)

			url, err := h.CreateURL(tc.cps)
			if tc.expectedErr {
				require.Error(t, err)
			} else {
				require.NoError(t, err)
				require.Equal(t, tc.url, url)
			}
		})
	}
}

func TestParseResponse(t *testing.T) {
	testCases := []struct {
		name     string
		cps      []mmtypes.Ticker
		response *http.Response
		expected providertypes.GetResponse[mmtypes.Ticker, *big.Int]
	}{
		{
			name: "valid",
			cps: []mmtypes.Ticker{
				constants.BITCOIN_USD,
			},
			response: testutils.CreateResponseFromJSON(
				`
{
	"data": {
		"amount": "1020.25",
		"currency": "USD"
	}
}
	`,
			),
			expected: providertypes.NewGetResponse(
				map[mmtypes.Ticker]providertypes.Result[*big.Int]{
					constants.BITCOIN_USD: {
						Value: big.NewInt(102025000000),
					},
				},
				map[mmtypes.Ticker]error{},
			),
		},
		{
			name: "malformed response",
			cps:  []mmtypes.Ticker{constants.BITCOIN_USD},
			response: testutils.CreateResponseFromJSON(
				`
{
	"data": {
		"amount": "1020.25",
		"currency": "USD",
	}
}
	`,
			),
			expected: providertypes.NewGetResponse(
				map[mmtypes.Ticker]providertypes.Result[*big.Int]{},
				map[mmtypes.Ticker]error{
					constants.BITCOIN_USD: fmt.Errorf("bad format"),
				},
			),
		},
		{
			name: "unable to parse float",
			cps:  []mmtypes.Ticker{constants.BITCOIN_USD},
			response: testutils.CreateResponseFromJSON(
				`
{
	"data": {
		"amount": "$1020.25",
		"currency": "USD"
	}
}
	`,
			),
			expected: providertypes.NewGetResponse(
				map[mmtypes.Ticker]providertypes.Result[*big.Int]{},
				map[mmtypes.Ticker]error{
					constants.BITCOIN_USD: fmt.Errorf("bad format"),
				},
			),
		},
		{
			name: "unable to parse json",
			cps:  []mmtypes.Ticker{constants.BITCOIN_USD},
			response: testutils.CreateResponseFromJSON(
				`
toms obvious but not minimal language
	`,
			),
			expected: providertypes.NewGetResponse(
				map[mmtypes.Ticker]providertypes.Result[*big.Int]{},
				map[mmtypes.Ticker]error{
					constants.BITCOIN_USD: fmt.Errorf("bad format"),
				},
			),
		},
		{
			name: "multiple currency pairs to parse response for",
			cps: []mmtypes.Ticker{
				constants.BITCOIN_USD,
				constants.ETHEREUM_USD,
			},
			response: testutils.CreateResponseFromJSON(
				`
{
	"data": {
		"amount": "1020.25",
		"currency": "USD"
	}
}
	`,
			),
			expected: providertypes.NewGetResponse(
				map[mmtypes.Ticker]providertypes.Result[*big.Int]{},
				map[mmtypes.Ticker]error{
					constants.BITCOIN_USD:  fmt.Errorf("multiple cps"),
					constants.ETHEREUM_USD: fmt.Errorf("multiple cps"),
				},
			),
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			h, err := coinbase.NewAPIHandler(coinbase.DefaultMarketConfig, coinbase.DefaultAPIConfig)
			require.NoError(t, err)

			now := time.Now()
			resp := h.ParseResponse(tc.cps, tc.response)

			require.Len(t, resp.Resolved, len(tc.expected.Resolved))
			require.Len(t, resp.UnResolved, len(tc.expected.UnResolved))

			for cp, result := range tc.expected.Resolved {
				require.Contains(t, resp.Resolved, cp)
				r := resp.Resolved[cp]
				require.Equal(t, result.Value, r.Value)
				require.True(t, r.Timestamp.After(now))
			}

			for cp := range tc.expected.UnResolved {
				require.Contains(t, resp.UnResolved, cp)
				require.Error(t, resp.UnResolved[cp])
			}
		})
	}
}
