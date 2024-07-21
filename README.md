# SIP002 解析器



### 什么是 SIP002？

- [SIP002 URI Scheme](https://github.com/shadowsocks/shadowsocks-org/wiki/SIP002-URI-Scheme)
- [SIP002 - Optional extension configurations as query strings in ss URLs](https://github.com/shadowsocks/shadowsocks-org/issues/27)
- [RFC3986](https://www.ietf.org/rfc/rfc3986.txt)



```
SS-URI = "ss://" userinfo "@" hostname ":" port [ "/" ] [ "?" plugin ] [ "#" tag ]
userinfo = websafe-base64-encode-utf8(method  ":" password)
           method ":" password
```



### 解析器的作用

网络工具较多，很多订阅提供方不能提供每种网络工具的支持。对于一些小众的工具可能也缺乏支持。但有些工具用起来确实方便，所以提供与 SIP002 的转换功能。

- Surge
- GOST
  - 命令行
  - 配置文件
  - 负载均衡模式

- sing-box
  - Mixed Inbound
  - URLTest Outbound




### 使用方法

```
Usage:
  sip002parser -it <ss | sip002> -ot <gost-cli | gost-lb | surge | sing-box> [-p <port>] [-tfo] data_source
	  Generate configuration from data.

  sip002parser
	  Show usage information.

Options:
  -it <ss | sip002>                                   Specify input type.
  -ot <gost-cli | gost-lb | surge | sing-box>         Specify output type.
  -p <port>                                           Specify local proxy port.
  -tfo                                                Specifies to use tfo mode.
```

