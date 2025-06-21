package cmd

import (
	"fmt"
	"os"

	"github.com/electricbubble/gadb"
	"github.com/fatih/color"
	"github.com/spf13/cobra"
)

// rootCmd represents the base command when called without any subcommands.
var rootCmd = &cobra.Command{
	Use:   "2dmvtool",
	Short: "2DMVTool is a CLI tool for managing Project SEKAI 2DMVs.",
	Long:  `2DMVTool is a CLI tool designed to help manage and manipulate 2D music videos (2DMVs) for Project SEKAI (PJSK).`,
	PersistentPreRun: func(cmd *cobra.Command, args []string) {
		gadb.SetDebug(verbose)
	},
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

// The flags.
var (
	verbose bool
)

func init() {
	rootCmd.PersistentFlags().BoolVarP(&verbose, "verbose", "v", false, "Print verbose output")
}

func checkErr(err error) {
	if err != nil {
		fmt.Println(color.HiRedString("Error:"), err)
		os.Exit(1)
	}
}
