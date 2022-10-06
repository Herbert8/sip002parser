package sip002

import (
	"fmt"
)

func MakeGOSTCommandLine(nPort uint, ssConfig ShadowsocksConfig) string {
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
