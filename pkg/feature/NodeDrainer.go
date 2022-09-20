package feature

import (
	"fmt"

	"github.com/mitzen/kubesausage/pkg/kube/util"
	apiv1 "k8s.io/api/core/v1"
)

type NodeDrainer struct {
	nsutil util.KubeObject
}

func (i *NodeDrainer) AssessDeploymentReadiness() {
	// Examine node status to identify faulty pods
	i.GetBadPods()
	// Get deployment
	i.GetDeployment()
	// Get pdb
}

func (i *NodeDrainer) GetDeployment() {
	list, err := i.nsutil.ListAllDeployment(apiv1.NamespaceAll)

	if err != nil {
		errorColor.Printf("Unable get deployment info.")
	}

	for _, deployment := range list.Items {

		podlabelName := deployment.Labels["app"]
		if podlabelName == "" {
			continue
		}

		pods, err := i.nsutil.ListAllPodsByLabel(deployment.Namespace, fmt.Sprintf("app=%s", podlabelName))

		if err != nil {
			errorColor.Printf("Unable get deployment info.")
		}

		var podNodeName string = ""
		var prevPod apiv1.Pod

		for _, pod := range pods.Items {

			if pod.Spec.NodeName == podNodeName {
				errorColor.Printf("Pod %s and %s are on the %s node.\n", pod.Name, prevPod.Name, pod.Spec.NodeName)
			}
			prevPod = pod
			podNodeName = pod.Spec.NodeName
		}
	}
}

func (i *NodeDrainer) GetBadPods() []apiv1.Pod {
	var badPodSlice []apiv1.Pod
	var isFreeOfError bool = true

	podlist, err := i.nsutil.ListAllPods(apiv1.NamespaceAll)
	if err != nil {
		errorColor.Printf("Unable get nodes info.")
	}

	for _, pod := range podlist.Items {
		if pod.Status.Phase != "Running" {
			isFreeOfError = false
			errorColor.Printf("Pod: %s status: %s", pod.Name, pod.Status.Phase)
			badPodSlice = append(badPodSlice, pod)
		}
	}

	if isFreeOfError {
		nodeInfoColor.Print("Pods working fine. \n")
	}

	return badPodSlice
}
