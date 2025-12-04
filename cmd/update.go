package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
	"github.com/takhv/expenseTracker/internal/storage"
)

var updateCmd = &cobra.Command{
	Use:   "update",
	Short: "Изменение записи о трате",
	Run: func(cmd *cobra.Command, args []string) {
		data, err := storage.ReadAndParse()
		if err != nil {
			return
		}

		if newAmount < 0 {
			fmt.Println("Amount problem")
			return
		}

		for i, elem := range data {
			if elem.Id == updateId {
				if newDesc != "" {
					fmt.Printf("New description: %s \n", newDesc)
					data[i].Description = newDesc
				}
				if newAmount != -1 {
					fmt.Printf("New amount: %d \n", newAmount)
					data[i].Amount = newAmount
				}
				fmt.Printf("Update expense with ID-%d \n", elem.Id)

			}
		}

		storage.ParseToJSONAndWrite(data)
	},
}

var updateId int
var newDesc string
var newAmount int

func init() {
	rootCmd.AddCommand(updateCmd)

	updateCmd.Flags().IntVarP(&updateId, "id", "i", -999, "ID for update")
	updateCmd.MarkFlagRequired("id")
	updateCmd.Flags().StringVarP(&newDesc, "desc", "d", "", "Your new description")
	updateCmd.Flags().IntVarP(&newAmount, "amount", "a", -1, "Your new amount")
}
