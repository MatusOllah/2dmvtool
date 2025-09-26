package cmd

import (
	"os"
	"os/exec"

	"github.com/spf13/cobra"
)

func isLocalFile(path string) bool {
	_, err := os.Stat(path)
	return err == nil
}

func NewPlayCommand() *cobra.Command {
	var (
		ffplayArgs string
	)

	c := &cobra.Command{
		Use:   "play",
		Short: "Play a 2DMV",
		Long:  `Play a 2DMV either from a local file or from an Android device by specifying the song ID.`,
		Args:  cobra.ExactArgs(1),
		Run: func(cmd *cobra.Command, args []string) {
			isFile := isLocalFile(args[0])

			if isFile {
				// Play local file
				cmd := exec.Command("ffplay", "-i", args[0])
				cmd.Args = append(cmd.Args, ffplayArgs)
			}
		},
	}

	// Flags
	c.Flags().StringVarP(&ffplayArgs, "ffplay-args", "a", "", "Additional arguments to pass to FFplay")

	return c
}

func init() {
	rootCmd.AddCommand(NewPlayCommand())
}
