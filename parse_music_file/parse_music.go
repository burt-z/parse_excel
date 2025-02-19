package parse_music_file

import (
	"bufio"
	"bytes"
	"fmt"
	"github.com/xuri/excelize/v2"
	"io"
	"io/ioutil"
	"net/http"
	"os"
	"time"
)

func ParseMusicFile(url string) {
	client := &http.Client{Timeout: time.Second * 5}
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		fmt.Println("err 1==>", err)
		return
	}
	resp, err := client.Do(req)
	if err != nil {
		fmt.Println("err 2==>", err)
		return
	}
	defer resp.Body.Close()

	dir, err := os.Getwd()
	if err != nil {
		fmt.Println("err 3==>", err)
		return
	}
	filePath := fmt.Sprintf("%s/music_file.xlsx", dir)
	fmt.Println("filePath", filePath)

	// 将 Excel 文件保存到本地
	file, err := os.Create(filePath)
	if err != nil {
		fmt.Println("err 4==>", err)
		return
	}
	defer func() {
		file.Close()
		//os.Remove("./music_file.xlsx")
	}()

	body := resp.Body
	write := bufio.NewWriter(file)
	bytes := make([]byte, 1024*1024)
	for {
		n, err := body.Read(bytes)
		if err != nil && err != io.EOF {
			fmt.Println("err 4==>", err)
			break
		}
		_, writeErr := write.Write(bytes[:n])
		if writeErr != nil {
			fmt.Println("err 5==>", writeErr)
			break
		}
		if err == io.EOF {
			break
		}
	}
	write.Flush()

	f, err := excelize.OpenFile(filePath)
	if err != nil {
		fmt.Println("open错误", err)
		return
	}
	defer func() {
		f.Close()
	}()

	rows, err := f.GetRows("song001")
	if err != nil {
		fmt.Println("row错误", err)
		return
	}

	index := 0
	for k, row := range rows {
		if k == 0 {
			continue
		}
		if k < 9 {
			fmt.Println("====>", fmt.Sprintf("%s:%s:%s:%s", row[0], row[1], row[2], row[3]))
		}
		index = k
	}
	fmt.Println("===index", index)
}

func ParseMusicFile2() {
	fileByte, err := ioutil.ReadFile("music_file.xlsx")
	if err != nil {
		fmt.Println("err 1==>", err)
		return
	}

	reader := bytes.NewReader(fileByte)
	f, err := excelize.OpenReader(reader)
	if err != nil {
		fmt.Println("err 2==>", err)
		return
	}
	defer func() {
		f.Close()
	}()

	rows, err := f.GetRows("song001")
	if err != nil {
		fmt.Println("err 3==>", err)
		return
	}
	index := 0
	for k, row := range rows {
		if k == 0 {
			continue
		}
		if k < 9 {
			fmt.Println("====>", fmt.Sprintf("%s:%s:%s:%s", row[0], row[1], row[2], row[3]))
		}
		index = k
	}
	fmt.Println("===index", index)
}
