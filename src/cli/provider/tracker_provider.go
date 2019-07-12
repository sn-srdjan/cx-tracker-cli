package provider

import (
	"fmt"
	"io/ioutil"
	"os"

	"github.com/sn-srdjan/cx-tracker-cli/src/api"
)

// TrackerProvider - provider object
type TrackerProvider struct {
	ServiceURL string
	apiClient  *api.Client
}

// DefaultCxTrackerURL - default cx tracker url
const DefaultCxTrackerURL = "https://tracker.skycoin.net"

// SaveToTrackerService - persist config on tracker service
func (t *TrackerProvider) SaveToTrackerService(configFilePath string) error {
	t.init()
	if _, err := os.Stat(configFilePath); os.IsNotExist(err) {
		return fmt.Errorf("config %s doesn't exist", configFilePath)
	}

	bs, err := ioutil.ReadFile(configFilePath)
	if err != nil {
		return fmt.Errorf("error while reading config %s", configFilePath)
	}

	if err := t.apiClient.Post(t.ServiceURL, bs, nil); err != nil {
		return fmt.Errorf("error while persisting config %s on service %s", configFilePath, t.ServiceURL)
	}

	return nil
}

func (t *TrackerProvider) init() {
	if len(t.ServiceURL) == 0 {
		t.ServiceURL = DefaultCxTrackerURL
	}

	if t.apiClient == nil {
		t.apiClient = api.NewClient(t.ServiceURL)
	}
}
