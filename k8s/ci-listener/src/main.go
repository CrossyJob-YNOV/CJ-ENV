package main

import (
	"context"
	"flag"
	"fmt"
	"log"
	"net/http"
	"strings"
	"sync"

	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/rest"
)

var jobMapping = map[string]string{
	"CrossyJob_Back":  "back-deployment",
	"CrossyJob_Front": "front-deployment",
}

var jobLocks = map[string]*sync.Mutex{}

func main() {
	port := flag.Uint("port", 8888, "listening port")
	flag.Parse()

	config, err := rest.InClusterConfig()
	if err != nil {
		panic(err.Error())
	}

	startServer(port, config)
}

func startServer(port *uint, config *rest.Config) {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		log.Println("=>", r.URL.Path)
	})

	for job, deploymentName := range jobMapping {
		jobLocks[deploymentName] = new(sync.Mutex)
		http.HandleFunc("/job/"+job, handleJob(deploymentName, config))
	}
	log.Printf("Starting CI listener server on :%d ...\n", *port)
	panic(http.ListenAndServe(fmt.Sprintf(":%d", *port), nil))
}

func handleJob(deploymentName string, config *rest.Config) func(http.ResponseWriter, *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		log.Printf("Updating %q containers ...\n", deploymentName)

		jobLocks[deploymentName].Lock()
		defer jobLocks[deploymentName].Unlock()

		clientset, err := kubernetes.NewForConfig(config)
		if err != nil {
			log.Fatalln(err)
		}

		pods := clientset.CoreV1().Pods("default")

		podList, err := pods.List(context.Background(), metav1.ListOptions{})
		if err != nil {
			log.Printf("Error: %v\n", err)
			w.WriteHeader(http.StatusInternalServerError)
			return
		}

		for _, pod := range podList.Items {
			if strings.HasPrefix(pod.Name, deploymentName+"-") {
				fmt.Println("Deleting pod", pod.Name, "...")
				err = pods.Delete(context.Background(), pod.Name, metav1.DeleteOptions{})
				if err != nil {
					log.Printf("error: %v\n", err)
				}
			}
		}

		fmt.Println("Update done !")
		w.WriteHeader(http.StatusOK)
	}
}
