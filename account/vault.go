package account

import (
	"encoding/json"
	"time"
)

type Vault struct {
	Accounts   []Account `json:"accounts"`
	UpdateddAt time.Time `json:"updateddAt"`
}

func NewVault() *Vault {
	return &Vault{
		Accounts:   []Account{},
		UpdateddAt: time.Now(),
	}
}

func (vault *Vault) AddAccount(acc Account) {
	vault.Accounts = append(vault.Accounts, acc)
	vault.UpdateddAt = time.Now()
}

func (vault *Vault) ToBytes() ([]byte, error) {
	file, err := json.Marshal(vault)
	if err != nil {
		return nil, err
	}
	return file, nil
}
