package controllers

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"
	"strconv"
	"strings"

	beego "github.com/beego/beego/v2/server/web"
	"google.golang.org/protobuf/proto"
	"google.golang.org/protobuf/reflect/protoreflect"

	"api/controllers/protobuf/gaode"
)

type TestProtoPmpController struct {
	beego.Controller
}

func (c *TestProtoPmpController) Post() {

	tag := c.GetString("tag", "")
	requestBody := c.Ctx.Input.RequestBody

	var err error
	var errMsg string
	var request interface{}
	//var url = "http://127.0.0.1/ad/bid?tag=" + tag
	var url = "http://test-api-adx.gmediation.com/gaode/ad/bid"

	//json 转 proto
	switch tag {
	case "gaode":
		request = &gaode.Request{}
		err = json.Unmarshal(requestBody, &request)
	default:
		c.Ctx.WriteString("invalid tag")
		return
	}

	if err != nil {
		errMsg = fmt.Sprintf("failed json to proto request, error=%v", err.Error())
		c.Ctx.WriteString(errMsg)
		return
	}

	headers := make(map[string]string)
	headers["Content-Type"] = "application/x-protobuf"
	client := &http.Client{}

	_, body, nbrError := HttpProtoRequest(request, url, headers, client)
	if nbrError != nil {
		errMsg = fmt.Sprintf("http request url=%v, errorNbr=%v", url, nbrError)
		c.Ctx.WriteString(errMsg)
		return
	}

	// proto 响应转换
	switch tag {
	case "gaode":
		jsonBody := &gaode.Response{}
		proto.Unmarshal(body, jsonBody)
		c.Data["json"] = jsonBody
		c.ServeJSON()
	default:
		c.Ctx.WriteString("invalid tag")
		return
	}

	//// 初始化jsonpb.Marshaler
	//marshaler := &jsonpb.Marshaler{
	//	OrigName:     true,  // 如果需要，可以保留Protobuf字段的原始名称
	//	EmitDefaults: true,  // 发射零值字段
	//	Indent:       "  ",  // 缩进字符串，用于美化输出
	//}
	//// 将Protobuf消息转换为JSON字符串
	//jsonStr, err := marshaler.MarshalToString(jsonBody)
	//if err != nil {
	//	errMsg = fmt.Sprintf("Failed to marshal to JSON: %v", err)
	//	c.Ctx.WriteString(errMsg)
	//	return
	//}
}

func HttpProtoRequest(request interface{}, url string, headers map[string]string, httpClient *http.Client) (*http.Response, []byte, error) {
	var (
		response *http.Response
		body     []byte
		err      error
	)

	var bytesDataProto []byte
	bytesDataProto, err = proto.Marshal(request.(protoreflect.ProtoMessage))
	if err != nil {
		return nil, nil, errors.New(fmt.Sprintf("proto Marshal error='%s'", err.Error()))
	}

	req, e := http.NewRequest("POST", url, bytes.NewBuffer(bytesDataProto))
	if e != nil {
		return nil, nil, errors.New(err.Error())
	}

	req.Header.Set("Content-Type", "application/x-protobuf")

	response, err = httpClient.Do(req)
	defer response.Body.Close()

	body, _ = io.ReadAll(response.Body)

	if err != nil {
		if strings.Contains(err.Error(), "Client.Timeout") {
			return nil, nil, errors.New("Client.Timeout")
		} else {
			return nil, nil, errors.New(err.Error())
		}
	}

	statusCode := response.StatusCode
	if statusCode != http.StatusOK {
		return nil, nil, errors.New(strconv.Itoa(statusCode))
	}

	return response, body, nil
}
