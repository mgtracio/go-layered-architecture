package db

import (
	"layered_architecture/pkg/config"
	"log"
	"os"
	"path/filepath"

	"github.com/syndtr/goleveldb/leveldb"
	"gorm.io/gorm"
)

var (
	DB      *gorm.DB
	LevelDB *leveldb.DB
)

func Init() {
	DB = config.InitDB()
	wd, _ := os.Getwd()
	dbPath := filepath.Join(wd, "../../db/testdb")
	var err error
	LevelDB, err = leveldb.OpenFile(dbPath, nil)
	if err != nil {
		log.Fatalf("Failed to open LevelDB: %v", err)
	}
}

func Close() {
	sqlDB, err := DB.DB()
	if err != nil {
		log.Fatalf("Error closing DB connection: %v", err)
	}
	sqlDB.Close()
	LevelDB.Close()
}
