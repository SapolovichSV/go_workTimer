package cmd

import (
	"my-start/workTimer/lib"

	"github.com/fatih/color"
	"github.com/spf13/cobra"
)

var RootCmd = &cobra.Command{
	Use:   "app",
	Short: "workTimer helps you to focus",
}

func init() {
	RootCmd.AddCommand(startCmd)
}

var startCmd = &cobra.Command{
	Use:   "start",
	Short: "start workTimer",
	Long:  `start <workDurationInMinutes> <breakDurationInMinutes> <breaksCounts>`,
	Args:  cobra.ExactArgs(3),
	Run: func(cmd *cobra.Command, args []string) {
		startWorkTimer(args)
	},
}

func startWorkTimer(args []string) {
	workDuration := args[0]
	breakDuration := args[1]
	breaksCounts := args[2]
	color.Green("Work Duration(minutes): %s", workDuration)
	color.Green("Break Duration(minutes): %s", breakDuration)
	color.Green("Breaks Counts: %s", breaksCounts)
	err := lib.Process(workDuration, breakDuration, breaksCounts)
	if err != nil {
		color.Red("Error: %s", err)
	}
}
