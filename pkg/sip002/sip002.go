package sip002

import (
	"encoding/base64"
	"errors"
	"log"
	"net/url"
	"strconv"
	"strings"
)

func ParseSIP002(sip002data []byte) []*ShadowsocksConfig {
	return ParseSIP002ShadowsocksURIs(ParseSIP002ToShadowsocksURIStrings(sip002data))
}

func ParseSIP002ToShadowsocksURIStrings(base64data []byte) []string {
	data, err := base64.StdEncoding.DecodeString(string(base64data))
	if err != nil {
		return nil
	}
	dataStr := strings.TrimSpace(string(data))
	ssUrlStrList := strings.Split(dataStr, "\n")
	return ssUrlStrList
}

func ParseSIP002ShadowsocksURIs(SIP002ShadowsocksURIStrList []string) []*ShadowsocksConfig {

	var retConfigList []*ShadowsocksConfig

	for _, ssUriStr := range SIP002ShadowsocksURIStrList {
		if cfg, err := NewShadowsocksConfigFromSIP002URI(ssUriStr); err == nil {
			retConfigList = append(retConfigList, cfg)
		}
	}
	return retConfigList
}

func NewShadowsocksConfigFromSIP002URI(ssUriStr string) (*ShadowsocksConfig, error) {

	retConfig := new(ShadowsocksConfig)

	// p判断是否为合法 URL
	ssEncodedUrl, err := url.Parse(ssUriStr)
	if err != nil {
		log.Fatalln(err.Error())
		return nil, err
	}

	// 获取用户信息
	// SIP002 格式的 URL，只有用户名，格式为 websafe-base64-encode-utf8(method  ":" password)
	userInfoStr := ssEncodedUrl.User.Username()
	userData, err := base64.RawURLEncoding.DecodeString(userInfoStr)
	if err != nil {
		log.Fatalln(err.Error())
		return nil, err
	}

	// 解析 websafe-base64-encode-utf8(method  ":" password)
	userInfoStr = string(userData)
	methodStr, password, _ := strings.Cut(userInfoStr, ":")

	retConfig.Method = methodStr
	retConfig.Password = password

	// 解析主机、端口
	retConfig.Server = ssEncodedUrl.Hostname()
	nPort, err := strconv.Atoi(ssEncodedUrl.Port())
	if err != nil {
		nPort = 0
	}
	retConfig.ServerPort = nPort

	// 获取 Plugin 信息
	pluginInfoValues := ssEncodedUrl.Query()["plugin"]
	if len(pluginInfoValues) < 1 {
		return nil, errors.New("not enough value")
	}
	pluginInfoStr := pluginInfoValues[0]

	pluginName, pluginOpts, _ := strings.Cut(pluginInfoStr, ";")
	retConfig.Plugin = pluginName
	retConfig.PluginOpts = pluginOpts

	retConfig.Remarks = ssEncodedUrl.Fragment

	return retConfig, nil
}
