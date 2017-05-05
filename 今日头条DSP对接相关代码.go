package main

import (
	"bytes"
	"crypto/hmac"
	"crypto/sha1"
	"encoding/base64"
	"encoding/binary"
	"encoding/json"
	"fmt"
	"reflect"
)

var I_key string = "ebcd1234efgh5678ebcd1234efgh5678" //原始值
var E_key string = "5678dcba1234abcd5678dcba1234abcd" //加密过的

type First_Data struct {
	RrequestId string
	ApiVersion string
}

type Adslots_Data struct {
	Id       string
	Banner   []map[string]interface{}
	AdType   []string
	BidFloor int
	KeyWords []string
}

type Adslots []Adslots_Data //保存的相关的类型

type App map[string]string //app的信息

//设备信息
type Device struct {
	Ip             string
	Geo            map[string]interface{}
	Deviceld       string
	Make           string
	Model          string
	Os             string
	Osv            string
	ConnectionType string
	DeviceType     string
	AndroidId      string
}

//用户的相关信息
type User struct {
	Id     string
	Yob    string
	Gender string
	Data   []interface{}
}

//---------------------------请求数据类型定义

func main() {

	fmt.Println("--->><<<<")
	//p_price := Decprice("VH2SogAO-2ZUfZKiAA77ZlONOMFBLsR53TOlOg")
	//fmt.Println(p_price)
	//CreatSign("04d0d988f3634f3f80147dc09a5bfd42", "http://adx.toutiao.com/adxbuyer/api/v1.0/creatives/put?dspid=12345&creative_num=3")
	//数据解析
	//DecodeData()
	EncodeData()
}

//返回金额是零表示失败
func Decprice(base64price string) uint64 {
	defer func() uint64 {
		if err := recover(); err != nil {
			return 0
		}
		return 0
	}()
	Lstr := base64price

	for {
		if len(Lstr)%4 == 0 {
			break
		}
		Lstr += "="
	}
	base64.RawStdEncoding.DecodeString(Lstr)
	enc_price, er := base64.URLEncoding.DecodeString(Lstr)

	if er != nil || len(enc_price) < 28 {
		return 0
	}

	iv := enc_price[0:16]
	price := enc_price[16:24]
	sig := enc_price[24:28]

	//fmt.Println(string(kla))
	e_key := []byte(E_key)
	e_key_d := hmac.New(sha1.New, e_key)
	e_key_d.Write(iv)
	e_key_end := e_key_d.Sum(nil)

	la := make([]byte, 8)
	for i := 0; i < 8; i++ {
		la[i] = price[i] ^ e_key_end[i]
	}

	//验证当前的签名
	laa := BytesCombine(la, iv)
	e_keya := []byte(I_key)
	e_keya_d := hmac.New(sha1.New, e_keya)
	e_keya_d.Write(laa)
	e_keya_end := e_keya_d.Sum(nil)

	sign_k := e_keya_end[0:4]
	istrue := bytes.Compare(sign_k, sig) //返回的值是零表示相等

	if istrue == 0 {
		v := BytesToInt64(la)
		return uint64(v)
	}
	return 0
}

func BytesCombine(pBytes ...[]byte) []byte {
	len := len(pBytes)
	s := make([][]byte, len)
	for index := 0; index < len; index++ {
		s[index] = pBytes[index]
	}
	sep := []byte("")
	return bytes.Join(s, sep)
}

func BytesToInt64(buf []byte) int64 {
	return int64(binary.BigEndian.Uint64(buf))
}

func Int64ToBytes(i int64) []byte {
	var buf = make([]byte, 8)
	binary.BigEndian.PutUint64(buf, uint64(i))
	return buf
}

//数据格式转化的操作

func StrToMap(data string) map[string]interface{} {
	var dat map[string]interface{}
	json.Unmarshal([]byte(data), &dat)
	return dat
}

// 结构转换成json对象
func JsonEncodeString(data interface{}) string {
	back, err := json.Marshal(data)
	if err != nil {
		return "encode error"
	}
	return string(back)
}

//map的类型转换成！

func MaptoJson(data map[string]interface{}) string {
	configJSON, err := json.MarshalIndent(data, "", "\t")
	if err != nil {
		return ""
	}
	return string(configJSON) //返回格式化后的字符串的内容0
}

//生成签名
func CreatSign(key, url string) string {
	e_keya := []byte(key)
	urls := []byte(url)
	e_keya_d := hmac.New(sha1.New, e_keya)
	e_keya_d.Write(urls)
	e_keya_end := e_keya_d.Sum(nil)
	return base64.StdEncoding.EncodeToString(e_keya_end)
}

//解析数据

func DecodeData() {
	data := "{\"requestId\":\"20170503145932172017201003300546\",\"apiVersion\":\"2.1\",\"adslots\":[{\"id\":\"7f9b804258734dad\",\"banner\":[{\"width\":580,\"height\":240,\"pos\":\"FEED\",\"sequence\":\"2\"}],\"adType\":[\"TOUTIAO_FEED_APP_LARGE\",\"TOUTIAO_FEED_LP_LARGE\",\"TOUTIAO_FEED_APP_SMALL\",\"TOUTIAO_FEED_LP_SMALL\",\"TOUTIAO_FEED_LP_GROUP\",\"TOUTIAO_FEED_APP_GROUP\"],\"bidFloor\":400,\"keywords\":[]}],\"app\":{\"id\":\"13\",\"name\":\"news_article\",\"ver\":\"589\"},\"device\":{\"ip\":\"111.204.243.7\",\"geo\":{\"lat\":40.001040000000003,\"lon\":116.38748,\"city\":\"北京\"},\"deviceId\":\"359596063768448\",\"make\":\"unknown\",\"model\":\"SM-G9250\",\"os\":\"android\",\"osv\":\"6.0.1\",\"connectionType\":\"WIFI\",\"deviceType\":\"PHONE\",\"androidId\":\"fb53a1101c34b557\"},\"user\":{\"id\":\"8555353760\",\"yob\":\"31\",\"gender\":\"MALE\",\"data\":[]}}"
	op := StrToMap(data)
	fmt.Println(op["requestId"])
	adsData := op["adslots"].([]interface{})

	mapData := adsData[0].(map[string]interface{})
	fmt.Println(mapData["id"]) //
	fmt.Println(adsData[0].(map[string]interface{}))

	ml := reflect.ValueOf(op)
	fmt.Println(ml.Type())

	//使用反射
	fmt.Println(reflect.ValueOf(op))

}

//生成请求的数据

type ImageBannerType struct {
	Width  uint8    `json:"width"`
	Height uint8    `json:"height"`
	Url    string   `json:"url"`
	Urls   []string `json:"urls"`
}

var ImageBannerData ImageBannerType

type CreativeDat struct {
	AdType          string            `json:"adtype"`
	Nurl            string            `json:"nurl"`
	Title           string            `json:"title"`
	Source          string            `json:"Source"`
	ImageBanner     ImageBannerType   `json:"imageBanner"`
	External        map[string]string `json:"External"`
	ShowUrl         []string          `json:"showUrl"`
	ClickUrl        []string          `json:"ClickUrl"`
	VideoPlayStart  []string          `json:"videoPlayStart"`
	VideoPlayFinish []string          `json:"videoPlayFinish"`
	SplashCreatives []string          `json:"splashCreatives"`
}

var CreativeData CreativeDat

type SendAdsData struct {
	Id       string      `json:"id"`
	AdslotId string      `json:"adslotId"`
	Price    uint        `json:"price"`
	Adid     string      `json:"adid"`
	Creative CreativeDat `json:"creative"`
	Cid      string      `json:"cid"`
}

var SendAdsDatas SendAdsData

type TodayEndData struct {
	Ads []SendAdsData `json:"ads"`
}

var TodayEndDatas TodayEndData

//请求今日头条的数据结构
type TodayData struct {
	RequestId string         `json:"requestId"`
	Seatbids  []TodayEndData `json:"seatbids"`
}

var TodayDataNow TodayData

func EncodeData() {
	//设置返回的banner图片信息
	ImageBannerData.Height = 150
	ImageBannerData.Width = 228
	ImageBannerData.Url = "http://jrtt.qcwanwan.com/1.jpg"
	img_baner_urls := []string{"http://jrtt.qcwanwan.com/1.jpg", "http://jrtt.qcwanwan.com/2.jpg", "http://jrtt.qcwanwan.com/3.jpg"}
	ImageBannerData.Urls = img_baner_urls

	//数据
	CreativeData.AdType = "TOUTIAO_FEED_LP_GROUP"
	CreativeData.Nurl = "http://jrtt.qcwanwan.com/win/notify?user_id={user_id}&request_id={request_id}&adid={adid}&bid_price={bid_price}&ip={ip}×tamp={timestamp}&did={did}"
	CreativeData.Title = "白野猪爆了个装备.换了小半个月工资"
	CreativeData.Source = "传奇无双"
	CreativeData.ImageBanner = ImageBannerData
	CreativeData.External = map[string]string{"url": "http://m.anfeng.cn/cqws_bbk-ios/12/"}
	CreativeData.ShowUrl = []string{"http://jrtt.qcwanwan.com/show/notify?user_id={user_id}&request_id={request_id}&adid={adid}&bid_price={bid_price}&ip={ip}×tamp={timestamp}&did={did}"}
	CreativeData.ClickUrl = []string{"http://jrtt.qcwanwan.com/click/notify?user_id={user_id}&request_id={request_id}&adid={adid}&bid_price={bid_price}&ip={ip}×tamp={timestamp}&did={did}"}
	CreativeData.VideoPlayFinish = []string{}
	CreativeData.SplashCreatives = []string{}

	//Ads data
	SendAdsDatas.Id = "59097fd429eedf29504df678"
	SendAdsDatas.AdslotId = "7f9b804258734dad"
	SendAdsDatas.Price = 401
	SendAdsDatas.Adid = "171186149393006560"
	SendAdsDatas.Creative = CreativeData
	SendAdsDatas.Cid = "171186149393006548"

	//
	TodayEndDatas.Ads = []SendAdsData{SendAdsDatas}
	TodayDataNow.RequestId = "20170503145932172017201003300546"
	TodayDataNow.Seatbids = []TodayEndData{TodayEndDatas}

	fmt.Println("-----")

	jsons, _ := json.Marshal(TodayDataNow)
	fmt.Println(string(jsons))
	//fmt.Println(MaptoJson(TodayDataNow))

}
