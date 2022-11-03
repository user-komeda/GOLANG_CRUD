package main

import (
	"user-komeda/GOLANG_SCRAIPE/src/sub/goquery"
	"user-komeda/GOLANG_SCRAIPE/src/sub/selenium"
)

func main() {
	html, _, page, driver := selenium.Selenium()
	goquery.GoqueryForCategory(html)
	goquery.GoqueryForMainContent(html, page)
	driver.Stop()

}
