package parser

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestParseHosts(t *testing.T) {
	hosts, err := ParseHosts("./testdata/1.csv")
	require.NoError(t, err)
	require.Len(t, hosts, 2)
	for i, host := range hosts {
		if i == 0 {
			require.Equal(t, "worker001", host.Name)
			require.Equal(t, "10.1.48.1", host.BMCIP)
			require.Equal(t, "10.1.0.1", host.ManagementIP)
		} else {
			require.Equal(t, "worker002", host.Name)
			require.Equal(t, "10.1.48.2", host.BMCIP)
			require.Equal(t, "10.1.0.2", host.ManagementIP)
		}
	}

	hosts, err = ParseHosts("./testdata/2.csv")
	require.NoError(t, err)
	require.Len(t, hosts, 0)
}
