package format

import "regexp"

const IdLength = 16

//noinspection GoUnusedGlobalVariable
var RegexpId = regexp.MustCompile("^[a-f0-9]{32}$")

//noinspection GoUnusedGlobalVariable
var UserNicknameRegexp = regexp.MustCompile("^[A-zА-яЁё0-9 ]{1,30}$")
