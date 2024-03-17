/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

// binListCmd represents the binList command
var binListCmd = &cobra.Command{
	Use:   "bin_list",
	Short: "Отображает список сохраненных бинарных данных",
	Long: `При вызове отображает список всех сохраненных бинарных данных пользователя.
	При наличии подключения к интернету, данные будут браться из удаленного сервера.`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("binList called")
		keepService, err := setupService(false)
		if err != nil {
			fmt.Printf("Ошибка при конфигурации сервиса %s", err.Error())
		}
		userModel, err := getUserID()
		if err != nil {
			fmt.Printf("Ошибка при получении данных %s", err.Error())
			return
		}
		bins, err := keepService.GetBins(userModel.UserID)
		if err != nil {
			fmt.Printf("Ошибка при получении данных: %s", err.Error())
			return
		}
		if len(bins) == 0 {
			fmt.Printf("Нет сохраненных данных")
			return
		}
		bList := ""
		for _, bin := range bins {
			binStr := fmt.Sprintf("\n\tBin name: %s\n",
				bin.Name)
			bList += binStr
		}
		fmt.Printf("Bin list: %s", bList)
	},
}

func init() {
	rootCmd.AddCommand(binListCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// binListCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// binListCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
