package cli

import (
	"fmt"
	"os"
	"testing"

	"github.com/stretchr/testify/require"

	"github.com/skycoin/skycoin/src/testutil"
)

func Example() {
	// In cmd/cli/cli.go:
	cfg, err := LoadConfig()
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	cli, err := NewCLI(cfg)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	if err := cli.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

func TestLoadConfig(t *testing.T) {
	t.Run("set CX_TRACKER_URL", func(t *testing.T) {
		val := "http://111.22.33.44:5555"
		os.Setenv("CX_TRACKER_URL", val)
		defer os.Unsetenv("CX_TRACKER_URL")

		cfg, err := LoadConfig()
		require.NoError(t, err)
		require.Equal(t, cfg.TrackerURL, val)
	})

	t.Run("set CX_TRACKER_URL invalid", func(t *testing.T) {
		val := "111.22.33.44:5555"
		os.Setenv("CX_TRACKER_URL", val)
		defer os.Unsetenv("CX_TRACKER_URL")

		_, err := LoadConfig()
		testutil.RequireError(t, err, "CX_TRACKER_URL must be in scheme://host format")
	})
}
