package main

import (
	"bufio"
	"encoding/csv"
	"encoding/json"
	"fmt"
	"github.com/dgrijalva/jwt-go"
	"github.com/xuri/excelize/v2"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"io"
	"net/http"
	"os"
	"strings"
	"time"
)

func main() {
	//parseRowsExcel()
	//parseColsExcel("/Users/zhuweijie/Documents/GoProject/parse_excel/agent.xlsx")
	//parseRowsExcel("/Users/zhuweijie/Documents/GoProject/parse_excel/first_recharge_records.xlsx")
	CheckFirstPay()
	//http.Handle("/", http.HandlerFunc(hello))

	//fmt.Println(Credential(100998687))
	//parseLogCSV("/Users/zhuweijie/Documents/GoProject/parse_excel/log.csv")
	//parse_music_file.ParseMusicFile("https://img.520yidui.com/upload/files/music_file.xlsx")
	//parse_music_file.ParseMusicFile2()
	//ParseText()
	//rangeMatch()
}

func FormatTarget() {
	path := ""
	parseLogCSV(path)
}

func Credential(memberId int64) (string, error) {
	return jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"id":           memberId,
		"expire_at":    time.Now().Add(time.Hour * 2).Format("2006-01-02 15:04:05"),
		"device_id":    "D629EE5A-B785-4AD5-BF5C-7D017D28595F",
		"gioid":        "",
		"ydid":         "8349b51118d0461aba124f9a636b32a4",
		"channel_name": "",
		"platform":     0,
		"noncestr":     "ALcEiUwGwVoHNNBvhhabcDjymAUcuhNG",
	}).SignedString([]byte("e5bc1bt093f9d39df09665714be98a2be93dc514f86914d81cf678a6f294291b11eeecc8bec091ab4cf7efb1b7a0920267ddbc042958bad730d6249a3ed2b786"))
}

func timeEscp(next http.Handler) http.Handler {
	return http.HandlerFunc(func(writer http.ResponseWriter, request *http.Request) {
		next.ServeHTTP(writer, request)

	})
}

func hello(wr http.ResponseWriter, r *http.Request) {
	wr.Write([]byte("hello"))
}

type BosomFriendGiftConfig struct {
	GiftId          int64  `json:"gift_id"`
	RarityLevel     int64  `json:"rarity_level"`
	RarityLevelName string `json:"rarity_level_name"`
	Source          int64  `json:"source"` // 1挚友礼物,2 福袋礼物
}

var matchmakers []int64 = []int64{10051207, 10016422, 10022886, 10035199, 10040023, 10046350, 10040654, 10042990, 10043866, 10049308, 10049500, 10049505, 10000966, 10051647, 10056574, 10057438, 10056098, 10060319, 10061445, 10085963, 10100107, 10106815, 10116622, 10117487, 10117662, 10118108, 10118373, 10121030, 10121027, 10138111, 10140068, 10153047, 10163084, 10175882, 10186432, 10194249, 10201492, 10206417, 10211410, 10204253, 10223942, 10249211, 10247626, 10250991, 10254007, 10259514, 10263362, 10261674, 10273480, 10278038, 10279346, 10279569, 10278872, 10288936, 10289058, 10347423, 10385303}
var matchmakers2 []int64 = []int64{10022886, 10035199, 10043866, 10049500, 10049505, 10060319, 10121030, 10138111, 10153047, 10186432, 10211410, 10259514, 10000966, 10016422, 10040023, 10040654, 10042990, 10046350, 10049308, 10051647, 10056098, 10056574, 10058503, 10118373, 10140068, 10163084, 10175882, 10194249, 10206417, 10249211, 10250991, 10254007, 10278083, 10279346, 10279569, 10288936, 10051207, 10057438, 10085963, 10100107, 10103283, 10106815, 10116622, 10117487, 10117662, 10121027, 10201492, 10204253, 10223942, 10261674, 10273480, 10278038, 10347423, 10052571, 10061445, 10118108, 10214911, 10247626, 10263362, 10278872, 10279582, 10289058, 10385303}

func rangeMatch() {
	matchMap1 := make(map[int64]int64)
	for _, v := range matchmakers {
		matchMap1[v] = v
	}

	matchMap2 := make(map[int64]int64)
	for _, v := range matchmakers2 {
		matchMap2[v] = v
	}

	for k, _ := range matchMap1 {
		if _, ok := matchMap2[k]; !ok {
			fmt.Println("1独有", k)
		}
	}

	fmt.Println("=========")

	for k, _ := range matchMap2 {
		if _, ok := matchMap1[k]; !ok {
			fmt.Println(k)
		}
	}

}

func parseRowsExcel(path string) {
	f, err := excelize.OpenFile(path)
	if err != nil {
		fmt.Println("错误", err)
		return
	}
	defer func() {
		f.Close()
	}()

	rows, err := f.GetRows("Sheet1")
	if err != nil {
		fmt.Println("row错误", err)
		return
	}

	for _, row := range rows {
		strs := make([]string, 0)
		for _, v := range row {
			strs = append(strs, v)
		}
		if len(strs) == 0 {
			continue
		}
		member := strs[1]
		extStr := strs[8]
		if extStr == "" {
			continue
		}
		ext := &GoldPayExtra{}
		err = json.Unmarshal([]byte(extStr), ext)
		if err != nil {
			fmt.Println("err==>", err.Error())
			continue
		}
		payMemberMpp := make(map[string]struct{})
		key := fmt.Sprintf("%d_%s", member, ext.ProductUid)
		payMemberMpp[key] = struct{}{}
	}

}

func parseRowsExcelRows(path string) (rows [][]string) {
	f, err := excelize.OpenFile(path)
	if err != nil {
		fmt.Println("错误", err)
		return
	}
	defer func() {
		f.Close()
	}()

	rows, err = f.GetRows("Sheet1")
	if err != nil {
		fmt.Println("row错误", err)
		return
	}
	return rows
}

func CheckFirstPay() {
	goldRows := parseRowsExcelRows("/Users/zhuweijie/Documents/GoProject/parse_excel/gold_pay_records.xlsx")
	goldMemberMap := make(map[string]struct{})
	for k, row := range goldRows {
		if k == 0 {
			continue
		}
		strs := make([]string, 0)
		for _, v := range row {
			strs = append(strs, v)
		}
		if len(strs) == 0 {
			continue
		}
		member := strs[1]
		extStr := strs[8]
		if extStr == "" {
			continue
		}
		ext := &GoldPayExtra{}
		err := json.Unmarshal([]byte(extStr), ext)
		if err != nil {
			fmt.Println("err==>", err.Error())
			fmt.Println("===>", extStr)
			continue
		}

		key := fmt.Sprintf("%s_%s", member, ext.ProductUid)
		goldMemberMap[key] = struct{}{}
	}

	//fmt.Println("goldMap==>", goldMemberMap)

	rechargeRows := parseRowsExcelRows("/Users/zhuweijie/Documents/GoProject/parse_excel/first_recharge_records.xlsx")
	rechargeMap := make(map[string]struct{})
	for k, row := range rechargeRows {
		if k == 0 {
			continue
		}
		strs := make([]string, 0)
		for _, v := range row {
			strs = append(strs, v)
		}
		if len(strs) == 0 {
			continue
		}
		member := strs[1]
		uid := strs[3]
		key := fmt.Sprintf("%s_%s", member, uid)
		rechargeMap[key] = struct{}{}
	}

	//fmt.Println("rechargeMap==>", rechargeMap)

	for k, _ := range goldMemberMap {
		_, has := rechargeMap[k]
		if !has {
			fmt.Println("==>", k)
		}
	}
}

type GoldPayExtra struct {
	FirstPay   bool   `json:"first_pay"` // 是否第一次支付
	PayId      int64  `json:"pay_id"`    // 支付id
	ProductUid string `json:"uid"`       // 商品uid
	TradeNo    string `json:"trade_no"`  // 支付订单号
	Extra      string `json:"extra"`     // 额外信息
}

func ParseText() {
	file, err := os.Open("/Users/zhuweijie/Documents/GoProject/parse_excel/file.text")
	if err != nil {
		fmt.Println("Failed to open file:", err)
		return
	}
	defer file.Close()
	db := InitEzMysql("")
	scanner := bufio.NewScanner(file)
	noUpdates := make([]string, 0)
	for scanner.Scan() {
		line := scanner.Text()

		// 在这里处理每一行的逻辑
		strs := strings.Split(line, ";")
		data := make(map[string]interface{})
		data["value"] = strs[1]
		data["extra"] = "push"
		origin := strings.TrimSpace(strs[0])
		value := strings.TrimSpace(strs[1])
		count := db.Table("system_tips_translate").Where("origin = ?", origin).Updates(data).RowsAffected
		if count == 0 {
			//fmt.Println(line)
			noUpdates = append(noUpdates, line)
			sql := fmt.Sprintf(`insert into system_tips_translate (md5, language, origin, value, created_at, updated_at, serve, type,extra) values (md5("%s"),"en","%s","%s",now(),now(),1,1,"push")`, origin, origin, value)
			err := db.Table("system_tips_translate").Exec(sql).Error
			if err != nil {
				fmt.Println("err==>", err.Error())
			}
		}
	}

	if err := scanner.Err(); err != nil {
		fmt.Println("Scanner error:", err)
	}
	for _, v := range noUpdates {
		fmt.Println(v)
	}
	//fmt.Println("===>",noUpdates)
}

func InitEzMysql(name string) *gorm.DB {
	port := 3306
	database := "cht_users_test"
	args := fmt.Sprintf("test_user:Yidui_1207!@tcp(test-internal.chattalive.com:%d)/%s?parseTime=true&loc=Local", port, database)

	db, err := gorm.Open(mysql.Open(args))

	if err != nil {
		fmt.Fprintln(os.Stderr, "mysql open error ", err)
	}
	return db
}

// 按列读
func parseColsExcel(str string) {
	f, err := excelize.OpenFile(str)
	if err != nil {
		fmt.Println("错误", err)
		return
	}
	defer func() {
		f.Close()
	}()

	cols, err := f.GetCols("Sheet1")
	if err != nil {
		fmt.Println("row错误", err)
		return
	}
	//fmt.Println("===v", cols)

	for _, cs := range cols[2] {
		s := `"` + cs + `",`
		fmt.Print(s)
	}
}

type LogInfo struct {
	Addr  string `json:"addr"`
	Count int64  `json:"count"`
}

type LogInfos []LogInfo

func (l LogInfos) Len() int {
	return len(l)
}

func (l LogInfos) Less(i, j int) bool {
	return l[i].Count > l[j].Count
}

func (s LogInfos) Swap(i, j int) {
	s[i], s[j] = s[j], s[i]
}

var GRate float64

type LogData struct {
	TestGroup string `json:"test_group"`
	Status    int64  `json:"status"`
	Action    string `json:"action"`
	MicType   string `json:"mic_type"`
	MicId     int64  `json:"mic_id"`
}

func parseLogCSV(str string) {
	file, err := os.Open(str)
	if err != nil {
		fmt.Println(err)
	}

	defer file.Close()

	reader := csv.NewReader(bufio.NewReader(file))
	reader.FieldsPerRecord = -1

	xlsx := excelize.NewFile()
	sheetName := "rate"
	xlsx.SetSheetName(xlsx.GetSheetName(1), sheetName)

	var test2, test3, test4 int64
	countMap := make(map[string]int64)
	index := 0
	row := 1
	for {
		index = index + 1
		csvdata, err := reader.Read() // 按行读取数据,可控制读取部分
		if err != nil {
			if err == io.EOF {
				break
			}
			fmt.Println("err==>", err)
			return
		}

		rateStr := csvdata[17]
		data := &LogData{}
		err = json.Unmarshal([]byte(rateStr), data)
		if err != nil {
			fmt.Println("err==>", err.Error())
			continue
		}
		if data.Action != "push" {
			continue
		}
		if data.Status != 1 {
			continue
		}
		if data.TestGroup == "2" {
			test2 = test2 + 1
		}
		if data.TestGroup == "3" {
			test3 = test3 + 1
		}
		if data.TestGroup == "4" {
			test4 = test4 + 1
		}
		key := fmt.Sprintf("%s", data.TestGroup)
		if val, ok := countMap[key]; ok {
			countMap[key] = val + 1
		} else {
			countMap[key] = 1
		}
		//str := fmt.Sprintf("update day_income_details set valid_live_minute =%s, live_seconds = %s, valid_live_seconds = %s where member_id = %s and days = '%s';", csvdata[2], csvdata[3], csvdata[3], csvdata[1], days)
		//AppendStringToFile(str, "sql.txt")
		row++
	}
	for k, v := range countMap {
		fmt.Println("==>", k, v)
	}

}

func AppendStringToFile(str string, filename string) error {
	file, err := os.OpenFile(filename, os.O_APPEND|os.O_WRONLY, 0644)
	if err != nil {
		return err
	}
	defer file.Close()

	scanner := bufio.NewScanner(strings.NewReader(str))
	for scanner.Scan() {
		line := scanner.Text()
		_, err := fmt.Fprintln(file, line)
		if err != nil {
			return err
		}
	}

	return nil
}

func WriteStringToFile(str string, filename string) error {
	file, err := os.Create(filename)
	if err != nil {
		return err
	}
	defer file.Close()

	scanner := bufio.NewScanner(strings.NewReader(str))
	for scanner.Scan() {
		line := scanner.Text()
		_, err := fmt.Fprintln(file, line)
		if err != nil {
			return err
		}
	}

	return nil
}
