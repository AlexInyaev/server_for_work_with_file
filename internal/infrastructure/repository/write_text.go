package repository

import (
	"context"
	"fmt"
	"os"
)

func (r *Repository) WriteTextToFile(ctx context.Context, name, email, text, directory string) error {
	_ = ctx

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
