package handler

import (
	"encoding/json"
	"fmt"
	"github.com/Catzero/astro/common"
	"github.com/Catzero/astro/db"
	"html/template"
	"io/ioutil"
	"log"
	"net/http"
	_ "net/url"
	"strings"
	"time"
	_ "unsafe"
)

type Template struct {
	Astro     string
	Date      string
	TodayDesc db.AstroTodayDesc
	WeekDesc  db.AstroWeekDesc
}

type GetAstroHandler struct {
	db *db.SqlClient
}

func unescaped(x string) interface{} { return template.HTML(x) }

func NewGetAstroHandler(conf *common.Config, db *db.SqlClient) *GetAstroHandler {
	handler := &GetAstroHandler{
		db: db,
	}

	return handler
}

func (d *GetAstroHandler) RespJson(w http.ResponseWriter, v interface{}, logInfo bool) {
	js, err := json.Marshal(v)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	if logInfo {
		log.Printf("resp=%s\n", js)
	}
	w.Header().Set("Content-Type", "application/json")
	w.Write(js)
}

type GetAstroReq struct {
	Id string `json:"id"`
}

type GetAstroResp struct {
	Description string `json:"description"`
}

func ParseTemplateToStr(tname string) string {
	b, err := ioutil.ReadFile(tname)
	if err != nil {
		log.Println(err)
	}
	s := string(b)
	return s
}

func (d *GetAstroHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	log.Printf("%+v\n", r.Body)

	var req GetAstroReq
	modDecoder := json.NewDecoder(r.Body)
	err := modDecoder.Decode(&req)
	if err != nil {
		log.Println(err.Error())
		//	http.Error(w, err.Error(), http.StatusInternalServerError)
		//	return
	}

	//uin := req.Uin
	//user := d.db.GetUserAttr(uin)

	//if user == nil {
	//	user = d.db.NewUserAttr(uin)
	//	if user == nil {
	//		http.Error(w, err.Error(), http.StatusInternalServerError)
	//		return
	//	}
	//}
	//d.db.SetUserAttr(user)

	//resp := ModUserAttrResp{
	//	NickName:   user.NickName,
	//	HeadImgUrl: user.HeadImgUrl,
	//	Sex:        user.Sex,
	//	Region:     user.Region,
	//}
	r.ParseForm()
	log.Println("go", r.Form["astro"])
	if len(r.Form["astro"]) == 0 {
		return
	}
	s := ParseTemplateToStr("jinriyunshi.html")
	t := template.Must(template.New("ex").Funcs(template.FuncMap{"unescaped": unescaped}).Parse(s))
	//t = t.Funcs(template.FuncMap{"unescaped": unescaped})
	//t, err := template.ParseFiles("jinriyunshi.html")

	if err != nil {
		fmt.Println("sb fail")
	}

	MachineDate := strings.Replace(time.Now().String()[0:10], "-", "", -1)
	user := d.db.GetAstroFortune(MachineDate, r.Form["astro"][0])
	if user != nil {
		var todaydesc db.AstroTodayDesc
		err := json.Unmarshal([]byte(user.TodayDesc), &todaydesc)
		if err != nil {
			fmt.Println("Cant decode json message", err)
		}

		var weekdesc db.AstroWeekDesc
		err = json.Unmarshal([]byte(user.WeekDesc), &weekdesc)
		if err != nil {
			fmt.Println("Cant decode json message", err)
		}

		tt := Template{
			Astro:     r.Form["astro"][0],
			Date:      MachineDate,
			TodayDesc: todaydesc,
			WeekDesc:  weekdesc,
		}
		err = t.Execute(w, tt)
	} else {
		fmt.Println("Cant find Data")
	}

	//d.RespJson(w, resp, true)
}
