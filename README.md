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



### 使用方法

```
Usage:
  sip002parser -it <ss | sip002> -ot <gost | surge> [-p <port>] [-tfo] data_source
	  Generate configuration from data.

  sip002parser
	  Show usage information.

Options:
  -it <ss | sip002>                         Specify input type.
  -ot <gost | surge>                        Specify output type.
  -p <port>                                 Specify local proxy port.
  -tfo                                      Specifies to use tfo mode.
```

