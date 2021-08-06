package crawler

import (
	"fmt"
	"gabor-boros/travaux-crawler/internal/pkg/ocim"
	"gabor-boros/travaux-crawler/internal/pkg/travaux"
	"github.com/gocolly/colly/v2"
	"net/url"
	"strings"
)

type CrawlResult map[ocim.AppServer][]string

// Crawl the given URL looking for the provided app server IDs.
func Crawl(u *url.URL, status travaux.TaskStatus, servers []ocim.AppServer, all bool, verbose bool) (CrawlResult, error) {
	results := make(CrawlResult)

	crawler := colly.NewCollector(
		colly.Async(true),
	)

	if err := crawler.Limit(&colly.LimitRule{DomainGlob: "*", Parallelism: 5}); err != nil {
		return results, err
	}

	crawler.OnHTML("table#tasklist_table tbody tr", func(element *colly.HTMLElement) {
		taskStatus := strings.ReplaceAll(element.ChildText("td.task_status"), "\u00A0", " ")

		if taskStatus == string(status) {
			visitErr := element.Request.Visit(element.ChildAttr("td.task_id a", "href"))
			if visitErr != nil {
				panic(visitErr)
			}
		}
	})

	crawler.OnHTML("div#taskdetails", func(element *colly.HTMLElement) {
		incidentName := element.ChildText("h2")

		for _, appServer := range servers {
			appServerID := appServer.Server.OpenstackId.String()
			if strings.Contains(element.ChildText("div#taskdetailsfull"), appServerID) {
				results[appServer] = append(results[appServer], incidentName)
			}
		}
	})

	if all {
		crawler.OnHTML("table#pagenumbers td#numbers a", func(element *colly.HTMLElement) {
			if element.Text == "Next >" {
				err := element.Request.Visit(element.Attr("href"))
				if err != nil {
					panic(err)
				}
			}
		})
	}

	if verbose {
		crawler.OnRequest(func(r *colly.Request) {
			fmt.Println("Visiting", r.URL)
		})
	}

	err := crawler.Visit(u.String())
	crawler.Wait()

	return results, err
}
