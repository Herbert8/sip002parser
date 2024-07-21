package singbox

import (
	"encoding/json"
	"log"
	"sip002parser/cmd/app/config"
	"sip002parser/pkg/sip002"
)

func CreateLogConfig() LogConfig {
	return LogConfig{
		Disabled:  false,
		Level:     "info",
		Output:    "box.log",
		Timestamp: true,
	}
}

func CreateMixedInbound() MixedInbound {
	return MixedInbound{
		BaseBound: BaseBound{
			Type: "mixed",
			Tag:  "mixed-in",
		},
		Listen:         "::",
		ListenPort:     config.AppConfig.Port,
		Users:          nil,
		SetSystemProxy: false,
	}
}

func CreateShadowsocksObfsOutbound(ssConfig sip002.ShadowsocksConfig) ShadowsocksObfsOutbound {
	return ShadowsocksObfsOutbound{
		BaseBound: BaseBound{
			Type: "shadowsocks",
			Tag:  ssConfig.Remarks,
		},
		Server:     ssConfig.Server,
		ServerPort: ssConfig.ServerPort,
		Method:     ssConfig.Method,
		Password:   ssConfig.Password,
		Plugin:     ssConfig.Plugin,
		PluginOpts: ssConfig.PluginOpts,
	}
}

func CreateShadowsocksObfsOutboundList(ssConfigList []sip002.ShadowsocksConfig) []Bound {
	var retShadowsocksObfsOutboundList []Bound
	for _, shadowsocksConfig := range ssConfigList {
		outboundCfg := CreateShadowsocksObfsOutbound(shadowsocksConfig)
		retShadowsocksObfsOutboundList = append(retShadowsocksObfsOutboundList, outboundCfg)
	}
	return retShadowsocksObfsOutboundList
}

func outboundListToTagList(outboundList []Bound) []string {
	var outboundTagList []string
	for _, outbound := range outboundList {
		outboundTagList = append(outboundTagList, outbound.GetTag())
	}
	return outboundTagList
}

func CreateSelectorOutbound(outbounds []Bound) SelectorOutbound {
	selectorOutbound := SelectorOutbound{
		BaseBound: BaseBound{
			Type: "selector",
			Tag:  "select",
		},
		InterruptExistConnections: false,
	}
	selectorOutbound.Outbounds = outboundListToTagList(outbounds)
	if len(selectorOutbound.Outbounds) > 0 {
		selectorOutbound.Default = selectorOutbound.Outbounds[0]
	}
	return selectorOutbound
}

func CreateURLTestOutbound(outbounds []Bound) URLTestOutbound {
	return URLTestOutbound{
		BaseBound: BaseBound{
			Type: "urltest",
			Tag:  "auto",
		},
		Outbounds: outboundListToTagList(outbounds),
		// The URL to test. https://www.gstatic.com/generate_204 will be used if empty.
		Url: "http://google.com",
		// The test interval. 3m will be used if empty.
		Interval: "3m",
		// The test tolerance in milliseconds. 50 will be used if empty.
		Tolerance: 50,
		// The idle timeout. 30m will be used if empty.
		IdleTimeout: "30m",
		// Interrupt existing connections when the selected outbound has changed.
		InterruptExistConnections: false,
	}
}

func CreateFullConfig(ssConfigList []sip002.ShadowsocksConfig) Config {
	// 日志配置字段
	logCfg := CreateLogConfig()
	// Inbound 的 Mixed
	mixedInboundCfgList := []Bound{CreateMixedInbound()}
	// Outbound 中的 SS
	ssObfsOutboundCfgList := []Bound(CreateShadowsocksObfsOutboundList(ssConfigList))
	// Outbound 中的 URLTest
	urlTestOutboundCfg := CreateURLTestOutbound(ssObfsOutboundCfgList)
	var allOutbounds []Bound
	// 将 URLTest 的 Outbound 作为第一项，默认使用
	allOutbounds = append(allOutbounds, urlTestOutboundCfg)
	// 添加其他 Outbound
	allOutbounds = append(allOutbounds, ssObfsOutboundCfgList...)
	return Config{
		Log:       logCfg,
		Inbounds:  mixedInboundCfgList,
		Outbounds: allOutbounds,
	}
}

func CreateFullConfigJSON(ssConfigList []sip002.ShadowsocksConfig) string {
	fullCfg := CreateFullConfig(ssConfigList)
	if data, err := json.MarshalIndent(fullCfg, "", "    "); err == nil {
		return string(data)
	} else {
		log.Fatalln(err)
	}
	return ""
}
