package provider

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestSaveToTrackerService(t *testing.T) {

	provider := TrackerProvider{
		ServiceURL: "http://localhost:8083/api/v1/config",
	}

	actualError := provider.SaveToTrackerService("config.json")
	require.Equal(t, nil, actualError)
}
