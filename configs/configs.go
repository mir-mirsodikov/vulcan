package configs

import "embed"

//go:embed *.yaml
var ConfigFiles embed.FS
