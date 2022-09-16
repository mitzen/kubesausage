package clusterinfo

import (
	"fmt"
	"os"

	"github.com/mitzen/kubesausage/pkg/feature"
	"github.com/spf13/cobra"
)

// Example of use case
// kubesausage.exe version
// kubesausage.exe evict

func Execute() {

	var informationType string

	var versionCmd = &cobra.Command{
		Use:   `version`,
		Short: `Get version info`,
		Long:  `Get version info`,
		Run: func(cmd *cobra.Command, args []string) {
			fmt.Println("version 1.0")
		},
	}

	var evictCmd = &cobra.Command{
		Use:   `evict`,
		Short: `Evit a pod from a namespace`,
		Long:  `Evit a pod from a namespace`,
		Run: func(cmd *cobra.Command, args []string) {

			// prepare eviction

			fmt.Println("evict 1.0, %s", args)
		},
	}

	var rootCmd = &cobra.Command{
		Use:   "info",
		Short: "info",
		Long:  `get cluster info!`,
		Run: func(cmd *cobra.Command, args []string) {

			fmt.Printf("%s\n", informationType)

			istioUpgrader := feature.ClusterManager{Cmd: cmd}
			istioUpgrader.GetNodeResourceLimits()
		},
	}

	//rootCmd.Flags().StringVarP(&informationType, "type", "t", "inplace", "Upgrade type")
	//rootCmd.MarkFlagRequired("type")

	rootCmd.AddCommand(versionCmd)
	rootCmd.AddCommand(evictCmd)

	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
