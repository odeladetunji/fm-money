package repository

import (
	"context"
	"gorm.io/gorm"
	Entity "money.com/entity"
	Migration "money.com/migration"
)

var dba Migration.Migration = &Migration.MigrationService{}

type WalletRepository interface {
	CreditOrDebitWallet(ctx context.Context, wallet Entity.Wallet, walletTransaction Entity.WalletTransactions, activity Entity.Activity) error
	GetWalletByAccountId(accountId int32) (*Entity.Wallet, error)
}

type WalletRepo struct {
	
}

func (emp *WalletRepo) CreditOrDebitWallet(ctx context.Context, wallet Entity.Wallet, walletTransaction Entity.WalletTransactions, activity Entity.Activity) error {
	var database *gorm.DB = dba.ConnectToDb()
	tx := database.WithContext(ctx).Begin()
	defer func() {
		if r := recover(); r != nil {
			tx.Rollback()
		}
	}()

	if err := tx.Error; err != nil {
		return err
	}

	//Update Wallet
	if err3 := tx.Model(&Entity.Wallet{}).Where(&Entity.Wallet{AccountId: wallet.AccountId}).Updates(map[string]interface{}{
		"amount": wallet.Balance}).Error; err3 != nil {
		tx.Rollback()
		return err3
	}

	//Create Transaction
	if err2 := tx.Create(&walletTransaction).Error; err2 != nil {
		tx.Rollback()
		return err2
	}

	//Log Activity
	if err4 := tx.Create(&activity).Error; err4 != nil {
		tx.Rollback()
		return err4
	}

	return tx.Commit().Error
}

func (emp *WalletRepo) GetWalletByAccountId(accountId int32) (*Entity.Wallet, error) {
	var database *gorm.DB = dba.ConnectToDb()
	var wallet Entity.Wallet
	err := database.Where(&Entity.Wallet{AccountId: accountId}).Find(&wallet).Error
	if err != nil {
		return nil, err
	}

	return &wallet, nil
}
