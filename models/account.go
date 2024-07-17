package models

import (
	"sync"

	"github.com/zykunov/bankAPI/storage"
)

type BankAccount interface {
	Deposit(amount float64) error  // пополнение баланса.
	Withdraw(amount float64) error // снятие средств.
	GetBalance() float64           // получение баланса

}

type Account struct {
	Id      uint       `json:"id" gorm:"primarykey"`
	Balance float64    `json:"balance"`
	Name    string     `json:"name"`
	Surname string     `json:"surname"`
	mutex   sync.Mutex // для потокобезопасности
}

func AddAccount(a *Account) error {
	return storage.DB.Create(a).Error
}

func (a *Account) Deposit(amount float64) error {
	a.mutex.Lock()
	defer a.mutex.Unlock()
	return storage.DB.Model(a).Where("id = ?", a.Id).Update("balance", amount).Error
}

func (a *Account) Withdraw(amount float64) error {
	a.mutex.Lock()
	defer a.mutex.Unlock()
	return storage.DB.Model(a).Where("id = ?", a.Id).Update("balance", amount).Error
}

func (a *Account) GetBalance() float64 {
	a.mutex.Lock()
	defer a.mutex.Unlock()
	storage.DB.Model(a).Select("balance").Where("id = ?", a.Id).Find(a)
	return a.Balance
}

// Вспомогательная функция, возвращает всю информацию по аккаунту
func GetAccountByID(a *Account, id int) error {
	return storage.DB.Where("id = ?", id).First(a).Error
}

func (a *Account) TableName() string {
	return "account"
}
