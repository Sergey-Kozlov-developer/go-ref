package main

import (
	"demo/app-demo/account"
	"fmt"
	"github.com/fatih/color"
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

func findAccount(vault *account.Vault) {
	url := promptData("Введите URL")
	accounts := vault.FindAccountsByUrl(url)
	if len(accounts) == 0 {
		color.Red("Аккаунтов не найдено")
	}
	for _, account := range accounts {
		account.Output()
	}
}
func deleteAccount(vault *account.Vault) {
	url := promptData("Введите URL")
	isDeleted := vault.DeleteAccount(url)
	if isDeleted {
		color.Green("Удалено")
	} else {
		color.Red("Не найден")
	}
}

func main() {
	fmt.Println("___Менеджер паролей___")
	vault := account.NewVault()
Menu:
	for {
		variant := getMenu()
		switch variant {
		case 1:
			createAcount(vault)
		case 2:
			findAccount(vault)
		case 3:
			deleteAccount(vault)
		default:
			break Menu
		}
	}
}

func createAcount(vault *account.Vault) {
	login := promptData("Введите логин")
	password := promptData("Введите пароль")
	url := promptData("Введите URL")

	myAccount, err := account.NewAccount(login, password, url)
	if err != nil {
		return
	}

	vault.AddAccount(*myAccount)

	//myAccount.OutputPassword()
}

func promptData(prompt string) string {
	fmt.Println(prompt + ": ")
	var res string
	fmt.Scan(&res)
	return res
}
