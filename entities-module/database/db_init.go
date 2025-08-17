package database

import (
	"entities-module/query"
	"fmt"
	"os"

	"gorm.io/driver/postgres"
	"gorm.io/gen"
	"gorm.io/gorm"
)

func DbInit() (DB *gorm.DB, Q *query.Query) {
	// Initialize the database connection
	host := os.Getenv("DB_HOST")
	port := os.Getenv("DB_PORT")
	username := os.Getenv("DB_USER")
	password := os.Getenv("DB_PASSWORD")
	dbName := os.Getenv("DB_NAME")
	sslMode := os.Getenv("DB_SSL")

	dsn := fmt.Sprintf("host=%v user=%v password=%v dbname=%v port=%v sslmode=%v", host, username, password, dbName, port, sslMode)
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}

	q := query.Use(db)

	return db, q
}

func GenModelQuery(db *gorm.DB) {
	// Generate model and query files using GORM Gen
	g := gen.NewGenerator(gen.Config{
		OutPath:      "entities-module/query",
		ModelPkgPath: "entities-module/model",
		Mode:         gen.WithoutContext | gen.WithDefaultQuery | gen.WithQueryInterface, // generate mode
	})

	g.UseDB(db)
	g.ApplyBasic(
		g.GenerateAllTable()...,
	)
	g.Execute()
}
