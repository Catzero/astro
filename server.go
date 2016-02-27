package server

import (
	"fmt"
	"github.com/astro/common"
	"github.com/astro/db"
	"github.com/astro/handler"
	"log"
	"net/http"
)

func ServerRun(conf_path string) {
	log.SetFlags(log.Lshortfile | log.LstdFlags | log.Lmicroseconds)

	conf := common.NewConfig(conf_path)
	addr := fmt.Sprintf("%s:%d", conf.IP, conf.Port)
	db := db.NewSqlClient(conf)

	//http.Handle("/echo.cgi", handler.NewEchoHandler(conf))
	http.Handle("/getastro.cgi", handler.NewGetAstroHandler(conf, db))
	err := http.ListenAndServe(addr, nil)
	if err != nil {
		log.Fatal("%s\n", err)
	}
}
