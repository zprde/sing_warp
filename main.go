package main

import (
	"github.com/ViRb3/wgcf/util"
	"github.com/zprde/sing_warp/commands"
	"log"
)

func main() {
	if err := commands.MainCmd.Execute(); err != nil {
		log.Fatal(util.GetErrorMessage(err))
	}
}
