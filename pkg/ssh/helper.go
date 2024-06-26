package scp

import (
	"bytes"
	"fmt"
	"os/exec"
)

func Execute(host, username, password, command string) (string, error) {
	if host == "" || username == "" || password == "" || command == "" {
		return "", fmt.Errorf("invalid arguments")
	}

	cmd := exec.Command("sshpass", "-p", password,
		"ssh", "-o", "StrictHostKeyChecking=accept-new",
		fmt.Sprintf("%s@%s", username, host), command)
	var outBuf, errBuf bytes.Buffer
	cmd.Stdout = &outBuf
	cmd.Stderr = &errBuf

	err := cmd.Run()
	if err != nil {
		return "", fmt.Errorf("failed to execute ssh command (%s): %w : %s", command, err, errBuf.String())
	}

	return outBuf.String(), nil
}
