package cli

import (
	"fmt"

	"github.com/spf13/cobra"

	"github.com/tessornetwork/fury/x/furymint/types"
	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/client/flags"
	"github.com/cosmos/cosmos-sdk/client/tx"
	sdk "github.com/cosmos/cosmos-sdk/types"

	uuid "github.com/satori/go.uuid"
)

const ()

func CmdMintFUSD() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "mint [amount]",
		Short: "Mints a given amount of FUSD\nAmount must be an integer number.",
		Args:  cobra.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) error {
			return mintFUSDCmdFunc(cmd, args)
		},
	}
	flags.AddTxFlagsToCmd(cmd)

	return cmd
}

func mintFUSDCmdFunc(cmd *cobra.Command, args []string) error {
	cliCtx, err := client.GetClientTxContext(cmd)
	if err != nil {
		return err
	}

	sender := cliCtx.GetFromAddress()
	deposit, ok := sdk.NewIntFromString(args[0])
	if !ok {
		return fmt.Errorf("amount must be an integer")
	}

	mintUUID := uuid.NewV4().String()
	if err != nil {
		return err
	}
	postion := types.Position{
		Owner:      sender.String(),
		Collateral: deposit.Int64(),
		ID:         mintUUID,
	}

	msg := types.NewMsgMintFUSD(postion)
	if err := msg.ValidateBasic(); err != nil {
		return err
	}
	return tx.GenerateOrBroadcastTxCLI(cliCtx, cmd.Flags(), msg)
}

func CmdBurnFUSD() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "burn [id] [amount]",
		Short: "Burns a given amount of tokens, associated with id.",
		Args:  cobra.ExactArgs(2),
		RunE: func(cmd *cobra.Command, args []string) error {
			return burnFUSDCmdFunc(cmd, args)
		},
	}
	flags.AddTxFlagsToCmd(cmd)

	return cmd
}

func burnFUSDCmdFunc(cmd *cobra.Command, args []string) error {
	cliCtx, err := client.GetClientTxContext(cmd)
	if err != nil {
		return err
	}

	sender := cliCtx.GetFromAddress()
	id := args[0]
	amount, err := sdk.ParseCoinNormalized(args[1])
	if err != nil {
		return err
	}

	msg := types.NewMsgBurnFUSD(sender, id, amount)
	if err := msg.ValidateBasic(); err != nil {
		return err
	}

	return tx.GenerateOrBroadcastTxCLI(cliCtx, cmd.Flags(), msg)
}
