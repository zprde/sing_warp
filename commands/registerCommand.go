package commands

import (
	"github.com/spf13/cobra"
	"log"
	"os"
	"os/exec"
)

var RegisterCommand = &cobra.Command{
	Use:   "register",
	Short: "reg",
	Long:  "register a new account",
	Run: func(command *cobra.Command, args []string) {
		if file, _ := os.Open("./wgcf-account.toml"); file != nil {
			log.Println("Config already exists, skip the register.")
			return
		}
		err := exec.Command(os.Args[0], "wgcf", "register", "--accept-tos").Run()
		if err != nil {
			log.Fatal("Register failed.")
		} else {
			log.Println("Register successfully.")
		}
	},
}
