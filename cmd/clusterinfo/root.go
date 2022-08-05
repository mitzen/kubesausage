package clusterinfo

import (
	"fmt"
	"os"

	"github.com/mitzen/kubesausage/pkg/feature"
	"github.com/spf13/cobra"
)

// Example of use case
// kubesausage.exe --type=cluster --version=10.20.30

func Execute() {

	var informationType string

	var versionCmd = &cobra.Command{
		Use:   "version",
		Short: "Get version info",
		Long:  `Get version info`,
		Run: func(cmd *cobra.Command, args []string) {
			fmt.Println("version 1.0")
		},
	}

	var rootCmd = &cobra.Command{
		Use:   "info",
		Short: "info",
		Long:  `get cluster info!`,
		Run: func(cmd *cobra.Command, args []string) {

			fmt.Printf("%s\n", informationType)

			istioUpgrader := feature.ClusterManager{Cmd: cmd}
			istioUpgrader.Execute()
		},
	}

	rootCmd.Flags().StringVarP(&informationType, "type", "t", "inplace", "Upgrade type")
	rootCmd.MarkFlagRequired("type")

	rootCmd.AddCommand(versionCmd)

	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
