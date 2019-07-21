package provider

import (
	"bytes"
	"encoding/json"
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

type cxApplicationConfig struct {
	GenesisHash string `json:"genesisHash"`
}

// DefaultCxTrackerURL - default cx tracker url
const DefaultCxTrackerURL = "https://cx-tracker.skycoin.net/api/v1/config"

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

	r := bytes.NewReader(bs)

	if err := t.apiClient.Put(t.ServiceURL, r, nil); err != nil {
		return fmt.Errorf("error while persisting config %s on service %s due to error: %s", configFilePath, t.ServiceURL, err)
	}

	return nil
}

func (t *TrackerProvider) GetConfigFromTrackerService(genesisHash, configFilePath string) error {
	t.init()
	t.ServiceURL = t.ServiceURL + "/" + genesisHash + "/file"
	configResp := cxApplicationConfig{}
	err := t.apiClient.Get(t.ServiceURL, &configResp)
	if err != nil {
		return fmt.Errorf("error while retreiving config with genesis hash: %s on service %s due to error: %s", genesisHash, t.ServiceURL, err)
	}

	data, err := json.MarshalIndent(configResp, "", " ")
	if err != nil {
		return fmt.Errorf("error while marshal config with genesis hash: %s due to error: %s", genesisHash, err)
	}

	f, err := os.Create(configFilePath)
	if err != nil {
		return fmt.Errorf("error while creating file for config with genesis hash: %s due to error: %s", genesisHash, err)
	}

	if _, err := f.Write(data); err != nil {
		return fmt.Errorf("error while saving config to file with genesis hash: %s due to error: %s", genesisHash, err)
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
