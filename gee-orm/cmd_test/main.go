package main

import (
	"fmt"
	geeorm "gee-orm"
	_ "github.com/mattn/go-sqlite3"
)

func main() {
	engine, _ := geeorm.NewEngine("sqlite3", "gee.db")
	defer engine.Close()
	s := engine.NewSession()
	result, _ := s.Raw("INSERT INTO User(`Name`) values (?), (?)", "boluo", "zhangsan").Exec()
	count, _ := result.RowsAffected()
	fmt.Printf("Exec success, %d affected\n", count)
}
