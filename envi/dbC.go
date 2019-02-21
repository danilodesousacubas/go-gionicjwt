package envi

import (
	"fmt"
	"os"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
)

// CreteStrConn for the connection protocol with the application file and its environment
func CreteStrConn() string {
	var config Config
	var err error

	srvPort := os.Getenv("PORT")

	if srvPort == "" {
		config, err = profile("application-dev.json")
	} else {
		fmt.Println("SET PROFILE PRODUCTION")
		config, err = profile("application-prod.json")
	}

	if err != nil {
		fmt.Println(err.Error())
		panic("error open file profile")
	}

	host := config.Database.Host
	port := config.Database.Port
	user := config.Database.User
	dbname := config.Database.Dbname
	password := config.Database.Password
	sslmode := config.Database.Sslmode

	strConn := fmt.Sprintf("host=%s port=%s user=%s dbname=%s password=%s sslmode=%s", host, port, user, dbname, password, sslmode)

	return strConn

}

// DbCon To create uri connection with db
func DbCon() *gorm.DB {

	uri := CreteStrConn()
	db, err := gorm.Open("postgres", uri)
	db.SingularTable(true)

	// fmt.Printf("--- method: DbCon -> db: {%d}", db)

	if err != nil {
		fmt.Println(err.Error())
		panic("Failed to connect to database")
	}

	db.DB().SetMaxIdleConns(0)

	db.LogMode(true)

	return db
}
