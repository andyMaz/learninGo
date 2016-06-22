package main
 import (
	 "net/http"
	 "io/ioutil"
	 "fmt"
 )

func main() {
	res, _ := http.Get("http://www.rightmove.co.uk/property-for-sale/property-54683818.html")
	page, _ := 	 ioutil.ReadAll(res.Body)
	res.Body.Close()
	fmt.Printf("%s", page)

}