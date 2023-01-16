package cli

import (
	"context"

	"github.com/tessornetwork/fury/x/did/types"
	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/client/flags"
	"github.com/spf13/cobra"
)

func CmdShowIdentity() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "show-identity [id]",
		Short: "Resolves the DID document for the specified id",
		Args:  cobra.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) error {
			clientCtx := client.GetClientContextFromCmd(cmd)

			queryClient := types.NewQueryClient(clientCtx)
			id := string(args[0])

			params := &types.QueryResolveIdentityRequest{
				ID: id,
			}

			res, err := queryClient.Identity(context.Background(), params)
			if err != nil {
				return err
			}

			return clientCtx.PrintObjectLegacy(res)
		},
	}

	flags.AddQueryFlagsToCmd(cmd)

	return cmd
}
