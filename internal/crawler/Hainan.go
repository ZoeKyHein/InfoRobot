package crawler

import (
	"InfoRobot/browser"
	"InfoRobot/models"
	"context"
	"github.com/go-rod/rod"
	"log"
	"time"
)

type HainanCrawler struct{}

func NewHainanCrawler() *HainanCrawler {
	return &HainanCrawler{}
}

func (b *HainanCrawler) FetchData() (models.Data, error) {
	lxrod := browser.Lxrod{}
	_, p, err := lxrod.NewLowBrowser(context.Background())
	if err != nil {
		log.Fatalf("浏览器初始化失败: %v", err)
	}
	defer lxrod.CloseBrowser()

	msg := make([]models.MsgSt, 0)
	rod.Try(func() {
		p.Timeout(20 * time.Second)
		p.MustNavigate(browser.HainanInfoUrl)
		p.MustWaitStable()
		table := p.MustElement(`#cona_TBABa_1`)
		for i, element := range table.MustElements("li") {
			if i > 4 {
				break
			}
			msg = append(msg, models.MsgSt{
				Title: *element.MustElement("a").MustAttribute("title"),
				Date:  element.MustElement("em").MustText(),
				Url:   browser.HainanBaseUrl + *element.MustElement("a").MustAttribute("href"),
			})
		}
	})
	data := models.Data{
		Region: "Hainan",
		Msgs:   msg,
	}
	return data, nil
}
