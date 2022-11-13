package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

var availableDaysByYear = map[string][]string{
	"2021": {
		"01",
		"02",
	},
}

var listCmd = &cobra.Command{
	Use:   "list",
	Short: "lists available days and years",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Print("Available years and days\n\n")
		for year, days := range availableDaysByYear {
			fmt.Printf("Year: %s\n", year)
			fmt.Println("----------")
			for _, day := range days {
				fmt.Printf("- %s\n", day)
			}
			fmt.Println()
		}
	},
}

func init() {
	rootCmd.AddCommand(listCmd)
}
