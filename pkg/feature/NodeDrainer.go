package feature

import (
	"github.com/mitzen/kubesausage/pkg/kube/util"
	apiv1 "k8s.io/api/core/v1"
)

type NodeDrainer struct {
	nsutil util.KubeObject
}

func (i *NodeDrainer) QueryPodStatus() {

	// Examine node status to identify faulty pods
	i.GetBadPods()
	// Get deployment

	// Get pdb
}

// func (i *NodeDrainer) GetDeployment() []apiv1.Pod {

// }

func (i *NodeDrainer) GetBadPods() []apiv1.Pod {

	var badPodSlice []apiv1.Pod

	podlist, err := i.nsutil.ListAllPods(apiv1.NamespaceAll)
	if err != nil {
		errorColor.Printf("Unable get nodes info.")
	}

	for _, pod := range podlist.Items {
		if pod.Status.Phase != "Running" {
			errorColor.Printf("Pod: %s status: %s", pod.Name, pod.Status.Phase)
			badPodSlice = append(badPodSlice, pod)
		}
	}
	return badPodSlice
}
