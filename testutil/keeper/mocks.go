// Code generated by MockGen. DO NOT EDIT.
// Source: x/alliance/types/keeper_interfaces.go

// Package mock_types is a generated GoMock package.
package keeper

import (
	reflect "reflect"
	time "time"

	math "cosmossdk.io/math"
	types "github.com/cosmos/cosmos-sdk/types"
	types0 "github.com/cosmos/cosmos-sdk/x/auth/types"
	types1 "github.com/cosmos/cosmos-sdk/x/staking/types"
	gomock "github.com/golang/mock/gomock"
)

// MockStakingKeeper is a mock of StakingKeeper interface.
type MockStakingKeeper struct {
	ctrl     *gomock.Controller
	recorder *MockStakingKeeperMockRecorder
}

// MockStakingKeeperMockRecorder is the mock recorder for MockStakingKeeper.
type MockStakingKeeperMockRecorder struct {
	mock *MockStakingKeeper
}

// NewMockStakingKeeper creates a new mock instance.
func NewMockStakingKeeper(ctrl *gomock.Controller) *MockStakingKeeper {
	mock := &MockStakingKeeper{ctrl: ctrl}
	mock.recorder = &MockStakingKeeperMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockStakingKeeper) EXPECT() *MockStakingKeeperMockRecorder {
	return m.recorder
}

// BeginRedelegation mocks base method.
func (m *MockStakingKeeper) BeginRedelegation(ctx types.Context, delAddr types.AccAddress, valSrcAddr, valDstAddr types.ValAddress, sharesAmount types.Dec) (time.Time, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "BeginRedelegation", ctx, delAddr, valSrcAddr, valDstAddr, sharesAmount)
	ret0, _ := ret[0].(time.Time)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// BeginRedelegation indicates an expected call of BeginRedelegation.
func (mr *MockStakingKeeperMockRecorder) BeginRedelegation(ctx, delAddr, valSrcAddr, valDstAddr, sharesAmount interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "BeginRedelegation", reflect.TypeOf((*MockStakingKeeper)(nil).BeginRedelegation), ctx, delAddr, valSrcAddr, valDstAddr, sharesAmount)
}

// BondDenom mocks base method.
func (m *MockStakingKeeper) BondDenom(ctx types.Context) string {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "BondDenom", ctx)
	ret0, _ := ret[0].(string)
	return ret0
}

// BondDenom indicates an expected call of BondDenom.
func (mr *MockStakingKeeperMockRecorder) BondDenom(ctx interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "BondDenom", reflect.TypeOf((*MockStakingKeeper)(nil).BondDenom), ctx)
}

// Delegate mocks base method.
func (m *MockStakingKeeper) Delegate(ctx types.Context, delAddr types.AccAddress, bondAmt math.Int, tokenSrc types1.BondStatus, validator types1.Validator, subtractAccount bool) (types.Dec, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Delegate", ctx, delAddr, bondAmt, tokenSrc, validator, subtractAccount)
	ret0, _ := ret[0].(types.Dec)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Delegate indicates an expected call of Delegate.
func (mr *MockStakingKeeperMockRecorder) Delegate(ctx, delAddr, bondAmt, tokenSrc, validator, subtractAccount interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Delegate", reflect.TypeOf((*MockStakingKeeper)(nil).Delegate), ctx, delAddr, bondAmt, tokenSrc, validator, subtractAccount)
}

// GetValidator mocks base method.
func (m *MockStakingKeeper) GetValidator(ctx types.Context, addr types.ValAddress) (types1.Validator, bool) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetValidator", ctx, addr)
	ret0, _ := ret[0].(types1.Validator)
	ret1, _ := ret[1].(bool)
	return ret0, ret1
}

// GetValidator indicates an expected call of GetValidator.
func (mr *MockStakingKeeperMockRecorder) GetValidator(ctx, addr interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetValidator", reflect.TypeOf((*MockStakingKeeper)(nil).GetValidator), ctx, addr)
}

// RemoveRedelegation mocks base method.
func (m *MockStakingKeeper) RemoveRedelegation(ctx types.Context, red types1.Redelegation) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "RemoveRedelegation", ctx, red)
}

// RemoveRedelegation indicates an expected call of RemoveRedelegation.
func (mr *MockStakingKeeperMockRecorder) RemoveRedelegation(ctx, red interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "RemoveRedelegation", reflect.TypeOf((*MockStakingKeeper)(nil).RemoveRedelegation), ctx, red)
}

// Unbond mocks base method.
func (m *MockStakingKeeper) Unbond(ctx types.Context, delAddr types.AccAddress, valAddr types.ValAddress, shares types.Dec) (math.Int, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Unbond", ctx, delAddr, valAddr, shares)
	ret0, _ := ret[0].(math.Int)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Unbond indicates an expected call of Unbond.
func (mr *MockStakingKeeperMockRecorder) Unbond(ctx, delAddr, valAddr, shares interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Unbond", reflect.TypeOf((*MockStakingKeeper)(nil).Unbond), ctx, delAddr, valAddr, shares)
}

// UnbondingTime mocks base method.
func (m *MockStakingKeeper) UnbondingTime(ctx types.Context) time.Duration {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UnbondingTime", ctx)
	ret0, _ := ret[0].(time.Duration)
	return ret0
}

// UnbondingTime indicates an expected call of UnbondingTime.
func (mr *MockStakingKeeperMockRecorder) UnbondingTime(ctx interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UnbondingTime", reflect.TypeOf((*MockStakingKeeper)(nil).UnbondingTime), ctx)
}

// ValidateUnbondAmount mocks base method.
func (m *MockStakingKeeper) ValidateUnbondAmount(ctx types.Context, delAddr types.AccAddress, valAddr types.ValAddress, amt math.Int) (types.Dec, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ValidateUnbondAmount", ctx, delAddr, valAddr, amt)
	ret0, _ := ret[0].(types.Dec)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ValidateUnbondAmount indicates an expected call of ValidateUnbondAmount.
func (mr *MockStakingKeeperMockRecorder) ValidateUnbondAmount(ctx, delAddr, valAddr, amt interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ValidateUnbondAmount", reflect.TypeOf((*MockStakingKeeper)(nil).ValidateUnbondAmount), ctx, delAddr, valAddr, amt)
}

// MockBankKeeper is a mock of BankKeeper interface.
type MockBankKeeper struct {
	ctrl     *gomock.Controller
	recorder *MockBankKeeperMockRecorder
}

// MockBankKeeperMockRecorder is the mock recorder for MockBankKeeper.
type MockBankKeeperMockRecorder struct {
	mock *MockBankKeeper
}

// NewMockBankKeeper creates a new mock instance.
func NewMockBankKeeper(ctrl *gomock.Controller) *MockBankKeeper {
	mock := &MockBankKeeper{ctrl: ctrl}
	mock.recorder = &MockBankKeeperMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockBankKeeper) EXPECT() *MockBankKeeperMockRecorder {
	return m.recorder
}

// BurnCoins mocks base method.
func (m *MockBankKeeper) BurnCoins(ctx types.Context, moduleName string, amt types.Coins) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "BurnCoins", ctx, moduleName, amt)
	ret0, _ := ret[0].(error)
	return ret0
}

// BurnCoins indicates an expected call of BurnCoins.
func (mr *MockBankKeeperMockRecorder) BurnCoins(ctx, moduleName, amt interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "BurnCoins", reflect.TypeOf((*MockBankKeeper)(nil).BurnCoins), ctx, moduleName, amt)
}

// GetAllBalances mocks base method.
func (m *MockBankKeeper) GetAllBalances(ctx types.Context, addr types.AccAddress) types.Coins {
        m.ctrl.T.Helper()
        ret := m.ctrl.Call(m, "GetAllBalances", ctx, addr)
        ret0, _ := ret[0].(types.Coins)
        return ret0
}

// GetAllBalances indicates an expected call of GetAllBalances.
func (mr *MockBankKeeperMockRecorder) GetAllBalances(ctx, addr interface{}) *gomock.Call {
        mr.mock.ctrl.T.Helper()
        return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetAllBalances", reflect.TypeOf((*MockBankKeeper)(nil).GetAllBalances), ctx, addr)
}

// MintCoins mocks base method.
func (m *MockBankKeeper) MintCoins(ctx types.Context, moduleName string, amt types.Coins) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "MintCoins", ctx, moduleName, amt)
	ret0, _ := ret[0].(error)
	return ret0
}

// MintCoins indicates an expected call of MintCoins.
func (mr *MockBankKeeperMockRecorder) MintCoins(ctx, moduleName, amt interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "MintCoins", reflect.TypeOf((*MockBankKeeper)(nil).MintCoins), ctx, moduleName, amt)
}

// SendCoinsFromAccountToModule mocks base method.
func (m *MockBankKeeper) SendCoinsFromAccountToModule(ctx types.Context, senderAddr types.AccAddress, recipientModule string, amt types.Coins) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "SendCoinsFromAccountToModule", ctx, senderAddr, recipientModule, amt)
	ret0, _ := ret[0].(error)
	return ret0
}

// SendCoinsFromAccountToModule indicates an expected call of SendCoinsFromAccountToModule.
func (mr *MockBankKeeperMockRecorder) SendCoinsFromAccountToModule(ctx, senderAddr, recipientModule, amt interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SendCoinsFromAccountToModule", reflect.TypeOf((*MockBankKeeper)(nil).SendCoinsFromAccountToModule), ctx, senderAddr, recipientModule, amt)
}

// SendCoinsFromModuleToAccount mocks base method.
func (m *MockBankKeeper) SendCoinsFromModuleToAccount(ctx types.Context, senderModule string, recipientAddr types.AccAddress, amt types.Coins) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "SendCoinsFromModuleToAccount", ctx, senderModule, recipientAddr, amt)
	ret0, _ := ret[0].(error)
	return ret0
}

// SendCoinsFromModuleToAccount indicates an expected call of SendCoinsFromModuleToAccount.
func (mr *MockBankKeeperMockRecorder) SendCoinsFromModuleToAccount(ctx, senderModule, recipientAddr, amt interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SendCoinsFromModuleToAccount", reflect.TypeOf((*MockBankKeeper)(nil).SendCoinsFromModuleToAccount), ctx, senderModule, recipientAddr, amt)
}

// MockAccountKeeper is a mock of AccountKeeper interface.
type MockAccountKeeper struct {
	ctrl     *gomock.Controller
	recorder *MockAccountKeeperMockRecorder
}

// MockAccountKeeperMockRecorder is the mock recorder for MockAccountKeeper.
type MockAccountKeeperMockRecorder struct {
	mock *MockAccountKeeper
}

// NewMockAccountKeeper creates a new mock instance.
func NewMockAccountKeeper(ctrl *gomock.Controller) *MockAccountKeeper {
	mock := &MockAccountKeeper{ctrl: ctrl}
	mock.recorder = &MockAccountKeeperMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockAccountKeeper) EXPECT() *MockAccountKeeperMockRecorder {
	return m.recorder
}

// GetAccount mocks base method.
func (m *MockAccountKeeper) GetAccount(ctx types.Context, addr types.AccAddress) types0.AccountI {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetAccount", ctx, addr)
	ret0, _ := ret[0].(types0.AccountI)
	return ret0
}

// GetAccount indicates an expected call of GetAccount.
func (mr *MockAccountKeeperMockRecorder) GetAccount(ctx, addr interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetAccount", reflect.TypeOf((*MockAccountKeeper)(nil).GetAccount), ctx, addr)
}

// GetModuleAccount mocks base method.
func (m *MockAccountKeeper) GetModuleAccount(ctx types.Context, name string) types0.ModuleAccountI {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetModuleAccount", ctx, name)
	ret0, _ := ret[0].(types0.ModuleAccountI)
	return ret0
}

// GetModuleAccount indicates an expected call of GetModuleAccount.
func (mr *MockAccountKeeperMockRecorder) GetModuleAccount(ctx, name interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetModuleAccount", reflect.TypeOf((*MockAccountKeeper)(nil).GetModuleAccount), ctx, name)
}

// GetModuleAddress mocks base method.
func (m *MockAccountKeeper) GetModuleAddress(name string) types.AccAddress {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetModuleAddress", name)
	ret0, _ := ret[0].(types.AccAddress)
	return ret0
}

// GetModuleAddress indicates an expected call of GetModuleAddress.
func (mr *MockAccountKeeperMockRecorder) GetModuleAddress(name interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetModuleAddress", reflect.TypeOf((*MockAccountKeeper)(nil).GetModuleAddress), name)
}
