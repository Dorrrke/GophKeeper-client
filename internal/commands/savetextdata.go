/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"bufio"
	"bytes"
	"errors"
	"fmt"
	"os"

	"github.com/Dorrrke/GophKeeper-client/internal/domain/models"
	"github.com/Dorrrke/GophKeeper-client/internal/storage"
	"github.com/spf13/cobra"
)

// savetextdataCmd represents the savetextdata command
var savetextdataCmd = &cobra.Command{
	Use:   "save_text_data",
	Short: "Сохраняет текстовые данные пользователя.",
	Long: `Сохраняет текстовые данные пользователя введенные при запуске комманды
	или данные из файла при использовании флага file.
	При подключении наличии подключения к сети данные отправляются на хранение на сервере, 
	в ином случае харнятся на личном ПК пользователя.
	Пример использование: 
	1) gophkeeper save_text_data data_name text_data
	2) gophkeeper save_text_data data_name path/to/text/file --file`,
	Args: cobra.ExactArgs(2),
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("savetextdata called")
		keepService, err := setupService(false)
		if err != nil {
			fmt.Printf("Ошибка при конфигурации сервиса %s", err.Error())
		}
		uFlag, err := cmd.Flags().GetBool("update")
		if err != nil {
			fmt.Printf("Ошибка при получении флага: %s", err.Error())
		}
		filePath, err := cmd.Flags().GetBool("file")
		if err != nil {
			fmt.Printf("Ошибка при получении флага: %s", err.Error())
		}
		userModel, err := getUserID()
		if err != nil {
			fmt.Printf("Ошибка при получении данных %s", err.Error())
			return
		}
		if filePath {
			data, err := parseFromeFile(args[1])
			if err != nil {
				fmt.Printf("Ошибка при получении данных из файла: %s", err.Error())
				return
			}
			textData := models.TextDataModel{
				Name: args[0],
				Data: data,
			}
			if uFlag {
				err := keepService.UpdateText(textData, userModel.UserID)
				if err != nil {
					fmt.Printf("Ошибка при обновлении данных: %s", err.Error())
					return
				}
				fmt.Printf("Данные успешно обновлены")
				return
			}
			_, err = keepService.SaveTextData(textData, userModel.UserID)
			if err != nil {
				if errors.Is(err, storage.ErrTextAlredyExist) {
					fmt.Printf("Текстовые данные с таким именем уже существуют")
					return
				}
				fmt.Printf("Ошибка при сохранении данных: %s", err.Error())
				return
			}
			fmt.Println("Успешно сохранено!")
		} else {
			textData := models.TextDataModel{
				Name: args[0],
				Data: args[1],
			}
			if uFlag {
				err := keepService.UpdateText(textData, userModel.UserID)
				if err != nil {
					fmt.Printf("Ошибка при обновлении данных: %s", err.Error())
					return
				}
				fmt.Printf("Данные успешно обновлены")
				return
			}
			_, err := keepService.SaveTextData(textData, userModel.UserID)
			if err != nil {
				if errors.Is(err, storage.ErrTextAlredyExist) {
					fmt.Printf("Текстовые данные с таким именем уже существуют")
					return
				}
				fmt.Printf("Ошибка при сохранении данных: %s", err.Error())
				return
			}
			fmt.Println("Успешно сохранено!")
		}
	},
}

func init() {
	rootCmd.AddCommand(savetextdataCmd)
	savetextdataCmd.Flags().Bool("file", false, "Файл с текстовыми данными для сохранения")
	savetextdataCmd.Flags().Bool("update", false, "Обновить существующие данные")

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// savetextdataCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// savetextdataCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}

func parseFromeFile(filePath string) (string, error) {
	if _, err := os.Stat(filePath); err != nil {
		if os.IsNotExist(err) {
			return "", fmt.Errorf("file does not exist")
		}
	}
	file, err := os.Open(filePath)
	if err != nil {
		return "", err
	}
	defer file.Close()

	dataBuf := bytes.Buffer{}
	sc := bufio.NewScanner(file)
	for sc.Scan() {
		dataBuf.WriteString(sc.Text())
	}
	return dataBuf.String(), err
}
