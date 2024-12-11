package crawler

import (
	"InfoRobot/browser"
	"InfoRobot/models"
	"context"
	"github.com/go-rod/rod"
	"log"
	"time"
)

type XiamenCrawler struct{}

func NewXiamenCrawler() *XiamenCrawler {
	return &XiamenCrawler{}
}

func (b *XiamenCrawler) FetchData() (models.Data, error) {
	lxrod := browser.Lxrod{}
	_, p, err := lxrod.NewLowBrowser(context.Background())
	if err != nil {
		log.Fatalf("浏览器初始化失败: %v", err)
	}
	defer lxrod.CloseBrowser()

	msg := make([]models.MsgSt, 0)
	rod.Try(func() {
		p.Timeout(20 * time.Second)
		p.MustNavigate(browser.XiamenInfoUrl)
		p.MustWaitStable()
		p.Timeout(5 * time.Second).MustSearch(`[href="/xmswcms/xxgk.html"]`).MustClick()
		p.MustWaitStable()
		table := p.MustElement(`[class="xxgk_left"]`)
		for i, element := range table.MustElements("li.tab_list_item") {
			if i > 4 {
				break
			}
			msg = append(msg, models.MsgSt{
				Title: *element.MustElement("a").MustAttribute("title"),
				Date:  element.MustElement("font").MustText(),
				Url:   browser.XiamenBaseUrl + *element.MustElement("a").MustAttribute("href"),
			})
		}
	})
	data := models.Data{
		Region: "厦门",
		Msgs:   msg,
	}
	return data, nil
}
