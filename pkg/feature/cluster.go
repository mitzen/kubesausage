package feature

import (
	"github.com/mitzen/kubeconfig/config"
	"github.com/mitzen/kubesausage/pkg/kube/util"
)

type IstioMeshDump struct {
	namespace string
}

func (i *IstioMeshDump) Execute(namespace string) {

	cfg := config.ClientConfig{}
	restConfig := cfg.NewRestConfig()

	ic := util.IstioClient{}
	ic.NewIstioClient(restConfig, namespace)

}
