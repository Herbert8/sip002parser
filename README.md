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





