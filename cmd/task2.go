/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"github.com/ecaron/rarecircles_tasks/task2"
	"github.com/spf13/cobra"
)

// task2Cmd represents the task2 command
var task2Cmd = &cobra.Command{
	Use:   "task2",
	Short: "Execute task2",
	Long: `Execute task2 where the parallelism is used. This command simulates the approval process of an authorization which previously
	needs of various validations that can be parallelized to optimize performance. This command has 4 optional flags to configure sleep time and 
	validation response of the card and balance repos.
	Example of usage: go run main.go task2 --cardSleep=2 --balanceSleep=3 --validBalance=true --validCard=true`,
	Run: func(cmd *cobra.Command, args []string) {
		task2.Task2()
	},
}

func init() {
	task2Cmd.Flags().BoolVar(&task2.ValidCard, "validCard", false, "card validation result")
	task2Cmd.Flags().BoolVar(&task2.ValidBalance, "validBalance", false, "balance validation result")
	task2Cmd.Flags().IntVar(&task2.CardSleep, "cardSleep", 0, "card validation sleep")
	task2Cmd.Flags().IntVar(&task2.BalanceSleep, "balanceSleep", 0, "balance validation sleep")
	rootCmd.AddCommand(task2Cmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// task2Cmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// task2Cmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
