package main

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"sip002parser/cmd/app/config"
	"sip002parser/pkg/gost"
	"sip002parser/pkg/sip002"
	"sort"
	"strings"
)

func main() {

	// 设置日志格式，显示代码行号
	log.SetFlags(log.Lshortfile | log.LstdFlags)

	// 解析命令行参数
	config.ParseConfig()

	// 判断数据源是否为空
	if config.AppConfig.DataSource == "" {
		log.Fatalln("Data source cannot be empty.")
	}

	var data []byte
	var err error

	// 判断数据源是网络请求还是本地文件
	// 如果开头是 http:// 或者 https:// 则认为是网络请求
	if strings.HasPrefix(config.AppConfig.DataSource, "http://") || strings.HasPrefix(config.AppConfig.DataSource, "https://") {
		var resp *http.Response
		resp, err = http.DefaultClient.Get(config.AppConfig.DataSource)
		if err != nil {
			log.Fatalln(err.Error())
		}
		data, err = io.ReadAll(resp.Body)
	} else { // 处理本地文件
		data, err = os.ReadFile(config.AppConfig.DataSource)
		if err != nil {
			log.Fatalln(err.Error())
		}
	}

	// Shadowsocks 配置的集合
	var ssCfgList []*sip002.ShadowsocksConfig

	// 处理 SIP002 格式数据
	if config.AppConfig.InputTypeIsSIP002() {
		ssCfgList = sip002.ParseSIP002(data)
		//log.Println(len(ssCfgList))
	}

	// 处理 JSON 格式的 Shadowsocks 设置
	if config.AppConfig.InputTypeIsShadowsocks() {
		if err = json.Unmarshal(data, &ssCfgList); err != nil {
			log.Fatalln(err.Error())
		}
	}

	// 排序：按 主机:端口 规则排序
	sort.Slice(ssCfgList, func(i, j int) bool {
		ssCfg1 := ssCfgList[i]
		ssCfg2 := ssCfgList[j]
		sHostAndPort1 := fmt.Sprintf("%s:%d", ssCfg1.Server, ssCfg1.ServerPort)
		sHostAndPort2 := fmt.Sprintf("%s:%d", ssCfg2.Server, ssCfg2.ServerPort)
		return sHostAndPort1 < sHostAndPort2
	})
	//log.Println(ssCfgList)

	if config.AppConfig.OutputTypeIsGost() {
		for _, shadowsocksConfig := range ssCfgList {
			fmt.Println(gost.MakeGOSTCommandLine(config.AppConfig.Port, *shadowsocksConfig))
		}
	}

	if config.AppConfig.OutputTypeIsGostLoadBalancing() {
		gost.GenerateGostLoadBalancingConfig(ssCfgList)
	}

	if config.AppConfig.OutputTypeIsSurge() {
		for _, shadowsocksConfig := range ssCfgList {
			fmt.Println(sip002.MakeSurgeProxyConfig(config.AppConfig.TFO, *shadowsocksConfig))
		}
	}

}
