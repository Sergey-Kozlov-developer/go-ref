package files

import (
	"fmt"
	"os"
)

func ReadFile(name string) ([]byte, error) {
	data, err := os.ReadFile(name)
	if err != nil {
		fmt.Println(err)
		return nil, err
	}
	return data, nil
}

func WriteFile(content []byte, name string) {
	// создать файл
	file, err := os.Create(name)
	if err != nil {
		fmt.Println(err)
	}
	defer file.Close()
	// записать в файл
	_, err = file.Write(content)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println("Запись упешна")

}
