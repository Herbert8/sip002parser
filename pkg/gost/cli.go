package gost

import (
	"fmt"
	"sip002parser/pkg/sip002"
)

func MakeGOSTCommandLine(nPort uint, ssConfig sip002.ShadowsocksConfig) string {
	const sTemplate = `gost -L :%d -F "ss+ohttp://%s:%s@%s:%d?host=%s"`

	return fmt.Sprintf(sTemplate,
		nPort,
		ssConfig.EscapedMethod(),
		ssConfig.EscapedPassword(),
		ssConfig.Server,
		ssConfig.ServerPort,
		ssConfig.EscapedObfsHost(),
	)
}
