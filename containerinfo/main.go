package main

import (
	"context"
	"encoding/json"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/rest"
	"log"
	"net/http"
	"strings"
)

type Container struct {
	Name      string `json:"container_name"`
	PodName   string `json:"pod_name"`
	Namespace string `json:"namespace"`
	MemReq    string `json:"mem_req,omitempty"`
	MemLimit  string `json:"mem_limit,omitempty"`
	CpuReq    string `json:"cpu_req,omitempty"`
	CpuLimit  string `json:"cpu_limit,omitempty"`
}

func initializeClientSet() (*kubernetes.Clientset, error) {
	// Initialize client to interact with kubernetes API
	config, err := rest.InClusterConfig()
	if err != nil {
		return nil, err
	}
	clientset, err := kubernetes.NewForConfig(config)
	if err != nil {
		return nil, err
	}
	return clientset, nil
}

func ContainerResourcesHandler(w http.ResponseWriter, r *http.Request, clientset *kubernetes.Clientset) {
	// Build label selector
	podLabelParams := r.URL.Query()["pod-label"]
	// If len(podlabelparams) > 1 e.g /container-resources?pod-label=foo=bar&pod-label=bar=foo, simple join all the params
	labelSelector := strings.Join(podLabelParams, ",")

	// List pods on cluster
	pods, err := clientset.CoreV1().Pods("").List(context.TODO(), metav1.ListOptions{
		LabelSelector: labelSelector,
	})
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		log.Fatal("Error listing pods")
		return
	}

	// Build response
	containers := []Container{}
	for _, currentPod := range pods.Items {
		for _, podContainer := range currentPod.Spec.Containers {
			containers = append(containers, Container{
				Name:      podContainer.Name,
				PodName:   currentPod.ObjectMeta.Name,
				Namespace: currentPod.ObjectMeta.Namespace,
				MemReq:    podContainer.Resources.Requests.Memory().String(),
				MemLimit:  podContainer.Resources.Limits.Memory().String(),
				CpuReq:    podContainer.Resources.Requests.Cpu().String(),
				CpuLimit:  podContainer.Resources.Limits.Cpu().String(),
			})
		}
	}
	json.NewEncoder(w).Encode(containers)
}

func main() {
	clientset, err := initializeClientSet()
	if err != nil {
		log.Fatal("Error initializing kubernetes config")
		return
	}

	http.HandleFunc("/container-resources", func(w http.ResponseWriter, r *http.Request) {
		ContainerResourcesHandler(w, r, clientSet)
	})

	http.HandleFunc("/health", func(w http.ResponseWriter, r *http.Request) {
		json.NewEncoder(w).Encode(map[string]string{"status": "healthy"})
	})
	log.Fatal(http.ListenAndServe(":8000", nil))
}
