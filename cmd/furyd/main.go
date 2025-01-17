package main

import (
	"os"

	"github.com/tessornetwork/fury/app"
	"github.com/tessornetwork/fury/cmd/furyd/cmd"
	svrcmd "github.com/cosmos/cosmos-sdk/server/cmd"
)

func main() {
	rootCmd, _ := cmd.NewRootCmd()
	if err := svrcmd.Execute(rootCmd, app.DefaultNodeHome); err != nil {
		os.Exit(1)
	}
}
