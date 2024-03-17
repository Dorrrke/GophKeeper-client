/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

// getloginsCmd represents the getlogins command
var getloginsCmd = &cobra.Command{
	Use:   "getauth",
	Short: "Отображает сохраненные пары логин пароль",
	Long: `При вызове отображает список всех сохраненных пар логин пароль.
	При наличии подключения к интернету, данные будут браться из удаленного сервера.`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("getlogins called")
		keepService, err := setupService(false)
		if err != nil {
			fmt.Printf("Ошибка при конфигурации сервиса %s", err.Error())
		}
		userModel, err := getUserID()
		if err != nil {
			fmt.Printf("Ошибка при получении данных %s", err.Error())
			return
		}
		res, err := keepService.GetLogins(userModel.UserID)
		if err != nil {
			fmt.Printf("Ошибка при получении данных: %s", err.Error())
		}
		if len(res) == 0 {
			fmt.Printf("Нет сохраненных данных")
			return
		}
		logins := ""
		for _, login := range res {
			cardStr := fmt.Sprintf("\nLogin name: %s \n\tLogin: %s \n\tPassword: %s\n",
				login.Name, login.Login, login.Password)
			logins += cardStr
		}
		fmt.Printf("Logins: %s", logins)
	},
}

func init() {
	rootCmd.AddCommand(getloginsCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// getloginsCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// getloginsCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
