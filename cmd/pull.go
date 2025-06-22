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

func pull(id int, dst string) error {
	// open device
	if verbose {
		fmt.Printf("Opening device at ADB %s with serial number %s...\n", adbAddress, serial)
	}
	device, err := adbutil.OpenDevice(adbAddress, serial)
	if err != nil {
		return fmt.Errorf("opening device: %v\nTry starting the ADB server using this command:\n\tadb start-server", err)
	}

	if verbose {
		fmt.Println("Device opened successfully.")
		adbutil.PrintDeviceInfo(device)
	}

	// get path
	path := mv.MVPath(id, kind, region)

	// get file size
	size, err := adbutil.GetRemoteFileSize(device, path)
	if err != nil {
		return err
	}

	// create output file
	if dst == "" {
		dst = fmt.Sprintf("%04d.usm", id)
	}

	if _, err := os.Stat(dst); err == nil {
		if force {
			yellow := color.New(color.FgYellow).SprintFunc()
			fmt.Fprintf(os.Stderr, "⚠️  %s %s %s\n", yellow("File"), dst, yellow("already exists and will be overwritten."))
		} else {
			return fmt.Errorf("file %s already exists. Use --force to overwrite it", dst)
		}
	}

	file, err := os.Create(dst)
	checkErr(err)
	defer func() {
		if err := file.Close(); err != nil {
			checkErr(err)
		}
	}()

	if verbose {
		fmt.Printf("Pulling remote file %s => %s (%d bytes)...\n", path, dst, size)
	}

	// create progress bar
	pb := progressbar.DefaultBytes(size, "Pulling CRI Sofdec video file...")

	// pull the video
	if err := device.Pull(path, io.MultiWriter(file, pb)); err != nil {
		return fmt.Errorf("failed to pull file %s: %w", path, err)
	}

	return nil
}

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

		checkErr(pull(id, output))

		// success message
		color.Green("✅ Successfully pulled raw 2DMV!")
	},
}

var (
	serial string
	kind   mv.MVKind       = mv.MVKindSEKAI
	region mv.ServerRegion = mv.ServerRegionEN
	output string
	force  bool
)

func init() {
	rootCmd.AddCommand(pullCmd)

	pullCmd.Flags().StringVarP(&serial, "serial", "s", "", "Device serial number")
	pullCmd.Flags().VarP(&kind, "kind", "k", "Type of 2DMV to prefer pulling (\"original\", \"sekai\")")
	pullCmd.Flags().VarP(&region, "region", "r", "Game server region (\"jp\", \"en\", \"tw\", \"kr\", \"cn\")")
	pullCmd.Flags().StringVarP(&output, "output", "o", "", "Output file path (default <id>.usm)")
	pullCmd.Flags().BoolVarP(&force, "force", "f", false, "Force overwrite existing output file")
}
