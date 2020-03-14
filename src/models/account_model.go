package models

import (
	"database/sql"
	"entities"
)

type AccountModel struct {
	Db *sql.DB
}

func (accountModel AccountModel) CreateAccount(account *entities.Account) (err error) {
	result, err := accountModel.Db.Exec("insert into account(id,username,password) value (?,?,?)", account.ID, account.Username, account.Password)
	if err != nil {
		return err
	} else {
		account.ID, _ = result.LastInsertId()
		return nil
	}
}

func (accountModel AccountModel) UpdateAccount(account *entities.Account) (int64, error) {
	result, err := accountModel.Db.Exec("update account set username = ?,password = ?  where id =? ", account.Username, account.Password,account.ID)
	if err != nil {
		return 0, err
	} else {
		return result.RowsAffected() //số dòng thay đổi
	}
}
