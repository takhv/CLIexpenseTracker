package cmd

import (
	"fmt"

	"github.com/takhv/expenseTracker/internal/storage"

	"github.com/spf13/cobra"
)

var summaryCmd = &cobra.Command{
	Use:   "summary",
	Short: "Затраты",
	Args:  cobra.NoArgs,
	Run: func(cmd *cobra.Command, args []string) {
		data, err := storage.ReadAndParse()
		if err != nil {
			return
		}

		var total int
		if timing != 0 {
			for _, elem := range data {
				if timing == int(elem.CreatedAt.Month()) {
					total += elem.Amount
				}
			}
		} else {
			for _, elem := range data {
				total += elem.Amount
			}
		}

		fmt.Println("Total expenses:", total)
	},
}

var timing int

func init() {
	rootCmd.AddCommand(summaryCmd)

	summaryCmd.Flags().IntVarP(&timing, "month", "m", 0, "Траты за месяц")
}
