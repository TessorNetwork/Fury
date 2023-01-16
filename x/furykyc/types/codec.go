package types

import (
	"github.com/cosmos/cosmos-sdk/codec"
	cdctypes "github.com/cosmos/cosmos-sdk/codec/types"
	sdk "github.com/cosmos/cosmos-sdk/types"

	"github.com/cosmos/cosmos-sdk/types/msgservice"
)

func RegisterCodec(cdc *codec.LegacyAmino) {
	cdc.RegisterConcrete(&MsgInviteUser{}, "fury/MsgInviteUser", nil)
	cdc.RegisterConcrete(&MsgDepositIntoLiquidityPool{}, "fury/MsgDepositIntoLiquidityPool", nil)
	cdc.RegisterConcrete(&MsgBuyMembership{}, "fury/MsgBuyMembership", nil)
	cdc.RegisterConcrete(&MsgAddTsp{}, "fury/MsgAddTsp", nil)
	cdc.RegisterConcrete(&MsgRemoveTsp{}, "fury/MsgRemoveTsp", nil)
	cdc.RegisterConcrete(&MsgRemoveMembership{}, "fury/MsgRemoveMembership", nil)
	cdc.RegisterConcrete(&MsgSetMembership{}, "fury/MsgSetMembership", nil)

}

func RegisterInterfaces(registry cdctypes.InterfaceRegistry) {
	registry.RegisterImplementations((*sdk.Msg)(nil),
		&MsgInviteUser{},
		&MsgDepositIntoLiquidityPool{},
		&MsgBuyMembership{},
		&MsgAddTsp{},
		&MsgRemoveTsp{},
		&MsgRemoveMembership{},
		&MsgSetMembership{},
	)

	msgservice.RegisterMsgServiceDesc(registry, &_Msg_serviceDesc)
}

var (
	ModuleCdc = codec.NewProtoCodec(cdctypes.NewInterfaceRegistry())
)
