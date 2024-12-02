package crawler

import (
	"InfoRobot/browser"
	"InfoRobot/models"
	"context"
	"github.com/go-rod/rod"
	"log"
)

type TianjinCrawler struct{}

func NewTianjinCrawler() *TianjinCrawler {
	return &TianjinCrawler{}
}

func (b *TianjinCrawler) FetchData() (models.Data, error) {
	lxrod := browser.Lxrod{}
	_, p, err := lxrod.NewLowBrowser(context.Background())
	if err != nil {
		log.Fatalf("浏览器初始化失败: %v", err)
	}
	defer lxrod.CloseBrowser()

	msg := make([]models.MsgSt, 0)
	rod.Try(func() {
		p.MustNavigate(browser.TianjinInfoUrl)
		p.MustWaitStable()
		table := p.MustElement(`div#tzgg>table>tbody>tr>td>div>table>tbody`)
		for i, element := range table.MustElements("tr") {
			if i == 0 {
				continue
			}
			if i > 5 {
				break
			}
			msg = append(msg, models.MsgSt{
				Title: *element.MustElement("td:nth-child(2)>a").MustAttribute("title"),
				Date:  element.MustElement("td:nth-child(3)").MustText(),
				Url:   browser.TianjinBaseUrl + *element.MustElement("td:nth-child(2)>a").MustAttribute("href"),
			})
		}
	})
	data := models.Data{
		Region: "天津",
		Msgs:   msg,
	}
	return data, nil
}
