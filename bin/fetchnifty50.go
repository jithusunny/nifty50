package main

import (
	"encoding/json"
	"log"
	"net/http"
	"os"
	"time"

	"database/sql"
	"github.com/lib/pq"
)

type Item struct {
	Symbol                   string
	Ltp                      string
	Netprice                 string
	TradedQuantity           string
	TurnoverInLakhs          string
	OpenPrice                string
	HighPrice                string
	LowPrice                 string
	PreviousPrice            string
	LastCorpAnnouncementDate string
	LastCorpAnnouncement     string
}

type Response struct {
	Time string
	Data []Item
}

func SetupDB() (txn *sql.Tx, db *sql.DB) {
	url := os.Getenv("DATABASE_URL")
	connection, _ := pq.ParseURL(url)
	connection += " sslmode=require"

	db, err := sql.Open("postgres", connection)

	txn, err = db.Begin()

	if err != nil {
		log.Fatal(err)
	}

	_, err = db.Query(`CREATE TABLE IF NOT EXISTS NIFTY50(
		ID SERIAL PRIMARY KEY,
		SYMBOL TEXT NOT NULL,
		LTP TEXT,
		NETPRICE TEXT,
		TRADEDQUANTITY TEXT,
		TURNOVERINLAKHS TEXT,
		OPENPRICE TEXT,
		HIGHPRICE TEXT,
		LOWPRICE TEXT,
		PREVIOUSPRICE TEXT,
		LASTCORPANNOUNCEMENTDATE TEXT,
		LASTCORPANNOUNCEMENT TEXT,
		GAINERORLOSER INTEGER)`)

	if err != nil {
		log.Fatal(err)
	}

	_, err = db.Query("DELETE FROM NIFTY50")

	if err != nil {
		log.Fatal(err)
	}

	_, err = db.Query("DROP TABLE NIFTYLASTUPDATED")

	_, err = db.Query("CREATE TABLE IF NOT EXISTS NIFTYLASTUPDATED(LAST_UPDATED INTEGER)")

	if err != nil {
		log.Fatal(err)
	}

	_, err = db.Query("DELETE FROM NIFTYLASTUPDATED")

	if err != nil {
		log.Fatal(err)
	}

	return txn, db
}

func GetData(URL string) *Response {
	client := &http.Client{}

	req, err := http.NewRequest("GET", URL, nil)

	if err != nil {
		log.Fatal(err)
	}

	req.Header.Set("Accept", "text/html,application/xhtml+xml,application/xml;q=0.9,    image/webp,*/*;q=0.8")

	resp, err := client.Do(req)

	if err != nil {
		log.Fatal(err)
	}

	if resp.StatusCode != http.StatusOK {
		log.Fatal(resp.Status)
	}

	r := new(Response)

	err = json.NewDecoder(resp.Body).Decode(r)

	if err != nil {
		log.Fatal(err)
	}

	return r
}

func WriteData(db *sql.DB, r *Response, Comptype int) {
	for _, item := range r.Data {
		_, err := db.Exec(`INSERT INTO nifty50 (Symbol, Ltp, Netprice, TradedQuantity, TurnoverInLakhs, OpenPrice, 
			HighPrice, LowPrice, PreviousPrice, LastCorpAnnouncementDate, LastCorpAnnouncement, Gainerorloser) 
			VALUES($1,$2,$3,$4,$5,$6,$7,$8,$9,$10,$11,$12);`, item.Symbol, item.Ltp, item.Netprice, item.TradedQuantity, item.TurnoverInLakhs,
			item.OpenPrice, item.HighPrice, item.LowPrice, item.PreviousPrice, item.LastCorpAnnouncementDate, item.LastCorpAnnouncement, Comptype)

		if err != nil {
			log.Fatal(err)
		}
	}
}

func UpdateTs(db *sql.DB) {
	_, err := db.Exec("INSERT INTO NIFTYLASTUPDATED (LAST_UPDATED) VALUES ($1)", int32(time.Now().Unix()))

	if err != nil {
		log.Fatal(err)
	}
}

func main() {

	txn, db := SetupDB()

	defer db.Close()

	var gainersURL = "https://www.nseindia.com/live_market/dynaContent/live_analysis/gainers/niftyGainers1.json"
	var losersURL = "https://www.nseindia.com/live_market/dynaContent/live_analysis/losers/niftyLosers1.json"

	r := GetData(gainersURL)
	WriteData(db, r, 1)

	r = GetData(losersURL)
	WriteData(db, r, 0)

	UpdateTs(db)

	err := txn.Commit()
	if err != nil {
		log.Fatal(err)
	}
}
