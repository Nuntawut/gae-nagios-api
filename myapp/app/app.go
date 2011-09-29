package app

import(
	"os"
	"io"
	"log"
	"fmt"
	"http"
	"template"
	//"appengine"
	//"appengine/urlfetch"
	//"appengine/user"
)

var (
	mainTemplate        *template.Template
	mainTemplateErr     os.Error
	w		    http.ResponseWriter
	r		    *http.Request
)

func Start(resp http.ResponseWriter,req *http.Request){
	w = resp
	r = req
	//c := appengine.NewContext(r)
	//client := urlfetch.Client(c)
}

func ParseTemp(file string){
	mainTemplate = template.New(nil)
	mainTemplate.SetDelims("{{", "}}")
	if err := mainTemplate.ParseFile(file); err != nil {
		mainTemplateErr = fmt.Errorf("tmpl.ParseFile failed: %v", err)
		return
	}
}

func ExecTemp(views  map[string]string){
	err := mainTemplate.Execute(w, views)
	if err != nil {
		log.Println(err)
	}
}

func GetHtml(url string) io.Reader{
	r, err := http.Get(url)
	if err != nil {
		fmt.Println("Error opening file:", err)
	}
	defer r.Body.Close()

	return r.Body
}

func GetValue(key string) string {
	return r.FormValue(key)
}

func Write(txt string) {
	fmt.Fprint(w,txt)
}

func Exit(status string) {
	switch  status {
		case "OK"	: Write("&status=OK")
		case "WARNING"	: Write("&status=WARNING")
		case "CRITICAL"	: Write("&status=CRITICAL")
		case "UNKNOW"	: Write("&status=UNKNOW")
	}
}
