package feature

import (
	"fmt"

	"github.com/mitzen/kubesausage/pkg/kube/util"
	"github.com/spf13/cobra"
	apiv1 "k8s.io/api/core/v1"

	"github.com/mitzen/kubeconfig/config"
)

type ClusterManager struct {
	Namespace                     string
	UpgradeType                   string
	VersionSelected               string
	isPreflightUpgradeCheckPassed bool
	Cmd                           *cobra.Command
	DryRun                        bool
}

// Get cluster cpu / memory configuration
// Get pods memory and cpu request and limits
// Total it all up

func (i *ClusterManager) Execute() {

	cfg := config.ClientConfig{}
	restConfig := cfg.NewRestConfig()
	clientset := cfg.NewClientSet(restConfig)

	nsutil := util.KubeObject{}
	nsutil.NewKubeObject(clientset)

	nodes, err := nsutil.ListAllNodes()

	if err != nil {
		fmt.Printf("Unable get nodes info.")
	}

	for _, node := range nodes.Items {

		fmt.Printf("----------------------------------------------\n")
		fmt.Printf("Node: %s \n", node.Name)
		fmt.Printf("CPU: %d \n", node.Status.Capacity.Cpu().ToDec().Value())
		fmt.Printf("Memory: %d \n", node.Status.Capacity.Memory().ToDec().Value())
		fmt.Printf("Pods: %d \n", node.Status.Capacity.Storage().ToDec().Value())
		fmt.Printf("----------------------------------------------\n")

		pods, err := nsutil.ListAllPods(apiv1.NamespaceAll)
		if err != nil {
			fmt.Printf("Unable get nodes info.")
		}

		var (
			totalMemoryRequested, totalCPURequested, totalCPULimit, totalMemoryLimit int64
		)

		totalMemoryRequested = 0
		totalCPURequested = 0
		totalCPULimit = 0
		totalMemoryLimit = 0

		for _, pod := range pods.Items {
			if pod.Spec.NodeName == node.Name {
				// same node //
				fmt.Printf("Namespace: %s \n", pod.Namespace)
				fmt.Printf("Pod name: %s \n", pod.Name)

				for _, container := range pod.Spec.Containers {

					CPURequested := container.Resources.Requests.Cpu().Value()
					MemoryRequested := container.Resources.Requests.Memory().Value()
					CPULimit := container.Resources.Limits.Cpu().Value()
					MemoryLimit := container.Resources.Limits.Memory().Value()

					fmt.Printf("Total cpu request for container: %d \n", CPURequested)
					fmt.Printf("Total cpu limits for container: %d \n", CPULimit)
					fmt.Printf("Total memory request for container: %d \n", MemoryRequested)
					fmt.Printf("Total memory limits for container: %d \n", MemoryLimit)

					totalCPURequested += CPURequested
					totalMemoryRequested += MemoryRequested
					totalCPULimit += CPULimit
					totalMemoryLimit += MemoryLimit
				}
			}
		}

		fmt.Printf("Total cpu request for container: %d \n", totalCPURequested)
		fmt.Printf("Total cpu limits for container: %d \n", totalCPULimit)
		fmt.Printf("Total memory request for container:%d \n", totalMemoryRequested)
		fmt.Printf("Total memory limits for container:%d \n", totalMemoryLimit)
	}
}
