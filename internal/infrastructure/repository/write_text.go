package repository

import (
	"context"
	"fmt"
	"log"
	"os"
	"strings"
)

func (r *Repository) WriteTextToFile(ctx context.Context, name, email, text, directory string) error {
	_ = ctx
	fmt.Println(email)
	//Проверка существования директории
	path := fmt.Sprintf("./app_files/%s", directory)
	info, err := os.Stat(path)
	if os.IsNotExist(err) {
		fmt.Printf("Директория %s не существует\n", path)
		makeDir(path)
	} else if info.IsDir() {
		fmt.Printf("Директория %s существует\n", path)
	} else {
		fmt.Printf("%s существует, но это не директория\n", path)
	}

	fmt.Printf("name: %s \n email %s \n text %s \n directory %s \n", name, email, text, directory)
	//Создание файла для записи
	emailForFileName := strings.Replace(email, ".", "_", 100)
	filePath := fmt.Sprintf("%s/%s.txt", path, emailForFileName)
	fmt.Println(filePath)
	//Проверка существования файла
	_, err = os.Stat(filePath)
	if os.IsNotExist(err) {
		fmt.Println("File does not exist")
		var file *os.File
		file, err = os.Create(filePath)
		if err != nil { // если возникла ошибка
			fmt.Println("Unable to create file:", err)
			os.Exit(1) // выходим из программы
		}
		defer func() {
			if err := file.Close(); err != nil {
				log.Printf("Ошибка закрытия файла: %v", err)
			}
		}()
		fmt.Println(file.Name())
	} else {
		fmt.Println("File exists")
	}
	return nil
}

// Функция создания директории
func makeDir(path string) {
	err := os.MkdirAll(path, 0755) // Права доступа
	if err != nil {
		fmt.Println("Ошибка создания директории:", err)
		return
	}
	fmt.Println("Директория успешно создана")
}
