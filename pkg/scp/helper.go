package scp

import (
	"bytes"
	"fmt"
	"os/exec"
)

func Copy(host, username, password, src, dest string) (string, error) {
	if host == "" || username == "" || password == "" || src == "" || dest == "" {
		return "", fmt.Errorf("invalid arguments")
	}

	cmd := exec.Command("sshpass", "-p", password,
		"scp", "-o", "StrictHostKeyChecking=accept-new",
		fmt.Sprintf("%s@%s:%s", username, host, src), dest)
	var outBuf, errBuf bytes.Buffer
	cmd.Stdout = &outBuf
	cmd.Stderr = &errBuf

	err := cmd.Run()
	if err != nil {
		return "", fmt.Errorf("failed to execute scp command: %w : %s", err, errBuf.String())
	}

	return outBuf.String(), nil
}
