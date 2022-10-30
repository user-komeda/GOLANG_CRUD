package goquery

import (
	"fmt"
	"strconv"
	"strings"
	"time"
	"user-komeda/GOLANG_SCRAIPE/src/sub/selenium"

	"github.com/PuerkitoBio/goquery"
	"github.com/sclevine/agouti"
)

func GoqueryForCategory(html string) {
	htmlContent := goqueryBase(html)
	contentA := htmlContent.Find(".cw-nav .parent .cw-list_nav_subcategory_title ")
	contentA.Each(func(i int, s *goquery.Selection) {
		fmt.Println(s.Text())
		fmt.Println(s.Parent().Find("ul").Text())
	})
}

func GoqueryForMainContent(html string, page *agouti.Page) {
	urlList := getCategoryURL(html)
	for _, url := range urlList {
		html, _ := selenium.NavigatePage(url+"?hide_expired=true", page)
		for i := 1; i < getPageSize(html); i++ {
			html, _ := selenium.NavigatePage(url+"?hide_expired=true&page="+strconv.Itoa(i), page)
			htmlContent := goqueryBase(html)
			title := htmlContent.Find(".item_title")
			title.Each(func(i int, s *goquery.Selection) {
				fmt.Println(s.Text())
				fmt.Println("-------")
				// time.Sleep(time.Second * 1)
			})
			fmt.Println("=======")
			time.Sleep(time.Second * 2)
		}
	}
}

func goqueryBase(html string) *goquery.Document {

	htmlContentReader := strings.NewReader(html)
	htmlContent, _ := goquery.NewDocumentFromReader(htmlContentReader)

	return htmlContent
}

func getCategoryURL(html string) []string {
	htmlContent := goqueryBase(html)
	categoryURL := htmlContent.Find(".cw-nav .parent .cw-list_nav_subcategory_title").Children()
	urlList := []string{}
	categoryURL.Each(func(i int, s *goquery.Selection) {
		urlList = append(urlList, s.AttrOr("href", ""))
	})
	return urlList
}

func getPageSize(html string) int {
	htmlContent := goqueryBase(html)
	resultContent := htmlContent.Find(".result_count").Text()
	pageSize, _ := strconv.Atoi(resultContent[:strings.Index(resultContent, "件")])
	return pageSize/50 + 1
}
