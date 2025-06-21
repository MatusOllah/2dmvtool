package cmd

import (
	"fmt"
	"runtime"
	"runtime/debug"

	"github.com/fatih/color"
	"github.com/spf13/cobra"
)

var The2DMVToolVersion string

func getVersion() string {
	if The2DMVToolVersion != "" {
		return The2DMVToolVersion
	}

	bi, ok := debug.ReadBuildInfo()
	if !ok {
		return ""
	}

	return bi.Main.Version
}

// versionCmd represents the version command.
var versionCmd = &cobra.Command{
	Use:   "version",
	Short: "Print the version of 2DMVTool.",
	Long: `All software has versions. This prints 2DMVTool's version.
Even Rui checks his before debugging show chaos.`,
	Run: func(cmd *cobra.Command, args []string) {
		cyan := color.New(color.FgCyan, color.Bold).SprintFunc()
		bold := color.New(color.Bold).SprintFunc()

		ver := getVersion()
		if ver == "" {
			fmt.Println("2DMVTool version is unavailable.")
		} else {
			fmt.Println(cyan("2DMVTool"), bold("version"), ver)
		}
		fmt.Println(cyan("Go"), bold("version"), runtime.Version(), runtime.GOOS, runtime.GOARCH)

		fmt.Printf("\nCopyright (c) 2025 Matúš Ollah; Licensed under MIT License\n")
	},
}

func init() {
	rootCmd.AddCommand(versionCmd)
}
