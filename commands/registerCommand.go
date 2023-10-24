package commands

import (
	"github.com/ViRb3/wgcf/config"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"log"
	"os"
	"os/exec"
)

var RegisterCommand = &cobra.Command{
	Use:   "register",
	Short: "reg",
	Long:  "register a new account",
	Run: func(command *cobra.Command, args []string) {
		if viper.GetString(config.PrivateKey) == "" {
			log.Println("Empty config, will register a new one.")
		} else {
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
