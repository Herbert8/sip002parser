package gost

import (
	"fmt"
	"log"
	"os"
	"sip002parser/cmd/app/config"
	"sip002parser/pkg/sip002"
	"text/template"
)

/*
此文件的目的是让 GOST 可以针对 SIP002 订阅进行处理。
SIP002 订阅中包含多个连接，这里的目的是把多个连接配置成 GOST 的负载均衡。
参见：https://v2.gost.run/load-balancing/

根据已有数据的情况看，SIP002 的订阅中，除了主机和端口，其他内容相同，这样的话，只要在 peer 文件中定义一个 peer，为这个 peer 定义一个 IP 列表
IP 列表中的内容可以是 IP:端口 或者 域名:端口

### 对于实际情况的处理
但实际情况中并不确定订阅中每一项的配置值是否一致。比如：是否存在以下情况：
- 相同的 主机:端口 对应不同的用户名和密码
- 不同的 主机:端口 对应不同的配置

对于这种情况，不同的配置（不包括主机和端口）对于 peer 文件中不同的 peer 项，可以为每个 peer 指定不同的 主机列表

考虑通过 Map 处理类似结构
Map 的 Key 为各种属性的集合，值为 IP:端口 的集合
*/

// 模板字符串
const peerTemplateText = `
strategy        random
max_fails       1
fail_timeout    30s

reload          10s

# peers
{{/* #peer   ss+ohttp://aes-128-gcm:6a86b460f2393df7@[host]?host=56786678.microsoft.com&ip=ips.txt */}}
{{range $index, $value := . -}}
peer   ss+ohttp://{{$value.Method}}:{{$value.Password}}@[host]?host={{$value.ObfsHost}}&ip=ips{{ if gt $index 0 }}{{ $index }}{{end}}.txt
{{end}}
`

// 创建模板
var peerTemplate = template.Must(template.New("peer_template").Funcs(template.FuncMap{
	"inc": func(i int) int {
		return i + 1
	},
}).Parse(peerTemplateText))

// 定义 用于存放 peer 和 IP 关系的 Map 的结构
type peerMap map[ssConfigForKey][]sip002.ShadowsocksConfig

// 用于存放 peer 和 IP 关系的 Map
var peerMapData peerMap

// 将 Config Item 添加到 Map 中
func addItem(key ssConfigForKey, value sip002.ShadowsocksConfig) {
	valSlice, ok := peerMapData[key]
	if !ok {
		valSlice = make([]sip002.ShadowsocksConfig, 0)
	}
	valSlice = append(valSlice, value)
	peerMapData[key] = valSlice
}

// 用于作为 Key 的结构体
// 类似 sip002.ShadowsocksConfig，但忽略掉会放入 ip 文件中的字段
// 这个结构体只存放 peer 中对应的内容
type ssConfigForKey struct {
	sip002.ShadowsocksConfig
	ObfsHost string
}

// 打印文件头
func printFileHeader(filename string) {
	fmt.Printf("===== File: %s =====\n", filename)
}

// 打印文件尾
func printFileFooter(filename string) {
	fmt.Println("=====================")
}

func printJSONSample() {
	printFileHeader("GOST JSON Configuration Sample")
	fmt.Printf(`{
    "Debug": true,
    "Retries": 10,
    "ServeNodes": [],
    "ChainNodes": [],
    "Routes": [
        {
            "ServeNodes": [
                ":%d"
            ],
            "ChainNodes": [
                "?peer=peers.txt"
            ]
        }
    ]
}
`, config.AppConfig.Port)
}

func printGostCLISample() {
	printFileHeader("GOST Command Line Sample")
	fmt.Printf("gost -L :%d -F '?peer=peers.txt'\n", config.AppConfig.Port)
}

// GenerateGostLoadBalancingConfig 处理已经加载的 Shadowsocks Config
func GenerateGostLoadBalancingConfig(ssConfigList []*sip002.ShadowsocksConfig) {
	// 创建存储的 Map
	peerMapData = make(peerMap)
	// 遍历加载的 Shadowsocks 并按 peer 和 IP 关系存入 Map
	for _, ssConfig := range ssConfigList {
		// 创建 Key
		ssCfgAsKey := ssConfigForKey{
			ShadowsocksConfig: *ssConfig,
			ObfsHost:          ssConfig.ObfsHost(),
		}
		// 清理掉不需要标识的字段
		// 服务器
		ssCfgAsKey.Server = ""
		// 端口
		ssCfgAsKey.ServerPort = 0
		// 配置中的标识
		ssCfgAsKey.Remarks = ""
		// 添加到 Map
		addItem(ssCfgAsKey, *ssConfig)
	}

	// 将 Map 的 Key 和 Value 分别读入数组
	var keys []ssConfigForKey
	var values [][]sip002.ShadowsocksConfig
	for key, val := range peerMapData {
		keys = append(keys, key)
		values = append(values, val)
	}

	printGostCLISample()
	printJSONSample()

	// 打印 peer 文件头
	printFileHeader("peers.txt")
	// 填充模板并输出 peer
	if err := peerTemplate.Execute(os.Stdout, keys); err != nil {
		log.Fatalln(err)
	}
	// 打印 peer 文件尾
	printFileFooter("")

	// 输出 IP
	for idx, value := range values {
		var filename string
		if idx > 0 {
			filename = fmt.Sprintf("ips%d.txt", idx)
		} else {
			filename = fmt.Sprint("ips.txt")
		}
		printFileHeader(filename)
		for _, config := range value {
			fmt.Printf("%s:%d\n", config.Server, config.ServerPort)
		}
		printFileFooter("")
	}

}
