package app

import(
	"http" 
	"io/ioutil" 
	"regexp"
)

//get source code html form url
func GetHtml(url string) string{
	var b []byte
	r, _ := http.Get(url) 

	b, _ = ioutil.ReadAll(r.Body) 
	regex:=regexp.MustCompile( "\\<[^\\>]*\\>" )
	r.Body.Close() 

	return regex.ReplaceAllString(string(b), "")
}
