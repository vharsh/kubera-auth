package k8s

import (
	"os"

	log "github.com/golang/glog"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/rest"
	"k8s.io/client-go/tools/clientcmd"
)

// ClientSet is the client set generated by in cluster config
var ClientSet *kubernetes.Clientset

func init() {
	var err error
	ClientSet, err = getGenericK8sClient()
	if err != nil {
		log.Errorln("Error creating client set", err)
		return
	}
}

// getKubeConfig gets the kubeconfig
func getKubeConfig() (*rest.Config, error) {
	KubeConfig := os.Getenv("KUBECONFIG")
	// Use in-cluster config if kubeconfig path is not specified
	if KubeConfig == "" {
		return rest.InClusterConfig()
	}

	return clientcmd.BuildConfigFromFlags("", KubeConfig)
}

// getGenericK8sClient gets the client set for the in cluster config
func getGenericK8sClient() (*kubernetes.Clientset, error) {
	config, err := getKubeConfig()
	if err != nil {
		return nil, err
	}

	return kubernetes.NewForConfig(config)
}