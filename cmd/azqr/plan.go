package azqr

import (
	"github.com/cmendible/azqr/internal/scanners"
	"github.com/cmendible/azqr/internal/scanners/plan"
	"github.com/spf13/cobra"
)

func init() {
	rootCmd.AddCommand(planCmd)
}

var planCmd = &cobra.Command{
	Use:   "plan",
	Short: "Scan Azure App Service",
	Long:  "Scan Azure App Service",
	Run: func(cmd *cobra.Command, args []string) {
		serviceScanners := []scanners.IAzureScanner{
			&plan.AppServiceScanner{},
		}

		scan(cmd, serviceScanners)
	},
}
