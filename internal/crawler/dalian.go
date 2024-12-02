package crawler

import (
	"InfoRobot/browser"
	"InfoRobot/models"
	"context"
	"github.com/go-rod/rod"
	"log"
)

type DalianCrawler struct{}

func NewDalianCrawler() *DalianCrawler {
	return &DalianCrawler{}
}

func (b *DalianCrawler) FetchData() (models.Data, error) {
	lxrod := browser.Lxrod{}
	_, p, err := lxrod.NewLowBrowser(context.Background())
	if err != nil {
		log.Fatalf("浏览器初始化失败: %v", err)
	}
	defer lxrod.CloseBrowser()

	msg := make([]models.MsgSt, 0)
	rod.Try(func() {
		p.MustNavigate(browser.DalianInfoUrl)
		p.MustWaitStable()
		table := p.MustElement(`[class="notice-list"]>ul`)
		for i, element := range table.MustElements("li") {
			if i > 4 {
				break
			}
			msg = append(msg, models.MsgSt{
				Title: *element.MustElement("a").MustAttribute("title"),
				Date:  element.MustElement("span").MustText(),
				Url:   *element.MustElement("a").MustAttribute("href"),
			})
		}
	})
	data := models.Data{
		Region: "大连",
		Msgs:   msg,
	}
	return data, nil
}
