package net

import (
	"net"
	"net/http"
	"sync"
	"time"

	"river/wedgets/common"
)

var (
	HttpClientDefault           *http.Client
	HttpClientDefaultBidTimeout int64

	HttpClient3000 *http.Client

	TimeoutHttpClient *timeoutHttpClient
)

type timeoutHttpClient struct {
	TimeoutHttpClientMap map[int64]*http.Client
	mutex                sync.RWMutex
}

func HttpInit() error {
	HttpClient3000 = createHTTPClient(3000)
	TimeoutHttpClient = &timeoutHttpClient{
		TimeoutHttpClientMap: map[int64]*http.Client{},
	}
	HttpClientDefault, HttpClientDefaultBidTimeout = TimeoutHttpClient.CreateHTTPClient(500)
	return nil
}

func createHTTPClient(bidTimeout int64) *http.Client {
	transport := &http.Transport{
		MaxIdleConns:        100, //Host最大连接数
		MaxConnsPerHost:     10,  //Host最大连接数
		MaxIdleConnsPerHost: 10,  //Host闲时最小连接数
		IdleConnTimeout:     time.Duration(30) * time.Second,
		TLSHandshakeTimeout: 5 * time.Second,

		DialContext: (&net.Dialer{
			Timeout:   30 * time.Second,
			KeepAlive: 30 * time.Second,
		}).DialContext,
	}

	client := &http.Client{
		Transport: transport,
		Timeout:   time.Duration(bidTimeout) * time.Millisecond,
	}
	return client
}

func (t *timeoutHttpClient) CreateHTTPClient(bidTimeout int64) (*http.Client, int64) {
	if bidTimeout <= 0 {
		return HttpClientDefault, HttpClientDefaultBidTimeout
	}

	//向下梯度50，最小100ms
	bidTimeout = common.MaxInt64(bidTimeout-bidTimeout%50, 100)

	t.mutex.RLock()
	if httpClient, ok := t.TimeoutHttpClientMap[bidTimeout]; ok {
		t.mutex.RUnlock()
		return httpClient, bidTimeout
	} else {
		t.mutex.RUnlock()
	}

	t.mutex.Lock()
	defer t.mutex.Unlock()
	newHttpClient := createHTTPClient(bidTimeout)
	t.TimeoutHttpClientMap[bidTimeout] = newHttpClient
	return newHttpClient, bidTimeout
}
