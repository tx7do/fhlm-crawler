package fhlm_test

import (
	"testing"

	. "."
)

/** 彩种列表
乐利六合彩	lllhc
乐利时时彩	llssc
乐利快三	llk3
乐利1.5分彩	ll90ssc

腾讯分分彩	txffc
比特币分分彩	btcffc
吉利分分彩	jlffc
吉利时时彩	jlssc

3D	f3d
排列3/5	plw
双色球	ssq
七乐彩		qlc
北京快乐8	bjkl8
上海时时乐	ssl
重庆时时彩	cqssc
黑龙江时时彩	hljssc
天津时时彩	tjssc
新疆时时彩	xjssc
安徽快三	ahk3
江苏快三	jsk3
江西快三	jxk3
湖北快三	hbk3
甘肃快三	gsk3
河南快三	hnk3
内蒙古快三	nmgk3
北京PK10	bjpk10
上海快三	shk3
北京快三	bjk3
广西快三	gxk3
广东快乐10分	gd10f
湖南快乐10分	hun10f
重庆幸运农场	cq10f

七星彩	qxc
大乐透	dlt
江西11选5	jx11y
广东11选5	gd11y
山东11选5	sd11y
山西11选5	sx11y
上海11选5	sh11y
湖北11选5	hb11y
安徽11选5	ah11y
江苏11选5	js11y
陕西11选5	shx11y
浙江11选5	zj11y

六合彩	lhc
幸运飞艇	xyft
*/

func TestJLFFC(t *testing.T) {
	StartCrawler("jlffc")
}
