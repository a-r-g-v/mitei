package main

import (
	"./pkg/punch"
	"fmt"
	"github.com/jinzhu/gorm"
	_ "github.com/mattn/go-sqlite3"
	"os"
	"time"
)

type Table struct {
	Id         uint `gorm:"primary_key"`
	TargetIP   string
	TargetPort string
	BoundPort  string
	CreatedAt  time.Time
}

var DB gorm.DB

func init() {
	db, err := gorm.Open("sqlite3", "database.db")
	DB = db
	if err != nil {
		fmt.Fprintf(os.Stderr, "Model Initialize error:%s \n", err)
		os.Exit(1)
	}

	db.DB().Ping()
	db.AutoMigrate(&Table{})

}

func Allocate(tIP string, tPort string, bPort string) bool {

	punch.Allocate(tIP, tPort, bPort)

	l := Table{TargetIP: tIP, TargetPort: tPort, BoundPort: bPort}
	err := DB.Create(&l).Error
	if err != nil {
		fmt.Fprintf(os.Stderr, "Allocate error:%s \n", err)
		return false
	}
	return true
}

func Release(tIP string, tPort string, bPort string) bool {

	punch.Release(tIP, tPort, bPort)

	err := DB.Debug().Where("target_ip = ?", tIP).Where("target_port = ?", tPort).Where("bound_port = ?", bPort).Delete(&Table{}).Error
	if err != nil {
		fmt.Fprintf(os.Stderr, "Release error:%s \n", err)
		return false
	}
	return true
}

func IsAllocatable(bPort string) int {
	var l Table
	count := 0
	DB.Where("BoundPort = ?", bPort).Find(&l).Count(&count)
	return count
}
