package main

import (
	"html/template"
	"io/ioutil"
	"log"
	"os"
)

const tmpl = `<html>
<head>
<title>{{.Title}}</title>
</head>
<body>
{{.Body | unescaped}}
</body>
</html>`

type tem struct {
	Body  string
	Title string
}

func unescaped(x string) interface{} { return template.HTML(x) }

func ParseTemplateToStr(tname string) string {
	b, err := ioutil.ReadFile(tname)
	if err != nil {
		log.Println(err)
	}
	s := string(b)
	return s
}
func main() {
	s := ParseTemplateToStr("a.html")
	t := template.Must(template.New("ex").Funcs(template.FuncMap{"unescaped": unescaped}).Parse(s))
	//t, _ = t.ParseFiles("a.html")
	//t, _ = template.ParseFiles("a.html")
	//t = t.Funcs(template.FuncMap{"unescaped": unescaped})
	var tt tem
	tt.Body = "Test <b>World</b>"
	tt.Title = "Hello <b>World</b>"
	/*
		v := map[string]interface{}{
			"Title": "Test <b>World</b>",
			"Body":  template.HTML("Hello <b>World</b>"),
		}
	*/
	t.Execute(os.Stdout, tt)
}
