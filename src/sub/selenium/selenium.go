package selenium

import (
	"fmt"
	"os"

	"github.com/sclevine/agouti"
)

func Selenium() (string, error) {
	driver := agouti.ChromeDriver(agouti.Browser("chrome"))
	defer driver.Stop()

	if err := driver.Start(); err != nil {
		fmt.Fprintf(os.Stderr, "%s\n", err)
		return err.Error(), err
	}

	page, err := driver.NewPage()
	if err != nil {
		fmt.Fprintf(os.Stderr, "%s\n", err)
		return err.Error(), err
	}

	// googleにアクセス
	if err := page.Navigate("https://crowdworks.jp/public/jobs?category=jobs&order=score"); err != nil {
		fmt.Fprintf(os.Stderr, "%s\n", err)
		return err.Error(), err
	}

	return page.HTML()
}
