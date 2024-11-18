package read

import _ "embed"

var (
	//go:embed queries/read.sql
	readQuery string
)
