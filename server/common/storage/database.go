package storage

import (
	"fmt"
	"log"
	"main/server/common/controller"
	"main/server/common/globals"
	"os"
	"time"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

var DB *gorm.DB

type Config struct {
	Host     string
	Port     string
	Password string
	User     string
	DBName   string
	SSLMode  string
}

func Default() *Config {
	return &Config{
		Host:     globals.Env.DB_HOST,
		Port:     globals.Env.DB_PORT,
		Password: globals.Env.DB_PASS,
		User:     globals.Env.DB_USER,
		SSLMode:  globals.Env.DB_SSLMODE,
		DBName:   globals.Env.DB_NAME,
	}
}

func Connect(config *Config) {
	dsn := fmt.Sprintf(
		"host=%s port=%s user=%s password=%s dbname=%s sslmode=%s",
		config.Host, config.Port, config.User, config.Password, config.DBName, config.SSLMode,
	)

	db, err := gorm.Open(postgres.New(postgres.Config{
		DSN: dsn,
		PreferSimpleProtocol: true,
		// DisableForeignKeyConstraintWhenMigrating: true,
	}), &gorm.Config{
		Logger: logger.New(
			log.New(os.Stdout, "\r\n", log.LstdFlags), // io writer
			logger.Config{
			  SlowThreshold:              time.Second,   // Slow SQL threshold
			  LogLevel:                   logger.Silent, // Log level
			  IgnoreRecordNotFoundError: false,           // Ignore ErrRecordNotFound error for logger
			  ParameterizedQueries:      false,           // Don't include params in the SQL log
			  Colorful:                  true,          // Disable color
			},
		  ),
	})

	if err != nil {
		// handle errors
		fmt.Println("Database connection error", err)
	}

	DB = db
}

func Paginate(ctx *controller.Context) func(db *gorm.DB) *gorm.DB {
	return func (db *gorm.DB) *gorm.DB {
		pageSize := ctx.PageSize()
		offset := (ctx.Page() - 1) * pageSize
		return db.Offset(offset).Limit(pageSize)
	}
}
