package index

import (
	"http"
	"appengine"
	"appengine/urlfetch"
	"app"
	"xml"
)

type Note struct{
	To	string
	From	string
	Heading	string
}

type Data struct {
        XMLName xml.Name "data"
        Note []Note
}

func init() {
	http.HandleFunc("/", main)
}

func main(w http.ResponseWriter, r *http.Request) {

	c := appengine.NewContext(r)
	client := urlfetch.Client(c)

	app.Start(w,r)

	var data *Data

	resp, err := client.Get("http://student.sut.ac.th/b5111299/test.xml")
	if err != nil {
		http.Error(w, err.String(), http.StatusInternalServerError)
		return
	}

        xml.Unmarshal (resp.Body, &data)

	for _,u := range data.Note{
		str := "to="+u.To+"&from="+u.From
		app.Write(str)
		app.Exit("OK")
	}
}
