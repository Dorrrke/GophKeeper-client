/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"errors"
	"fmt"
	"os"

	"github.com/Dorrrke/GophKeeper-client/internal/storage"
	"github.com/spf13/cobra"
)

// binCmd represents the bin command
var binCmd = &cobra.Command{
	Use:   "bin",
	Short: "Отображает данные сохраненых бинарных данных с указанным именем",
	Long: `При вызове отображает бинарные данные с указанным именем.
	При наличии подключения к интернету, данные будут браться из удаленного сервера`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("bin called")
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
			err = keepService.DeleteBinByName(args[0], userModel.UserID)
			if err != nil {
				fmt.Printf("Ошибка при удалении данных %s", err.Error())
				return
			}
			fmt.Printf("Успешное удаление\n")
			return
		}
		bin, err := keepService.GetBinByName(args[0], userModel.UserID)
		if err != nil {
			if errors.Is(err, storage.ErrBinDataNotExist) {
				fmt.Printf("Бинарных данных сохраненных с таким именем не существует.")
				return
			}
			fmt.Printf("Ошибка при получении данных %s", err.Error())
			return
		}
		writeBinFile(bin.Name, bin.Data)
		fmt.Println("Created saved bin file with name " + bin.Name)
	},
}

func init() {
	rootCmd.AddCommand(binCmd)
	binCmd.Flags().Bool("delete", false, "Удаление данных")

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// binCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// binCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}

func writeBinFile(name string, bData []byte) error {
	f, err := os.Open(name + ".bin")
	if err != nil {
		return err
	}
	defer f.Close()

	_, err = f.Write(bData)
	if err != nil {
		return err
	}
	return nil
}
