package cmd

import (
	"fmt"
	"os"

	"github.com/MatusOllah/2dmvtool/internal/rui"
	"github.com/fatih/color"
	"github.com/spf13/cobra"
)

// ruiCmd represents the rui command.
var ruiCmd = &cobra.Command{
	Use:   "rui",
	Short: "Prints Rui Kamishiro ANSI art.",
	Long: `You weren't supposed to find this. And yet... here you are.

This is no ordinary command.
When executed, this command summons the enigmatic entity named Rui Kamishiro, in glorious ANSI art,
painstakingly rendered in monospaced pixels and terminal-safe colors, together with a
questionable quote about coding, Project SEKAI, or household appliances best left unexplained.

Why? Because sometimes, when working on 2DMVs at 3 AM on a Tuesday, running solely on Red Bull and spite,
you just need a little Rui to judge your code and cheer you up.

Let it be known: this command does not help you push or pull 2DMVs from Android devices.
It will not solve your encoding bugs.
It will not offer advice about CRI Sofdec video files.

But it will bless your shell with existential flair and a quote that absolutely no one asked for - yet everyone needed.

Use responsibly.

`,
	Hidden: true,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Fprintf(os.Stderr, rui.ANSIArt)
		fmt.Fprintln(os.Stderr)
		fmt.Fprint(os.Stderr, color.New(color.Bold).Sprintf("\"%s\"\n\t- Rui Kamishiro\n", rui.RandomQuote()))
		if rui.IsBirthday() {
			color.New(color.FgHiMagenta, color.Bold).Fprintln(os.Stderr, "\nHappy Birthday, Rui Kamishiro! 🎉❤️")
		}
	},
}

func init() {
	rootCmd.AddCommand(ruiCmd)
}
