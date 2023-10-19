package main

import (
	"github.com/ViRb3/wgcf/cmd"
	"github.com/ViRb3/wgcf/cmd/shared"
	"github.com/ViRb3/wgcf/util"
	"github.com/spf13/cobra"
	"log"
)

var MainCmd = &cobra.Command{
	Use:   "test",
	Short: "test",
	Long: shared.FormatMessage("short message", `
long message`),
	Run: func(cmd *cobra.Command, args []string) {
		if err := cmd.Help(); err != nil {
			log.Fatal(util.GetErrorMessage(err))
		}
	},
}

func main() {
	MainCmd.AddCommand(cmd.RootCmd)
	if err := cmd.Execute(); err != nil {
		log.Fatal(util.GetErrorMessage(err))
	}
}
