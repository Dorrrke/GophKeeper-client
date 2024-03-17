/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"bufio"
	"errors"
	"fmt"
	"os"

	"github.com/Dorrrke/GophKeeper-client/internal/domain/models"
	"github.com/Dorrrke/GophKeeper-client/internal/storage"
	"github.com/spf13/cobra"
)

// saveBinCmd represents the saveBin command
var saveBinCmd = &cobra.Command{
	Use:   "save_bin",
	Short: "Сохраняет бинарный файл под указанным именем.",
	Long:  `Сохраняет бинраный файл находяйщийся по указанному пути. В базе данных данные будут сохранены под указанным именем.`,
	Args:  cobra.ExactArgs(2),
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("saveBin called")
		bData, err := readBinFile(args[1])
		if err != nil {
			fmt.Printf("Ошибка чтения бинарного файла %s", err.Error())
			return
		}
		userModel, err := getUserID()
		if err != nil {
			fmt.Printf("Ошибка при получении данных %s", err.Error())
			return
		}
		keepService, err := setupService(false)
		if err != nil {
			fmt.Printf("Ошибка при конфигурации сервиса %s", err.Error())
			return
		}
		bModel := models.BinaryDataModel{
			Name: args[0],
			Data: bData,
		}
		updateFlag, err := cmd.Flags().GetBool("update")
		if err != nil {
			fmt.Printf("Ошибка при получении флага: %s", err.Error())
		}
		if updateFlag {
			err := keepService.UpdateBin(bModel, userModel.UserID)
			if err != nil {
				fmt.Printf("Ошибка при обновлении данных: %s", err.Error())
				return
			}
			fmt.Printf("Данные успешно обновлены")
			return
		}

		_, err = keepService.SaveBinaryData(bModel, userModel.UserID)
		if err != nil {
			if errors.Is(err, storage.ErrBinAlredyExist) {
				fmt.Printf("Бинарные данные с таким именем уже существуют")
				return
			}
			fmt.Printf("Ошибка при сохранении данных: %s", err.Error())
			return
		}
		fmt.Println("Успешно сохранено!")
	},
}

func init() {
	rootCmd.AddCommand(saveBinCmd)
	saveBinCmd.Flags().Bool("update", false, "Обновить существующие данные")
	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// saveBinCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// saveBinCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}

func readBinFile(fName string) ([]byte, error) {
	f, err := os.Open(fName)
	if err != nil {
		return nil, err
	}
	defer f.Close()

	stat, err := f.Stat()
	if err != nil {
		return nil, err
	}
	size := stat.Size()
	bData := make([]byte, size)
	bufr := bufio.NewReader(f)
	_, err = bufr.Read(bData)
	if err != nil {
		return nil, err
	}
	return bData, nil
}
