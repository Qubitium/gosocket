package gosocket

import (
	"net"
	"strings"
)

func GetTestResult() (s string) {
	const ERR = "protocol not available"

	_, err := net.Dial("tcp", "127.0.0.1:80")
	if strings.Contains(err.Error(), ERR) {
		s += "MISSING TCP: " + err.Error() + "\n\n"
	} else {
		s += "GOT TCP! " + err.Error() + "\n\n"
	}

	_, err = net.Dial("tcp4", "127.0.0.1:80")
	if strings.Contains(err.Error(), ERR) {
		s += "MISSING TCP4: " + err.Error() + "\n\n"

	} else {
		s += "GOT TCP4! " + err.Error() + "\n\n"
	}

	_, err = net.Dial("tcp6", "127.0.0.1:80")
	if strings.Contains(err.Error(), ERR) {
		s += "MISSING TCP6: " + err.Error() + "\n\n"

	} else {
		s += "GOT TCP6! " + err.Error() + "\n\n"
	}

	_, err = net.Dial("udp", "127.0.0.1:80")
	if strings.Contains(err.Error(), ERR) {
		s += "MISSING UDP: " + err.Error() + "\n\n"

	} else {
		s += "GOT UDP! " + err.Error() + "\n\n"
	}

	return

}
