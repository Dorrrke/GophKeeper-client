/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"errors"
	"fmt"

	"github.com/Dorrrke/GophKeeper-client/internal/domain/models"
	"github.com/Dorrrke/GophKeeper-client/internal/storage"
	"github.com/spf13/cobra"
)

// saveauthdataCmd represents the saveauthdata command
var saveauthdataCmd = &cobra.Command{
	Use:   "save_auth_data",
	Short: "Сохраняет пару логин пароль введенные пользователем",
	Long: `Сохраняет пару логин пароль введенные пользователем.
	При подключении наличии подключения к сети данные отправляются на хранение на сервере, 
	в ином случае харнятся на личном ПК пользователя.
	Пример использование: gophkeeper saveauthdata login_name login password`,
	Args: cobra.ExactArgs(3),
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("saveauthdata called")

		keepService, err := setupService(false)
		if err != nil {
			fmt.Printf("Ошибка при конфигурации сервиса %s", err.Error())
		}
		authData := models.LoginModel{
			Name:     args[0],
			Login:    args[1],
			Password: args[2],
		}
		userModel, err := getUserID()
		if err != nil {
			fmt.Printf("Ошибка при получении данных %s", err.Error())
			return
		}
		updateFlag, err := cmd.Flags().GetBool("update")
		if err != nil {
			fmt.Printf("Ошибка при получении флага: %s", err.Error())
		}
		if updateFlag {
			err := keepService.UpdateLogin(authData, userModel.UserID)
			if err != nil {
				fmt.Printf("Ошибка при обновлении данных: %s", err.Error())
				return
			}
			fmt.Printf("Данные успешно обновлены")
			return
		}
		_, err = keepService.SaveLogin(authData, userModel.UserID)
		if err != nil {
			if errors.Is(err, storage.ErrLoginAlredyExist) {
				fmt.Printf("Пара логин/пароль с таким именем уже сохранена.")
				return
			}
			fmt.Printf("Ошибка при сохранении данных: %s", err.Error())
			return
		}
		fmt.Println("Успешно сохранено!")
	},
}

func init() {
	rootCmd.AddCommand(saveauthdataCmd)
	saveauthdataCmd.Flags().Bool("update", false, "Обновить существующие данные")

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// saveauthdataCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// saveauthdataCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
