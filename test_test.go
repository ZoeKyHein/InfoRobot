package main

import (
	"InfoRobot/internal"
	"testing"
)

func TestSendMultiRegions(t *testing.T) {
	regions := []string{
		//"北京",
		//"上海",
		//"天津",
		//"大连",
		//"重庆",
		//"青岛",
		//"宁波",
		//"厦门",
		//"深圳",
		//"内蒙古",
		//"安徽",
		//"福建",
		//"甘肃",
		//"广东",
		//"广西",
		//"贵州",
		//"海南",
		//"河北",
		//"黑龙江",
		//"河南",
		//"湖北",
		//"湖南",
		//"江苏",
		//"江西",
		//"吉林",
		"辽宁",
	}
	internal.SpiderAndSend(regions)
}
