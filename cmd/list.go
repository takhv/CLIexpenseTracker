package cmd

import (
	"fmt"

	"github.com/takhv/expenseTracker/internal/storage"

	"github.com/spf13/cobra"
)

var listCmd = &cobra.Command{
	Use:   "list",
	Short: "Список всех трат",
	Args:  cobra.NoArgs,
	Run: func(cmd *cobra.Command, args []string) {
		data, err := storage.ReadAndParse()
		if err != nil {
			return
		}
		for _, elem := range data {
			fmt.Printf("(ID-%d) Description: %s, Amount: %d\n", elem.Id, elem.Description, elem.Amount)
		}
	},
}

func init() {
	rootCmd.AddCommand(listCmd)
}
