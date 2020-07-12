package cli

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

// Docbook defines the `docbook` subcommand
func Docbook() *cobra.Command {
	cmd := &cobra.Command{
		Use:           "docbook",
		Short:         "output specification in Docbook format",
		Long:          "output the specification in Docbook format",
		SilenceErrors: true,
		SilenceUsage:  true,
		RunE: func(cmd *cobra.Command, args []string) error {
			toc, err := prepareTOC(cmd)
			if err != nil {
				return fmt.Errorf("Unable to load specs and/or toc config: %v", err)
			}

			err = toc.ToDocbook(os.Stdout)
			if err != nil {
				return err
			}
			return nil
		},
	}
	cmd.Flags().StringP(configDirOption, "c", "", "Directory containing documentation configuration")
	cmd.MarkFlagRequired(configDirOption)
	return cmd
}
