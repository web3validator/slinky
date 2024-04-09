package types

import (
	"fmt"
)

// ValidateBasic validates the market map configuration and its expected configuration.
//
//	In particular, this will
//
//	1. Ensure that the market map is valid (ValidateBasic). This ensures that each of the provider's
//	   markets are supported by the market map.
//	2. Ensure that each provider config has a valid corresponding ticker.
func (mm *MarketMap) ValidateBasic() error {
	for _, market := range mm.Markets {
		if err := market.ValidateBasic(); err != nil {
			return err
		}

		for _, providerConfig := range market.ProviderConfigs {
			if providerConfig.NormalizeByPair != nil {
				if _, found := mm.Markets[providerConfig.NormalizeByPair.String()]; !found {
					return fmt.Errorf("provider's (%s) pair for normalization (%s) was not found in the marketmap", providerConfig.Name, providerConfig.NormalizeByPair.String())
				}
			}
		}
	}

	return nil
}

// String returns the string representation of the market map.
func (mm *MarketMap) String() string {
	return fmt.Sprintf(
		"MarketMap: {Markets %v}",
		mm.Markets,
	)
}

// ValidateBasic performs stateless validation of a Market.
func (m *Market) ValidateBasic() error {
	if err := m.Ticker.ValidateBasic(); err != nil {
		return err
	}

	if uint64(len(m.ProviderConfigs)) < m.Ticker.MinProviderCount {
		return fmt.Errorf("this ticker must have at least %d providers; got %d",
			m.Ticker.MinProviderCount,
			len(m.ProviderConfigs),
		)
	}

	seenProviders := make(map[string]struct{})
	for _, providerConfig := range m.ProviderConfigs {
		if err := providerConfig.ValidateBasic(); err != nil {
			return err
		}

		// check for duplicate providers
		key := providerConfig.Name + providerConfig.OffChainTicker
		if _, seen := seenProviders[key]; seen {
			return fmt.Errorf("duplicate (provider, off-chain-ticker) found: %s", key)
		}
		seenProviders[key] = struct{}{}

	}

	return nil
}

// String returns the string representation of the market.
func (m *Market) String() string {
	return fmt.Sprintf(
		"Market: {Ticker %v Providers: %v}", m.Ticker, m.ProviderConfigs,
	)
}

// Equal returns true if the MarketMap is equal to the given MarketMap.
func (mm *MarketMap) Equal(other MarketMap) bool {
	if len(mm.Markets) != len(other.Markets) {
		return false
	}

	for tickerStr, market := range mm.Markets {
		otherMarket, found := other.Markets[tickerStr]
		if !found {
			return false
		}

		if !market.Equal(otherMarket) {
			return false
		}
	}

	return true
}

// Equal returns true if the Market is equal to the given Market.
func (m *Market) Equal(other Market) bool {
	if !m.Ticker.Equal(other.Ticker) {
		return false
	}

	if len(m.ProviderConfigs) != len(other.ProviderConfigs) {
		return false
	}

	for i, providerConfig := range m.ProviderConfigs {
		if !providerConfig.Equal(other.ProviderConfigs[i]) {
			return false
		}
	}

	return true
}