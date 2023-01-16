package cli

import (
	"fmt"

	"github.com/spf13/cobra"

	"github.com/cosmos/cosmos-sdk/client"

	"github.com/tessornetwork/fury/x/documents/types"
)

// GetQueryCmd returns the cli query commands for this module
func GetQueryCmd(queryRoute string) *cobra.Command {
	// Group documents queries under a subcommand
	cmd := &cobra.Command{
		Use:                        types.ModuleName,
		Short:                      fmt.Sprintf("Querying commands for the %s module", types.ModuleName),
		DisableFlagParsing:         true,
		SuggestionsMinimumDistance: 2,
		RunE:                       client.ValidateCmd,
	}

	// the following comment can be removed
	//	cmd.AddCommand(CmdListDocument())
	cmd.AddCommand(
		CmdShowDocument(),
		CmdSentDocuments(),
		CmdReceivedDocuments(),

		// implement CmdShowDocumentReceipt() ?
		CmdSentReceipts(),
		CmdReceivedReceipts(),
		CmdDocumentsReceipts(),
	)

	return cmd
}
