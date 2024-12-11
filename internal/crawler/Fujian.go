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

type FujianCrawler struct{}

func NewFujianCrawler() *FujianCrawler {
	return &FujianCrawler{}
}

func (b *FujianCrawler) FetchData() (models.Data, error) {
	lxrod := browser.Lxrod{}
	_, p, err := lxrod.NewLowBrowser(context.Background())
	if err != nil {
		log.Fatalf("浏览器初始化失败: %v", err)
	}
	defer lxrod.CloseBrowser()

	msg := make([]models.MsgSt, 0)
	rod.Try(func() {
		p.Timeout(20 * time.Second)
		p.MustNavigate(browser.FujianInfoUrl)
		p.MustWaitLoad()
		p.MustElementR(".nav_icon", "信息公开").MustHover()
		table := p.MustElements(`[class="arts_gg fr"]`)[2]
		for i, element := range table.MustElements("li") {
			if i > 4 {
				break
			}
			msg = append(msg, models.MsgSt{
				Title: *element.MustElement("a").MustAttribute("title"),
				Date:  element.MustElement("span").MustText(),
				Url:   browser.FujianBaseUrl + strings.Replace(*element.MustElement("a").MustAttribute("href"), ".", "", 1),
			})
		}
	})
	data := models.Data{
		Region: "Fujian",
		Msgs:   msg,
	}
	return data, nil
}
