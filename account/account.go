package account

import (
	"errors"
	"fmt"
	"math/rand/v2"
	"net/url"
	"time"

	"github.com/fatih/color"
)

var letterRunes = []rune("abcdefghijklmnopABCDEFGHIJKLMNOP1234567890#%^")

// struct
type Account struct {
	Login      string    `json:"login"`
	Password   string    `json:"password"`
	Url        string    `json:"url"`
	CreatedAt  time.Time `json:"createdAt"`
	UpdateddAt time.Time `json:"updateddAt"`
}

func (acc *Account) Output() {
	color.Cyan(acc.Login)
	color.Cyan(acc.Password)
	color.Cyan(acc.Url)
}

func (acc *Account) generatePassword(n int) {
	res := make([]rune, n)
	for i := range res {
		res[i] = letterRunes[rand.IntN(len(letterRunes))]
	}
	acc.Password = string(res)
}

// constructor
func NewAccount(login, password, urlString string) (*Account, error) {

	if login == "" {
		return nil, errors.New("INVALID_LOGIN")
	}
	_, err := url.ParseRequestURI(urlString)
	if err != nil {
		fmt.Println("Неверный формат")
		return nil, errors.New("INVALID_URL")
	}

	newAcc := &Account{
		CreatedAt:  time.Now(),
		UpdateddAt: time.Now(),
		Url:        urlString,
		Login:      login,
		Password:   password,
	}

	if password == "" {
		newAcc.generatePassword(12)
	}
	return newAcc, nil
}
