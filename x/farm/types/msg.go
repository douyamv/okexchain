package types

import (
	"strconv"

	sdk "github.com/cosmos/cosmos-sdk/types"
)

type MsgCreatePool struct {
	Address    sdk.AccAddress `json:"address" yaml:"address"`
	PoolName   string         `json:"pool_name" yaml:"pool_name"`
	LockToken  string         `json:"lock_token", yaml:"lock_token"`
	YieldToken string         `json:"yield_token", yaml:"yield_token"`
}

var _ sdk.Msg = MsgCreatePool{}

func NewMsgCreatePool(address sdk.AccAddress, poolName, lockToken, yieldToken string) MsgCreatePool {
	return MsgCreatePool{
		Address:    address,
		PoolName:   poolName,
		LockToken:  lockToken,
		YieldToken: yieldToken,
	}
}

func (m MsgCreatePool) Route() string {
	return RouterKey
}

func (m MsgCreatePool) Type() string {
	return "create_pool"
}

func (m MsgCreatePool) ValidateBasic() sdk.Error {
	if m.Address.Empty() {
		return ErrNilAddress(DefaultCodespace)
	}
	if m.PoolName == "" {
		return ErrInvalidInput(DefaultCodespace, m.PoolName)
	}
	if m.LockToken == "" {
		return ErrInvalidInput(DefaultCodespace, m.LockToken)
	}
	if m.YieldToken == "" {
		return ErrInvalidInput(DefaultCodespace, m.YieldToken)
	}
	return nil
}

func (m MsgCreatePool) GetSignBytes() []byte {
	bz := ModuleCdc.MustMarshalJSON(m)
	return sdk.MustSortJSON(bz)
}

func (m MsgCreatePool) GetSigners() []sdk.AccAddress {
	return []sdk.AccAddress{m.Address}
}

type MsgDestroyPool struct {
	Address  sdk.AccAddress `json:"address" yaml:"address"`
	PoolName string         `json:"pool_name" yaml:"pool_name"`
}

var _ sdk.Msg = MsgDestroyPool{}

func NewMsgDestroyPool(address sdk.AccAddress, poolName string) MsgDestroyPool {
	return MsgDestroyPool{
		Address:  address,
		PoolName: poolName,
	}
}

func (m MsgDestroyPool) Route() string {
	return RouterKey
}

func (m MsgDestroyPool) Type() string {
	return "destroy_pool"
}

func (m MsgDestroyPool) ValidateBasic() sdk.Error {
	if m.Address.Empty() {
		return ErrNilAddress(DefaultCodespace)
	}
	if m.PoolName == "" {
		return ErrInvalidInput(DefaultCodespace, m.PoolName)
	}
	return nil
}

func (m MsgDestroyPool) GetSignBytes() []byte {
	bz := ModuleCdc.MustMarshalJSON(m)
	return sdk.MustSortJSON(bz)
}

func (m MsgDestroyPool) GetSigners() []sdk.AccAddress {
	return []sdk.AccAddress{m.Address}
}

type MsgProvide struct {
	PoolName           string         `json:"pool_name" yaml:"pool_name"`
	Address            sdk.AccAddress `json:"address" yaml:"address"`
	Amount             sdk.DecCoin    `json:"amount" yaml:"amount"`
	YieldPerBlock      sdk.Dec        `json:"yield_per_block" yaml:"yield_per_block"`
	StartHeightToYield int64          `json:"start_height_to_yield" yaml:"start_height_to_yield"`
}

func NewMsgProvide(poolName string, address sdk.AccAddress, amount sdk.DecCoin,
	yieldPerBlock sdk.Dec, startHeightToYield int64) MsgProvide {
	return MsgProvide{
		PoolName:           poolName,
		Address:            address,
		Amount:             amount,
		YieldPerBlock:      yieldPerBlock,
		StartHeightToYield: startHeightToYield,
	}
}

var _ sdk.Msg = MsgProvide{}

func (m MsgProvide) Route() string {
	return RouterKey
}

func (m MsgProvide) Type() string {
	return "provide"
}

func (m MsgProvide) ValidateBasic() sdk.Error {
	if m.Address.Empty() {
		return ErrNilAddress(DefaultCodespace)
	}
	if m.Amount.Amount.LTE(sdk.ZeroDec()) || !m.Amount.IsValid() {
		return ErrInvalidAmount(DefaultCodespace, m.Amount.String())
	}
	if m.YieldPerBlock.LTE(sdk.ZeroDec()) {
		return ErrInvalidInput(DefaultCodespace, m.YieldPerBlock.String())
	}
	if m.YieldPerBlock.GT(m.Amount.Amount) {
		return ErrInsufficientAmount(DefaultCodespace, m.Amount.String())
	}
	if m.StartHeightToYield <= 0 {
		return ErrInvalidInput(DefaultCodespace, strconv.FormatInt(m.StartHeightToYield, 10))
	}
	return nil
}

func (m MsgProvide) GetSignBytes() []byte {
	bz := ModuleCdc.MustMarshalJSON(m)
	return sdk.MustSortJSON(bz)
}

func (m MsgProvide) GetSigners() []sdk.AccAddress {
	return []sdk.AccAddress{m.Address}
}

type MsgLock struct {
	PoolName string         `json:"pool_name" yaml:"pool_name"`
	Address  sdk.AccAddress `json:"address" yaml:"address"`
	Amount   sdk.DecCoin    `json:"amount" yaml:"amount"`
}

func NewMsgLock(poolName string, address sdk.AccAddress, amount sdk.DecCoin) MsgLock {
	return MsgLock{
		PoolName: poolName,
		Address:  address,
		Amount:   amount,
	}
}

var _ sdk.Msg = MsgLock{}

func (m MsgLock) Route() string {
	return RouterKey
}

func (m MsgLock) Type() string {
	return "lock"
}

func (m MsgLock) ValidateBasic() sdk.Error {
	if m.Address.Empty() {
		return ErrNilAddress(DefaultCodespace)
	}
	if m.Amount.Amount.LTE(sdk.ZeroDec()) || !m.Amount.IsValid() {
		return ErrInvalidAmount(DefaultCodespace, m.Amount.Amount.String())
	}
	return nil
}

func (m MsgLock) GetSignBytes() []byte {
	bz := ModuleCdc.MustMarshalJSON(m)
	return sdk.MustSortJSON(bz)
}

func (m MsgLock) GetSigners() []sdk.AccAddress {
	return []sdk.AccAddress{m.Address}
}

type MsgUnlock struct {
	PoolName string         `json:"pool_name" yaml:"pool_name"`
	Address  sdk.AccAddress `json:"address" yaml:"address"`
	Amount   sdk.DecCoin    `json:"amount" yaml:"amount"`
}

func NewMsgUnlock(poolName string, Address sdk.AccAddress, amount sdk.DecCoin) MsgUnlock {
	return MsgUnlock{
		PoolName: poolName,
		Address:  Address,
		Amount:   amount,
	}
}

var _ sdk.Msg = MsgUnlock{}

func (m MsgUnlock) Route() string {
	return RouterKey
}

func (m MsgUnlock) Type() string {
	return "unlock"
}

func (m MsgUnlock) ValidateBasic() sdk.Error {
	if m.Address.Empty() {
		return ErrNilAddress(DefaultCodespace)
	}
	if m.Amount.Amount.LTE(sdk.ZeroDec()) || !m.Amount.IsValid() {
		return ErrInvalidAmount(DefaultCodespace, m.Amount.Amount.String())
	}
	return nil
}

func (m MsgUnlock) GetSignBytes() []byte {
	bz := ModuleCdc.MustMarshalJSON(m)
	return sdk.MustSortJSON(bz)
}

func (m MsgUnlock) GetSigners() []sdk.AccAddress {
	return []sdk.AccAddress{m.Address}
}

type MsgClaim struct {
	PoolName string         `json:"pool_name" yaml:"pool_name"`
	Address  sdk.AccAddress `json:"address" yaml:"address"`
}

func NewMsgClaim(poolName string, Address sdk.AccAddress) MsgClaim {
	return MsgClaim{
		PoolName: poolName,
		Address:  Address,
	}
}

var _ sdk.Msg = MsgClaim{}

func (m MsgClaim) Route() string {
	return RouterKey
}

func (m MsgClaim) Type() string {
	return "claim"
}

func (m MsgClaim) ValidateBasic() sdk.Error {
	if m.Address.Empty() {
		return ErrNilAddress(DefaultCodespace)
	}
	return nil
}

func (m MsgClaim) GetSignBytes() []byte {
	bz := ModuleCdc.MustMarshalJSON(m)
	return sdk.MustSortJSON(bz)
}

func (m MsgClaim) GetSigners() []sdk.AccAddress {
	return []sdk.AccAddress{m.Address}
}
