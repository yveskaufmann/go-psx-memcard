package root

import (
	"com.yvka.memcard/pkg/ui"
	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "psx-memcard",
	Short: "psx-memcard is a PlayStation memory card file manager",
	Long: `psx-memcard is a PlayStation memory card file manager.

It allows you to view, extract, and modify the contents of PlayStation memory card files (.mcr).`,
	Run: func(cmd *cobra.Command, args []string) {
		if err := ui.Start(); err != nil {
			panic(err)
		}
	},
}

func Execute() error {
	return rootCmd.Execute()
}
