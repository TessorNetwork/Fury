package cli

import (
	"github.com/spf13/cobra"

	"github.com/tessornetwork/fury/x/furykyc/types"
	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/client/flags"
	"github.com/cosmos/cosmos-sdk/client/tx"
	sdk "github.com/cosmos/cosmos-sdk/types"
)

func CmdBuy() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "buy [subscriber] [membership-type]",
		Short: "Tsp buy a membership for subscriber",
		Args:  cobra.ExactArgs(2),
		RunE: func(cmd *cobra.Command, args []string) error {
			return buyMembershipFunc(cmd, args)
		},
	}
	flags.AddTxFlagsToCmd(cmd)

	return cmd
}

func buyMembershipFunc(cmd *cobra.Command, args []string) error {
	cliCtx, err := client.GetClientTxContext(cmd)
	if err != nil {
		return err
	}

	tsp := cliCtx.GetFromAddress()

	buyer, err := sdk.AccAddressFromBech32(args[0])
	if err != nil {
		return err
	}
	membershipType := args[1]

	msg := types.NewMsgBuyMembership(membershipType, buyer, tsp)
	err = msg.ValidateBasic()
	if err != nil {
		return err
	}

	return tx.GenerateOrBroadcastTxCLI(cliCtx, cmd.Flags(), msg)

}
