package cmd

import (
	"fmt"
	"os/exec"

	"github.com/fatih/color"
	"github.com/spf13/cobra"
)

func checkTool(name, cmd string) error {
	fmt.Printf("Checking %s... ", name)

	path, err := exec.LookPath(cmd)
	if err != nil {
		fmt.Printf("not found\n")
		return fmt.Errorf("tool %s not found in PATH: %w", name, err)
	}

	fmt.Print(color.HiGreenString("OK"))
	if verbose {
		fmt.Printf(" (%s)", path)
	}
	fmt.Println()

	return nil
}

// doctorCmd represents the doctor command.
var doctorCmd = &cobra.Command{
	Use:   "doctor",
	Short: "Check whether the dependencies are installed and configured correctly.",
	Run: func(cmd *cobra.Command, args []string) {
		checkErr(checkTool("ADB", "adb"))
		checkErr(checkTool("FFmpeg", "ffmpeg"))
		checkErr(checkTool("FFplay", "ffplay"))
	},
}

func init() {
	rootCmd.AddCommand(doctorCmd)
}
