package main

import (
	"encoding/json"
	"fmt"
	"github.com/PuerkitoBio/goquery"
	//"github.com/opesun/goquery"
	//	"io/ioutil"
	"log"
	"star.org/common"
	"star.org/db"
	"strings"
	"time"
	//	"net/http"
)

func CallToday(starname string) {
	var url = "http://astro.sina.com.cn/fate_day_" + starname
	var array [20][2]string
	doc, _ := goquery.NewDocument(url)
	cnt := 0
	doc.Find(".content .tb td").Each(func(i int, contentSelection *goquery.Selection) {
		title := contentSelection.Text()
		if title != "" {
			if array[cnt][0] == "" {
				array[cnt][0] = title
			} else {
				array[cnt][1] = title
				cnt = cnt + 1
			}
		} else {
			str, _ := contentSelection.Html()
			if str != "" {
				sb := strings.TrimPrefix(str, "<i class=\"")
				sb2 := strings.TrimSuffix(sb, "\"></i>")
				array[cnt][1] = sb2
				cnt = cnt + 1
			}
		}
		//		fmt.Println("xxoo", i+1, "title is :", title)
	})
	/*
		for i := 0; i < cnt; i++ {
			fmt.Println(array[i][0], ":", array[i][1])
		}
	*/

	titname := strings.TrimSuffix(doc.Find(".head .tit .tit_n").Text(), "关注")
	titday := doc.Find(".head .tit .tit_d").Text()
	fmt.Println("sb", titname, " ", titday)

	detail, _ := doc.Find(".content .words").Html()

	todaydesc := db.AstroTodayDesc{
		Astro:   starname,
		Zonghe:  array[0][1],
		Love:    array[1][1],
		Work:    array[2][1],
		Money:   array[3][1],
		Health:  array[4][1],
		Kaiyun:  array[5][1],
		Color:   array[6][1],
		Number:  array[7][1],
		Friend:  array[8][1],
		Detail:  detail,
		Chinese: titname,
		Time:    titday,
	}
	MachineDate := strings.Replace(time.Now().String()[0:10], "-", "", -1)
	stodaydesc, _ := json.Marshal(todaydesc)
	//fmt.Println(string(stodaydesc))
	conf := common.NewConfig("../cmd/config.json")
	db := db.NewSqlClient(conf)
	hehe := db.GetAstroFortune(MachineDate, starname)
	if hehe == nil {
		log.Println("not exist")
		db.NewAstroFortune(MachineDate, starname)
	}

	db.SetAstroFortuneToday(MachineDate, starname, string(stodaydesc))

	//fmt.Println(hehe.Astro)
}

func CallWeek(starname string) {
	var url = "http://astro.sina.com.cn/fate_week_" + starname
	var array [20][2]string
	doc2, _ := goquery.NewDocument(url)
	fmt.Println(starname)

	cnt := 0
	doc2.Find(".content .subtit").Each(func(i int, contentSelection *goquery.Selection) {
		title := contentSelection.Find(".sp3").Text()
		title2 := contentSelection.Find(".sp4").Text()
		//fmt.Println("xxoo", i+1, "title is :", title)
		//fmt.Println("xxoo", i+1, "title is :", title2)
		array[cnt][0] = title
		if title2 != "" {
			array[cnt][1] = title2
		}
		cnt = cnt + 1
	})
	cnt2 := 0
	doc2.Find(".content .words").Each(func(i int, contentSelection *goquery.Selection) {
		title := contentSelection.Find("p").Text()
		if title != "" {
			detail, _ := contentSelection.Html()
			array[cnt2][1] = detail
		}
		cnt2 = cnt2 + 1
		//		fmt.Println("xxoo", i+1, "title is :", title)
	})

	array[8][1] = doc2.Find(".time").Text()
	/*
		for i := 0; i < cnt; i++ {
			fmt.Println(array[i][0], ":a", array[i][1])
		}
	*/

	weekdesc := db.AstroWeekDesc{
		Astro:    starname,
		General:  array[0][1],
		Love:     array[1][1],
		Work:     array[2][1],
		RedAlarm: array[3][1],
		Lucky:    array[4][1],
		Position: array[5][1],
		Try:      array[6][1],
		Friend:   array[7][1],
		Range:    array[8][1],
	}
	MachineDate := strings.Replace(time.Now().String()[0:10], "-", "", -1)
	sweekdesc, _ := json.Marshal(weekdesc)
	//fmt.Println(string(sweekdesc))
	conf := common.NewConfig("../cmd/config.json")
	db := db.NewSqlClient(conf)
	hehe := db.GetAstroFortune(MachineDate, starname)
	if hehe == nil {
		log.Println("not exist")
		db.NewAstroFortune(MachineDate, starname)
	}

	db.SetAstroFortuneWeek(MachineDate, starname, string(sweekdesc))

}

func main() {
	log.Printf("xoo")
	//resp, err := http.Get("http://astro.sina.com.cn/fate_week_Capricorn/")
	//defer resp.Body.Close()

	//if err != nil {
	//	fmt.Println("error: ", err)
	//} else {
	//	b, _ := ioutil.ReadAll(resp.Body)
	//	fmt.Println(string(b))
	//}

	//	var url = "http://astro.sina.com.cn/fate_week_Capricorn/"
	//	p, err := goquery.ParseUrl(url)
	//	if err != nil {
	//		panic(err)
	//	} else {
	//		pTitle := p.Find("content clearfix").Text()
	//		fmt.Println(pTitle)
	//		fmt.Println(p.Find("div").HasClass("clearfix"))
	//
	//		sb := p.Find(".clearfix")
	//		for i := 0; i < sb.Length(); i++ {
	//			d := sb.Eq(i)
	//			fmt.Println(d)
	//		}
	//		//fmt.Println(sb)
	//
	//		bs := p.Find(".clearfix")
	//		//ee := bs.Find(".words")
	//		//		fmt.Println("gogo")
	//		//		fmt.Println(bs)
	//		//		fmt.Println("gogo")
	//		//		fmt.Println(bs.Html())
	//		//		fmt.Println("gogo2")
	//		//		fmt.Println(ee.Html())
	//		//		fmt.Println("gogo3")
	//
	//		for i := 0; i < bs.Length(); i++ {
	//			fmt.Println(bs.Eq(i).Html())
	//			//vv := bs.Eq(i).Find(".subtit")
	//			//fmt.Println(vv.Html())
	//			fmt.Println("gogogo1")
	//		}
	//		wenzi := p.Find(".clearfix.subtit")
	//		fmt.Println(wenzi.Text())
	//	}
	CallWeek("Aries")
	CallWeek("Taurus")
	CallWeek("Gemini")
	CallWeek("Cancer")
	CallWeek("leo")
	CallWeek("Virgo")
	CallWeek("Libra")
	CallWeek("Scorpio")
	CallWeek("Sagittarius")
	CallWeek("Capricorn")
	CallWeek("Aquarius")
	CallWeek("Pisces")
	CallToday("Aries")
	CallToday("Taurus")
	CallToday("Gemini")
	CallToday("Cancer")
	CallToday("leo")
	CallToday("Virgo")
	CallToday("Libra")
	CallToday("Scorpio")
	CallToday("Sagittarius")
	CallToday("Capricorn")
	CallToday("Aquarius")
	CallToday("Pisces")
}
