package scp

import (
	"os"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestCopy(t *testing.T) {
	output, err := Copy("192.168.4.58", "mccxadmin", "mccxadmin", "/etc/docker/certs.d/dockerhub.kubekey.local/ca.crt", ".")
	require.NoError(t, err)
	require.Empty(t, output)
	_, err = os.Stat("ca.crt")
	require.NoError(t, err)
	// Clean up
	os.Remove("ca.crt")
}
