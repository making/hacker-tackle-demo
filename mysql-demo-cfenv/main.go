package main

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"
	"net/url"
	"os"
	"time"

	"github.com/cloudfoundry-community/go-cfenv"
	_ "github.com/go-sql-driver/mysql"
)

var db *sql.DB

type visit struct {
	timestamp int64
	userIP    string
}

func main() {
	var port string
	if port = os.Getenv("PORT"); len(port) == 0 {
		port = "4000"
	}

	var err error
	db, err = sql.Open("mysql", connectionURL())
	if err != nil {
		log.Fatal(err)
	}

	if err := createTable(); err != nil {
		log.Fatal(err)
	}

	http.HandleFunc("/", hello)
	fmt.Println("listening at " + port)
	if err := http.ListenAndServe(":"+port, nil); err != nil {
		log.Fatal("ListenAndServe:", err)
	}
}

func hello(w http.ResponseWriter, r *http.Request) {
	if err := recordVisit(time.Now().UnixNano(), r.RemoteAddr); err != nil {
		http.Error(w, fmt.Sprintf("Could not save visit: %v", err), http.StatusInternalServerError)
		return
	}
	visits, err := queryVisits()
	if err != nil {
		http.Error(w, fmt.Sprintf("Could not get recent visits: %v", err), http.StatusInternalServerError)
		return
	}
	for _, v := range visits {
		fmt.Fprintf(w, "[%s] %s\n", time.Unix(0, v.timestamp), v.userIP)
	}
}

func connectionURL() string {
	appEnv, err := cfenv.Current()
	if err != nil {
		log.Fatal("VCAP_APPLICATION and VCAP_SERVICES must be set!")
	}
	mysqlService, err := appEnv.Services.WithName("mysql-db")
	if err != nil {
		log.Fatal(err)
	}
	uri, ok := mysqlService.Credentials["uri"].(string)
	if !ok {
		log.Fatal("No valid MariabDB uri\n")
	}
	u, err := url.Parse(uri)
	if err != nil {
		log.Fatal("No valid MariabDB uri\n")
	}
	return fmt.Sprintf("%s@tcp(%s)%s", u.User.String(), u.Host, u.Path)
}

func createTable() error {
	stmt := `CREATE TABLE IF NOT EXISTS visits (
			timestamp  BIGINT,
			userip     VARCHAR(255)
		)`
	_, err := db.Exec(stmt)
	return err
}

func recordVisit(timestamp int64, userIP string) error {
	stmt := "INSERT INTO visits (timestamp, userip) VALUES (?, ?)"
	_, err := db.Exec(stmt, timestamp, userIP)
	return err
}

func queryVisits() ([]visit, error) {
	rows, err := db.Query("SELECT timestamp, userip FROM visits ORDER BY timestamp DESC LIMIT 10")
	if err != nil {
		return nil, fmt.Errorf("Could not get recent visits: %v", err)
	}
	defer rows.Close()
	var visits []visit
	for rows.Next() {
		var v visit
		if err := rows.Scan(&v.timestamp, &v.userIP); err != nil {
			return nil, fmt.Errorf("Could not get timestamp/user IP out of row: %v", err)
		}
		visits = append(visits, v)
	}
	return visits, rows.Err()
}
