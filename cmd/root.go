package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

var RootCmd = &cobra.Command{
	Use:   "task",
	Short: "Task is a CLI task manager",
	Long: `A CLI Testing Tool
		still a work in progress
		Thanks`,
	// Run: func(cmd *cobra.Command, args []string) {
	// 	// Do Stuff Here
	// },
}

func Execute() {
	if err := RootCmd.Execute(); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}
