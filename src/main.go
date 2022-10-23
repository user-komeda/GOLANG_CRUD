package main

import (
	"user-komeda/GOLANG_SCRAIPE/src/sub/goquery"
	"user-komeda/GOLANG_SCRAIPE/src/sub/selenium"
)

func main() {
	html, _ := selenium.Selenium()
	goquery.Goquery(html)
}
