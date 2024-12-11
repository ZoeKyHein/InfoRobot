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

type ChongqingCrawler struct{}

func NewChongqingCrawler() *ChongqingCrawler {
	return &ChongqingCrawler{}
}

func (b *ChongqingCrawler) FetchData() (models.Data, error) {
	lxrod := browser.Lxrod{}
	_, p, err := lxrod.NewLowBrowser(context.Background())
	if err != nil {
		log.Fatalf("浏览器初始化失败: %v", err)
	}
	defer lxrod.CloseBrowser()

	msg := make([]models.MsgSt, 0)
	rod.Try(func() {
		p.Timeout(20 * time.Second)
		p.MustNavigate(browser.ChongqingInfoUrl)
		p.MustWaitStable()
		table := p.MustElement(`#xxgk_1`)
		for i, element := range table.MustElements("li") {
			if i > 4 {
				break
			}
			msg = append(msg, models.MsgSt{
				Title: element.MustElement("dl>dd>a").MustText(),
				Date:  element.MustElement("dl>dt").MustText(),
				Url:   browser.ChongqingBaseInfoUrl + strings.Replace(*element.MustElement("dl>dd>a").MustAttribute("href"), ".", "", 1),
			})
		}
	})
	data := models.Data{
		Region: "重庆",
		Msgs:   msg,
	}
	return data, nil
}
