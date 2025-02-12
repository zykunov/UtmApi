package handlers

import (
	"log"
	"strconv"
	"sync"

	"github.com/gin-gonic/gin"
	"github.com/zykunov/bankAPI/helpers"
	"github.com/zykunov/bankAPI/models"
)

var Wg sync.WaitGroup // Для обуздания рутин

// @Summary AddAccount
// @Tags Account
// @Description Создание аккаунта. Добавил имя и фамилию, чтоб по интереснее было, но это не обязательные параметры.
// @Accept json
// @Produce json
// @Param input body helpers.AccountAdd true "account add"
// @Success 201 {object} helpers.AccountSuccess
// @Failure 404 {object} models.Account
// @Router /accounts [post]
func AddAccount(c *gin.Context) {
	var account models.Account

	if c.BindJSON(&account) != nil {
		c.String(400, "parameter error")
		return
	}

	err := models.AddAccount(&account)
	if err != nil {
		helpers.RespondJSON(c, 404, account)
		return
	}
	helpers.RespondJSON(c, 201, account)

	log.Printf("user adding operation, accountId: %d", account.Id)
}

// @Summary Deposit
// @Tags Account
// @Description Внесение денег
// @Accept json
// @Produce json
// @Param id  path int true "User ID"
// @Param input body  helpers.SummAmount true "deposit"
// @Success 200 {object} helpers.AccountSuccess
// @Failure 404 {object} models.Account
// @Router /accounts/{id}/deposit [post]
func Deposit(c *gin.Context) {

	var summ helpers.SummAmount
	var account models.Account
	idRaw := c.Params.ByName("id")
	id, err := strconv.Atoi(idRaw)
	if err != nil {
		log.Println("user ID error")
		return
	}
	log.Printf("Deposit operation, accountId: %d", account.Id)

	if c.BindJSON(&summ) != nil {
		c.String(400, "parameter error")
		return
	}

	err = models.GetAccountByID(&account, id)
	if err != nil {
		helpers.RespondJSON(c, 404, account)
		return
	}

	newBalance := account.Balance + summ.Amount

	Wg.Add(1)
	go func(c *gin.Context, a models.Account, newBalance float64) {
		defer Wg.Done()
		err = account.Deposit(newBalance)
		if err != nil {
			helpers.RespondJSON(c, 404, account)
		}
		helpers.RespondJSON(c, 200, account)
	}(c, account, newBalance)

	Wg.Wait()

}

// @Summary Withdrow
// @Tags Account
// @Description Снятие денег
// @Accept json
// @Produce json
// @Param id  path int true "User ID"
// @Param input body helpers.SummAmount true "withdraw"
// @Success 200 {object} helpers.AccountSuccess
// @Failure 404 {object} models.Account
// @Router /accounts/{id}/withdraw [post]
func Withdraw(c *gin.Context) {

	var summ helpers.SummAmount
	idRaw := c.Params.ByName("id")
	id, err := strconv.Atoi(idRaw)
	if err != nil {
		log.Println("user ID error")
		return
	}

	var account models.Account

	if c.BindJSON(&summ) != nil {
		c.String(400, "parameter error")
		return
	}

	err = models.GetAccountByID(&account, id)
	if err != nil {
		helpers.RespondJSON(c, 404, account)
		return
	}

	newBalance := account.Balance - summ.Amount

	// Проверка на наличие достаточной суммы для вывода
	if newBalance < 0 {
		log.Println("Недостаточно средств")
		helpers.RespondJSON(c, 404, "not enough money to withdrow")
		return
	}

	Wg.Add(1)
	go func(c *gin.Context, a models.Account, newBalance float64) {
		defer Wg.Done()
		err = account.Withdraw(newBalance)
		if err != nil {
			helpers.RespondJSON(c, 404, account)
		}
		helpers.RespondJSON(c, 200, account)

		log.Printf("Withdraw operation, accountId: %d", account.Id)
	}(c, account, newBalance)
	Wg.Wait()
}

// @Summary GetBalance
// @Tags Account
// @Description Получение баланса
// @Accept json
// @Produce json
// @Param id  path int true "User ID"
// @Success 200 {object} helpers.BalanceSuccess
// @Failure 404 {object} helpers.BalanceFail
// @Router /accounts/{id}/balance [get]
func GetBalance(c *gin.Context) {

	var account models.Account

	idRaw := c.Params.ByName("id")
	id, err := strconv.Atoi(idRaw)

	if err != nil {
		log.Println("user ID error")
		return
	}

	account.Id = uint(id)
	Wg.Add(1)
	go func(c *gin.Context, a models.Account) {
		defer Wg.Done()
		balance := account.GetBalance()
		helpers.RespondJSON(c, 200, balance)
	}(c, account)
	Wg.Wait()

	log.Printf("Get balance command, accountId: %d", account.Id)

}
