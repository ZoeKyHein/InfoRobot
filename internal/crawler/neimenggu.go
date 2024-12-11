package crawler

import (
	"InfoRobot/browser"
	"InfoRobot/models"
	"context"
	"github.com/go-rod/rod"
	"log"
	"strings"
	"time"
)

type NeimengguCrawler struct{}

func NewNeimengguCrawler() *NeimengguCrawler {
	return &NeimengguCrawler{}
}

func (b *NeimengguCrawler) FetchData() (models.Data, error) {
	lxrod := browser.Lxrod{}
	_, p, err := lxrod.NewLowBrowser(context.Background())
	if err != nil {
		log.Fatalf("浏览器初始化失败: %v", err)
	}
	defer lxrod.CloseBrowser()

	msg := make([]models.MsgSt, 0)
	rod.Try(func() {
		p.Timeout(20 * time.Second)
		p.MustNavigate(browser.NeimengguInfoUrl)
		p.MustWaitStable()
		table := p.MustElement(`[class="tzgg_list"]`)
		for i, element := range table.MustElements("li") {
			if i > 4 {
				break
			}
			msg = append(msg, models.MsgSt{
				Title: *element.MustElement("a").MustAttribute("title"),
				Date:  element.MustElement("span").MustText(),
				Url:   browser.NeimengguBaseUrl + strings.Replace(*element.MustElement("a").MustAttribute("href"), ".", "", 1),
			})
		}
	})
	data := models.Data{
		Region: "重庆",
		Msgs:   msg,
	}
	return data, nil
}
