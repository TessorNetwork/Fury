package v3_0_0

import (
	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/codec"
	sdk "github.com/cosmos/cosmos-sdk/types"
	capabilitytypes "github.com/cosmos/cosmos-sdk/x/capability/types"
	evidencetypes "github.com/cosmos/cosmos-sdk/x/evidence/types"
	sdkLegacy "github.com/cosmos/cosmos-sdk/x/genutil/legacy/v040"
	"github.com/cosmos/cosmos-sdk/x/genutil/types"
	govtypes "github.com/cosmos/cosmos-sdk/x/gov/types"
	ibctransfertypes "github.com/cosmos/ibc-go/v3/modules/apps/transfer/types"
	ibc "github.com/cosmos/ibc-go/v3/modules/core/types"

	v220government "github.com/tessornetwork/fury/x/government/legacy/v2.2.0"
	v300government "github.com/tessornetwork/fury/x/government/legacy/v3.0.0"

	v220docs "github.com/tessornetwork/fury/x/documents/legacy/v2.2.0"
	v300docs "github.com/tessornetwork/fury/x/documents/legacy/v3.0.0"

	v220did "github.com/tessornetwork/fury/x/did/legacy/v2.2.0"
	v300did "github.com/tessornetwork/fury/x/did/legacy/v3.0.0"

	v220furymint "github.com/tessornetwork/fury/x/furymint/legacy/v2.2.0"
	v300furymint "github.com/tessornetwork/fury/x/furymint/legacy/v3.0.0"

	v220furykyc "github.com/tessornetwork/fury/x/furykyc/legacy/v2.2.0"
	v300furykyc "github.com/tessornetwork/fury/x/furykyc/legacy/v3.0.0"

	v220vbr "github.com/tessornetwork/fury/x/vbr/legacy/v2.2.0"
	v300vbr "github.com/tessornetwork/fury/x/vbr/legacy/v3.0.0"

	v300epochs "github.com/tessornetwork/fury/x/epochs/types"

	"github.com/CosmWasm/wasmd/x/wasm"
	wasmTypes "github.com/CosmWasm/wasmd/x/wasm/types"
)

func Migrate(appState types.AppMap, clientCtx client.Context) types.AppMap {
	v039Codec := codec.NewLegacyAmino()
	v040Codec := clientCtx.JSONCodec

	appState = sdkLegacy.Migrate(appState, clientCtx)

	if appState[v220government.ModuleName] != nil {
		var govGenState v220government.GenesisState
		v039Codec.MustUnmarshalJSON(appState[v220government.ModuleName], &govGenState)
		appState[v300government.ModuleName] = v040Codec.MustMarshalJSON(v300government.Migrate(govGenState))
	}

	if appState[v220did.ModuleName] != nil {
		var didGenState v220did.GenesisState
		v039Codec.MustUnmarshalJSON(appState[v220did.ModuleName], &didGenState)
		appState[v300did.ModuleName] = v040Codec.MustMarshalJSON(v300did.Migrate(didGenState))
		delete(appState, v220did.ModuleName)
	}

	if appState[v220docs.ModuleName] != nil {
		var docGenState v220docs.GenesisState
		v039Codec.MustUnmarshalJSON(appState[v220docs.ModuleName], &docGenState)
		appState[v300docs.ModuleName] = v040Codec.MustMarshalJSON(v300docs.Migrate(docGenState))
	}
	//appState[v300docs.ModuleName] = appState[v220docs.ModuleName]

	if appState[v220furymint.ModuleName] != nil {
		var furymintGenState v220furymint.GenesisState
		v039Codec.MustUnmarshalJSON(appState[v220furymint.ModuleName], &furymintGenState)
		appState[v300furymint.ModuleName] = v040Codec.MustMarshalJSON(v300furymint.Migrate(furymintGenState))

	}

	if appState[v220furykyc.ModuleName] != nil {
		var furykycGenState v220furykyc.GenesisState
		v039Codec.MustUnmarshalJSON(appState[v220furykyc.ModuleName], &furykycGenState)
		appState[v300furykyc.ModuleName] = v040Codec.MustMarshalJSON(v300furykyc.Migrate(furykycGenState))

	}

	if appState[v220vbr.ModuleName] != nil {
		var vbrGenState v220vbr.GenesisState
		v039Codec.MustUnmarshalJSON(appState[v220vbr.ModuleName], &vbrGenState)
		appState[v300vbr.ModuleName] = v040Codec.MustMarshalJSON(v300vbr.Migrate(vbrGenState))
	}

	if appState[v300epochs.ModuleName] == nil {
		appState[v300epochs.ModuleName] = v040Codec.MustMarshalJSON(v300epochs.DefaultGenesis())
	}

	//appState[wasm.ModuleName] = wasmKeeper.InitGenesis()
	wasmModule := &wasmTypes.GenesisState{}

	wasmModule.Params.InstantiateDefaultPermission = 3
	wasmModule.Params.CodeUploadAccess.Permission = 3
	//wasmModule.Params.MaxWasmCodeSize = 1228800

	appState[wasm.ModuleName] = v040Codec.MustMarshalJSON(wasmModule)
	appState[ibctransfertypes.ModuleName] = v040Codec.MustMarshalJSON(ibctransfertypes.DefaultGenesisState())

	appState["ibc"] = v040Codec.MustMarshalJSON(ibc.DefaultGenesisState())
	appState[capabilitytypes.ModuleName] = v040Codec.MustMarshalJSON(capabilitytypes.DefaultGenesis())
	appState[evidencetypes.ModuleName] = v040Codec.MustMarshalJSON(evidencetypes.DefaultGenesisState())
	appState[govtypes.ModuleName] = v040Codec.MustMarshalJSON(govtypes.DefaultGenesisState())
	if appState[govtypes.ModuleName] != nil {
		var govGenState govtypes.GenesisState
		v040Codec.MustUnmarshalJSON(appState[govtypes.ModuleName], &govGenState)
		coins := sdk.NewCoins(sdk.NewCoin("ufury", sdk.NewInt(50000000000)))
		govGenState.DepositParams.MinDeposit = coins
		appState[govtypes.ModuleName] = v040Codec.MustMarshalJSON(&govGenState)
	}

	return appState

}
