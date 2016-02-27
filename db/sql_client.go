package db

import (
	"database/sql"
	"fmt"
	"github.com/astro/common"
	_ "github.com/go-sql-driver/mysql"
	"log"
	//"strconv"
)

type SqlClient struct {
	db *sql.DB
}

func NewSqlClient(conf *common.Config) *SqlClient {
	dbStr := fmt.Sprintf("%s:%s@/%s", conf.DBUser, conf.DBPassword, conf.DBName)
	db, err := sql.Open("mysql", dbStr)
	if err != nil {
		log.Println(err)
		return nil
	}

	err = db.Ping()
	if err != nil {
		log.Println(err)
		return nil
	}

	return &SqlClient{
		db: db,
	}
}

func (s *SqlClient) GetAstroFortune(ds string, astro string) *AstroFortune {
	rows, err := s.db.Query("select Id, Date, TodayDesc, WeekDesc from AstroFortune where Id = ? and Date = ?", astro, ds)
	if err != nil {
		log.Println(err)
		return nil
	}
	defer rows.Close()
	for rows.Next() {
		var astro []byte
		var ds []byte
		var todaydesc []byte
		var weekdesc []byte
		err := rows.Scan(&astro, &ds, &todaydesc, &weekdesc)
		if err != nil {
			log.Println(err)
			return nil
		}
		return &AstroFortune{
			Astro:     string(astro),
			Date:      string(ds),
			TodayDesc: string(todaydesc),
			WeekDesc:  string(weekdesc),
		}
	}
	return nil
}

func (s *SqlClient) NewAstroFortune(ds string, astro string) bool {
	stmt, err := s.db.Prepare("Insert Into AstroFortune(Id, Date, TodayDesc, WeekDesc) VALUES(?, ?, ?, ?)")
	if err != nil {
		log.Println(err)
		return false
	}
	defer stmt.Close()

	_, err = stmt.Exec(astro, ds, "", "")
	if err != nil {
		log.Println(err)
		return false
	}
	log.Println("insert AstroFortune Date=%s Astro=%s\n", ds, astro)
	return true
}

func (s *SqlClient) SetAstroFortuneWeek(ds string, astro string, weekdesc string) bool {
	stmt, err := s.db.Prepare("UPDATE AstroFortune set WeekDesc=? where Date=? and Id=?")
	if err != nil {
		log.Println(err)
		return false
	}
	defer stmt.Close()

	_, err = stmt.Exec(weekdesc, ds, astro)
	if err != nil {
		log.Println(err)
		return false
	}
	log.Printf("update AstroFortuneWeek Date=%s Astro=%s\n", ds, astro)
	return true
}

func (s *SqlClient) SetAstroFortuneToday(ds string, astro string, todaydesc string) bool {
	stmt, err := s.db.Prepare("UPDATE AstroFortune set TodayDesc=? where Date=? and Id=?")
	if err != nil {
		log.Println(err)
		return false
	}
	defer stmt.Close()

	_, err = stmt.Exec(todaydesc, ds, astro)
	if err != nil {
		log.Println(err)
		return false
	}
	log.Printf("update AstroFortuneToday Date=%s Astro=%s\n", ds, astro)
	return true
}

/*
func (s *SqlClient) GetAccount(wbID uint32) *Account {
	rows, err := s.db.Query("select Uin,UserName,WbID,SessionKey,SessionExpired from Account where WbID=?", wbID)
	if err != nil {
		log.Println(err)
		return nil
	}
	defer rows.Close()

	for rows.Next() {
		var uin int64
		var username []byte
		var wbid []byte
		var key []byte
		var exp []byte
		err := rows.Scan(&uin, &username, &wbid, &key, &exp)
		if err != nil {
			log.Println(err)
			return nil
		}
		WbID, err := strconv.ParseInt(string(wbid), 10, 64)
		Exp, err := strconv.ParseInt(string(exp), 10, 64)
		return &Account{
			Uin:            uin,
			UserName:       string(username),
			WbID:           WbID,
			SessionKey:     string(key),
			SessionExpired: Exp,
		}
	}

	return nil
}

func (s *SqlClient) SetAccount(wbID uint32, acct *Account) bool {
	stmt, err := s.db.Prepare("UPDATE Account set SessionKey=?,SessionExpired=? where WbID=?")
	if err != nil {
		log.Println(err)
		return false
	}
	defer stmt.Close()

	_, err = stmt.Exec(acct.SessionKey, acct.SessionExpired, wbID)
	if err != nil {
		log.Println(err)
		return false
	}
	log.Printf("update Account wbID=%d session=%s\n", wbID, acct.SessionKey)

	return true
}
*/
