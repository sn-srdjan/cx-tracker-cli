package provider

import (
	"testing"

	"github.com/stretchr/testify/require"
)

const (
	testTrackerURL         = "http://localhost:8083/api/v1/config"
	testConfigFilePath     = "configTest.json"
	testConfigSaveFilePath = "configSaveTest.json"
	testGenesisHash        = "testHash1234"
)

func TestSaveToTrackerService(t *testing.T) {

	provider := TrackerProvider{
		ServiceURL: testTrackerURL,
	}

	actualError := provider.SaveToTrackerService(testConfigFilePath)
	require.Equal(t, nil, actualError)
}

func TestGetConfigFromTrackerService(t *testing.T) {

	provider := TrackerProvider{
		ServiceURL: testTrackerURL,
	}

	actualError := provider.GetConfigFromTrackerService(testGenesisHash, testConfigSaveFilePath)
	require.Equal(t, nil, actualError)
}
