package goquery

import (
	"fmt"
	"strings"

	"github.com/PuerkitoBio/goquery"
)

func Goquery(html string) {
	htmlContentReader := strings.NewReader(html)
	htmlContent, _ := goquery.NewDocumentFromReader(htmlContentReader)
	content_a := htmlContent.Find(".cw-nav .parent .cw-list_nav_subcategory_title ")
	content_a.Each(func(i int, s *goquery.Selection) {
		fmt.Println(s.Text())
		fmt.Println(s.Parent().Find("ul").Text())
	})
}
