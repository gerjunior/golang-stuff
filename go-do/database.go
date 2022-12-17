package main

import (
	"fmt"

	"github.com/jinzhu/gorm"
)

var (
	gormDB *gorm.DB
)

func init() {
	d, err := gorm.Open("mysql", "root:root@(localhost)/tododb?charset=utf8&parseTime=True&loc=Local")
	fmt.Println("CONNECTED")
	if err != nil {
		panic(err)
	}

	gormDB = d
}

func GetDB() *gorm.DB {
	return gormDB
}
