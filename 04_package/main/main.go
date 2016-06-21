package main
 import (
	 "net/http"
	 "io/ioutil"
	 "fmt"
 )

func main() {
	res, _ := http.Get("http://www.bbc.co.uk/")
	page, _ := 	 ioutil.ReadAll(res.Body)
	res.Body.Close()
	fmt.Printf("%s", page)

}