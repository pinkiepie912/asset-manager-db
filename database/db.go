package database

import (
	"asset-manager.com/asset-manager-db/config"
	"asset-manager.com/asset-manager-db/stock"
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func GetConnection(config *config.MysqlConfig, withMigration bool) (*gorm.DB, error) {
	db, err := gorm.Open(mysql.Open(config.GenerateDSN()), &gorm.Config{})
	if err != nil {
		return nil, err
	}

	// Set connection pool
	sqlDB, sqlErr := db.DB()
	if sqlErr != nil {
		return nil, sqlErr
	}
	sqlDB.SetMaxIdleConns(config.MaxIdleConns)
	sqlDB.SetMaxOpenConns(config.MaxOpenConns)
	sqlDB.SetConnMaxLifetime(config.ConnMaxLifetime)

	if withMigration {
		fmt.Println("Start migration")
		_, migErr := migration(db)
		if migErr != nil {
			return nil, migErr
		}
		fmt.Println("Success migration")
	}
	return db, nil
}

func migration(db *gorm.DB) (*gorm.DB, error) {
	migErr := db.AutoMigrate(
		&stock.Firm{},
		&stock.Stock{},
		&stock.Chart{},
	)

	if migErr != nil {
		return nil, migErr
	}

	return db, nil
}