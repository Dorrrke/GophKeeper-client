/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"

	errText "github.com/Dorrrke/GophKeeper-client/internal/domain/errors"
	"github.com/Dorrrke/GophKeeper-client/internal/domain/models"
	"github.com/Dorrrke/GophKeeper-client/internal/services"
	"github.com/spf13/cobra"
	"google.golang.org/grpc/status"
)

// updateCmd represents the update command
var updateCmd = &cobra.Command{
	Use:   "update",
	Short: "Обновление базы данных на удаленном сервере.",
	Long: `При выполнении команды, все сохраненные данные пользователя отправляются на удаленный сервер и происходит синхронизация.
	Если на сервере оказались более новые данные, они вернутся и будут занесены в локальную базу данных.`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("update called")
		keepService, err := setupService(true)
		if err != nil {
			fmt.Printf("Ошибка при конфигурации сервиса %s", err.Error())
		}
		userModel, err := getUserID()
		if err != nil {
			fmt.Printf("Ошибка при получении данных %s", err.Error())
			return
		}
		authOnServer(keepService, userModel)
		if err != nil {
			fmt.Printf("Ошибка при входе в сервер сервиса %s", err.Error())
			fmt.Println(err.Error())
			return
		}
		err = keepService.SyncBD(userModel.UserID)
		if err != nil {
			fmt.Printf("Ошибка при получении синхронизации базы данных данных %s", err.Error())
			return
		}
		fmt.Println("Данные синхронизированны!")
	},
}

func init() {
	rootCmd.AddCommand(updateCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// updateCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// updateCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}

func authOnServer(keepService *services.KeepService, userModel models.UserModel) error {
	err := keepService.ServerLogin(userModel)
	if err != nil {
		rpcStatus, ok := status.FromError(err)
		if !ok {
			return err
		}
		if rpcStatus.Message() != errText.NoUserOnServerError {
			return err
		}
		err := keepService.ServerRegister(userModel)
		if err != nil {
			return err
		}
		return nil
	}
	return nil
}
