package samplesql

import (
	"fmt"
	"log"

	"gorm.io/gorm"
	"gorm.io/gorm/clause"
	"main.go/gorm/samplesql/core"
)

type Accounts struct {
	Id      int
	Balance float64
}

type Transactions struct {
	Id            int
	FromAccountId int
	ToAccountId   int
	Amount        float64
}

func TransferAccounts(fromId int, toId int, amount float64) error {
	if fromId == toId {
		log.Printf("源账户和目标账户相同")
		return fmt.Errorf("源账户和目标账户相同")
	}
	fromAccount := Accounts{}
	toAccount := Accounts{}
	tx := core.DB.Begin()
	if err := tx.Error; err != nil {
		return err
	}
	defer func() {
		if r := recover(); r != nil {
			tx.Rollback()
			panic(r)
		}
	}()

	if err := tx.Clauses(clause.Locking{Strength: clause.LockingStrengthUpdate}).Where("id = ?", fromId).First(&fromAccount).Error; err != nil {
		tx.Rollback()
		log.Printf("查询账户 %d 失败：%v", fromId, err)
		return fmt.Errorf("找不到源账户")
	}
	if fromAccount.Balance < amount {
		tx.Rollback()
		log.Printf("账户:%d, 余额不足,余额:%.2f", fromId, fromAccount.Balance)
		return fmt.Errorf("账户:%d, 余额不足,余额:%.2f", fromId, fromAccount.Balance)
	}

	if err := tx.Clauses(clause.Locking{Strength: clause.LockingStrengthUpdate}).Where("id = ?", toId).First(&toAccount).Error; err != nil {
		tx.Rollback()
		log.Panicf("找不到目标账户")
		return fmt.Errorf("找不到目标账户")
	}

	if err := tx.Model(&Accounts{}).Where("id = ?", fromId).Update("balance", gorm.Expr("balance - ?", amount)).Error; err != nil {
		tx.Rollback()
		return fmt.Errorf("扣减余额失败,原因:%s", err)
	}
	if err := tx.Model(&Accounts{}).Where("id = ?", toId).Update("balance", gorm.Expr("balance + ?", amount)).Error; err != nil {
		tx.Rollback()
		return fmt.Errorf("增加余额失败,原因:%s", err)
	}
	transactions := Transactions{
		FromAccountId: fromId,
		ToAccountId:   toId,
		Amount:        amount,
	}
	if err := tx.Create(&transactions).Error; err != nil {
		tx.Rollback()
		return fmt.Errorf("记录转账记录失败,失败原因:%s", err)
	}

	return tx.Commit().Error
}
