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
		fmt.Printf("Создана директория  %s \n", path)
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
		fmt.Println("Файл не существует")
		var file *os.File
		file, err = os.Create(filePath)
		fmt.Println(filePath)
		if err != nil { // если возникла ошибка
			fmt.Println("Не удается создать файл:", err)
			os.Exit(1) // выходим из программы
		}
		defer func() {
			if err := file.Close(); err != nil {
				log.Printf("Ошибка закрытия файла: %v", err)
			}
		}()
		fmt.Println(file.Name())
	} else {
		fmt.Println("Файл существует")
	}
	writToFile(name, text, filePath)
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
func writToFile(authorName, quote, filePath string) {
	file, err := os.OpenFile(filePath, os.O_WRONLY|os.O_CREATE, 0644)
	if err != nil {
		log.Fatal(err)
	}
	defer func() {
		if err := file.Close(); err != nil {
			log.Printf("Ошибка закрытия файла: %v", err)
		}
	}()

	authorNameMod := fmt.Sprintf("%s \n", authorName)
	_, err = file.WriteString(authorNameMod)
	if err != nil {
		log.Fatal(err)
	}
	_, err = file.WriteString(quote)
	if err != nil {
		log.Fatal(err)
	}

}
