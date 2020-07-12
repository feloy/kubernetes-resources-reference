package cli

import (
	"fmt"

	"github.com/spf13/cobra"
)

// Hugo defines the `hugo` subcommand
func Hugo() *cobra.Command {
	cmd := &cobra.Command{
		Use:           "hugo",
		Short:         "output specification for Hugo website",
		Long:          "output the specification in a format usable for a Hugo website",
		SilenceErrors: true,
		SilenceUsage:  true,
		RunE: func(cmd *cobra.Command, args []string) error {
			toc, err := prepareTOC(cmd)
			if err != nil {
				return fmt.Errorf("Unable to load specs and/or toc config: %v", err)
			}

			outputDir := cmd.Flag(outputDirOption).Value.String()
			err = toc.ToHugo(outputDir)
			if err != nil {
				return err
			}

			show, err := cmd.Flags().GetBool(showDefinitionsOption)
			if err != nil {
				return err
			}
			if show {
				toc.OutputDocumentedDefinitions()
			}
			return nil
		},
	}
	cmd.Flags().StringP(configDirOption, "c", "", "Directory containing documentation configuration")
	cmd.MarkFlagRequired(configDirOption)
	cmd.Flags().StringP(outputDirOption, "o", "", "Directory to write markdown files")
	cmd.MarkFlagRequired(outputDirOption)
	cmd.Flags().Bool(showDefinitionsOption, false, "Show where definitions are defined on output")
	return cmd
}
