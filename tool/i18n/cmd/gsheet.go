package cmd

import (
	"github.com/getcouragenow/bootstrap/tool/i18n/services"
	"github.com/spf13/cobra"
)

var (
	option *string
)

// gsheetCmd represents the gsheet command
var gsheetCmd = &cobra.Command{
	Use:   "gsheet",
	Short: "Download data from google sheet.",
	Long: `Download data from google sheet. For example:
			gsheet -o lang`,
	RunE: func(cmd *cobra.Command, args []string) error {
		return services.Service(*option)
	},
}

func init() {
	rootCmd.AddCommand(gsheetCmd)
	option = gsheetCmd.Flags().StringP("option", "o", "lang", "Download from google sheet")
}
