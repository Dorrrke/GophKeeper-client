/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

// gettextsCmd represents the gettexts command
var gettextsCmd = &cobra.Command{
	Use:   "gettexts",
	Short: "Отображает сохраненные текстовые данные пользователя",
	Long: `При вызове отображает список всех сохраненных текстовых данных пользователя.
	При наличии подключения к интернету, данные будут браться из удаленного сервера.`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("gettexts called")
		keepService, err := setupService(false)
		if err != nil {
			fmt.Printf("Ошибка при конфигурации сервиса %s", err.Error())
		}
		userModel, err := getUserID()
		if err != nil {
			fmt.Printf("Ошибка при получении данных %s", err.Error())
			return
		}
		res, err := keepService.GetTextData(userModel.UserID)
		if err != nil {
			fmt.Printf("Ошибка при получении данных: %s", err.Error())
			return
		}
		if len(res) == 0 {
			fmt.Printf("Нет сохраненных данных")
			return
		}
		texts := ""
		for _, text := range res {
			cardStr := fmt.Sprintf("\nText name: %s \n\tData: %s\n",
				text.Name, text.Data)
			texts += cardStr
		}
		fmt.Printf("Texts: %s", texts)
	},
}

func init() {
	rootCmd.AddCommand(gettextsCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// gettextsCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// gettextsCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
