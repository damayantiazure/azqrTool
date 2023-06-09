package azqr

import (
	"github.com/cmendible/azqr/internal/scanners"
	"github.com/cmendible/azqr/internal/scanners/cae"
	"github.com/spf13/cobra"
)

func init() {
	rootCmd.AddCommand(caeCmd)
}

var caeCmd = &cobra.Command{
	Use:   "cae",
	Short: "Scan Azure Container Apps",
	Long:  "Scan Azure Container Apps",
	Run: func(cmd *cobra.Command, args []string) {
		serviceScanners := []scanners.IAzureScanner{
			&cae.ContainerAppsScanner{},
		}

		scan(cmd, serviceScanners)
	},
}
