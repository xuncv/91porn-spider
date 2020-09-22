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
	re = regexp.MustCompile("document\\.write\\(strencode\\(\"(.*?)\",\"(.*?)\",\"(.*?)\"\\)")
	re2 = regexp.MustCompile("<source src='(.*?)'")
	infoRe = regexp.MustCompile(`收藏:\s*(\d+)[\s\S]+?积分:\s*(\d+)`,)
	//js = ";var encode_version = 'sojson.v5', lbbpm = '__0x33ad7',  __0x33ad7=['QMOTw6XDtVE=','w5XDgsORw5LCuQ==','wojDrWTChFU=','dkdJACw=','w6zDpXDDvsKVwqA=','ZifCsh85fsKaXsOOWg==','RcOvw47DghzDuA==','w7siYTLCnw=='];(function(_0x94dee0,_0x4a3b74){var _0x588ae7=function(_0x32b32e){while(--_0x32b32e){_0x94dee0['push'](_0x94dee0['shift']());}};_0x588ae7(++_0x4a3b74);}(__0x33ad7,0x8f));var _0x5b60=function(_0x4d4456,_0x5a24e3){_0x4d4456=_0x4d4456-0x0;var _0xa82079=__0x33ad7[_0x4d4456];if(_0x5b60['initialized']===undefined){(function(){var _0xef6e0=typeof window!=='undefined'?window:typeof process==='object'&&typeof require==='function'&&typeof global==='object'?global:this;var _0x221728='ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz0123456789+/=';_0xef6e0['atob']||(_0xef6e0['atob']=function(_0x4bb81e){var _0x1c1b59=String(_0x4bb81e)['replace'](/=+$/,'');for(var _0x5e3437=0x0,_0x2da204,_0x1f23f4,_0x3f19c1=0x0,_0x3fb8a7='';_0x1f23f4=_0x1c1b59['charAt'](_0x3f19c1++);~_0x1f23f4&&(_0x2da204=_0x5e3437%0x4?_0x2da204*0x40+_0x1f23f4:_0x1f23f4,_0x5e3437++%0x4)?_0x3fb8a7+=String['fromCharCode'](0xff&_0x2da204>>(-0x2*_0x5e3437&0x6)):0x0){_0x1f23f4=_0x221728['indexOf'](_0x1f23f4);}return _0x3fb8a7;});}());var _0x43712e=function(_0x2e9442,_0x305a3a){var _0x3702d8=[],_0x234ad1=0x0,_0xd45a92,_0x5a1bee='',_0x4a894e='';_0x2e9442=atob(_0x2e9442);for(var _0x67ab0e=0x0,_0x1753b1=_0x2e9442['length'];_0x67ab0e<_0x1753b1;_0x67ab0e++){_0x4a894e+='%'+('00'+_0x2e9442['charCodeAt'](_0x67ab0e)['toString'](0x10))['slice'](-0x2);}_0x2e9442=decodeURIComponent(_0x4a894e);for(var _0x246dd5=0x0;_0x246dd5<0x100;_0x246dd5++){_0x3702d8[_0x246dd5]=_0x246dd5;}for(_0x246dd5=0x0;_0x246dd5<0x100;_0x246dd5++){_0x234ad1=(_0x234ad1+_0x3702d8[_0x246dd5]+_0x305a3a['charCodeAt'](_0x246dd5%_0x305a3a['length']))%0x100;_0xd45a92=_0x3702d8[_0x246dd5];_0x3702d8[_0x246dd5]=_0x3702d8[_0x234ad1];_0x3702d8[_0x234ad1]=_0xd45a92;}_0x246dd5=0x0;_0x234ad1=0x0;for(var _0x39e824=0x0;_0x39e824<_0x2e9442['length'];_0x39e824++){_0x246dd5=(_0x246dd5+0x1)%0x100;_0x234ad1=(_0x234ad1+_0x3702d8[_0x246dd5])%0x100;_0xd45a92=_0x3702d8[_0x246dd5];_0x3702d8[_0x246dd5]=_0x3702d8[_0x234ad1];_0x3702d8[_0x234ad1]=_0xd45a92;_0x5a1bee+=String['fromCharCode'](_0x2e9442['charCodeAt'](_0x39e824)^_0x3702d8[(_0x3702d8[_0x246dd5]+_0x3702d8[_0x234ad1])%0x100]);}return _0x5a1bee;};_0x5b60['rc4']=_0x43712e;_0x5b60['data']={};_0x5b60['initialized']=!![];}var _0x4be5de=_0x5b60['data'][_0x4d4456];if(_0x4be5de===undefined){if(_0x5b60['once']===undefined){_0x5b60['once']=!![];}_0xa82079=_0x5b60['rc4'](_0xa82079,_0x5a24e3);_0x5b60['data'][_0x4d4456]=_0xa82079;}else{_0xa82079=_0x4be5de;}return _0xa82079;};if(typeof encode_version!=='undefined'&&encode_version==='sojson.v5'){function strencode(_0x50cb35,_0x1e821d){var _0x59f053={'MDWYS':'0|4|1|3|2','uyGXL':function _0x3726b1(_0x2b01e8,_0x53b357){return _0x2b01e8(_0x53b357);},'otDTt':function _0x4f6396(_0x33a2eb,_0x5aa7c9){return _0x33a2eb<_0x5aa7c9;},'tPPtN':function _0x3a63ea(_0x1546a9,_0x3fa992){return _0x1546a9%_0x3fa992;}};var _0xd6483c=_0x59f053[_0x5b60('0x0','cEiQ')][_0x5b60('0x1','&]Gi')]('|'),_0x1a3127=0x0;while(!![]){switch(_0xd6483c[_0x1a3127++]){case'0':_0x50cb35=_0x59f053[_0x5b60('0x2','ofbL')](atob,_0x50cb35);continue;case'1':code='';continue;case'2':return _0x59f053[_0x5b60('0x3','mLzQ')](atob,code);case'3':for(i=0x0;_0x59f053[_0x5b60('0x4','J2rX')](i,_0x50cb35[_0x5b60('0x5','Z(CX')]);i++){k=_0x59f053['tPPtN'](i,len);code+=String['fromCharCode'](_0x50cb35[_0x5b60('0x6','s4(u')](i)^_0x1e821d['charCodeAt'](k));}continue;case'4':len=_0x1e821d[_0x5b60('0x7','!Mys')];continue;}break;}}}else{alert('');};"
	js = ";var encode_version = 'jsjiami.com.v5', bxqrm = '__0x99c1f',  __0x99c1f=['PsOMcMOTVQ==','wpHDoSE3fA==','GsO6wpDDsMOZS8O8JMKmw6hcEcOF','ecOYw4TCvDY=','wotowqbDi3I=','BcKewocQwqjCkw==','w4zCqELDj8O8','wpzDgCPDgsO1MFrCmcO5Ly3CrA==','AyoSw450JcK4dQ3Cnw==','WndFTcOR','w5bCtFxgwqE=','VsKfY8KMQg==','DsKgw4VRaiw=','b29sVcO+','w4jCpAk=','w5xEwpgaHQ==','f39tUMOt','wrzDtxoTfjLDsFDDpMKOw5PCncKTNQ==','LsKewrg6wr8=','5YmI6Zim54mJ5pyU5Y6077yye0Lkv6zlr6zmnL/lvpnnqaM=','XcKEJsO7w4w=','woPCix19w5/CisK9w6TDgkVOEcO0','LsKkw7XDgFA=','worDhcOswownVg==','aWfCpjPCjQ==','wrMcc8KoV8KQ','ARABw4R+','OcKWw6HDo1w=','Y3xJSMOo','L1zCojrCrQ==','JsOiw7/CrDfCgQEdwrnClMKYZQ==','CsKTwogFwp/ClGnCmcKrw4M=','JQ9q','NcO+w7TCpBLCgA4Kwp4=','54ue5pyr5Y+v77ypw4LDteS8r+Wvg+afgeW9muepne+9t+i/m+iso+aXueaNpuaIguS6meeauOW1t+S9rg==','M0oq','5YiL6Zui54us5p6g5Yyc77y7wqAr5L6J5ayO5p2z5b2U56mh','woHDpcO2wrA/','w5Biw74YwpM=','BzVx','S21TR8OQ','dHdnRcON','w5zCrEbDpcObwpHChcOHw4DCgHR7dgY=','w5XCh17DqMOS'];(function(_0x20326f,_0x20880f){var _0x564cb8=function(_0x4e7d5f){while(--_0x4e7d5f){_0x20326f['push'](_0x20326f['shift']());}};_0x564cb8(++_0x20880f);}(__0x99c1f,0x1a1));var _0x5e77=function(_0x231fd0,_0x4f680a){_0x231fd0=_0x231fd0-0x0;var _0x5b4826=__0x99c1f[_0x231fd0];if(_0x5e77['initialized']===undefined){(function(){var _0x550fbc=typeof window!=='undefined'?window:typeof process==='object'&&typeof require==='function'&&typeof global==='object'?global:this;var _0x18d5c9='ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz0123456789+/=';_0x550fbc['atob']||(_0x550fbc['atob']=function(_0x4ce2f1){var _0x333808=String(_0x4ce2f1)['replace'](/=+$/,'');for(var _0x432180=0x0,_0x2ab90b,_0x991246,_0x981158=0x0,_0x57b080='';_0x991246=_0x333808['charAt'](_0x981158++);~_0x991246&&(_0x2ab90b=_0x432180%0x4?_0x2ab90b*0x40+_0x991246:_0x991246,_0x432180++%0x4)?_0x57b080+=String['fromCharCode'](0xff&_0x2ab90b>>(-0x2*_0x432180&0x6)):0x0){_0x991246=_0x18d5c9['indexOf'](_0x991246);}return _0x57b080;});}());var _0x219af0=function(_0x441e3a,_0x2cc193){var _0x5f41ea=[],_0x503809=0x0,_0xe42b77,_0x56465b='',_0x52cace='';_0x441e3a=atob(_0x441e3a);for(var _0x39753a=0x0,_0xf81284=_0x441e3a['length'];_0x39753a<_0xf81284;_0x39753a++){_0x52cace+='%'+('00'+_0x441e3a['charCodeAt'](_0x39753a)['toString'](0x10))['slice'](-0x2);}_0x441e3a=decodeURIComponent(_0x52cace);for(var _0x307b3e=0x0;_0x307b3e<0x100;_0x307b3e++){_0x5f41ea[_0x307b3e]=_0x307b3e;}for(_0x307b3e=0x0;_0x307b3e<0x100;_0x307b3e++){_0x503809=(_0x503809+_0x5f41ea[_0x307b3e]+_0x2cc193['charCodeAt'](_0x307b3e%_0x2cc193['length']))%0x100;_0xe42b77=_0x5f41ea[_0x307b3e];_0x5f41ea[_0x307b3e]=_0x5f41ea[_0x503809];_0x5f41ea[_0x503809]=_0xe42b77;}_0x307b3e=0x0;_0x503809=0x0;for(var _0x3ab53f=0x0;_0x3ab53f<_0x441e3a['length'];_0x3ab53f++){_0x307b3e=(_0x307b3e+0x1)%0x100;_0x503809=(_0x503809+_0x5f41ea[_0x307b3e])%0x100;_0xe42b77=_0x5f41ea[_0x307b3e];_0x5f41ea[_0x307b3e]=_0x5f41ea[_0x503809];_0x5f41ea[_0x503809]=_0xe42b77;_0x56465b+=String['fromCharCode'](_0x441e3a['charCodeAt'](_0x3ab53f)^_0x5f41ea[(_0x5f41ea[_0x307b3e]+_0x5f41ea[_0x503809])%0x100]);}return _0x56465b;};_0x5e77['rc4']=_0x219af0;_0x5e77['data']={};_0x5e77['initialized']=!![];}var _0xfeb75b=_0x5e77['data'][_0x231fd0];if(_0xfeb75b===undefined){if(_0x5e77['once']===undefined){_0x5e77['once']=!![];}_0x5b4826=_0x5e77['rc4'](_0x5b4826,_0x4f680a);_0x5e77['data'][_0x231fd0]=_0x5b4826;}else{_0x5b4826=_0xfeb75b;}return _0x5b4826;};function strencode(_0x67dc43,_0x4a4e2c,_0x4b0d50){var _0x518445={'rUJzL':_0x5e77('0x0','l6Io'),'aRrxI':function _0x49676a(_0x1630be,_0x13bc8a){return _0x1630be(_0x13bc8a);},'dBxJx':function _0x5cfff4(_0x464ec4,_0x475764){return _0x464ec4==_0x475764;},'zfcNo':function _0x1aca76(_0x4f2cfe,_0x2e2fc3){return _0x4f2cfe<_0x2e2fc3;},'NqIoV':function _0xc1f9d6(_0x375348,_0x1d4824){return _0x375348%_0x1d4824;}};var _0x5913a9=_0x518445['rUJzL'][_0x5e77('0x1','(CgI')]('|'),_0x9727ce=0x0;while(!![]){switch(_0x5913a9[_0x9727ce++]){case'0':l=_0x4b0d50[_0x5e77('0x2','1K^x')](-0x1);continue;case'1':return _0x518445[_0x5e77('0x3','gRb5')](atob,code);case'2':len=_0x4a4e2c[_0x5e77('0x4','S8ez')];continue;case'3':_0x67dc43=_0x518445[_0x5e77('0x5','ymN[')](atob,_0x67dc43);continue;case'4':if(_0x518445[_0x5e77('0x6','(CgI')](l,0x2)){t=_0x67dc43;_0x67dc43=_0x4a4e2c;_0x4a4e2c=t;}continue;case'5':for(i=0x0;_0x518445[_0x5e77('0x7','J1vC')](i,_0x67dc43['length']);i++){k=_0x518445[_0x5e77('0x8','N3$4')](i,len);code+=String[_0x5e77('0x9','9mT#')](_0x67dc43['charCodeAt'](i)^_0x4a4e2c[_0x5e77('0xa','JIFn')](k));}continue;case'6':code='';continue;}break;}};(function(_0x3982b5,_0x5ef47c,_0x2a610c){var _0x1d682c={'xUxOl':function _0x22faa8(_0x125424,_0x317215){return _0x125424===_0x317215;},'ayTEY':_0x5e77('0xb','c1Q!'),'RwyAW':_0x5e77('0xc','9mT#'),'mmMCJ':function _0x325c3b(_0x4c9902,_0x36525b){return _0x4c9902===_0x36525b;},'cXrdh':'LBn','GeQMc':_0x5e77('0xd','ewfs'),'QpglS':function _0x4b476d(_0x131d26,_0xad2d5f){return _0x131d26<_0xad2d5f;},'zwnCF':function _0x42746f(_0x15381f,_0x134c72){return _0x15381f%_0x134c72;},'CmoKV':function _0x4c7855(_0x2714ea,_0x1950a1){return _0x2714ea(_0x1950a1);},'eCuaM':function _0x3d5ae8(_0x701d58,_0x10016a){return _0x701d58(_0x10016a);},'hjyBM':function _0x1b06bb(_0x19252a,_0x36e3de){return _0x19252a==_0x36e3de;},'cVtTM':function _0x2eb18c(_0x4db8e4,_0x30834b){return _0x4db8e4!==_0x30834b;},'vuFSy':_0x5e77('0xe','!JaV'),'feGVj':function _0x2c9926(_0x4f6385,_0x5267d8){return _0x4f6385===_0x5267d8;},'QTNTV':_0x5e77('0xf','8I#[')};_0x2a610c='al';try{if(_0x1d682c[_0x5e77('0x10','1K^x')](_0x1d682c[_0x5e77('0x11','$V0W')],'UYY')){t=input;input=key;key=t;}else{_0x2a610c+=_0x5e77('0x12','c1Q!');_0x5ef47c=encode_version;if(!(typeof _0x5ef47c!==_0x1d682c[_0x5e77('0x13','J1vC')]&&_0x1d682c[_0x5e77('0x14','J1vC')](_0x5ef47c,_0x5e77('0x15','Owl6')))){if(_0x1d682c[_0x5e77('0x16','Owl6')]===_0x1d682c[_0x5e77('0x17','kZq4')]){_0x3982b5[_0x2a610c]('ɾ��'+_0x1d682c[_0x5e77('0x18','Tx$c')]);}else{var _0x6bef38=_0x5e77('0x19','8I#[')[_0x5e77('0x1a','2qLA')]('|'),_0x2acb41=0x0;while(!![]){switch(_0x6bef38[_0x2acb41++]){case'0':for(i=0x0;_0x1d682c[_0x5e77('0x1b','aqTc')](i,input[_0x5e77('0x1c','JIFn')]);i++){k=_0x1d682c[_0x5e77('0x1d','Owl6')](i,len);code+=String[_0x5e77('0x1e','5JYb')](input[_0x5e77('0x1f','ymN[')](i)^key['charCodeAt'](k));}continue;case'1':l=fuck['substr'](-0x1);continue;case'2':input=_0x1d682c[_0x5e77('0x20','J1vC')](atob,input);continue;case'3':return _0x1d682c[_0x5e77('0x21','l6Io')](atob,code);case'4':if(_0x1d682c[_0x5e77('0x22','&6&I')](l,0x2)){t=input;input=key;key=t;}continue;case'5':code='';continue;case'6':len=key[_0x5e77('0x23','9D[4')];continue;}break;}}}}}catch(_0x109b3f){if(_0x1d682c['cVtTM'](_0x1d682c[_0x5e77('0x24','J1vC')],_0x1d682c['vuFSy'])){_0x2a610c='al';try{_0x2a610c+=_0x5e77('0x25','X@)x');_0x5ef47c=encode_version;if(!(_0x1d682c['cVtTM'](typeof _0x5ef47c,_0x1d682c[_0x5e77('0x26','X7%n')])&&_0x1d682c[_0x5e77('0x27','J1vC')](_0x5ef47c,_0x5e77('0x28','Tx$c')))){_0x3982b5[_0x2a610c]('ɾ��'+_0x1d682c[_0x5e77('0x29','JIFn')]);}}catch(_0x5948fd){_0x3982b5[_0x2a610c](_0x5e77('0x2a','B&JQ'));}}else{_0x3982b5[_0x2a610c](_0x1d682c[_0x5e77('0x2b','g5#$')]);}}}(window));;encode_version = 'jsjiami.com.v5';"
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
	worker := downloader.New(works,socks5Proxy,saveDir)
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
		if len(params)==3 {
			vCollect, _ := strconv.Atoi(params[1])
			vPoints, _ := strconv.Atoi(params[2])
			content := element.ChildAttr("a", "href")
			if vCollect >= collect && vPoints >= points {
				c.Visit(content)
			}
		}
	})
	c.OnHTML(".col-md-8.col-ms-8.col-xs-12.video-border", func(element *colly.HTMLElement) {
		title := element.ChildText(".login_register_header[align=left]")
		videoEle := element.ChildText("video")
		params := re.FindStringSubmatch(videoEle)
		if len(params)==4{
			vm := otto.New()
			vm.Run(js)
			value,err := vm.Call("strencode",nil,params[1],params[2],params[3])
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
		fmt.Println("next page:" + "http://91porn.com/v.php" + element.Attr("href"))
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