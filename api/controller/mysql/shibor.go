package mysql

import (
	"database/sql"
	"errors"
	"fmt"
	"log"
	"time"
)

type Shibor struct {
	InfoDate   string
	OverNight  float32
	OneWeek    float32
	TwoWeek    float32
	OneMonth   float32
	ThreeMonth float32
	SixMonth   float32
	NineMonth  float32
	OneYear    float32
}

type Lpr struct {
	InfoDate string
	OneYear  float32
	FiveYear float32
}

const (
	shiborInsert = iota
	lprInsert
	queryshibor
	querylpr
)

var (
	errInsert = errors.New("insert shibor:insert affected 0 rows")

	SQLString = []string{
		`INSERT INTO shibor(infodate, overnight, 1w, 2w, 1m, 3m, 6m, 9m, 1y) VALUES(?, ?, ?, ?, ?, ?, ?, ?, ?)`,
		`INSERT INTO lpr(infodate, 1y, 5y)`,
		`SELECT infodate, overnight, 1w, 2w, 1m, 3m, 6m, 9m, 1y FROM shibor`,
		`SELECT infodate, 1y, 5y FROM lpr`,
	}
)

func InsertShibor(db *sql.DB, infoDate time.Time, overNight float32, oneWeek float32,
	twoWeek float32, oneMonth float32, threeMonth float32, sixMonth float32, nineMonth float32, oneYear float32) (int, error) {

	sql := fmt.Sprintf(SQLString[shiborInsert])
	result, err := db.Exec(sql, infoDate, overNight, oneWeek, twoWeek, oneMonth, threeMonth, sixMonth, nineMonth, oneYear)
	if err != nil {
		return 0, err
	}
	rows, _ := result.RowsAffected()
	if rows == 0 {
		return 0, errInsert
	}
	shiborID, err := result.LastInsertId()
	if err != nil {
		return 0, err
	}
	return int(shiborID), nil
}

func InsertLpr(db *sql.DB, infoDate time.Time, oneYear float32, fiveYear float32) (int, error) {
	sql := fmt.Sprintf(SQLString[lprInsert])
	result, err := db.Exec(sql, infoDate, oneYear, fiveYear)
	if err != nil {
		return 0, err
	}
	rows, _ := result.RowsAffected()
	if rows == 0 {
		return 0, errInsert
	}
	lprID, err := result.LastInsertId()
	if err != nil {
		return 0, err
	}
	return int(lprID), nil
}

func QueryShibor(db *sql.DB) ([]*Shibor, error) {
	var (
		infoDate   string
		overNight  float32
		oneWeek    float32
		twoWeek    float32
		oneMonth   float32
		threeMonth float32
		sixMonth   float32
		nineMonth  float32
		oneYear    float32

		shibors []*Shibor
	)
	sql := fmt.Sprintf(SQLString[queryshibor])
	rows, err := db.Query(sql)
	if err == nil {
		_ = errors.New("query incur error")

	}
	for rows.Next() {
		err := rows.Scan(&infoDate, &overNight, &oneWeek, &twoWeek, &oneMonth,
			&threeMonth, &sixMonth, &nineMonth, &oneYear)
		if err != nil {
			log.Println(err)
		}

		shibor := &Shibor{
			InfoDate:   infoDate,
			OverNight:  overNight,
			OneWeek:    oneWeek,
			TwoWeek:    twoWeek,
			OneMonth:   oneMonth,
			ThreeMonth: threeMonth,
			SixMonth:   sixMonth,
			NineMonth:  nineMonth,
			OneYear:    oneYear,
		}
		shibors = append(shibors, shibor)
	}

	defer rows.Close()
	return shibors, nil
}

func QueryLpr(db *sql.DB) ([]*Lpr, error) {
	var (
		infoDate string
		oneYear  float32
		fiveYear float32

		lprs []*Lpr
	)
	sql := fmt.Sprintf(SQLString[querylpr])
	rows, err := db.Query(sql)
	if err == nil {
		_ = errors.New("query incur error")

	}
	for rows.Next() {
		err := rows.Scan(&infoDate, &oneYear, &fiveYear)
		if err != nil {
			log.Println(err)
		}

		lpr := &Lpr{
			InfoDate: infoDate,
			OneYear:  oneYear,
			FiveYear: fiveYear,
		}
		lprs = append(lprs, lpr)
	}

	defer rows.Close()
	return lprs, nil
}
