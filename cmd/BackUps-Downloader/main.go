package backupsdownloader

import (
	"fmt"
	"strings"

	"github.com/JorgeSDiaz/BackUps-Downloader/internal/scrapper"
)

func Run() {
	baseURL := ""

	directoryCSSSelector := "tr td.link a[href]"
	filesDirectory := scrapper.ScrapDataWithAttribute(baseURL, directoryCSSSelector, "href")
	for _, pathParam := range filesDirectory {
		if !strings.HasPrefix(pathParam, ".") {
			url := fmt.Sprintf("%s/%s", baseURL, pathParam)
			fmt.Println(scrapper.ScrapDataWithAttribute(url, directoryCSSSelector, "href"))
		}
	}
}
