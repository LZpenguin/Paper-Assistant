package model

import (
	"fmt"
	"git.bingyan.net/doc-aid-re-go/config"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"log"
)

var db_gorm *gorm.DB

func init() {
	dsn := fmt.Sprintf(
		"host=%s user=%s password=%s dbname=%s port=%s sslmode=disable TimeZone=Asia/Shanghai",
		config.C.Postgres.Host,
		config.C.Postgres.Username,
		config.C.Postgres.Password,
		config.C.Postgres.DB,
		config.C.Postgres.Port,
	)
	DB, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		panic(err)
	}
	log.Println("DB open Success!")
	db_gorm = DB

	// AutoMigrate
	//err = db_gorm.AutoMigrate(
	//	&Magazine{},
	//	&Paper{},
	//	&SharePaper{},
	//	&Keyword{},
	//	&Favorite{},
	//	&FavoritesFolder{},
	//	&User{},
	//)

	err = db_gorm.AutoMigrate(&Magazine{})
	if err != nil {
		panic(err)
	}
	err = db_gorm.AutoMigrate(&Keyword{})
	if err != nil {
		panic(err)
	}
	//err = db_gorm.Migrator().CreateConstraint(&Paper{}, "MagazineID")
	//if err != nil {
	//	panic(err)
	//}
	err = db_gorm.AutoMigrate(&Paper{})
	if err != nil {
		panic(err)
	}
	err = db_gorm.AutoMigrate(&SharePaper{})
	if err != nil {
		panic(err)
	}

	err = db_gorm.AutoMigrate(&Favorite{})
	if err != nil {
		panic(err)
	}
	err = db_gorm.AutoMigrate(&FavoritesFolder{})
	if err != nil {
		panic(err)
	}
	err = db_gorm.AutoMigrate(&User{})
	if err != nil {
		panic(err)
	}
	err = db_gorm.AutoMigrate(&Feedback{})
	if err != nil {
		panic(err)
	}
}
