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

// Lxrod 只允许task.go 初始化一次
type Lxrod struct {
	Browser  *rod.Browser
	Launcher *launcher.Launcher
	WsURL    string // rod的ws地址
}

// CloseBrowser 关闭浏览器 统一控制
func (lr *Lxrod) CloseBrowser() (err error) {
	if lr.Browser != nil {
		return rod.Try(func() {
			if lr.Browser != nil {
				err = lr.Browser.Close()
				if err != nil {
					return
				}
				go func() {
					_ = rod.Try(func() {
						lr.Launcher.Cleanup()
						lr.Launcher.Kill()
					})
				}()
			}
		})
	}
	return nil
}

func (lr *Lxrod) NewLowBrowser(ctx context.Context) (b *rod.Browser, p *rod.Page, err error) {
	_ = rod.Try(func() {
		err = lr.CloseBrowser()
		if err != nil {
			return
		}
	})
	l := launcher.New().
		Headless(false).
		Set("high-dpi-support", "1").
		Set("disable-features", "CalculateNativeWinOcclusion").
		Set("force-device-scale-factor", "1")
	l.Flags["disable-blink-features"] = []string{"AutomationControlled"}
	lr.Launcher = l

	wsURL, err := l.Launch()
	lr.WsURL = wsURL
	if err != nil {
		return nil, nil, fmt.Errorf("launch browser error: %v", err)
	}
	b = rod.New().ControlURL(wsURL).Context(ctx).MustConnect()
	utils.Sleep(1)
	b = b.MustSetCookies()
	p = b.MustPage().MustSetWindow(0, 0, 1600, 900).MustSetViewport(1600, 900, 1, false)
	lr.Browser = b
	return b, p, nil
}
