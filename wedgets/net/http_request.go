package net

import (
	"bytes"
	"compress/gzip"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"net/http"
	"strconv"
	"strings"

	"google.golang.org/protobuf/proto"
	"google.golang.org/protobuf/reflect/protoreflect"
)

const (
	HttpEncodingGzip = "gzip" //http数据传输编码类型：gzip

	AllowDebugLog = true

	METHOD_TYPE_POST = "POST"
	METHOD_TYPE_GET  = "GET"

	BODY_TYPE_STRUCT = "JSON"
	BODY_TYPE_PROTO  = "PROTO"
	BODY_TYPE_BYTE   = "BYTE"
)

func basicHttpRequest(methodType string, bodyType string, request interface{}, url string, headers map[string]string, httpClient *http.Client) (*http.Response, []byte, error) {
	var (
		response  *http.Response
		bytesData []byte
		body      []byte
		err       error
	)

	//POST 请求
	if methodType == METHOD_TYPE_POST {

		switch bodyType {
		case BODY_TYPE_STRUCT:
			bytesData, err = json.Marshal(request)
		case BODY_TYPE_PROTO:
			bytesData, err = proto.Marshal(request.(protoreflect.ProtoMessage))
		case BODY_TYPE_BYTE:
			bytesData = request.([]byte)
		}

		if err != nil {
			return nil, nil, errors.New(fmt.Sprintf("json Marshal error='%s'", err.Error()))
		}

		if AllowDebugLog {
			if bodyType == BODY_TYPE_PROTO {
				var bytesDataJson []byte
				bytesDataJson, err = json.Marshal(request)
				if err != nil {
					return nil, nil, errors.New(fmt.Sprintf("json Marshal error='%s'", err.Error()))
				}
				log.Printf("sendRequest of POST: url = %v, headers =%v, request = %v", url, headers, string(bytesDataJson))
			} else {
				log.Printf("sendRequest of POST: url = %v, headers =%v, request = %v", url, headers, string(bytesData))
			}
		}

		response, body, err = HttpPostRequest(url, bytesData, headers, httpClient)
	}

	//GET 请求
	if methodType == METHOD_TYPE_GET {
		if AllowDebugLog {
			log.Printf("sendRequest of GET: headers =%v, url = %v", headers, url)
		}
		response, body, err = HttpGetRequest(url, headers, httpClient)
	}

	if err != nil {
		if strings.Contains(err.Error(), "Client.Timeout") {
			return nil, nil, errors.New("504")
		} else {
			return nil, nil, errors.New(err.Error())
		}
	}

	if AllowDebugLog {
		log.Printf("getResponse : header = %v, response = %v", response.Header, string(body))
	}

	statusCode := response.StatusCode
	if statusCode != http.StatusOK {
		return nil, nil, errors.New(strconv.Itoa(statusCode))
	}

	return response, body, nil
}

func HttpPostRequest(apiUrl string, bytesData []byte, headers map[string]string, httpClient *http.Client) (*http.Response, []byte, error) {
	var (
		err    error
		buffer *bytes.Buffer
	)

	//gzip压缩
	if encoding, ok := headers["Accept-Encoding"]; ok && encoding == HttpEncodingGzip {
		buf := BufferPool.Get()
		zw := GzipWriterPool.Get()
		zw.Reset(buf)
		defer func() {
			BufferPool.Put(buf)
			GzipWriterPool.Put(zw)
		}()
		if _, err = zw.Write(bytesData); err != nil {
			return nil, []byte{}, errors.New(fmt.Sprintf("zw.Write error='%s'", err.Error()))
		}
		if err = zw.Flush(); err != nil {
			return nil, []byte{}, errors.New(fmt.Sprintf("zw.Flush error='%s'", err.Error()))
		}
		if err = zw.Close(); err != nil {
			return nil, []byte{}, errors.New(fmt.Sprintf("zw.Close error='%s'", err.Error()))
		}
		buffer = buf
	} else {
		buffer = bytes.NewBuffer(bytesData)
	}

	request, err := http.NewRequest("POST", apiUrl, buffer)
	if err != nil {
		return nil, []byte{}, errors.New(fmt.Sprintf("newRequest error='%s'", err))
	}

	if len(headers) > 0 {
		for key, item := range headers {
			request.Header.Set(key, item)
		}
	}

	response, err := httpClient.Do(request)
	if err != nil {
		return nil, []byte{}, errors.New(fmt.Sprintf("clientDo error='%s'", err))
	}
	defer response.Body.Close()

	var body = response.Body
	if response.Header.Get("Content-Encoding") == HttpEncodingGzip {
		body2, e := gzip.NewReader(response.Body)
		if e == nil {
			body = body2
		} else if e != io.EOF {
			return nil, []byte{}, errors.New(fmt.Sprintf("unzip error='%s'", e))
		}
	}

	data, err := ioutil.ReadAll(body)
	io.Copy(ioutil.Discard, response.Body)
	if err != nil {
		return nil, []byte{}, errors.New(fmt.Sprintf("ioutilReadAll error='%s'", err))
	}

	return response, data, nil
}

func HttpGetRequest(apiUrl string, headers map[string]string, httpClient *http.Client) (*http.Response, []byte, error) {

	request, err := http.NewRequest("GET", apiUrl, nil)
	if err != nil {
		return nil, []byte{}, errors.New(fmt.Sprintf("newRequest error='%s'", err))
	}

	if len(headers) > 0 {
		for key, item := range headers {
			request.Header.Set(key, item)
		}
	}

	response, err := httpClient.Do(request)
	if err != nil {
		return nil, []byte{}, errors.New(fmt.Sprintf("clientDo error='%s'", err))
	}
	defer response.Body.Close()

	var body = response.Body
	if response.Header.Get("Content-Encoding") == HttpEncodingGzip {
		body2, e := gzip.NewReader(response.Body)
		if e == nil {
			body = body2
		} else if e != io.EOF {
			return nil, []byte{}, errors.New(fmt.Sprintf("unzip error='%s'", e))
		}
	}

	data, err := ioutil.ReadAll(body)
	io.Copy(ioutil.Discard, response.Body)
	if err != nil {
		return nil, []byte{}, errors.New(fmt.Sprintf("ioutilReadAll error='%s'", err))
	}

	return response, data, nil
}
