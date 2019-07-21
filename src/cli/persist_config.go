package cli

import (
	"fmt"

	"github.com/spf13/cobra"
)

func persistConfigCmd() *cobra.Command {
	saveCmd := &cobra.Command{
		Short: "Save config to CX Tracker service",
		Use:   "save [flags] [path_to_config_file]",
		Long: fmt.Sprintf(`Save config file on remote CX Tracker service, the default
    service URL is %s.`, cxtracker.DefaultCxTrackerURL),
		SilenceUsage: true,
		Args:         cobra.MinimumNArgs(1),
		DisableFlagsInUseLine: true,
		RunE: func(c *cobra.Command, args []string) error {
			configFilePath := args[0]
			if configFilePath == "" {
				return c.Help()
			}

			serviceURL, err := c.Flags().GetString("url")
			if err != nil {
				serviceURL = cxtracker.DefaultCxTrackerURL
			}

			tracker := cxtracker.TrackerProvider{
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
