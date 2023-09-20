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
		{MigrationFileName: "migrations/20230920162845.sql", Version: 20230920162845},
		{MigrationFileName: "migrations/20230920171126.sql", Version: 20230920171126},
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

	// Finished migration
	fmt.Println("Finished migrations")

}
