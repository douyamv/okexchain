package keeper

import (
	abci "github.com/tendermint/tendermint/abci/types"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/okex/okexchain/x/ammswap/types"
)

// NewQuerier creates a new querier for swap clients.
func NewQuerier(k Keeper) sdk.Querier {
	return func(ctx sdk.Context, path []string, req abci.RequestQuery) (res []byte, err sdk.Error) {
		switch path[0] {
		case types.QuerySwapTokenPair:
			return querySwapTokenPair(ctx, path[1:], req, k)
		case types.QueryParams:
			return queryParams(ctx, path[1:], req, k)
		case types.QuerySwapTokenPairs:
			return querySwapTokenPairs(ctx, path[1:], req, k)
		case types.QueryRedeemableAssets:
			return queryRedeemableAssets(ctx, path[1:], req, k)
		case types.QueryBuyAmount:
			return queryBuyAmount(ctx, path[1:], req, k)
		default:
			return nil, sdk.ErrUnknownRequest("unknown swap query endpoint")
		}
	}
}

// nolint
func querySwapTokenPair(
	ctx sdk.Context, path []string, req abci.RequestQuery, keeper Keeper,
) (res []byte, err sdk.Error) {
	baseTokenName := path[0]
	quoteTokenName := path[1]
	errToken := types.ValidateBaseAndQuoteTokenName(baseTokenName, quoteTokenName)
	if errToken != nil {
		return nil, sdk.ErrUnknownRequest(errToken.Error())
	}
	tokenPairName := types.GetSwapTokenPairName(baseTokenName, quoteTokenName)
	tokenPair, errSwapTokenPair := keeper.GetSwapTokenPair(ctx, tokenPairName)
	if errSwapTokenPair != nil {
		return nil, sdk.ErrUnknownRequest(errSwapTokenPair.Error())
	}
	bz := keeper.cdc.MustMarshalJSON(tokenPair)
	return bz, nil
}

// nolint
func queryBuyAmount(
	ctx sdk.Context, path []string, req abci.RequestQuery, keeper Keeper,
) ([]byte, sdk.Error) {
	var queryParams types.QueryBuyAmountParams
	err := keeper.cdc.UnmarshalJSON(req.Data, &queryParams)
	if err != nil {
		return nil, sdk.ErrUnknownRequest(sdk.AppendMsgToErr("incorrectly formatted request data", err.Error()))
	}
	errToken := types.ValidateSwapTokenName(queryParams.TokenToBuy)
	if errToken != nil {
		return nil, sdk.ErrUnknownRequest(errToken.Error())
	}
	errToken = types.ValidateSwapTokenName(queryParams.SoldToken.Denom)
	if errToken != nil {
		return nil, sdk.ErrUnknownRequest(errToken.Error())
	}
	params := keeper.GetParams(ctx)
	var buyAmount sdk.Dec
	swapTokenPair := types.GetSwapTokenPairName(queryParams.SoldToken.Denom, queryParams.TokenToBuy)
	tokenPair, errTokenPair := keeper.GetSwapTokenPair(ctx, swapTokenPair)
	if errTokenPair != nil {
		return nil, sdk.ErrUnknownRequest(errTokenPair.Error())
	}
	buyAmount = CalculateTokenToBuy(tokenPair, queryParams.SoldToken, queryParams.TokenToBuy, params).Amount

	bz := keeper.cdc.MustMarshalJSON(buyAmount)

	return bz, nil
}

func queryParams(ctx sdk.Context, path []string, req abci.RequestQuery, keeper Keeper) ([]byte, sdk.Error) {
	return keeper.cdc.MustMarshalJSON(keeper.GetParams(ctx)), nil
}

// nolint
func querySwapTokenPairs(ctx sdk.Context, path []string, req abci.RequestQuery, keeper Keeper) (res []byte,
	err sdk.Error) {
	return keeper.cdc.MustMarshalJSON(keeper.GetSwapTokenPairs(ctx)), nil
}


// nolint
func queryRedeemableAssets(ctx sdk.Context, path []string, req abci.RequestQuery, keeper Keeper) (res []byte,
	err sdk.Error) {
	baseTokenName := path[0]
	quoteTokenName := path[1]
	errToken := types.ValidateBaseAndQuoteTokenName(baseTokenName, quoteTokenName)
	if errToken != nil {
		return nil, sdk.ErrUnknownRequest(errToken.Error())
	}
	liquidity, decErr := sdk.NewDecFromStr(path[2])
	if decErr != nil {
		return nil, sdk.ErrUnknownRequest("invalid params: liquidity")
	}
	var tokenList sdk.DecCoins
	baseToken, quoteToken, redeemErr := keeper.GetRedeemableAssets(ctx, baseTokenName, quoteTokenName, liquidity)
	if redeemErr != nil {
		return nil, sdk.ErrUnknownRequest(redeemErr.Error())
	}
	tokenList = append(tokenList, baseToken, quoteToken)
	bz := keeper.cdc.MustMarshalJSON(tokenList)
	return bz, nil
}