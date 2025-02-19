package account

import (
	"demo/app-demo/files"
	"encoding/json"
	"github.com/fatih/color"
	"strings"
	"time"
)

type Vault struct {
	Accounts   []Account `json:"accounts"`
	UpdateddAt time.Time `json:"updateddAt"`
}

func NewVault() *Vault {
	file, err := files.ReadFile("data.json")
	if err != nil {
		return &Vault{
			Accounts:   []Account{},
			UpdateddAt: time.Now(),
		}
	}
	var vault Vault
	err = json.Unmarshal(file, &vault)
	if err != nil {
		color.Red("Не удалось разобрать файл JSON")
	}
	return &vault
}

func (vault *Vault) DeleteAccount(url string) bool {
	var accounts []Account
	isDeleted := false
	for _, account := range vault.Accounts {
		isMatched := strings.Contains(account.Url, url)
		if !isMatched {
			accounts = append(accounts, account)
			continue
		}
		isDeleted = true
	}
	vault.Accounts = accounts
	vault.save()

	return isDeleted
}

func (vault *Vault) FindAccountsByUrl(url string) []Account {
	var accounts []Account
	for _, account := range vault.Accounts {
		isMatched := strings.Contains(account.Url, url)
		if !isMatched {
			accounts = append(accounts, account)
		}
	}
	return accounts
}

func (vault *Vault) AddAccount(acc Account) {
	vault.Accounts = append(vault.Accounts, acc)
	vault.save()
}

func (vault *Vault) ToBytes() ([]byte, error) {
	file, err := json.Marshal(vault)
	if err != nil {
		return nil, err
	}
	return file, nil
}
func (vault *Vault) save() {
	vault.UpdateddAt = time.Now()
	data, err := vault.ToBytes()
	if err != nil {
		color.Red("Не удалось преобразовать файл JSON")
	}
	files.WriteFile(data, "data.json")
}
