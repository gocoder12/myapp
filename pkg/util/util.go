package util

import (
    "database/sql"
    "fmt"
    "log"
    "github.com/spf13/viper"
    _ "github.com/lib/pq"
)

type Config struct {
    Server struct {
        Port string
    }
    Database struct {
        Host     string
        Port     int
        User     string
        Password string
        Dbname   string
    }
}

func LoadConfig() Config {
    viper.SetConfigName("config")
    viper.SetConfigType("yaml")
    viper.AddConfigPath("./configs")
    
    if err := viper.ReadInConfig(); err != nil {
        log.Fatalf("Error reading config file: %v", err)
    }

    var config Config
    if err := viper.Unmarshal(&config); err != nil {
        log.Fatalf("Unable to decode config into struct: %v", err)
    }

    return config
}

func ConnectDB() *sql.DB {
    config := LoadConfig()
    connStr := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable",
        config.Database.Host, config.Database.Port, config.Database.User, config.Database.Password, config.Database.Dbname)
    db, err := sql.Open("postgres", connStr)
    if err != nil {
        log.Fatal(err)
    }
    if err := db.Ping(); err != nil {
        log.Fatal(err)
    }
    return db
}
