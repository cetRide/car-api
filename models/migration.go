package models

import (
	"github.com/jinzhu/gorm"
	gormigrate "gopkg.in/gormigrate.v1"
)

func MigrateDatabase(db *gorm.DB) error {
	m := gormigrate.New(db, gormigrate.DefaultOptions, []*gormigrate.Migration{
		{
			ID: "2019092600",
			Migrate: func(tx *gorm.DB) error {
				return tx.AutoMigrate(&Accounts{}).Error
			},
			Rollback: func(tx *gorm.DB) error {
				return tx.DropTable("accounts").Error
			},
		},
		{
			ID: "2019092601",
			Migrate: func(tx *gorm.DB) error {
				return tx.AutoMigrate(&Followings{}).Error
			},
			Rollback: func(tx *gorm.DB) error {
				return tx.DropTable("following").Error
			},
		},
		{
			ID: "2019092602",
			Migrate: func(tx *gorm.DB) error {
				return tx.AutoMigrate(&Posts{}).Error
			},
			Rollback: func(tx *gorm.DB) error {
				return tx.DropTable("posts").Error
			},
		},
		{
			ID: "2019092603",
			Migrate: func(tx *gorm.DB) error {
				return tx.AutoMigrate(&Likes{}).Error
			},
			Rollback: func(tx *gorm.DB) error {
				return tx.DropTable("likes").Error
			},
		},
		{
			ID: "2019092604",
			Migrate: func(tx *gorm.DB) error {
				return tx.AutoMigrate(&Comments{}).Error
			},
			Rollback: func(tx *gorm.DB) error {
				return tx.DropTable("comments").Error
			},
		},
		{
			ID: "2019092605",
			Migrate: func(tx *gorm.DB) error {
				return tx.AutoMigrate(&Reposts{}).Error
			},
			Rollback: func(tx *gorm.DB) error {
				return tx.DropTable("reposts").Error
			},
		},
	})
	m.InitSchema(func(tx *gorm.DB) error {
		err := tx.AutoMigrate(
			&Accounts{},
			&Followings{},
			&Posts{},
			&Likes{},
			&Comments{},
			&Reposts{},
		).Error
		if err != nil {
			return err
		}

		//add foreign keys
		if err := db.Model(&Followings{}).AddForeignKey("following_id", "accounts(account_id)", "RESTRICT", "RESTRICT").Error; err != nil {
			return err
		}
		if err := db.Model(&Followings{}).AddForeignKey("follower_id", "accounts(account_id)", "RESTRICT", "RESTRICT").Error; err != nil {
			return err
		}
		if err := db.Model(&Posts{}).AddForeignKey("user_id", "accounts(account_id)", "RESTRICT", "RESTRICT").Error; err != nil {
			return err
		}
		if err := db.Model(&Likes{}).AddForeignKey("user_id", "accounts(account_id)", "RESTRICT", "RESTRICT").Error; err != nil {
			return err
		}
		if err := db.Model(&Likes{}).AddForeignKey("post_id", "posts(post_id)", "RESTRICT", "RESTRICT").Error; err != nil {
			return err
		}
		if err := db.Model(&Comments{}).AddForeignKey("user_id", "accounts(account_id)", "RESTRICT", "RESTRICT").Error; err != nil {
			return err
		}
		if err := db.Model(&Comments{}).AddForeignKey("post_id", "posts(post_id)", "RESTRICT", "RESTRICT").Error; err != nil {
			return err
		}
		if err := db.Model(&Reposts{}).AddForeignKey("user_id", "accounts(account_id)", "RESTRICT", "RESTRICT").Error; err != nil {
			return err
		}
		if err := db.Model(&Reposts{}).AddForeignKey("post_id", "posts(post_id)", "RESTRICT", "RESTRICT").Error; err != nil {
			return err
		}
		return nil
	})
	return m.Migrate()
}
