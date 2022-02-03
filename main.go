package CheckPWD

import "regexp"

func CheckPWD(pwd string) bool {
	match, _ := regexp.MatchString("((?=.*/d)(?=.*[A-Z])(?=.*/W).{8,8}) ", pwd)
	return match
}
