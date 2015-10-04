package punch

import (
	"bytes"
	"fmt"
	"net"
	"os"
	"os/exec"
)

const (
	my_chain = "poe"
)

// wrapper os.exec
func run(args ...string) (string, string, error) {
	cmd := exec.Command("iptables", args...)

	var stdout bytes.Buffer
	var stderr bytes.Buffer
	cmd.Stdout = &stdout
	cmd.Stderr = &stderr

	err := cmd.Run()

	if err != nil {
		fmt.Fprintf(os.Stderr, "os.exec error : %s : %s", err, stderr.String())
	}

	return stdout.String(), stderr.String(), err
}

func Setup() {
	run("-t", "nat", "-N", my_chain)
	run("-t", "nat", "-I", "PREROUTING", "-j", my_chain)
}

func Allocate(targetIPv4addr string, targetPort string, boundPort string) {
	run("-t", "nat", "-I", my_chain, "-p", "tcp", "--dport", boundPort, "-j", "DNAT", "--to-destination", net.JoinHostPort(targetIPv4addr, targetPort))

}

func Release(targetIPv4addr string, targetPort string, boundPort string) {
	run("-t", "nat", "-D", my_chain, "-p", "tcp", "--dport", boundPort, "-j", "DNAT", "--to-destination", net.JoinHostPort(targetIPv4addr, targetPort))
}
