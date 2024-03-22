/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"github.com/ecaron/rarecircles_tasks/task1"
	"github.com/spf13/cobra"
)

// task1Cmd represents the task1 command
var task1Cmd = &cobra.Command{
	Use:   "task1",
	Short: "Execute task1",
	Long: `Execute task1 where the adapter pattern is used. This command uses this pattern to process a payment from a capture.
	The command has no flags or additional args.
	Example of usage: go run main.go task1`,
	Run: func(cmd *cobra.Command, args []string) {
		task1.Task1()
	},
}

func init() {
	rootCmd.AddCommand(task1Cmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// task1Cmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// task1Cmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
