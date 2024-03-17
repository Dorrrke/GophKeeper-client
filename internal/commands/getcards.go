/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"
	"time"

	"github.com/spf13/cobra"
)

// getcardsCmd represents the getcards command
var getcardsCmd = &cobra.Command{
	Use:   "getcards",
	Short: "Отображает сохраненные данные карт",
	Long: `При вызове отображает список всех сохраненных пользователем карт.
	При наличии подключения к интернету, данные будут браться из удаленного сервера`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("getcards called" + time.Now().Format(time.RFC3339))
		keepService, err := setupService(false)
		if err != nil {
			fmt.Printf("Ошибка при конфигурации сервиса %s", err.Error())
		}
		userModel, err := getUserID()
		if err != nil {
			fmt.Printf("Ошибка при получении данных %s", err.Error())
			return
		}
		res, err := keepService.GetCards(userModel.UserID)
		if err != nil {
			fmt.Printf("Ошибка при получении данных: %s", err.Error())
		}
		if len(res) == 0 {
			fmt.Printf("Нет сохраненных данных")
			return
		}
		cards := ""
		for _, card := range res {
			cardStr := fmt.Sprintf("\nCard name: %s \n\tNumber: %s \n\tDate: %s\n\tCVV: %v\n",
				card.Name, card.Number, card.Date, card.CVVCode)
			cards += cardStr
		}
		fmt.Printf("Cards: %s", cards)
	},
}

func init() {
	rootCmd.AddCommand(getcardsCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// getcardsCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// getcardsCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
