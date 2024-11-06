package common

import (
	"encoding/base64"
	"errors"
	"net/url"
	"regexp"
	"strings"
)

// urlEncode
func UrlEscape(str string) string {
	return url.QueryEscape(str)
}

// urlPathEncode
func UrlPathEscape(str string) string {
	return url.PathEscape(str)
}

// urlUncode
func UrlUnescape(str string) string {
	unescape, err := url.QueryUnescape(str)
	if err != nil {
		return ""
	}
	return unescape
}

// base64解码
func Base64URLDecode(data string) ([]byte, error) {
	var missing = (4 - len(data)%4) % 4
	data += strings.Repeat("=", missing)
	res, err := base64.URLEncoding.DecodeString(data)
	if err != nil {
		return nil, err
	}
	return res, nil
}

// 安全Base64编码
func Base64UrlSafeEncode(source []byte) string {
	// Base64 Url Safe is the same as Base64 but does not contain '/' and '+' (replaced by '_' and '-') and trailing '=' are removed.
	bytearr := base64.StdEncoding.EncodeToString(source)
	safeurl := strings.Replace(string(bytearr), "/", "_", -1)
	safeurl = strings.Replace(safeurl, "+", "-", -1)
	safeurl = strings.Replace(safeurl, "=", "", -1)
	return safeurl
}

// 域名解析主域名
func GetHostMainDomain(urlString string) (string, error) {
	if len(urlString) == 0 {
		return "", nil
	}

	if !strings.HasPrefix(urlString, "http") {
		urlString = "https://" + urlString
	}
	urlData, err := url.Parse(urlString)
	if err != nil {
		return "", err
	}

	mainDomain := strings.ToLower(urlData.Host)

	var domainRegex = regexp.MustCompile(`^[a-zA-Z0-9]+([\-\.]{1}[a-zA-Z0-9]+)*\.[a-zA-Z]{2,}$`)
	if !domainRegex.MatchString(mainDomain) {
		return "", errors.New("failed match mainDomain")
	}

	if len(mainDomain) > 4 && strings.HasPrefix(mainDomain, "www.") {
		mainDomain = mainDomain[4:]
	}

	return mainDomain, nil
}
