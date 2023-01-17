package cli

import (
	"encoding/json"
	"fmt"
	"time"

	govTypes "github.com/tessornetwork/fury/x/government/types"

	furykycTypes "github.com/tessornetwork/fury/x/furykyc/types"
	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/x/genutil"
	"github.com/tendermint/tendermint/libs/cli"

	"github.com/spf13/cobra"
	tmtypes "github.com/tendermint/tendermint/types"

	"github.com/cosmos/cosmos-sdk/types/errors"

	"github.com/cosmos/cosmos-sdk/server"
	sdk "github.com/cosmos/cosmos-sdk/types"
	tmjson "github.com/tendermint/tendermint/libs/json"
)

// SetGenesisGovernmentAddressCmd returns set-genesis-government-address cobra Command.
func SetGenesisGovernmentAddressCmd(defaultNodeHome string) *cobra.Command {
	cmd := &cobra.Command{
		Use:   "set-genesis-government-address [government_address_or_key]",
		Short: "Sets the given address as the government address inside genesis.json, and assings a black membership to it",
		Args:  cobra.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) error {
			clientCtx := client.GetClientContextFromCmd(cmd)

			serverCtx := server.GetServerContextFromCmd(cmd)
			config := serverCtx.Config

			config.SetRoot(clientCtx.HomeDir)

			address, err := sdk.AccAddressFromBech32(args[0])
			if err != nil {
				return err
			}

			// retrieve the app state
			genFile := config.GenesisFile()
			genDoc, err := tmtypes.GenesisDocFromFile(genFile)
			if err != nil {
				return err
			}

			genState, err := SetGovernmentAddress(clientCtx, genDoc.AppState, address)
			if err != nil {
				return err
			}

			genDoc.AppState, err = json.Marshal(genState)

			return genutil.ExportGenesisFile(genDoc, genFile)
		},
	}

	cmd.Flags().String(cli.HomeFlag, defaultNodeHome, "node's home directory")

	return cmd
}

// SetGovernmentAddress set government address in the genesis
func SetGovernmentAddress(clientCtx client.Context, appState json.RawMessage, address sdk.AccAddress) (map[string]json.RawMessage, error) {
	var genState map[string]json.RawMessage
	if err := json.Unmarshal(appState, &genState); err != nil {
		return nil, fmt.Errorf("error unmarshalling genesis doc for government address: %s", err.Error())
	}
	var genStateGovernment govTypes.GenesisState
	json.Unmarshal(genState[govTypes.ModuleName], &genStateGovernment)

	if genStateGovernment.GovernmentAddress != "" {
		return nil, fmt.Errorf("cannot replace existing government address")
	}

	genStateGovernment.GovernmentAddress = address.String()

	genesisStateBzGovernment, err := tmjson.Marshal(genStateGovernment)
	if err != nil {
		return nil, errors.Wrap(err, "failed to marshal genesis doc")
	}
	genState[govTypes.ModuleName] = genesisStateBzGovernment

	// set a black membership to the government address
	// add a membership to the genesis state
	var genStateMemberships furykycTypes.GenesisState
	json.Unmarshal(genState[furykycTypes.ModuleName], &genStateMemberships)

	initSecondsPerYear := time.Hour * 24 * 365
	initExpirationDate := time.Now().Add(initSecondsPerYear) // It's safe becouse command is executed in one machine

	membership := furykycTypes.NewMembership(furykycTypes.MembershipTypeBlack, address, address, initExpirationDate)
	genStateMemberships.Memberships = append(genStateMemberships.Memberships, &membership)

	genesisStateBzMemberships, err := tmjson.Marshal(genStateMemberships)
	if err != nil {
		return nil, errors.Wrap(err, "failed to marshal genesis doc")
	}
	genState[furykycTypes.ModuleName] = genesisStateBzMemberships

	return genState, nil
}
