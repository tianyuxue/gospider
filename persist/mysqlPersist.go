package persist

import (
	"net/http"
	"net/url"
	"io/ioutil"
	"database/sql"
    _ "github.com/go-sql-driver/mysql"
	"gospider/config"
)

func MysqlPersist(response *http.Response, addr *url.URL) {
	html, err := ioutil.ReadAll(response.Body)
	if err != nil {
		panic(err)
	}
	defer response.Body.Close()

	user := config.Conf["dbconfig"].(map[string]interface{})["user"].(string)
	password := config.Conf["dbconfig"].(map[string]interface{})["password"].(string)
	url := config.Conf["dbconfig"].(map[string]interface{})["url"].(string)
	dbname := config.Conf["dbconfig"].(map[string]interface{})["dbname"].(string)
    remote := user+":"+password+"@tcp("+url+")/"+dbname
	db, err := sql.Open("mysql", remote)

	if err != nil {
		panic(err.Error())
	}
	defer db.Close()

	// Open doesn't open a connection. Validate DSN data:
	err = db.Ping()
	if err != nil {
		panic(err.Error())
	}

	// Prepare statement for inserting data
	stmtIns, err := db.Prepare("INSERT INTO t1(content, tags) VALUES(?, ?)") // ? = placeholder
	if err != nil {
		panic(err.Error())
	}
	defer stmtIns.Close()


	_, err = stmtIns.Exec(html[0: 500], addr.Host)
	if err != nil {
		panic(err.Error())
	}
}