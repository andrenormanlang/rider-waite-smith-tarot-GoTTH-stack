// package database

// import (
//     "gorm.io/driver/postgres"
//     "gorm.io/gorm"
//     "log"
// )

// var DB *gorm.DB

// func ConnectDatabase() {
//     dsn := "host=localhost user=postgres password=shiva7 dbname=tarotdb port=5432 sslmode=disable TimeZone=Europe/Stockholm"
//     var err error
//     DB, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
//     if err != nil {
//         log.Fatal("Failed to connect to database: ", err)
//     }
// }

// package database

// import (
//     // "fmt"
//     "log"
//     "os"
//     "gorm.io/driver/postgres"
//     "gorm.io/gorm"
// )

// var DB *gorm.DB

// func ConnectDatabase() {
    // host := os.Getenv("DB_HOST")
    // user := os.Getenv("DB_USER")
    // password := os.Getenv("DB_PASSWORD")
    // dbname := os.Getenv("DB_NAME")
    // port := os.Getenv("DB_PORT")

    // if host == "" {
    //     // Fallback to default values if environment variables are not set
    //     host = "localhost"
    //     user = "postgres"
    //     password = "shiva7"
    //     dbname = "tarotdb"
    //     port = "5432"
    // }

    // dsn := os.Getenv("DATABASE_URL")

    // dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=disable TimeZone=Europe/Stockholm", host, user, password, dbname, port)

//     var err error
//     DB, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
//     if err != nil {
//         log.Fatal("Failed to connect to database: ", err)
//     }
// }

package database

import (
    "log"
    "gorm.io/driver/postgres"
    "gorm.io/gorm"
)

var DB *gorm.DB

func ConnectDatabase() {
    // Hardcoded PostgreSQL connection string (this is not recommended for production)
    dsn := "postgresql://tarotdb_881o_user:heQdnAfYKGsgtHfdIAEnuhno2LSMWdQK@dpg-crc2gfi6l47c73dafui0-a/tarotdb_881o"
    
    var err error
    // Connect to the database using the connection string
    DB, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
    if err != nil {
        log.Fatal("Failed to connect to database: ", err)
    }
}
