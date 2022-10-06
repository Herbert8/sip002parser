package sip002

import (
	"fmt"
	"net/url"
	"strings"
)

type ShadowsocksConfig struct {
	Server     string `json:"server"`
	ServerPort int    `json:"server_port"`
	Password   string `json:"password"`
	Method     string `json:"method"`
	Plugin     string `json:"plugin"`
	PluginOpts string `json:"plugin_opts"`
	Remarks    string `json:"remarks"`
}

func (receiver *ShadowsocksConfig) String() string {
	const sTemplate = `--->
    Server: %s
    ServerPort: %d
    Password: %s
    Method: %s
    Plugin: %s
    PluginOpts: %s
    Remarks: %s	
<---`
	return fmt.Sprintf(sTemplate,
		receiver.Server,
		receiver.ServerPort,
		receiver.Password,
		receiver.Method,
		receiver.Plugin,
		receiver.PluginOpts,
		receiver.Remarks,
	)
}

type PluginOption map[string]string

func (receiver *ShadowsocksConfig) PluginOptionsMap() PluginOption {
	retOption := make(PluginOption)

	for _, optStr := range strings.Split(receiver.PluginOpts, ";") {
		optKey, optVal, _ := strings.Cut(optStr, "=")
		retOption[optKey] = optVal
	}

	return retOption
}

func (receiver *ShadowsocksConfig) EscapedMethod() string {
	return url.QueryEscape(receiver.Method)
}

func (receiver *ShadowsocksConfig) EscapedPassword() string {
	return url.QueryEscape(receiver.Password)
}

func (receiver *ShadowsocksConfig) ObfsHost() string {
	return receiver.PluginOptionsMap()["obfs-host"]
}

func (receiver *ShadowsocksConfig) EscapedObfsHost() string {
	return url.QueryEscape(receiver.ObfsHost())
}
