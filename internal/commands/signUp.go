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

// signUpCmd represents the signUp command
var signUpCmd = &cobra.Command{
	Use:   "sign_up",
	Short: "Регистарция пользователя",
	Args:  cobra.ExactArgs(2),
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("signUp called")
		keepService, err := setupService(false)
		if err != nil {
			fmt.Printf("Ошибка при конфигурации сервиса %s", err.Error())
		}
		uId, err := keepService.RegisterUser(args[0], args[1])
		if err != nil {
			if errors.Is(err, storage.ErrUserAlredyExist) {
				fmt.Printf("Такой пользователь уже существует.")
				return
			}
			fmt.Printf("Ошибка получения данных: %s\n", err.Error())
			return
		}
		authData := models.UserModel{
			UserID: uId,
			Login:  args[0],
			Hash:   args[1],
		}
		if err := createConfigFile(authData); err != nil {
			fmt.Printf("Ошибка входа в систему: %s\n", err.Error())
			return
		}
		fmt.Println("Успешная регистрация и вход в систему")
	},
}

func init() {
	rootCmd.AddCommand(signUpCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// signUpCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// signUpCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
