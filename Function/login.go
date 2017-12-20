package xiaomi

import (
	"crypto/tls"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"mi/Data"
	"net/http"
	"net/http/cookiejar"
	"net/url"
	"os"
	"strings"
)

//日志初始化
func LogInit() {
	var logFile *os.File
	var err error
	// 定义一个文件
	fileName := smartisanData.UserName + ".log"
	if checkFileIsExist(fileName) { //如果文件存在
		logFile, err = os.OpenFile(fileName, os.O_APPEND, 0666) //打开文件
	} else {
		logFile, err = os.Create(fileName) //创建文件
	}

	if err != nil {
		log.Fatalln("open file error !")
	}
	// 创建一个日志对象
	smartisanData.Log = log.New(logFile, "[Log]", log.LstdFlags)
	//配置log的Flag参数
	smartisanData.Log.SetFlags(smartisanData.Log.Flags() | log.LstdFlags)

}

func LogStr(a ...interface{}) {
	fmt.Println(a)
	smartisanData.Log.Println(a)
}

/**
 * 判断文件是否存在  存在返回 true 不存在返回false
 */
func checkFileIsExist(filename string) bool {
	var exist = true
	if _, err := os.Stat(filename); os.IsNotExist(err) {
		exist = false
	}
	return exist
}

func Draw() {
	v := url.Values{}
	v.Set("batchId", "3555")
	v.Set("subkey", "1502971200")
	body := ioutil.NopCloser(strings.NewReader(v.Encode())) //把form数据编下码
	client := &http.Client{Transport: smartisanData.Tr, Jar: smartisanData.CookieJar}
	req, _ := http.NewRequest("POST", "http://www.smartisan.com/store/index.php?r=lottery/winlist", body)
	req.Header.Set("Referer", "https://account.smartisan.com/")
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded; param=value") //这个一定要加，不加form的值post不过去，被坑了两小时                                                       //看下发送的结构

	resp, _ := client.Do(req) //发送

	defer resp.Body.Close() //一定要关闭resp.Body
	data, _ := ioutil.ReadAll(resp.Body)

	fmt.Println(string(data))
}

func Hit() {
	v := url.Values{}
	body := ioutil.NopCloser(strings.NewReader(v.Encode())) //把form数据编下码
	client := &http.Client{Transport: smartisanData.Tr, Jar: smartisanData.CookieJar}
	req, _ := http.NewRequest("POST", "http://www.smartisan.com/store/index.php?r=activity/getJianGuoPro100DayCoupon&action=get", body)
	req.Header.Set("Referer", "https://account.smartisan.com/")
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded; param=value") //这个一定要加，不加form的值post不过去，被坑了两小时                                                       //看下发送的结构

	resp, _ := client.Do(req) //发送

	defer resp.Body.Close() //一定要关闭resp.Body
	data, _ := ioutil.ReadAll(resp.Body)

	fmt.Println(string(data))
}

func BuyGood() {
	v := url.Values{}

	body := ioutil.NopCloser(strings.NewReader(v.Encode())) //把form数据编下码
	client := &http.Client{Transport: smartisanData.Tr, Jar: smartisanData.CookieJar}
	req, _ := http.NewRequest("GET", "http://tp.hd.mi.com/hdget/cn?product=1175000016&m=2&addcart=1&jsonpcallback=cn1175000016&source=flashsale", body)
	//req.Header.Set("Referer", "https://account.xiaomi.com/pass/serviceLogin?callback=https%3A%2F%2Faccount.xiaomi.com%2Fsts%3Fsign%3DZvAtJIzsDsFe60LdaPa76nNNP58%253D%26followup%3Dhttps%253A%252F%252Faccount.xiaomi.com%252Fpass%252Fauth%252Fsecurity%252Fhome%26sid%3Dpassport&sid=passport")
	//req.Header.Set("Content-Type", "application/x-www-form-urlencoded; param=value") //这个一定要加，不加form的值post不过去，被坑了两小时
	req.Header.Set("Accept-Encoding", "gzip")
	req.Header.Set("User-Agent", "okhttp/3.4.1")

	req.Header.Set("device", "jSD3YNkogSzDgV0wOwe3AA==")
	req.Header.Set("Device-Id", "ffffffff-e836-345b-ffff-ffffe89c3d90")
	req.Header.Set("Mishop-Auth", "57b6d971fe5c35b7;3014971118")
	req.Header.Set("Mishop-Client-Id", "180100031052")
	req.Header.Set("Mishop-Client-VersionCode", "20171213")
	req.Header.Set("Uuid", "84c8b9a5-bcfb-523d-9e65-271c07d9c8a1")
	req.Header.Set("Host", "tp.hd.mi.com")

	resp, _ := client.Do(req) //发送

	defer resp.Body.Close() //一定要关闭resp.Body
	data, _ := ioutil.ReadAll(resp.Body)

	fmt.Println("BuyGood", data)
}

func Remind() {
	v := url.Values{}
	v.Set("start_time", "1513677600")
	v.Set("goods_id", "1175000018")
	v.Set("is_remind", "1")

	body := ioutil.NopCloser(strings.NewReader(v.Encode())) //把form数据编下码
	client := &http.Client{Transport: smartisanData.Tr, Jar: smartisanData.CookieJar}
	req, _ := http.NewRequest("POST", "http://api.m.mi.com/v1/seckill/remind3", body)
	//req.Header.Set("Referer", "https://account.xiaomi.com/pass/serviceLogin?callback=https%3A%2F%2Faccount.xiaomi.com%2Fsts%3Fsign%3DZvAtJIzsDsFe60LdaPa76nNNP58%253D%26followup%3Dhttps%253A%252F%252Faccount.xiaomi.com%252Fpass%252Fauth%252Fsecurity%252Fhome%26sid%3Dpassport&sid=passport")
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded; param=value") //这个一定要加，不加form的值post不过去，被坑了两小时

	req.Header.Set("device", "jSD3YNkogSzDgV0wOwe3AA==")
	req.Header.Set("Device-Id", "ffffffff-e836-345b-ffff-ffffe89c3d90")
	req.Header.Set("Mishop-Auth", "57b6d971fe5c35b7;3014971118")
	req.Header.Set("Mishop-Client-Id", "180100031052")
	req.Header.Set("Mishop-Client-VersionCode", "20171213")

	resp, _ := client.Do(req) //发送

	defer resp.Body.Close() //一定要关闭resp.Body
	data, _ := ioutil.ReadAll(resp.Body)

	fmt.Println(string(data))
}

func CallBack(callback string) {
	v := url.Values{}

	body := ioutil.NopCloser(strings.NewReader(v.Encode())) //把form数据编下码
	client := &http.Client{Transport: smartisanData.Tr, Jar: smartisanData.CookieJar}
	req, _ := http.NewRequest("GET", callback, body)
	req.Header.Set("Referer", "https://account.xiaomi.com/pass/serviceLogin?callback=https%3A%2F%2Faccount.xiaomi.com%2Fsts%3Fsign%3DZvAtJIzsDsFe60LdaPa76nNNP58%253D%26followup%3Dhttps%253A%252F%252Faccount.xiaomi.com%252Fpass%252Fauth%252Fsecurity%252Fhome%26sid%3Dpassport&sid=passport")
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded; param=value") //这个一定要加，不加form的值post不过去，被坑了两小时

	req.Header.Set("device", "jSD3YNkogSzDgV0wOwe3AA==")
	req.Header.Set("Device-Id", "ffffffff-e836-345b-ffff-ffffe89c3d90")
	req.Header.Set("Mishop-Auth", "57b6d971fe5c35b7;3014971118")
	req.Header.Set("Mishop-Client-Id", "180100031052")
	req.Header.Set("Mishop-Client-VersionCode", "20171213")

	resp, _ := client.Do(req) //发送

	defer resp.Body.Close() //一定要关闭resp.Body
	data, _ := ioutil.ReadAll(resp.Body)

	fmt.Println("callBack:", string(data))
}

func AHL() {
	v := url.Values{}
	v.Set("_json", "true")
	v.Set("sid", "mipaypr")

	body := ioutil.NopCloser(strings.NewReader(v.Encode())) //把form数据编下码
	client := &http.Client{Transport: smartisanData.Tr, Jar: smartisanData.CookieJar}
	req, _ := http.NewRequest("POST", "http://a.hl.mi.com/m", body)
	req.Header.Set("Referer", "https://account.xiaomi.com/pass/serviceLogin?callback=https%3A%2F%2Faccount.xiaomi.com%2Fsts%3Fsign%3DZvAtJIzsDsFe60LdaPa76nNNP58%253D%26followup%3Dhttps%253A%252F%252Faccount.xiaomi.com%252Fpass%252Fauth%252Fsecurity%252Fhome%26sid%3Dpassport&sid=passport")
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded; param=value") //这个一定要加，不加form的值post不过去，被坑了两小时

	req.Header.Set("Android-Ver", "19")
	req.Header.Set("device", "jSD3YNkogSzDgV0wOwe3AA==")
	req.Header.Set("Device-Id", "ffffffff-e836-345b-ffff-ffffe89c3d90")
	req.Header.Set("Mishop-Auth", "57b6d971fe5c35b7;3014971118")
	req.Header.Set("Mishop-Client-Id", "180100031052")
	req.Header.Set("Mishop-Client-VersionCode", "20171213")
	req.Header.Set("Uuid", "a31c8e32-8a74-c122-1486-bb941ba780fb")

	resp, _ := client.Do(req) //发送

	defer resp.Body.Close() //一定要关闭resp.Body
	data, _ := ioutil.ReadAll(resp.Body)

	fmt.Println(string(data))
}

func HuoDongLogin() {
	v := url.Values{}
	v.Set("_json", "true")
	v.Set("sid", "mipaypr")

	body := ioutil.NopCloser(strings.NewReader(v.Encode())) //把form数据编下码
	client := &http.Client{Transport: smartisanData.Tr, Jar: smartisanData.CookieJar}
	req, _ := http.NewRequest("GET", "https://account.xiaomi.com/pass/serviceLogin", body)
	req.Header.Set("Referer", "https://account.xiaomi.com/pass/serviceLogin?callback=https%3A%2F%2Faccount.xiaomi.com%2Fsts%3Fsign%3DZvAtJIzsDsFe60LdaPa76nNNP58%253D%26followup%3Dhttps%253A%252F%252Faccount.xiaomi.com%252Fpass%252Fauth%252Fsecurity%252Fhome%26sid%3Dpassport&sid=passport")
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded; param=value") //这个一定要加，不加form的值post不过去，被坑了两小时

	req.Header.Set("device", "jSD3YNkogSzDgV0wOwe3AA==")
	req.Header.Set("Device-Id", "ffffffff-e836-345b-ffff-ffffe89c3d90")
	req.Header.Set("Mishop-Auth", "57b6d971fe5c35b7;3014971118")
	req.Header.Set("Mishop-Client-Id", "180100031052")
	req.Header.Set("Mishop-Client-VersionCode", "20171213")
	req.Header.Set("Uuid", "a31c8e32-8a74-c122-1486-bb941ba780fb")

	resp, _ := client.Do(req) //发送

	defer resp.Body.Close() //一定要关闭resp.Body
	data, _ := ioutil.ReadAll(resp.Body)

	fmt.Println(string(data))
}

//登陆函数
func Login() {
	smartisanData.CookieJar, _ = cookiejar.New(nil)
	smartisanData.Tr = &http.Transport{
		TLSClientConfig:    &tls.Config{InsecureSkipVerify: true},
		DisableCompression: true,
	}

	v := url.Values{}
	v.Set("envKey", "Cw7oQKiJ7Qx8hDTeduTA5FSOwCSz0FwzWiqrrDqaIXEJWKl9SL5hk6tYX1kzQCkE4CHgMT2PMHJZH0AcNCcvxoDBfHoNcrg8bU-hFUuaRd5ZJAd1onQJt_Ym2YvUta2LuA1GydAuH2E9WKXzrIhA0_I7wS8OR3ai761R-5SCgIQ=")
	v.Set("sid", "eshopmobile")
	v.Set("hash", "B1884621478CB6CE7AFD067EDE8AD2B7")
	v.Set("callback", "http://api.m.mi.com/v1/passport/callback")
	v.Set("_json", "true")
	v.Set("env", "D9iuRcc8wrOIg24rBWyjuQHCuij6Rp6rAvb97af8ktek65tsxcQsUGsMhgQXgpGJLw0v3Lf7clX7EqebZtGZAZWvFwZ3X0A97VJ1fjjwfNE8Q7VA/e7hNQLp2lpPe4zV7XrOaIGgjDRmcztHcpGX9M+aYIdJaTNjrKZlSkl1eAdIjSmzvI6MqfqMQlfevXR0uoFYvBEjl3bDrFXRnmXMwg==")
	v.Set("qs", "%3F_json%3Dtrue%26sid%3Deshopmobile")
	v.Set("_sign", "vSE3gT4hGtNnZzE0amfGBZtXmZ0=")
	v.Set("user", "15111048113")
	body := ioutil.NopCloser(strings.NewReader(v.Encode())) //把form数据编下码
	client := &http.Client{Transport: smartisanData.Tr, Jar: smartisanData.CookieJar}
	req, _ := http.NewRequest("POST", "https://account.xiaomi.com/pass/serviceLoginAuth2?_dc=1513663850442", body)
	req.Header.Set("Referer", "https://account.xiaomi.com/pass/serviceLogin?callback=https%3A%2F%2Faccount.xiaomi.com%2Fsts%3Fsign%3DZvAtJIzsDsFe60LdaPa76nNNP58%253D%26followup%3Dhttps%253A%252F%252Faccount.xiaomi.com%252Fpass%252Fauth%252Fsecurity%252Fhome%26sid%3Dpassport&sid=passport")
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded; param=value") //这个一定要加，不加form的值post不过去，被坑了两小时
	req.Header.Set("Host", "api.m.mi.com")

	req.Header.Set("device", "jSD3YNkogSzDgV0wOwe3AA==")
	req.Header.Set("Device-Id", "ffffffff-e836-345b-ffff-ffffe89c3d90")
	req.Header.Set("Mishop-Auth", "57b6d971fe5c35b7;3014971118")
	req.Header.Set("Mishop-Client-Id", "180100031052")
	req.Header.Set("Mishop-Client-VersionCode", "20171213") //看下发送的结构

	resp, _ := client.Do(req) //发送

	defer resp.Body.Close() //一定要关闭resp.Body
	data, _ := ioutil.ReadAll(resp.Body)

	strData := string(data)
	strData = strData[11:len(strData)]
	fmt.Println(strData)
	var loginstruct smartisanData.LoginStruct

	if err := json.Unmarshal([]byte(strData), &loginstruct); err == nil {
		fmt.Println(loginstruct.Location)
		CallBack(loginstruct.Location)
	}

	//获取地址信息
	//getAddress()
}

func getAddress() {
	postDataStr := ""

	body := ioutil.NopCloser(strings.NewReader(postDataStr)) //把form数据编下码
	client := &http.Client{Transport: smartisanData.Tr, Jar: smartisanData.CookieJar}
	req, _ := http.NewRequest("POST", "http://store.smartisan.com/serv/v1/address/list", body)

	req.Header.Set("Referer", "http://store.smartisan.com/")
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded; param=value") //这个一定要加，不加form的值post不过去，被坑了两小时                                                       //看下发送的结构

	resp, _ := client.Do(req) //发送

	defer resp.Body.Close() //一定要关闭resp.Body
	data, _ := ioutil.ReadAll(resp.Body)

	err := json.Unmarshal(data, &smartisanData.AddressData)
	if err != nil {
		return
	}

	LogStr("获取收货地址成功:", smartisanData.AddressData.Data.List[0])
}
