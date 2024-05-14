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

// loginCmd represents the login command
var loginCmd = &cobra.Command{
	Use:   "login",
	Short: "Отображает данные сохраненой пары логин пароль с указанным именем",
	Long: `При вызове отображает данные сохраненой пары логин пароль с указанным именем.
	При наличии подключения к интернету, данные будут браться из удаленного сервера`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("login called")
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
			err = keepService.DeleteLoginByName(args[0], userModel.UserID)
			if err != nil {
				fmt.Printf("Ошибка при удалении данных %s", err.Error())
				return
			}
			fmt.Printf("Успешное удаление\n")
			return
		}
		login, err := keepService.GetLoginByName(args[0], userModel.UserID)
		if err != nil {
			if errors.Is(err, storage.ErrLoginNotExist) {
				fmt.Printf("Пары логин/пароль сохраненных с таким именем не существует.")
				return
			}
			fmt.Printf("Ошибка при получении данных %s", err.Error())
			return
		}
		fmt.Printf("\nLogin name: %s \n\tLogin: %s \n\tPassword: %s\n",
			login.Name, login.Login, login.Password)
	},
}

func init() {
	rootCmd.AddCommand(loginCmd)
	loginCmd.Flags().Bool("delete", false, "Удаление данных")

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// loginCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// loginCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
