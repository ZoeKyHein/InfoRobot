package crawler

import (
	"fmt"
)

func GetCrawler(city string) (Crawler, error) {
	switch city {
	case "北京":
		return NewBeijingCrawler(), nil
	case "天津":
		return NewTianjinCrawler(), nil
	case "上海":
		return NewShanghaiCrawler(), nil
	case "重庆":
		return NewChongqingCrawler(), nil
	case "大连":
		return NewDalianCrawler(), nil
	case "青岛":
		return NewQingdaoCrawler(), nil
	case "宁波":
		return NewNingboCrawler(), nil
	case "厦门":
		return NewXiamenCrawler(), nil
	case "深圳":
		return NewShenzhenCrawler(), nil
	case "内蒙古":
		return NewNeimengguCrawler(), nil
	//case "广西":
	//	return NewGuangxiCrawler(), nil
	//case "西藏":
	//	return NewXizangCrawler(), nil
	//case "宁夏":
	//	return NewNingxiaCrawler(), nil
	//case "新疆":
	//	return NewXinjiangCrawler(), nil
	//case "河北":
	//	return NewHebeiCrawler(), nil
	//case "山西":
	//	return NewShanxiCrawler(), nil
	//case "辽宁":
	//	return NewLiaoningCrawler(), nil
	//case "吉林":
	//	return NewJilinCrawler(), nil
	//case "黑龙江":
	//	return NewHeilongjiangCrawler(), nil
	//case "江苏":
	//	return NewJiangsuCrawler(), nil
	//case "浙江":
	//	return NewZhejiangCrawler(), nil
	//case "安徽":
	//	return NewAnhuiCrawler(), nil
	//case "福建":
	//	return NewFujianCrawler(), nil
	//case "江西":
	//	return NewJiangxiCrawler(), nil
	//case "山东":
	//	return NewShandongCrawler(), nil
	//case "河南":
	//	return NewHenanCrawler(), nil
	//case "湖北":
	//	return NewHubeiCrawler(), nil
	//case "湖南":
	//	return NewHunanCrawler(), nil
	//case "广东":
	//	return NewGuangdongCrawler(), nil
	//case "海南":
	//	return NewHainanCrawler(), nil
	//case "四川":
	//	return NewSichuanCrawler(), nil
	//case "贵州":
	//	return NewGuizhouCrawler(), nil
	//case "云南":
	//	return NewYunnanCrawler(), nil
	//case "陕西":
	//	return NewShaanxiCrawler(), nil
	//case "甘肃":
	//	return NewGansuCrawler(), nil
	//case "青海":
	//	return NewQinghaiCrawler(), nil
	default:
		return nil, fmt.Errorf("unsupported city: %s", city)
	}
}
