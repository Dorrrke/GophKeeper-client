/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"errors"
	"fmt"

	"github.com/Dorrrke/GophKeeper-client/internal/storage"
	"github.com/spf13/cobra"
)

// textCmd represents the text command
var textCmd = &cobra.Command{
	Use:   "text",
	Short: "Отображает текстовые данные сохраненные под указанным именем.",
	Long: `При вызове отображает текстовые данные сохраненные под указанным именем.
	При наличии подключения к интернету, данные будут браться из удаленного сервера`,
	Args: cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("text called")

		keepService, err := setupService(false)
		if err != nil {
			fmt.Printf("Ошибка при конфигурации сервиса %s", err.Error())
		}
		userModel, err := getUserID()
		if err != nil {
			fmt.Printf("Ошибка при получении данных %s", err.Error())
			return
		}
		delFlag, err := cmd.Flags().GetBool("delete")
		if err != nil {
			fmt.Printf("Ошибка при получении флага: %s", err.Error())
		}
		if delFlag {
			err = keepService.DeleteTextByName(args[0], userModel.UserID)
			if err != nil {
				fmt.Printf("Ошибка при удалении данных %s", err.Error())
				return
			}
			fmt.Printf("Успешное удаление\n")
			return
		}
		tData, err := keepService.GetTextDataByName(args[0], userModel.UserID)
		if err != nil {
			if errors.Is(err, storage.ErrTextNotExist) {
				fmt.Printf("Текстовых данных сохраненных с таким именем не существует.")
				return
			}
			fmt.Printf("Ошибка при получении данных %s", err.Error())
			return
		}
		fmt.Printf("\nText name: %s \n\tData: %s\n",
			tData.Name, tData.Data)
	},
}

func init() {
	rootCmd.AddCommand(textCmd)
	textCmd.Flags().Bool("delete", false, "Удаление данных")

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// textCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// textCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
