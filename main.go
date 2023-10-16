package main

import (
	"fmt"
	"log"
	"os"
	"sort"

	"example/models"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

type DatabaseVersionMigrationFile struct {
	MigrationFileName string
	Version           uint
}

func main() {

	DatabaseVersions := []DatabaseVersionMigrationFile{
		{MigrationFileName: "migrations/20230920174315.sql", Version: 20230920174315},
		{MigrationFileName: "migrations/20230920174516.sql", Version: 20230920174516},
		{MigrationFileName: "migrations/20230920174604.sql", Version: 20230920174604},
		{MigrationFileName: "migrations/20230920174735.sql", Version: 20230920174735},
		{MigrationFileName: "migrations/20230926122021.sql", Version: 20230926122021},
		{MigrationFileName: "migrations/20230926142733.sql", Version: 20230926142733},
	}

	// Just to make sure the versions get applied in the correct order.
	sort.Slice(DatabaseVersions, func(i, j int) bool {
		return DatabaseVersions[i].Version < DatabaseVersions[j].Version
	})

	// Get a database connection.
	db, err := gorm.Open(sqlite.Open("data.db"), &gorm.Config{})
	if err != nil {
		panic("failed to connect database create it.")
	}

	// Get the current database version
	var dbversion models.DBVersion
	if err := db.Last(&dbversion).Error; err != nil {
		// If none assume version 0 -> initial migration
		// db.AutoMigrate(&Product{}, &DBVersion{})
		dbversion = models.DBVersion{Version: 0}
	}
	fmt.Println("Current Database-Version is:", dbversion.Version)

	// Run versioned migrations until head

	fmt.Println("Database-Version is ", dbversion.Version)
	for _, version := range DatabaseVersions {
		fmt.Println("Database-Version is ", dbversion.Version, " < ", version.Version, " -> Do migration? ", dbversion.Version < version.Version)
		if dbversion.Version < version.Version {
			content, err := os.ReadFile(version.MigrationFileName)
			if err != nil {
				log.Fatal(err)
			}
			//fmt.Println(string(content))
			db.Exec(string(content))
			dbversion = models.DBVersion{Version: version.Version}
			db.Create(&dbversion)
			fmt.Println("Migrated to version: ", dbversion.Version)
		}
	}

	db.Create(&models.User{Name: "John", Age: 42})

	fmt.Println("Last added user: ", db.Last(&models.User{}))

	// Finished migration
	fmt.Println("Finished migrations")

}
