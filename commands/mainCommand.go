package commands

import (
	"github.com/ViRb3/wgcf/cmd"
	"github.com/ViRb3/wgcf/cmd/shared"
	"github.com/ViRb3/wgcf/util"
	sing_box "github.com/sagernet/sing-box/cmd/sing-box"
	"github.com/spf13/cobra"
	"log"
)

var MainCmd = &cobra.Command{
	Use:   "sing_warp",
	Short: "sing_warp",
	Long: shared.FormatMessage("short message", `
long message`),
	Run: func(cmd *cobra.Command, args []string) {
		if err := cmd.Help(); err != nil {
			log.Fatal(util.GetErrorMessage(err))
		}
	},
}

func init() {

	MainCmd.AddCommand(cmd.RootCmd)
	MainCmd.AddCommand(sing_box.SingBoxCommand)
	MainCmd.AddCommand(RegisterCommand)
}
