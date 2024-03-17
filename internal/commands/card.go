/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"os"

	"github.com/Dorrrke/GophKeeper-client/internal/client"
	"github.com/Dorrrke/GophKeeper-client/internal/coder"
	"github.com/Dorrrke/GophKeeper-client/internal/config"
	"github.com/Dorrrke/GophKeeper-client/internal/domain/models"
	"github.com/Dorrrke/GophKeeper-client/internal/services"
	"github.com/Dorrrke/GophKeeper-client/internal/storage"
	"github.com/spf13/cobra"
)

// cardCmd represents the card command
var cardCmd = &cobra.Command{
	Use:   "card",
	Short: "Отображает данные сохраненных карт с указанным именем",
	Long: `При вызове отображает данные карты сохраненных с указаным именем.
	При наличии подключения к интернету, данные будут браться из удаленного сервера`,
	Args: cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("card called")
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
			err = keepService.DeleteCardByName(args[0], userModel.UserID)
			if err != nil {
				fmt.Printf("Ошибка при удалении данных %s", err.Error())
				return
			}
			fmt.Printf("Успешное удаление\n")
			return
		}
		card, err := keepService.GetCardByName(args[0], userModel.UserID)
		if err != nil {
			if errors.Is(err, storage.ErrCardNotExist) {
				fmt.Printf("Карты сохраненной с таким именем не существует.")
				return
			}
			fmt.Printf("Ошибка при получении данных %s", err.Error())
			return
		}
		fmt.Printf("\nCard name: %s \n\tNumber: %s \n\tDate: %s\n\tCVV: %v\n",
			card.Name, card.Number, card.Date, card.CVVCode)
	},
}

func init() {
	rootCmd.AddCommand(cardCmd)
	cardCmd.Flags().Bool("delete", false, "Удаление данных")

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// cardCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// cardCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}

func setupService(sync bool) (*services.KeepService, error) {
	cfg := config.ReadConfig()
	storage, err := storage.New(cfg.DBPath)
	if err != nil {
		return nil, err
	}
	if sync {
		clietn, err := client.New(context.Background(), cfg.ServerAddr)
		if err != nil {
			return nil, fmt.Errorf("clietn init error: %w", err)
		}
		keepService := services.New(clietn, storage, storage, storage, storage, storage, storage)
		return keepService, nil
	}
	client := client.KeeperClient{}
	keepService := services.New(&client, storage, storage, storage, storage, storage, storage)
	return keepService, nil
}

func getUserID() (models.UserModel, error) {
	f, err := os.ReadFile("auth_conf")
	if err != nil {
		return models.UserModel{}, err
	}
	data, err := coder.Decoder(f)
	if err != nil {
		return models.UserModel{}, err
	}
	var userModel models.UserModel
	err = json.Unmarshal(data, &userModel)
	if err != nil {
		return models.UserModel{}, err
	}
	return userModel, nil
}
