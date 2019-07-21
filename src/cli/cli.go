/*
Package cli implements the CLI cmd's methods.

Includes methods for interacting with the REST API to query a cx-tracker data.
*/
package cli

import (
	"encoding/json"
	"errors"
	"flag"
	"fmt"
	"net/url"

	"os"

	"github.com/spf13/cobra"
	"github.com/spf13/pflag"
	"github.com/sn-srdjan/cx-tracker-cli/src/cli/cxtracker"
	// "github.com/skycoin/skycoin/src/util/file"
)

var (
	// Version is the CLI Version
	Version = "0.1.1"
)

var (
	envVarsHelp = fmt.Sprintf(`ENVIRONMENT VARIABLES:
    CX_TRACKER_ADDR: Address of CX Tracker service. Must be in scheme://host format. Default "%s"`, cxtracker.DefaultCxTrackerURL)

	helpTemplate = fmt.Sprintf(`USAGE:{{if .Runnable}}
  {{.UseLine}}{{end}}{{if .HasAvailableSubCommands}}
  {{.CommandPath}} [command] [flags] [arguments...]{{end}}{{with (or .Long .Short)}}

DESCRIPTION:
    {{. | trimTrailingWhitespaces}}{{end}}{{if .HasExample}}

EXAMPLES:
{{.Example}}{{end}}{{if .HasAvailableSubCommands}}

COMMANDS:{{range .Commands}}{{if (or .IsAvailableCommand (eq .Name "help"))}}
  {{rpad .Name .NamePadding }} {{.Short}}{{end}}{{end}}{{end}}{{if .HasAvailableLocalFlags}}

FLAGS:
{{.LocalFlags.FlagUsages | trimTrailingWhitespaces}}{{end}}{{if .HasAvailableInheritedFlags}}

GLOBAL FLAGS:
{{.InheritedFlags.FlagUsages | trimTrailingWhitespaces}}{{end}}{{if .HasHelpSubCommands}}

Additional help topics:{{range .Commands}}{{if .IsAdditionalHelpTopicCommand}}
  {{rpad .CommandPath .CommandPathPadding}} {{.Short}}{{end}}{{end}}{{end}}{{if .HasAvailableSubCommands}}

Use "{{.CommandPath}} [command] --help" for more information about a command.{{end}}

%s
`, envVarsHelp)

	// ErrJSONMarshal is returned if JSON marshaling failed
	ErrJSONMarshal = errors.New("json marshal failed")
)

var (
	cliConfig Config
	quitChan  = make(chan struct{})
)

// Config cli's configuration struct
type Config struct {
	TrackerURL string `json:"tracker_url"`
}

// LoadConfig loads config from environment, prior to parsing CLI flags
func LoadConfig() (Config, error) {
	trackerURL := os.Getenv("CX_TRACKER_URL")
	if trackerURL == "" {
		trackerURL = cxtracker.DefaultCxTrackerURL
	}

	if _, err := url.Parse(trackerURL); err != nil {
		return Config{}, errors.New("CX_TRACKER_URL must be in scheme://host format")
	}

	return Config{
		TrackerURL: trackerURL,
	}, nil
}

// NewCLI creates a cli instance
func NewCLI(cfg Config) (*cobra.Command, error) {
	cliConfig = cfg

	cxTrackerCLI := &cobra.Command{
		Short: fmt.Sprintf("The cx-tracker command line interface"),
		Use:   fmt.Sprintf("cx-tracker-cli"),
	}

	commands := []*cobra.Command{
		persistConfigCmd(),
		versionCmd(),
	}

	cxTrackerCLI.Version = Version
	cxTrackerCLI.SuggestionsMinimumDistance = 1
	cxTrackerCLI.AddCommand(commands...)

	cxTrackerCLI.SetHelpTemplate(helpTemplate)
	cxTrackerCLI.SetUsageTemplate(helpTemplate)
	pflag.CommandLine.AddGoFlagSet(flag.CommandLine)

	return cxTrackerCLI, nil
}

func printHelp(c *cobra.Command) {
	c.Printf("See '%s %s --help'\n", c.Parent().Name(), c.Name())
}

func formatJSON(obj interface{}) ([]byte, error) {
	d, err := json.MarshalIndent(obj, "", "    ")
	if err != nil {
		return nil, ErrJSONMarshal
	}
	return d, nil
}

func printJSON(obj interface{}) error {
	d, err := formatJSON(obj)
	if err != nil {
		return err
	}

	fmt.Println(string(d))

	return nil
}
