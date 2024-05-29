package localizable

import _ "embed"

var(
	//go:embed en.json
	EN []byte
	//go:embed zh.json
	ZH []byte
)
