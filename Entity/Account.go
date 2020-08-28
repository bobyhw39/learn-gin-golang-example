package Entity

import (
	"github.com/bobyhw39/rest_mvc/Config"
	"math/big"
	//_ "github.com/go-sql-driver/mysql"
	_ "database/sql"
	_ "github.com/lib/pq"
)

type Account struct {
	Id int `json:"id"`
	NumAcc string `json:"account"`
	Password string `json:"password"`
	Name string `json:"name"`
	Balance big.Float `json:"balance"`
	Point big.Float `json:"point"`
	Status int `json:"status"`
}

func GetAllAccounts(account *[]Account) (err error) {
	if err = Config.DB.Find(account).Error; err != nil {
		return err
	}
	return nil
}
