package main

import (
	"91porn-spider/utils"
	"fmt"
	"github.com/gocolly/colly/v2"
	"github.com/robertkrimen/otto"
	"github.com/spf13/viper"
	"math/rand"
	"regexp"
	"strconv"
	"time"
)

var(
	config *viper.Viper
	socks5Proxy string
	saveDir string
	startUrl string
	works int
	points int
	collect int
	re = regexp.MustCompile("document\\.write\\(strencode\\(\"(.*?)\",\"(.*?)\"\\)")
	re2 = regexp.MustCompile("<source src='(.*?)'")
	infoRe = regexp.MustCompile(`收藏:\s*(\d+)[\s\S]+?积分:\s*(\d+)`,)
	js = ";var encode_version = 'sojson.v5', lbbpm = '__0x33ad7',  __0x33ad7=['QMOTw6XDtVE=','w5XDgsORw5LCuQ==','wojDrWTChFU=','dkdJACw=','w6zDpXDDvsKVwqA=','ZifCsh85fsKaXsOOWg==','RcOvw47DghzDuA==','w7siYTLCnw=='];(function(_0x94dee0,_0x4a3b74){var _0x588ae7=function(_0x32b32e){while(--_0x32b32e){_0x94dee0['push'](_0x94dee0['shift']());}};_0x588ae7(++_0x4a3b74);}(__0x33ad7,0x8f));var _0x5b60=function(_0x4d4456,_0x5a24e3){_0x4d4456=_0x4d4456-0x0;var _0xa82079=__0x33ad7[_0x4d4456];if(_0x5b60['initialized']===undefined){(function(){var _0xef6e0=typeof window!=='undefined'?window:typeof process==='object'&&typeof require==='function'&&typeof global==='object'?global:this;var _0x221728='ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz0123456789+/=';_0xef6e0['atob']||(_0xef6e0['atob']=function(_0x4bb81e){var _0x1c1b59=String(_0x4bb81e)['replace'](/=+$/,'');for(var _0x5e3437=0x0,_0x2da204,_0x1f23f4,_0x3f19c1=0x0,_0x3fb8a7='';_0x1f23f4=_0x1c1b59['charAt'](_0x3f19c1++);~_0x1f23f4&&(_0x2da204=_0x5e3437%0x4?_0x2da204*0x40+_0x1f23f4:_0x1f23f4,_0x5e3437++%0x4)?_0x3fb8a7+=String['fromCharCode'](0xff&_0x2da204>>(-0x2*_0x5e3437&0x6)):0x0){_0x1f23f4=_0x221728['indexOf'](_0x1f23f4);}return _0x3fb8a7;});}());var _0x43712e=function(_0x2e9442,_0x305a3a){var _0x3702d8=[],_0x234ad1=0x0,_0xd45a92,_0x5a1bee='',_0x4a894e='';_0x2e9442=atob(_0x2e9442);for(var _0x67ab0e=0x0,_0x1753b1=_0x2e9442['length'];_0x67ab0e<_0x1753b1;_0x67ab0e++){_0x4a894e+='%'+('00'+_0x2e9442['charCodeAt'](_0x67ab0e)['toString'](0x10))['slice'](-0x2);}_0x2e9442=decodeURIComponent(_0x4a894e);for(var _0x246dd5=0x0;_0x246dd5<0x100;_0x246dd5++){_0x3702d8[_0x246dd5]=_0x246dd5;}for(_0x246dd5=0x0;_0x246dd5<0x100;_0x246dd5++){_0x234ad1=(_0x234ad1+_0x3702d8[_0x246dd5]+_0x305a3a['charCodeAt'](_0x246dd5%_0x305a3a['length']))%0x100;_0xd45a92=_0x3702d8[_0x246dd5];_0x3702d8[_0x246dd5]=_0x3702d8[_0x234ad1];_0x3702d8[_0x234ad1]=_0xd45a92;}_0x246dd5=0x0;_0x234ad1=0x0;for(var _0x39e824=0x0;_0x39e824<_0x2e9442['length'];_0x39e824++){_0x246dd5=(_0x246dd5+0x1)%0x100;_0x234ad1=(_0x234ad1+_0x3702d8[_0x246dd5])%0x100;_0xd45a92=_0x3702d8[_0x246dd5];_0x3702d8[_0x246dd5]=_0x3702d8[_0x234ad1];_0x3702d8[_0x234ad1]=_0xd45a92;_0x5a1bee+=String['fromCharCode'](_0x2e9442['charCodeAt'](_0x39e824)^_0x3702d8[(_0x3702d8[_0x246dd5]+_0x3702d8[_0x234ad1])%0x100]);}return _0x5a1bee;};_0x5b60['rc4']=_0x43712e;_0x5b60['data']={};_0x5b60['initialized']=!![];}var _0x4be5de=_0x5b60['data'][_0x4d4456];if(_0x4be5de===undefined){if(_0x5b60['once']===undefined){_0x5b60['once']=!![];}_0xa82079=_0x5b60['rc4'](_0xa82079,_0x5a24e3);_0x5b60['data'][_0x4d4456]=_0xa82079;}else{_0xa82079=_0x4be5de;}return _0xa82079;};if(typeof encode_version!=='undefined'&&encode_version==='sojson.v5'){function strencode(_0x50cb35,_0x1e821d){var _0x59f053={'MDWYS':'0|4|1|3|2','uyGXL':function _0x3726b1(_0x2b01e8,_0x53b357){return _0x2b01e8(_0x53b357);},'otDTt':function _0x4f6396(_0x33a2eb,_0x5aa7c9){return _0x33a2eb<_0x5aa7c9;},'tPPtN':function _0x3a63ea(_0x1546a9,_0x3fa992){return _0x1546a9%_0x3fa992;}};var _0xd6483c=_0x59f053[_0x5b60('0x0','cEiQ')][_0x5b60('0x1','&]Gi')]('|'),_0x1a3127=0x0;while(!![]){switch(_0xd6483c[_0x1a3127++]){case'0':_0x50cb35=_0x59f053[_0x5b60('0x2','ofbL')](atob,_0x50cb35);continue;case'1':code='';continue;case'2':return _0x59f053[_0x5b60('0x3','mLzQ')](atob,code);case'3':for(i=0x0;_0x59f053[_0x5b60('0x4','J2rX')](i,_0x50cb35[_0x5b60('0x5','Z(CX')]);i++){k=_0x59f053['tPPtN'](i,len);code+=String['fromCharCode'](_0x50cb35[_0x5b60('0x6','s4(u')](i)^_0x1e821d['charCodeAt'](k));}continue;case'4':len=_0x1e821d[_0x5b60('0x7','!Mys')];continue;}break;}}}else{alert('');};"
)
func init()  {
	rand.Seed(time.Now().Unix())
	config = viper.New()
	config.SetConfigName("config")
	config.SetConfigType("yaml")
	config.AddConfigPath(".")
	if err :=config.ReadInConfig();err!=nil{
		panic("can not load config")
	}
	saveDir = config.GetString("saveDir")
	socks5Proxy = config.GetString("socks5Proxy")
	works = config.GetInt("workers")
	points = config.GetInt("points")
	collect = config.GetInt("collect")
}

/// 区域随机整型数字
func random_int(min, max int) int {
	randNum := rand.Intn(max-min) + min
	return randNum
}
/// 生成随机ip
func random_ip() string {
	return fmt.Sprintf("%d.%d.%d.%d",
		random_int(1, 255), random_int(1, 255), random_int(1, 255), random_int(1, 255))
}

func main() {
	worker := downloader.New(works,socks5Proxy,"download")
	c := colly.NewCollector(
			colly.AllowedDomains("91porn.com","www.91porn.com"),
			//colly.Async(true),
		)
	c.UserAgent = "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/83.0.4103.116 Safari/537.36"
	c.MaxDepth = 2
	if socks5Proxy != ""{
		c.SetProxy("socks5://" + socks5Proxy)
	}
	c.OnRequest(func(request *colly.Request) {
		request.Headers.Set("accept-language","zh-CN,zh;q=0.9,en-US;q=0.8,en;q=0.7")
		request.Headers.Set("X-Forwarded-For",random_ip())
	})
	c.OnHTML(".videos-text-align", func(element *colly.HTMLElement) {
		//title:=element.ChildText(".video-title")
		//fmt.Println(title)
		params := infoRe.FindStringSubmatch(element.Text)
		vCollect,_ := strconv.Atoi(params[1])
		vPoints,_ := strconv.Atoi(params[2])
		content := element.ChildAttr("a","href")
		if vCollect >= collect && vPoints >= points {
			c.Visit(content)
		}
	})
	c.OnHTML(".col-md-8.col-ms-8.col-xs-12.video-border", func(element *colly.HTMLElement) {
		title := element.ChildText(".login_register_header[align=left]")
		videoEle := element.ChildText("video")
		params := re.FindStringSubmatch(videoEle)
		if len(params)==3{
			vm := otto.New()
			vm.Run(js)
			value,err := vm.Call("strencode",nil,params[1],params[2])
			if err==nil{
				params = re2.FindStringSubmatch(value.String())
				worker.Push(title,params[1])
			}
		}else{
			src := element.ChildAttr("video source","src")
			if len(src)>0{
				worker.Push(title,src)
			}
		}
	})

	c.OnHTML("span[class=pagingnav] + a[href]", func(element *colly.HTMLElement) {
		c.Visit("http://91porn.com/v.php" + element.Attr("href"))
	})

	c.OnError(func(response *colly.Response, err error) {
		fmt.Println(err)
	})
	fmt.Println( config.GetString("startUrl") )
	c.Visit( config.GetString("startUrl") )
	//c.Visit("http://91porn.com/view_video.php?viewkey=5f6e46d818f646ed5d4c&page=1&viewtype=basic&category=hot")
	c.Wait()
	worker.Wait()
}