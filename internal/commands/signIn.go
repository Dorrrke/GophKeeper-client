/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"encoding/json"
	"errors"
	"fmt"
	"os"

	"github.com/Dorrrke/GophKeeper-client/internal/coder"
	"github.com/Dorrrke/GophKeeper-client/internal/domain/models"
	"github.com/Dorrrke/GophKeeper-client/internal/services"
	"github.com/Dorrrke/GophKeeper-client/internal/storage"
	"github.com/spf13/cobra"
)

// signInCmd represents the signIn command
var signInCmd = &cobra.Command{
	Use:   "sign_in",
	Short: "Вход в систему.",
	Args:  cobra.ExactArgs(2),
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("signIn called")
		keepService, err := setupService(false)
		if err != nil {
			fmt.Printf("Ошибка при конфигурации сервиса %s", err.Error())
		}
		userModel, err := keepService.LoginUser(args[0], args[1])
		if err != nil {
			if errors.Is(err, storage.ErrUserNotExist) {
				fmt.Printf("Неверный логин или пароль. Такого пользователя не существует.")
				return
			}
			if errors.Is(err, services.ErrInvalidPassword) {
				fmt.Printf("Неверный пароль.")
				return
			}
			fmt.Printf("Ошибка получения данных: %s\n", err.Error())
			return
		}
		err = createConfigFile(userModel)
		if err != nil {
			fmt.Printf("Ошибка входа в систему: %s\n", err.Error())
			return
		}
		fmt.Println("Успешный вход в систему")
	},
}

func init() {
	rootCmd.AddCommand(signInCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// signInCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// signInCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}

func createConfigFile(auth models.UserModel) error {
	f, err := os.Create("auth_conf")
	if err != nil {
		return err
	}
	defer f.Close()
	authData, err := json.Marshal(auth)
	if err != nil {
		return err
	}
	encData, err := coder.Encoder(authData)
	if err != nil {
		return err
	}
	_, err = f.WriteString(string(encData))
	if err != nil {
		return err
	}
	return nil
}
