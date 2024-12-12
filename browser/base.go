package browser

var Today string

const (
	BeijingBaseUrl = `beijing.chinatax.gov.cn`
	BeijingInfoUrl = `http://beijing.chinatax.gov.cn/bjswj/c104176/Information.shtml`

	TianjinBaseUrl = `tianjin.chinatax.gov.cn`
	TianjinInfoUrl = `https://tianjin.chinatax.gov.cn/sjsy_xxgk.jsp`

	ShanghaiBaseUrl = `shanghai.chinatax.gov.cn/xxgk`
	ShanghaiInfoUrl = `https://shanghai.chinatax.gov.cn/xxgk/`

	ChongqingBaseUrl     = `chongqing.chinatax.gov.cn/xxgk`
	ChongqingBaseInfoUrl = `chongqing.chinatax.gov.cn/cqtax/xxgk`
	ChongqingInfoUrl     = `http://chongqing.chinatax.gov.cn/cqtax/xxgk/`

	DalianBaseUrl = `https://dalian.chinatax.gov.cn`
	DalianInfoUrl = `https://dalian.chinatax.gov.cn/col/col1681/index.html`

	QingdaoBaseUrl = `https://qingdao.chinatax.gov.cn`
	QingdaoInfoUrl = `http://qingdao.chinatax.gov.cn/xxgk2019/`

	NingboBaseUrl = `https://ningbo.chinatax.gov.cn`
	NingboInfoUrl = `https://ningbo.chinatax.gov.cn/xxgk/index.html`

	XiamenBaseUrl = `http://xiamen.chinatax.gov.cn`
	XiamenInfoUrl = `http://xiamen.chinatax.gov.cn/xmswcms/xxgk.html`

	ShenzhenBaseUrl = `shenzhen.chinatax.gov.cn`
	ShenzhenInfoUrl = `https://shenzhen.chinatax.gov.cn/sztax/xxgk/xxgk.shtml`

	NeimengguBaseUrl = `neimenggu.chinatax.gov.cn/xxgk`
	NeimengguInfoUrl = `https://neimenggu.chinatax.gov.cn/xxgk/`

	GuangxiBaseUrl = `guangxi.chinatax.gov.cn/xxgk`
	GuangxiInfoUrl = `https://guangxi.chinatax.gov.cn/xxgk/`

	XizangBaseUrl = `https://xizang.chinatax.gov.cn`
	XizangInfoUrl = `https://xizang.chinatax.gov.cn/col/col5330/index.html`

	NingxiaBaseUrl = `http://ningxia.chinatax.gov.cn`
	NingxiaInfoUrl = `http://ningxia.chinatax.gov.cn/col/col10877/index.html`

	XinjiangBaseUrl = `https://xinjiang.chinatax.gov.cn/zwgk`
	XinjiangInfoUrl = `https://xinjiang.chinatax.gov.cn/zwgk/`

	HebeiBaseUrl = `http://hebei.chinatax.gov.cn/hbsw/xxgk`
	HebeiInfoUrl = `http://hebei.chinatax.gov.cn/hbsw/xxgk/`

	ShanxiBaseUrl = `http://shanxi.chinatax.gov.cn`
	ShanxiInfoUrl = `http://shanxi.chinatax.gov.cn/xxgk`

	LiaoningBaseUrl = `https://liaoning.chinatax.gov.cn/`
	LiaoningInfoUrl = `https://liaoning.chinatax.gov.cn/col/col6/index.html`

	JilinBaseUrl = `jilin.chinatax.gov.cn`
	JilinInfoUrl = `http://jilin.chinatax.gov.cn/col/col296/index.html`

	HeilongjiangBaseUrl = `heilongjiang.chinatax.gov.cn`
	HeilongjiangInfoUrl = `http://heilongjiang.chinatax.gov.cn/col/col7757/index.html`

	JiangsuBaseUrl = `jiangsu.chinatax.gov.cn`
	JiangsuInfoUrl = `https://jiangsu.chinatax.gov.cn/col/col8197/index.html`

	ZhejiangBaseUrl = `https://zhejiang.chinatax.gov.cn`
	ZhejiangInfoUrl = `https://zhejiang.chinatax.gov.cn/col/col13133/index.html`

	AnhuiBaseUrl = `anhui.chinatax.gov.cn`
	AnhuiInfoUrl = `https://anhui.chinatax.gov.cn/col/col5419/index.html`

	FujianBaseUrl = `http://fujian.chinatax.gov.cn`
	FujianInfoUrl = `http://fujian.chinatax.gov.cn/`

	JiangxiBaseUrl = `jiangxi.chinatax.gov.cn`
	JiangxiInfoUrl = `https://jiangxi.chinatax.gov.cn/col/col31013/index.html`

	ShandongBaseUrl = `https://shandong.chinatax.gov.cn`
	ShandongInfoUrl = `https://shandong.chinatax.gov.cn/col/col3/index.html`

	HenanBaseUrl = `https://henan.chinatax.gov.cn`
	HenanInfoUrl = `https://henan.chinatax.gov.cn/henanchinatax/xxgk/index.html`

	HubeiBaseUrl = `http://hubei.chinatax.gov.cn/`
	HubeiInfoUrl = `http://hubei.chinatax.gov.cn/hbsw/xxgk/index.html`

	HunanBaseUrl = `https://hunan.chinatax.gov.cn`
	HunanInfoUrl = `https://hunan.chinatax.gov.cn/category/20181220000678`

	GuangdongBaseUrl = `guangdong.chinatax.gov.cn`
	GuangdongInfoUrl = `https://guangdong.chinatax.gov.cn/gdsw/xxgk/xxgk.shtml`

	HainanBaseUrl = `hainan.chinatax.gov.cn`
	HainanInfoUrl = `https://hainan.chinatax.gov.cn/xxgk/`

	SichuanBaseUrl = `https://sichuan.chinatax.gov.cn`
	SichuanInfoUrl = `https://sichuan.chinatax.gov.cn/col/col278/index.html`

	GuizhouBaseUrl = `guizhou.chinatax.gov.cn`
	GuizhouInfoUrl = `https://guizhou.chinatax.gov.cn/xxgk/`

	YunnanBaseUrl = `https://yunnan.chinatax.gov.cn`
	YunnanInfoUrl = `https://yunnan.chinatax.gov.cn/col/col3829/index.html`

	ShaanxiBaseUrl = `shaanxi.chinatax.gov.cn`
	ShaanxiInfoUrl = `https://shaanxi.chinatax.gov.cn/col/col7284/index.html`

	GansuBaseUrl = `http://gansu.chinatax.gov.cn`
	GansuInfoUrl = `http://gansu.chinatax.gov.cn/col/col2543/index.html`

	QinghaiBaseUrl = `qinghai.chinatax.gov.cn`
	QinghaiInfoUrl = `http://qinghai.chinatax.gov.cn/web/xxgk/xxgk.shtml`
)
