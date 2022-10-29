package selenium

import (
	"fmt"
	"os"

	"github.com/sclevine/agouti"
)

func Selenium() (string, error, *agouti.Page, *agouti.WebDriver) {
	const url string = "/public/jobs?category=jobs&order=score"
	driver := agouti.ChromeDriver(agouti.Browser("chrome"))
	// defer driver.Stop()

	if err := driver.Start(); err != nil {
		fmt.Fprintf(os.Stderr, "%s\n", err)
		return err.Error(), err, nil, nil
	}

	page, err := driver.NewPage()
	if err != nil {
		fmt.Fprintf(os.Stderr, "%s\n", err)
		return err.Error(), err, nil, nil
	}
	html, err := NavigatePage(url, page)
	return html, err, page, driver
}

func NavigatePage(url string, page *agouti.Page) (string, error) {
	// googleにアクセス
	if err := page.Navigate("https://crowdworks.jp" + url); err != nil {
		fmt.Fprintf(os.Stderr, "%s\n", err)
		return err.Error(), err
	}

	return page.HTML()

}
