package types

import sdk "github.com/cosmos/cosmos-sdk/types"

type CodeType = sdk.CodeType

const (
	DefaultCodespace sdk.CodespaceType = ModuleName

	CodeInvalidFarmPool  CodeType = 101
	CodeInvalidLockInfo  CodeType = 102
	CodeInvalidInput     CodeType = 103
	CodePoolAlreadyExist CodeType = 104
	CodeTokenNotExist    CodeType = 105
	CodeInvalidAddress            = sdk.CodeInvalidAddress
	CodeUnknownRequest            = sdk.CodeUnknownRequest
)

// ErrNoFarmPoolFound returns an error when a farm pool doesn't exist
func ErrNoFarmPoolFound(codespace sdk.CodespaceType, poolName string) sdk.Error {
	return sdk.NewError(codespace, CodeInvalidFarmPool, "failed. Farm pool %s does not exist", poolName)
}

// ErrPoolAlreadyExist returns an error when a pool exist
func ErrPoolAlreadyExist(codespace sdk.CodespaceType, poolName string) sdk.Error {
	return sdk.NewError(codespace, CodeTokenNotExist, "failed. farm pool %s already exists", poolName)
}

// ErrTokenNotExist returns an error when a token not exists
func ErrTokenNotExist(codespace sdk.CodespaceType, tokenName string) sdk.Error {
	return sdk.NewError(codespace, CodeTokenNotExist, "failed. lock token %s not exists", tokenName)
}

// ErrNoLockInfoFound returns an error when an address doesn't have any lock infos
func ErrNoLockInfoFound(codespace sdk.CodespaceType, addr string) sdk.Error {
	return sdk.NewError(codespace, CodeInvalidLockInfo, "failed. %s doesn't have any lock infos", addr)
}

// ErrInvalidProvidedDenom returns an error when the remaining amount in yieldedTokenInfo is not zero
func ErrRemainingAmountNotZero(codespace sdk.CodespaceType, amount string) sdk.Error {
	return sdk.NewError(codespace, CodeInvalidInput,
		"failed. The remaining amount is %s, so it's not enable to provide token repeatedly util amount become zero",
		amount)
}


// ErrInvalidTokenOwner returns an error when an input address is not the owner of token
func ErrInvalidTokenOwner(codespace sdk.CodespaceType, addr string, token string) sdk.Error {
	return sdk.NewError(codespace, CodeInvalidInput, "failed. %s isn't the owner of token %s", addr, token)
}

// ErrInvalidProvidedDenom returns an error when it provides an unmatched token name
func ErrInvalidProvidedDenom(codespace sdk.CodespaceType, symbolLocked string, token string) sdk.Error {
	return sdk.NewError(codespace, CodeInvalidInput,
		"failed. The locked coin name in the farm is %s, not %s", symbolLocked, token)
}

// ErrInvalidAmount returns an error when an input amount is invaild
func ErrInvalidAmount(codespace sdk.CodespaceType, amount string) sdk.Error {
	return sdk.NewError(codespace, CodeInvalidInput, "failed. The input amount %s is invaild", amount)
}

// ErrInvalidStartHeight returns an error when the start_height_to_yield parameter is invaild
func ErrInvalidStartHeight(codespace sdk.CodespaceType) sdk.Error {
	return sdk.NewError(codespace, CodeInvalidInput, "failed. The start height to yield is less than current height")
}

// ErrInvalidInput returns an error when an input parameter is invaild
func ErrInvalidInput(codespace sdk.CodespaceType, input string) sdk.Error {
	return sdk.NewError(codespace, CodeInvalidInput, "failed. The input parameter %s is invaild", input)
}

// ErrNilAddress returns an error when an empty address appears
func ErrNilAddress(codespace sdk.CodespaceType) sdk.Error {
	return sdk.NewError(codespace, CodeInvalidAddress, "failed. Address is nil")
}
