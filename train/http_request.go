package train

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

type adx_report_param_map struct {
	Report_type string `json:"report_type"`
	Dimension   string `json:"dimension"`
	Start_date  int    `json:"start_date"`
	End_date    int    `json:"end_date"`
	Data_source int    `json:"data_source"`
	Group_type  int    `json:"group_type"`
}

type adx_report_data_map struct {
	Code    string `json:"code"`
	Message string `json:"message"`
	Data    struct {
		Total     int      `json:"total"`
		Max_total int      `json:"max_total"`
		Sdk_list  []string `json:"sdk_list"`
		Size_list []string `json:"size_list"`
		Report    map[string][]struct {
			Bid_ecpm float32 `json:"bid_ecpm"`
		} `json:"report"`
	} `json:"data"`
}

// 返回参数格式化存储
var adx_report_data adx_report_data_map

func Post() (res []byte, err error) {
	client := &http.Client{}

	//设置请求体，json
	//song := make(map[string]string)
	//song["mldm"] = "01"
	//bytesData, _ := json.Marshal(song)

	adx_params := adx_report_param_map{"dsp", "day", 1792028800, 1792633599, 1, 1}
	adx_json, _ := json.Marshal(adx_params)

	req, err := http.NewRequest("POST", "http://saas.adx.com/api/v1/report/dsp/data", bytes.NewBuffer([]byte(adx_json)))
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
