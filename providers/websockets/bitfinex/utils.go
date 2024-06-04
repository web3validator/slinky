package bitfinex

import (
	"github.com/skip-mev/slinky/oracle/config"
)

const (
	// Name is the name of the BitFinex provider.
	Name = "bitfinex_ws"

	// URLProd is the public BitFinex Websocket URL.
	URLProd = "wss://api-pub.bitfinex.com/ws/2"

	// MaxSubscriptionsPerConnection is the maximum number of subscriptions that
	// can be assigned to a single connection for this provider.
	MaxSubscriptionsPerConnection = 30
)

// DefaultWebSocketConfig is the default configuration for the BitFinex Websocket.
var DefaultWebSocketConfig = config.WebSocketConfig{
	Name:                          Name,
	Enabled:                       true,
	MaxBufferSize:                 1000,
	ReconnectionTimeout:           config.DefaultReconnectionTimeout,
	PostConnectionTimeout:         config.DefaultPostConnectionTimeout,
	Endpoints:                     []config.Endpoint{{URL: URLProd}},
	ReadBufferSize:                config.DefaultReadBufferSize,
	WriteBufferSize:               config.DefaultWriteBufferSize,
	HandshakeTimeout:              config.DefaultHandshakeTimeout,
	EnableCompression:             config.DefaultEnableCompression,
	ReadTimeout:                   config.DefaultReadTimeout,
	WriteInterval:                 config.DefaultWriteInterval,
	WriteTimeout:                  config.DefaultWriteTimeout,
	PingInterval:                  config.DefaultPingInterval,
	MaxReadErrorCount:             config.DefaultMaxReadErrorCount,
	MaxSubscriptionsPerConnection: MaxSubscriptionsPerConnection,
}
