package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"

	"github.com/amjadjibon/dkv/server"
)

var rootCmd = &cobra.Command{
	Use:   "dkv",
	Short: "dkv root cmd",
	Long:  `dkv root cmd`,
	Run: func(cmd *cobra.Command, args []string) {

	},
}

var runCmd = &cobra.Command{
	Use:   "run",
	Short: "run dkv server",
	Long:  `run dkv server`,
	Run: func(cmd *cobra.Command, args []string) {
		server.BadgerRaftSetup()
	},
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		_, _ = fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}

func init() {
	rootCmd.AddCommand(runCmd)
}
