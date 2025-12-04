package cmd

import (
	"fmt"
	"time"

	"github.com/takhv/expenseTracker/internal/model"
	"github.com/takhv/expenseTracker/internal/storage"

	"github.com/spf13/cobra"
)

var addCmd = &cobra.Command{
	Use:   "add",
	Short: "Создание записи о трате",
	Run: func(cmd *cobra.Command, args []string) {
		data, err := storage.ReadAndParse()
		if err != nil {
			return
		}

		if amount < 0 {
			fmt.Println("Amount problem")
			return
		}

		newExpense := model.Expense{
			Id:          idCalc(data),
			Description: desc,
			CreatedAt:   time.Now(),
			Amount:      amount,
		}
		data = append(data, newExpense)

		storage.ParseToJSONAndWrite(data)
		fmt.Printf("added with ID-%d\n", newExpense.Id)
	},
}

var desc string
var amount int

func init() {
	rootCmd.AddCommand(addCmd)

	addCmd.Flags().StringVarP(&desc, "description", "d", "", "Описание")
	addCmd.Flags().IntVarP(&amount, "amount", "a", 0, "Цена")
	addCmd.MarkFlagRequired("description")
	addCmd.MarkFlagRequired("amount")
}

func idCalc(data []model.Expense) int {
	var id int
	for _, elem := range data {
		if elem.Id >= id {
			id = elem.Id
		}
	}
	return id + 1
}
