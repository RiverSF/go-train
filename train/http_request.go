package train

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

func Post() (request interface{}, err error) {
	client := &http.Client{}

	//设置请求体，json
	//song := make(map[string]string)
	//song["mldm"] = "01"
	//bytesData, _ := json.Marshal(song)

	body, _ := json.Marshal(request)

	req, err := http.NewRequest("POST", "http://saas.adx.com/api/v1/report/dsp/data", bytes.NewBuffer(body))
	if err != nil {
		return nil, fmt.Errorf("post error：【%s】", err.Error())
	}

	//req.Header.Set("content-type", "multipart/form-data")
	req.Header.Set("content-type", "application/json")
	req.Header.Set("Authorization", "Bearer eyJ0eXAiOiJKV1QiLCJhbGciOiJIUzI1NiJ9.eyJpc3MiOiJTQUFTX0FEWCIsInN1YiI6MSwiaWF0IjoxNjkyNzcxMTczLCJleHAiOjE2OTI4NTc1NzMsInNjb3BlcyI6ImFjY2Vzc190b2tlbiJ9.VkmOOi5iLq3n0LQtn5CsWZF_LlaV-YdcHSt2YX5WVb0")
	req.Header.Set("User-Agent", "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/115.0.0.0 Safari/537.36")

	resp, err := client.Do(req)
	defer resp.Body.Close()

	if resp.StatusCode != 200 {
		return nil, fmt.Errorf("body error: 【%s】", err.Error())
	}

	content, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("read error:【%s】", err.Error())
	}

	//str := *(*string)(unsafe.Pointer(&content))
	//fmt.Println(str)

	//返回参数格式化存储
	//json.Unmarshal(content, &adx_report_data)

	return content, nil
}

func Get() (res []byte, err error) {
	client := &http.Client{}

	url := "http://saas.adx.com/api/v1/dsp/account/filter/list"
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return nil, fmt.Errorf("get error: 【%s】", err.Error())
	}

	req.Header.Add("Authorization", "Bearer eyJ0eXAiOiJKV1QiLCJhbGciOiJIUzI1NiJ9.eyJpc3MiOiJTQUFTX0FEWCIsInN1YiI6MSwiaWF0IjoxNjkyNzcxMTczLCJleHAiOjE2OTI4NTc1NzMsInNjb3BlcyI6ImFjY2Vzc190b2tlbiJ9.VkmOOi5iLq3n0LQtn5CsWZF_LlaV-YdcHSt2YX5WVb0")

	resp, err := client.Do(req)
	defer resp.Body.Close()

	if err != nil {
		return nil, fmt.Errorf("resp error:【%s】", err.Error())
	}

	content, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("read error:【%s】", err.Error())
	}

	//fmt.Println(string(content))

	return content, nil
}
