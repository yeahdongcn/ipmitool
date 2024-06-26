package scp

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestExecute(t *testing.T) {
	output, err := Execute("192.168.4.58", "mccxadmin", "mccxadmin", "whoami")
	require.NoError(t, err)
	require.Equal(t, "mccxadmin\n", output)
}
