package sip002

import "fmt"

func MakeSurgeProxyConfig(tfo bool, ssConfig ShadowsocksConfig) string {
	const sTemplate = `%s= ss, %s, %d, encrypt-method=%s, password=%s, obfs=http, obfs-host=%s, udp-relay=true, tfo=%s`

	var tfoValue string
	if tfo {
		tfoValue = "true"
	} else {
		tfoValue = "false"
	}

	return fmt.Sprintf(sTemplate,
		ssConfig.Remarks,
		ssConfig.Server,
		ssConfig.ServerPort,
		ssConfig.Method,
		ssConfig.Password,
		ssConfig.ObfsHost(),
		tfoValue)
}
