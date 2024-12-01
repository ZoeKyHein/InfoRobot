package browser

import (
	"context"
	"fmt"
	"github.com/go-rod/rod"
	"github.com/go-rod/rod/lib/launcher"
)

import (
	"github.com/go-rod/rod/lib/utils"
)

// 只允许task.go 初始化一次
type Lxrod struct {
	Browser               *rod.Browser
	Launcher              *launcher.Launcher
	NoByPass              bool //添加反反爬虫代码 原理就是执行一段js去掉window属性中关于自动化的代码
	Headless              bool //headless 默认显示浏览器 为true是不显示浏览器
	MonitorDialog         bool //默认监控alert然后去点击true 如果要自己处理 设置为true不添加监控
	Leakless              bool
	NoDefaultDevice       bool
	Proxy                 string
	ProxyAuth             string
	WsURL                 string // rod的ws地址
	NoEachEventForCreated bool
}

// 关闭浏览器 统一控制
func (lr *Lxrod) CloseBrowser() (err error) {
	if lr.Browser != nil {
		return rod.Try(func() {
			if lr.Browser != nil {
				lr.Browser.Close()
				go rod.Try(func() {
					lr.Launcher.Cleanup()
					lr.Launcher.Kill()
				})
			}
		})
	}
	return nil
}

func (lxrod *Lxrod) NewLowBrowser(ctx context.Context) (b *rod.Browser, p *rod.Page, err error) {
	rod.Try(func() {
		lxrod.CloseBrowser()
	})
	l := launcher.New().
		Headless(false).
		Set("high-dpi-support", "1").
		Set("disable-features", "CalculateNativeWinOcclusion").
		Set("force-device-scale-factor", "1")
	l.Flags["disable-blink-features"] = []string{"AutomationControlled"}
	lxrod.Launcher = l
	//lxrod.Launcher.UserDataDir("user-data")
	//if lxrod.Area == "hebei" {
	//	lxrod.Launcher.UserDataDir("user-data")
	//}

	wsURL, err := l.Launch()
	lxrod.WsURL = wsURL
	if err != nil {
		return nil, nil, fmt.Errorf("launch browser error: %v", err)
	}
	b = rod.New().ControlURL(wsURL).Context(ctx).MustConnect()
	utils.Sleep(1)
	b = b.MustSetCookies()
	p = b.MustPage().MustSetWindow(0, 0, 1600, 900).MustSetViewport(1600, 900, 1, false)
	lxrod.Browser = b
	return b, p, nil
}
