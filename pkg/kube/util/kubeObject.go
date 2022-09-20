package util

import (
	"context"

	appsv1 "k8s.io/api/apps/v1"
	v1 "k8s.io/api/core/v1"
	po "k8s.io/api/policy/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/kubernetes"
)

type KubeObject struct {
	c *kubernetes.Clientset
}

func (n *KubeObject) NewKubeObject(clientSet *kubernetes.Clientset) {
	n.c = clientSet
}

func (n *KubeObject) ListAllNamespace() (*v1.NamespaceList, error) {
	return n.c.CoreV1().Namespaces().List(context.TODO(), metav1.ListOptions{})
}

func (n *KubeObject) ListAllNodes() (*v1.NodeList, error) {
	return n.c.CoreV1().Nodes().List(context.TODO(), metav1.ListOptions{})
}

func (n *KubeObject) ListDaemonset(namespace string) (*appsv1.DaemonSetList, error) {
	return n.c.AppsV1().DaemonSets(namespace).List(context.TODO(), metav1.ListOptions{})
}

func (n *KubeObject) ListAllPods(namespace string) (*v1.PodList, error) {
	return n.c.CoreV1().Pods(namespace).List(context.TODO(), metav1.ListOptions{})
}

func (n *KubeObject) ListAllPodsByLabel(namespace string, label string) (*v1.PodList, error) {
	return n.c.CoreV1().Pods(namespace).List(context.TODO(), metav1.ListOptions{
		LabelSelector: label,
	})
}

func (n *KubeObject) ListAllDeployment(namespace string) (*appsv1.DeploymentList, error) {
	return n.c.AppsV1().Deployments(namespace).List(context.TODO(), metav1.ListOptions{})
}

func (n *KubeObject) GetPdb(namespace string) (*po.PodDisruptionBudgetList, error) {
	return n.c.PolicyV1().PodDisruptionBudgets(namespace).List(context.TODO(), metav1.ListOptions{})
}

func (n *KubeObject) EvictPod(podName string, namespace string) error {
	eviction := po.Eviction{
		ObjectMeta: metav1.ObjectMeta{
			Name:      podName,
			Namespace: namespace,
		},
		DeleteOptions: &metav1.DeleteOptions{}}
	return n.c.PolicyV1().Evictions(namespace).Evict(context.TODO(), &eviction)
}
