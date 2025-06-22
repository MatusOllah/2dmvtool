package cmd

import (
	"fmt"
	"io"
	"os"
	"strconv"

	"github.com/MatusOllah/2dmvtool/internal/adbutil"
	"github.com/MatusOllah/2dmvtool/internal/mv"
	"github.com/fatih/color"
	"github.com/schollz/progressbar/v3"
	"github.com/spf13/cobra"
)

// pullCmd represents the pull command
var pullCmd = &cobra.Command{
	Use:     "pull id",
	Short:   "Pull a raw 2DMV from an Android device",
	Long:    `Pull a raw CRI Sofdec 2DMV from an Android device using the specified song ID.`,
	Example: `2dmvtool pull 514`, // example for FAKE HEART by KIRA feat. Kagamine Rin/Len (#514)
	Args:    cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		// parse song ID integer
		id, err := strconv.Atoi(args[0])
		if err != nil {
			checkErr(fmt.Errorf("failed to parse song ID %s: %w", args[0], err))
		}

		if id <= 0 {
			checkErr(fmt.Errorf("song ID must be a positive integer"))
		}

		// open device
		device, err := adbutil.OpenDevice(adbAdddress, serial)
		if err != nil {
			checkErr(fmt.Errorf("opening device: %v\nTry starting the ADB server using this command:\n\tadb start-server", err))
		}

		// get path
		path := mv.MVPath(id, kind, region)

		// get file size
		size, err := adbutil.GetRemoteFileSize(device, path)
		checkErr(err)

		// create output file
		if output == "" {
			output = fmt.Sprintf("%04d.usm", id)
		}

		if _, err := os.Stat(output); err == nil {
			fmt.Fprint(os.Stderr, color.HiYellowString("⚠️ File %s already exists and will be overwritten.\n", output))
		}

		file, err := os.Create(output)
		checkErr(err)
		defer checkErr(file.Close())

		// create progress bar
		pb := progressbar.DefaultBytes(size, "Pulling CRI Sofdec video file...")

		// pull the video
		if err := device.Pull(path, io.MultiWriter(file, pb)); err != nil {
			checkErr(fmt.Errorf("failed to pull file %s: %w", path, err))
		}

		// success message
		fmt.Printf("✅ Successfully pulled CRI Sofdec video file for song ID %d to %s!\n", id, output)
	},
}

var (
	adbAdddress string
	serial      string
	kind        mv.MVKind       = mv.MVKindSEKAI
	region      mv.ServerRegion = mv.ServerRegionEN
	output      string
)

func init() {
	rootCmd.AddCommand(pullCmd)

	pullCmd.Flags().StringVar(&adbAdddress, "adb", "localhost:5037", "ADB server address")
	pullCmd.Flags().StringVarP(&serial, "serial", "s", "", "Device serial number")
	pullCmd.Flags().VarP(&kind, "kind", "k", "Type of 2DMV to prefer pulling (\"original\", \"sekai\")")
	pullCmd.Flags().VarP(&region, "region", "r", "Game server region (\"jp\", \"en\", \"tw\", \"kr\", \"cn\")")
	pullCmd.Flags().StringVarP(&output, "output", "o", "", "Output file path (default <id>.usm)")
}
