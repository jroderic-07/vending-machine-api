package error_catching

import (
	"log"
	"net"
	"regexp"
)

func CheckPort(port string) (bool, error) {
	ln, err := net.Listen("tcp", port)

	if err != nil {
		return false, err
	}

	_ = ln.Close()
	return true, err
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
