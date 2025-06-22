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
		if superVerbose {
			fmt.Println("Just use a debugger at this point 🤣")
			os.Exit(69)
		}

		gadb.SetDebug(adbDebug)
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
	verbose      bool
	superVerbose bool
	adbDebug     bool
	adbAddress   string
)

func init() {
	rootCmd.PersistentFlags().BoolVarP(&verbose, "verbose", "v", false, "Print verbose output")

	rootCmd.PersistentFlags().BoolVar(&superVerbose, "super-verbose", false, "Print super verbose output")
	rootCmd.PersistentFlags().Lookup("super-verbose").Hidden = true // Hide this flag from help output

	rootCmd.PersistentFlags().BoolVar(&adbDebug, "adb-debug", false, "Enable ADB debug output")
	rootCmd.PersistentFlags().StringVar(&adbAddress, "adb-address", "localhost:5037", "ADB server address")
}

func checkErr(err error) {
	if err != nil {
		fmt.Println(color.HiRedString("Error:"), err)
		os.Exit(1)
	}
}
