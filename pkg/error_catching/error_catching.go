package error_catching

import (
	"log"
	"net"
	"regexp"
)

func CheckPort(port string) bool {
	ln, err := net.Listen("tcp", port)

	if err != nil {
		return false
	}

	_ = ln.Close()
	return true
}

func CheckPattern(element string) bool {
	pattern := ".*:::.*"

	res, err := regexp.MatchString(pattern, element)

	if err != nil {
		log.Println(err)
		return false
	}

	return res
}
