package migration

import (
	Entity "money.com/entity"
	"time"
	"fmt"
	"log"
    "gorm.io/driver/postgres"
    "gorm.io/gorm"
)

type Migration interface {
	MigrateTables() *gorm.DB
	ConnectToDb() *gorm.DB
}

type MigrationService struct {
   
}

func (migration *MigrationService) MigrateTables() *gorm.DB {

	dsn := "host=db-postgresql-sfo3-21964-apr-3-backup-do-user-9772821-0.b.db.ondigitalocean.com user=motherhonestly password=AVNS_skbZbzrbZFIHAPG8C-i dbname=fm-wallet port=25061 sslmode=require TimeZone=UTC"
	db, err := gorm.Open(postgres.New(postgres.Config{
		DSN: dsn,
		PreferSimpleProtocol: true, // disables implicit prepared statement usage
    }), &gorm.Config{});

    if err != nil {
        log.Fatal("Cannot connect to DB at this time, please try again");
    }

	db.AutoMigrate(&Entity.Wallet{});
	db.AutoMigrate(&Entity.WalletTransactions{});
	db.AutoMigrate(&Entity.Activity{});

	
	postgresDB, err1 := db.DB();
    if err1 == nil {
        postgresDB.SetConnMaxLifetime(time.Millisecond * 200)
        fmt.Println("Database connection timeout has been set to 3 mins")
    }

	return db;

}

func (migration *MigrationService) ConnectToDb() *gorm.DB {
	
	dsn := "host=db-postgresql-sfo3-21964-apr-3-backup-do-user-9772821-0.b.db.ondigitalocean.com user=motherhonestly password=AVNS_skbZbzrbZFIHAPG8C-i dbname=fm-wallet port=25061 sslmode=require TimeZone=UTC"
	db, err := gorm.Open(postgres.New(postgres.Config{
		DSN: dsn,
		PreferSimpleProtocol: true, // disables implicit prepared statement usage
    }), &gorm.Config{})

	if err != nil {
        log.Fatal("Cannot connect to DB at this time, please try again");
    }

	postgresDB, err1 := db.DB();
    if err1 == nil {
        postgresDB.SetConnMaxLifetime(time.Millisecond * 200)
        fmt.Println("Database connection timeout has been set to 3 mins")
    }

    return db;
}
















