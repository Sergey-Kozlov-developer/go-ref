package main

import (
	"demo/app-demo/account"
	"demo/app-demo/files"
	"fmt"
)

func getMenu() int {
	var variant int
	fmt.Println("Выберите вариант")
	fmt.Println("1. Создать аккаунт")
	fmt.Println("2. Найти аккаунт")
	fmt.Println("3. Удалить аккаунт")
	fmt.Println("4. Выход аккаунт")
	fmt.Scan(&variant)
	return variant
}

func findAccount() {

}
func deleteAccount() {

}

func main() {
	fmt.Println("___Менеджер паролей___")
Menu:
	for {
		variant := getMenu()
		switch variant {
		case 1:
			createAcount()
		case 2:
			findAccount()
		case 3:
			deleteAccount()
		default:
			break Menu
		}
	}
	createAcount()
}

func createAcount() {
	login := promptData("Введите логин")
	password := promptData("Введите пароль")
	url := promptData("Введите URL")

	myAccount, err := account.NewAccount(login, password, url)
	if err != nil {
		return
	}

	vault := account.NewVault()
	vault.AddAccount(*myAccount)
	data, err := vault.ToBytes()
	if err != nil {
		fmt.Println("Ошибка преобразования в JSON")
		return
	}
	files.WriteFile(data, "data.json")

	//myAccount.OutputPassword()
}

func promptData(prompt string) string {
	fmt.Println(prompt + ": ")
	var res string
	fmt.Scan(&res)
	return res
}
