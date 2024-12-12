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

type GuangxiCrawler struct{}

func NewGuangxiCrawler() *GuangxiCrawler {
	return &GuangxiCrawler{}
}

func (b *GuangxiCrawler) FetchData() (models.Data, error) {
	lxrod := browser.Lxrod{}
	_, p, err := lxrod.NewLowBrowser(context.Background())
	if err != nil {
		log.Fatalf("浏览器初始化失败: %v", err)
	}
	defer lxrod.CloseBrowser()

	msg := make([]models.MsgSt, 0)
	rod.Try(func() {
		p.Timeout(20 * time.Second)
		p.MustNavigate(browser.GuangxiInfoUrl)
		p.MustWaitStable()
		table := p.MustElements(`[class="xxgk_list"]`)[0]
		for i, element := range table.MustElements("li") {
			if i > 4 {
				break
			}
			msg = append(msg, models.MsgSt{
				Title: *element.MustElement("a").MustAttribute("title"),
				Date:  element.MustElement("span").MustText(),
				Url:   browser.GuangxiBaseUrl + strings.Replace(*element.MustElement("a").MustAttribute("href"), ".", "", 1),
			})
		}
	})
	data := models.Data{
		Region: "Guangxi",
		Msgs:   msg,
	}
	return data, nil
}
