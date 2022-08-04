package util

import (
	"context"

	v1 "k8s.io/api/core/v1"
	"k8s.io/client-go/kubernetes"
	"k8s.io/kubectl/pkg/drain"
)

type NodeClient struct {
	clientset kubernetes.Interface
}

func (nc *NodeClient) NewNodeClient(clientset kubernetes.Interface) {
	nc.clientset = clientset
}

func (nc *NodeClient) DrainNode(node *v1.Node) {
	helper := nc.createHelper()
	drain.RunNodeDrain(&helper, node.Name)
}

func (nc *NodeClient) Cordon(node *v1.Node) {
	helper := nc.createHelper()
	drain.RunCordonOrUncordon(&helper, node, true)
}

func (nc *NodeClient) UnCordon(node *v1.Node) {
	helper := nc.createHelper()
	drain.RunCordonOrUncordon(&helper, node, false)
}

func (nc *NodeClient) createHelper() drain.Helper {

	return drain.Helper{
		Ctx:                             context.TODO(),
		Client:                          nc.clientset,
		Force:                           false,
		GracePeriodSeconds:              0,
		IgnoreAllDaemonSets:             false,
		Timeout:                         0,
		DeleteEmptyDirData:              false,
		Selector:                        "",
		PodSelector:                     "",
		ChunkSize:                       0,
		DisableEviction:                 false,
		SkipWaitForDeleteTimeoutSeconds: 0,
		AdditionalFilters:               []drain.PodFilter{},
		Out:                             nil,
		ErrOut:                          nil,
		DryRunStrategy:                  0,
		//DryRunVerifier:                  &resource.QueryParamVerifier{},
		//OnPodDeletedOrEvicted: func(pod *v1.Pod, usingEviction bool) {
	}
}
