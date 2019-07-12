package cli

import (
	"fmt"

	"github.com/spf13/cobra"
	"github.com/sn-srdjan/cx-tracker-cli/src/cli/provider"
)

func persistConfigCmd() *cobra.Command {
	saveCmd := &cobra.Command{
		Short: "Save config to CX Tracker service",
		Use:   "save [flags] [path_to_config_file]",
		Long: fmt.Sprintf(`Save config file on remote CX Tracker service, the default
    service URL is %s.`, provider.DefaultCxTrackerURL),
		SilenceUsage:          true,
		Args:                  cobra.MinimumNArgs(1),
		DisableFlagsInUseLine: true,
		RunE: func(c *cobra.Command, args []string) error {
			configFilePath := args[0]
			if configFilePath == "" {
				return c.Help()
			}

			serviceURL, err := c.Flags().GetString("url")
			if err != nil {
				serviceURL = provider.DefaultCxTrackerURL
			}

			tracker := provider.TrackerProvider{
				ServiceURL: serviceURL,
			}

			err = tracker.SaveToTrackerService(configFilePath)

			switch err.(type) {
			case nil:
				fmt.Println("success")
				return nil
			default:
				return err
			}
		},
	}

	saveCmd.Flags().StringP("url", "u", "", "CX Tracker service URL. If no URL is specified default CX Tracker service URL will be used.")

	return saveCmd
}
