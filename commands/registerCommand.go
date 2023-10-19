package commands

import (
	"fmt"
	"github.com/spf13/cobra"
	"os"
	"os/exec"
)

var RegisterCommand = &cobra.Command{
	Use:   "register",
	Short: "reg",
	Long:  "register a new account",
	Run: func(command *cobra.Command, args []string) {

		err := exec.Command(os.Args[0], "wgcf", "register", "--accept-tos").Run()
		if err != nil {
			fmt.Println("Register failed.")
		}
	},
}
