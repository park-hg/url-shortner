package main

import (
	"log"

	"traffic-reporter/config"
	"traffic-reporter/internal/pkg"
	"traffic-reporter/internal/shortener/adapter"
)

func main() {
	c := config.MustNewConfig()
	db := pkg.MustConnectMySQL(c.MySQLConfig)

	models := []interface{}{
		adapter.URLMappingTable{},
	}
	err := db.Set("gorm:table_options", "ENGINE=InnoDB CHARSET=ascii COLLATE=ascii_bin").
		AutoMigrate(models...)
	if err != nil {
		log.Fatal(err)
	}
}
