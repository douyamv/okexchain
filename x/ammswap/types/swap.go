package types

import (
	"errors"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/x/supply"
	token "github.com/okex/okexchain/x/token/types"

	"fmt"
	"strings"
)

// PoolTokenPrefix defines pool token prefix name
const PoolTokenPrefix = "ammswap_"

// SwapTokenPair defines token pair exchange
type SwapTokenPair struct {
	QuotePooledCoin sdk.DecCoin `json:"quote_pooled_coin"` // The volume of quote token in the token pair exchange pool
	BasePooledCoin  sdk.DecCoin `json:"base_pooled_coin"`  // The volume of base token in the token pair exchange pool
	PoolTokenName   string      `json:"pool_token_name"`   // The name of pool token
}

// NewSwapTokenPair is a constructor function for SwapTokenPair
func NewSwapTokenPair(quotePooledCoin sdk.DecCoin, basePooledCoin sdk.DecCoin, poolTokenName string) *SwapTokenPair {
	swapTokenPair := &SwapTokenPair{
		QuotePooledCoin: quotePooledCoin,
		BasePooledCoin:  basePooledCoin,
		PoolTokenName:   poolTokenName,
	}
	return swapTokenPair
}

// String implement fmt.Stringer
func (s SwapTokenPair) String() string {
	return strings.TrimSpace(fmt.Sprintf(`QuotePooledCoin: %s
BasePooledCoin: %s
PoolTokenName: %s`, s.QuotePooledCoin.String(), s.BasePooledCoin.String(), s.PoolTokenName))
}

// TokenPairName defines token pair
func (s SwapTokenPair) TokenPairName() string {
	return s.BasePooledCoin.Denom + "_" + s.QuotePooledCoin.Denom
}

// InitPoolToken default pool token
func InitPoolToken(poolTokenName string) token.Token {
	return token.Token{
		Description:         poolTokenName,
		Symbol:              poolTokenName,
		OriginalSymbol:      poolTokenName,
		WholeName:           poolTokenName,
		OriginalTotalSupply: sdk.NewDec(0),
		Owner:               supply.NewModuleAddress(ModuleName),
		Type:                GenerateTokenType,
		Mintable:            true,
	}
}

func GetSwapTokenPairName(token1, token2 string) string {
	if token1 < token2 {
		return token1 + "_" + token2
	}
	return token2 + "_" + token1
}

func ValidateBaseAndQuoteTokenName(baseTokenName, quoteTokenName string) error {
	if baseTokenName > quoteTokenName {
		return errors.New("The lexicographic order of BaseTokenName must be less than QuoteTokenName, it may be ok to reverse BaseTokenName and QuoteTokenName")
	}else if baseTokenName == quoteTokenName {
		return errors.New("BaseTokenName should not equal to QuoteTokenName")
	}
	if err := ValidateSwapTokenName(baseTokenName); err != nil {
		return err
	}

	if err := ValidateSwapTokenName(quoteTokenName); err != nil {
		return err
	}
	return nil
}

func ValidateSwapTokenName(amountName string) error {
	if sdk.ValidateDenom(amountName) != nil {
		return errors.New(fmt.Sprintf("invalid token name: %s", amountName))
	}
	if token.NotAllowedOriginSymbol(amountName) {
		return errors.New(fmt.Sprintf("liquidity-pool-token(with prefix \"%s\") is not allowed to be a base or quote token", PoolTokenPrefix))
	}
	return nil
}

func GetPoolTokenName(token1, token2 string) string {
	return PoolTokenPrefix + GetSwapTokenPairName(token1, token2)
}