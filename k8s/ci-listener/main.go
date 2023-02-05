package main

import (
	"context"
	"flag"
	"fmt"
	"log"
	"net/http"
	"path/filepath"
	"time"

	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/rest"
	"k8s.io/client-go/tools/clientcmd"
	"k8s.io/client-go/util/homedir"
)

var jobMapping = map[string]string{
	"CrossyJob_Back":  "back-deployment",
	"CrossyJob_Front": "front-deployment",
}

func main() {
	var kubeconfig *string
	if home := homedir.HomeDir(); home != "" {
		kubeconfig = flag.String("kubeconfig", filepath.Join(home, ".kube", "config"), "(optional) absolute path to the kubeconfig file")
	} else {
		kubeconfig = flag.String("kubeconfig", "", "absolute path to the kubeconfig file")
	}

	port := flag.Uint("port", 8888, "listening port")

	flag.Parse()

	config, err := clientcmd.BuildConfigFromFlags("", *kubeconfig)
	if err != nil {
		log.Fatalln("Can't build config:", err)
	}

	startServer(port, config)
}

func startServer(port *uint, config *rest.Config) {
	for job, pod := range jobMapping {
		http.HandleFunc("/job/"+job, handleJob(pod, config))
	}
	log.Printf("Starting CI listener server on :%d ...\n", *port)
	panic(http.ListenAndServe(fmt.Sprintf(":%d", *port), nil))
}

func handleJob(deploymentName string, config *rest.Config) func(http.ResponseWriter, *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		clientset, err := kubernetes.NewForConfig(config)
		if err != nil {
			log.Fatalln(err)
		}

		deployments := clientset.AppsV1().Deployments("default")

		initialScale, err := deployments.GetScale(context.Background(), deploymentName, metav1.GetOptions{})
		if err != nil {
			log.Printf("Error: %v\n", err)
			w.WriteHeader(http.StatusInternalServerError)
			return
		}
		// Scale the deployment to zero to delete all pods
		log.Printf("Scaling deployment %q to 0 ...\n", deploymentName)
		zeroScale := *initialScale
		zeroScale.Spec.Replicas = 0
		_, err = deployments.UpdateScale(context.Background(), deploymentName, &zeroScale, metav1.UpdateOptions{})
		if err != nil {
			log.Printf("Error: %v\n", err)
			w.WriteHeader(http.StatusInternalServerError)
			return
		}

		// Add tempo between deployment downscaling and upscaling
		time.Sleep(10 * time.Second)

		// Scale back the deployment to create pods with the lastest image
		log.Printf("Scaling back deployment %q to %d\n", deploymentName, initialScale.Spec.Replicas)
		_, err = deployments.UpdateScale(context.Background(), deploymentName, initialScale, metav1.UpdateOptions{})
		if err != nil {
			log.Printf("Error: %v\n", err)
			w.WriteHeader(http.StatusInternalServerError)
			return
		}
		w.WriteHeader(http.StatusOK)
	}
}
