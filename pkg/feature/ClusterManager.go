package feature

import (
	"fmt"

	"github.com/fatih/color"
	"github.com/mitzen/kubeconfig/config"
	"github.com/mitzen/kubesausage/pkg/kube/util"
	"github.com/spf13/cobra"
	apiv1 "k8s.io/api/core/v1"
)

type ClusterManager struct {
	Namespace string
	Cmd       *cobra.Command
}

const (
	unitMegabytes int64 = 1000000
)

func (i *ClusterManager) PrepareClusterDrain() {

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
		color.Red(node.Name)
	}
}

// Get cluster cpu / memory configuration
// Get pods memory and cpu request and limits
// Total it all up

func (i *ClusterManager) GetNodeResourceLimits() {

	infoColor := color.New(color.FgHiYellow)
	podInfoColor := color.New(color.FgGreen)

	cfg := config.ClientConfig{}
	restConfig := cfg.NewRestConfig()
	clientset := cfg.NewClientSet(restConfig)

	nsutil := util.KubeObject{}
	nsutil.NewKubeObject(clientset)

	nodes, err := nsutil.ListAllNodes()

	if err != nil {
		color.Red("Unable get nodes info.")
	}

	for _, node := range nodes.Items {

		color.White("----------------------------------------------\n")
		infoColor.Printf("Node: %s \n", node.Name)
		infoColor.Printf("OS: %s \n", node.Status.NodeInfo.OperatingSystem)
		infoColor.Printf("Version: %s \n", node.Status.NodeInfo.KubeletVersion)
		infoColor.Printf("Arch: %s \n", &node.Status.NodeInfo.Architecture)
		fmt.Printf("----------------------------------------------\n")

		pods, err := nsutil.ListAllPods(apiv1.NamespaceAll)
		if err != nil {
			color.Red("Unable get nodes info.")
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
				podInfoColor.Printf("Namespace: %s Name %s Status %s \n", pod.Namespace, pod.Name, pod.Status.Phase)
				for _, container := range pod.Spec.Containers {
					CPURequested := container.Resources.Requests.Cpu().Value()
					MemoryRequested := container.Resources.Requests.Memory().Value()
					CPULimit := container.Resources.Limits.Cpu().Value()
					MemoryLimit := container.Resources.Limits.Memory().Value()

					fmt.Printf("Container Name: %s \n", container.Name)
					fmt.Printf("Image Name: %s \n", container.Image)

					if CPURequested <= 0 {
						fmt.Printf("Cpu request for container: %d \n", CPURequested)
						fmt.Printf("Cpu limits for container: %d \n", CPULimit)
					}

					if MemoryRequested <= 0 {
						fmt.Printf("Memory request for container (M): %d \n", MemoryRequested/unitMegabytes)
						fmt.Printf("Memory limits for container (M): %d \n", MemoryLimit/unitMegabytes)
					}

					totalCPURequested += CPURequested
					totalMemoryRequested += MemoryRequested
					totalCPULimit += CPULimit
					totalMemoryLimit += MemoryLimit
				}
			}
		}

		fmt.Printf("Total cpu request for container: %d \n", totalCPURequested)
		fmt.Printf("Total cpu limits for container: %d \n", totalCPULimit)
		fmt.Printf("Total memory request for container:%d \n", totalMemoryRequested/unitMegabytes)
		fmt.Printf("Total memory limits for container:%d \n", totalMemoryLimit/unitMegabytes)
		fmt.Printf("Node CPU: %d \n", node.Status.Capacity.Cpu().ToDec().Value())
		fmt.Printf("Node Memory: %d \n", node.Status.Capacity.Memory().ToDec().Value()/unitMegabytes)
	}
}
