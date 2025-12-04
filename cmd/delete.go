package cmd

import (
	"fmt"

	"github.com/takhv/expenseTracker/internal/storage"

	"github.com/spf13/cobra"
)

var deleteCmd = &cobra.Command{
	Use:   "delete",
	Short: "Удаление записи о трате",
	Run: func(cmd *cobra.Command, args []string) {
		data, err := storage.ReadAndParse()
		if err != nil {
			return
		}
		for i, elem := range data {
			if elem.Id == deleteId {
				fmt.Printf("delete expense with ID-%d\n", elem.Id)
				data = append(data[:i], data[i+1:]...)
			}
		}

		storage.ParseToJSONAndWrite(data)
	},
}

var deleteId int

func init() {
	rootCmd.AddCommand(deleteCmd)

	deleteCmd.Flags().IntVarP(&deleteId, "id", "i", -999, "Delete by ID")
	deleteCmd.MarkFlagRequired("id")
}
