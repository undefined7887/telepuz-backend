package common

import "regexp"

const IdLength = 16

//noinspection GoUnusedGlobalVariable
var IdRegexp = regexp.MustCompile("^[a-f0-9]{32}$")
