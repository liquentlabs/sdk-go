package types

import (
	"github.com/cosmos/cosmos-sdk/codec"
	cdctypes "github.com/cosmos/cosmos-sdk/codec/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	authzcdc "github.com/cosmos/cosmos-sdk/x/authz/codec"
	// this line is used by starport scaffolding # 1
	"github.com/cosmos/cosmos-sdk/types/msgservice"
)

func RegisterCodec(cdc *codec.LegacyAmino) {
	cdc.RegisterConcrete(&MsgCreateDenom{}, "liquent/tokenfactory/create-denom", nil)
	cdc.RegisterConcrete(&MsgMint{}, "liquent/tokenfactory/mint", nil)
	cdc.RegisterConcrete(&MsgBurn{}, "liquent/tokenfactory/burn", nil)
	// nolint:all
	// cdc.RegisterConcrete(&MsgForceTransfer{}, "liquent/tokenfactory/force-transfer", nil)
	cdc.RegisterConcrete(&MsgChangeAdmin{}, "liquent/tokenfactory/change-admin", nil)
	cdc.RegisterConcrete(&MsgUpdateParams{}, "liquent/tokenfactory/update-params", nil)
	cdc.RegisterConcrete(&MsgSetDenomMetadata{}, "liquent/tokenfactory/set-denom-metadata", nil)
	cdc.RegisterConcrete(&Params{}, "liquent/tokenfactory/Params", nil)

}

func RegisterInterfaces(registry cdctypes.InterfaceRegistry) {
	registry.RegisterImplementations((*sdk.Msg)(nil),
		&MsgCreateDenom{},
		&MsgMint{},
		&MsgBurn{},
		// &MsgForceTransfer{},
		&MsgChangeAdmin{},
		&MsgUpdateParams{},
		&MsgSetDenomMetadata{},
	)

	msgservice.RegisterMsgServiceDesc(registry, &_Msg_serviceDesc)
}

var (
	ModuleCdc = codec.NewLegacyAmino()
)

func init() {
	RegisterCodec(ModuleCdc)
	// Register all Amino interfaces and concrete types on the authz Amino codec so that this can later be
	// used to properly serialize MsgGrant and MsgExec instances
	sdk.RegisterLegacyAminoCodec(ModuleCdc)
	RegisterCodec(authzcdc.Amino)

	ModuleCdc.Seal()
}
