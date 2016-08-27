package controllers

import (
	"os"
	"log"

	"database/sql"
	"github.com/lib/pq"

	"github.com/astaxie/beego"

	// "github.com/jithusunny/nifty50/niftyutils"
)

type MainController struct {
	beego.Controller
}

type Item struct {
	Symbol string
	Ltp string
	Netprice string
	TradedQuantity string
	TurnoverInLakhs string
	OpenPrice string
	HighPrice string
	LowPrice string
	PreviousPrice string
	LastCorpAnnouncementDate string
	LastCorpAnnouncement string
}

const (
	AUTHOR  = "Jithu Sunny <jithusunnyk@gmail.com>"
	VERSION = "0.0.1"

	DEBUG      = false
	DB_USERNAME = "postgres"
	DB_PWD = "postgres123"
	DB_NAME = "postgres"
)

func GetRecords (db *sql.DB, comptype int) []Item {
	rows, err := db.Query(`SELECT SYMBOL, LTP, NETPRICE, TRADEDQUANTITY, TURNOVERINLAKHS, 
		OPENPRICE , HIGHPRICE, LOWPRICE, PREVIOUSPRICE, LASTCORPANNOUNCEMENTDATE, 
		LASTCORPANNOUNCEMENT FROM NIFTY50 where GAINERORLOSER=$1`, comptype)

	if err != nil {
		if err != sql.ErrNoRows {
			log.Fatal(err)
		}
	}

	defer rows.Close()

	records := []Item{}

	for rows.Next() {
            var r Item

            err = rows.Scan(&r.Symbol, &r.Ltp, &r.Netprice, &r.TradedQuantity, 
            	&r.TurnoverInLakhs,	&r.OpenPrice, &r.HighPrice, &r.LowPrice, &r.PreviousPrice, 
            	&r.LastCorpAnnouncementDate, &r.LastCorpAnnouncement)

            if err != nil {
                log.Fatal(err)
            }
            records = append(records, r)
    }
    err = rows.Err()

    if err != nil {
		log.Fatal(err)
	}

    return records
}

func GetLastUpdatedTs (db *sql.DB) (ts int) {
	err := db.QueryRow("SELECT LAST_UPDATED FROM NIFTYLASTUPDATED LIMIT 1").Scan(&ts)
	
	if err != nil {
		if err != sql.ErrNoRows {
			log.Fatal(err)
		}
	}
	return ts
}


func (main *MainController) Get() {
	url := os.Getenv("DATABASE_URL")
    connection, _ := pq.ParseURL(url)
    connection += " sslmode=require"

    db, err := sql.Open("postgres", connection)
	
	if err != nil {
		log.Fatal(err)
	}

	defer db.Close()

    main.Data["grecords"] = GetRecords(db, 1)
    main.Data["lrecords"] = GetRecords(db, 0)

   	ts := GetLastUpdatedTs(db)

   	main.Data["when"] = ts

	main.TplName = "index.tpl"
}