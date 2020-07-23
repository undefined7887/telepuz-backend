package format

import "regexp"

//noinspection GoUnusedGlobalVariable
var UserNicknameRegexp = regexp.MustCompile("^[A-zА-яЁё0-9 ]{1,30}$")
