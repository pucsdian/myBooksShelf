package context

import (
	"log"
	"os"
	"testing"

	"github.com/shubhamjagdhane/myBooksShelf/server/users/config"
)

var testConn *Database

func TestMain(m *testing.M) {
	ENV, err := config.InitConfig("../")
	if err != nil {
		log.Fatal("Config file is not successfully loaded: ", err)
	}

	testConn = InitDB(ENV.DBUser, ENV.DBPassword, ENV.DBName, ENV.Host, ENV.Port)

	os.Exit(m.Run())
}
