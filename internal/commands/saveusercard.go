/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"errors"
	"fmt"
	"strconv"

	"github.com/Dorrrke/GophKeeper-client/internal/domain/models"
	"github.com/Dorrrke/GophKeeper-client/internal/storage"
	"github.com/spf13/cobra"
)

// saveusercardCmd represents the saveusercard command
var saveusercardCmd = &cobra.Command{
	Use:   "save_card",
	Short: "Сохраняет данные банковской карты",
	Long: `Сохраняет данные банковской карты пользователя.
	При подключении наличии подключения к сети данные отправляются на хранение на сервере, 
	в ином случае харнятся на личном ПК пользователя.`,
	Args: cobra.ExactArgs(4),
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("saveusercard called")

		keepService, err := setupService(false)
		if err != nil {
			fmt.Printf("Ошибка при конфигурации сервиса %s", err.Error())
		}
		cvvCode, err := strconv.Atoi(args[3])
		if err != nil {
			fmt.Printf("Внутренняя ошибка: %s", err.Error())
		}
		card := models.CardModel{
			Name:    args[0],
			Number:  args[1],
			Date:    args[2],
			CVVCode: cvvCode,
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
			err := keepService.UpdateCard(card, userModel.UserID)
			if err != nil {
				fmt.Printf("Ошибка при обновлении данных: %s", err.Error())
				return
			}
			fmt.Printf("Данные успешно обновлены")
			return
		}
		_, err = keepService.SaveCard(card, userModel.UserID)
		if err != nil {
			if errors.Is(err, storage.ErrCardAlredyExist) {
				fmt.Printf("Карта с таким именем уже сохранена.")
				return
			}
			fmt.Printf("Ошибка при сохранении карты: %s", err.Error())
			return
		}

		fmt.Println("Успешно сохранено!")
	},
}

func init() {
	rootCmd.AddCommand(saveusercardCmd)
	saveusercardCmd.Flags().Bool("update", false, "Обновить существующие данные")

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// saveusercardCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// saveusercardCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
