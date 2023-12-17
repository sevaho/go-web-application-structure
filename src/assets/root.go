package assets

import "embed"

//go:embed all:templates all:public
var Assets embed.FS

