package main

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"sip002parser/cmd/app/config"
	"sip002parser/pkg/sip002"
	"strings"
)

func main() {

	config.ParseConfig()

	if config.AppConfig.DataSource == "" {
		log.Fatalln("Data source cannot be empty.")
	}

	var data []byte
	var err error

	if strings.HasPrefix(config.AppConfig.DataSource, "http://") || strings.HasPrefix(config.AppConfig.DataSource, "https://") {
		var resp *http.Response
		resp, err = http.DefaultClient.Get(config.AppConfig.DataSource)
		if err != nil {
			log.Fatalln(err.Error())
		}
		data, err = io.ReadAll(resp.Body)
	} else {
		data, err = os.ReadFile(config.AppConfig.DataSource)
		if err != nil {
			log.Fatalln(err.Error())
		}
	}

	var ssCfgList []*sip002.ShadowsocksConfig

	if config.AppConfig.InputTypeIsSIP002() {
		ssCfgList = sip002.ParseSIP002(data)
	}

	if config.AppConfig.InputTypeIsShadowsocks() {
		if err = json.Unmarshal(data, &ssCfgList); err != nil {
			log.Fatalln(err.Error())
		}
	}

	//log.Println(ssCfgList)

	if config.AppConfig.OutputTypeIsGost() {
		for _, shadowsocksConfig := range ssCfgList {
			fmt.Println(sip002.MakeGOSTCommandLine(config.AppConfig.Port, *shadowsocksConfig))
		}
	}

	if config.AppConfig.OutputTypeIsSurge() {
		for _, shadowsocksConfig := range ssCfgList {
			fmt.Println(sip002.MakeSurgeProxyConfig(config.AppConfig.TFO, *shadowsocksConfig))
		}
	}
}
