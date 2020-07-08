package main

import (
	"deck-api/cmd"
	"math/rand"
	"time"

	"github.com/spf13/cobra"
)

func main() {
	rand.Seed(time.Now().UnixNano())

	rootCmd := &cobra.Command{}
	rootCmd.AddCommand(cmd.Server())
	rootCmd.Execute()
}
