package util

import (
	"strings"

	log "github.com/sirupsen/logrus"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/kubernetes"
)

var resources = []string{
	"Pods",
	"ComponentStatuses",
	"ConfigMaps",
	"Endpoints",
	"Events",
	"LimitRanges",
	"Namespaces",
	"PersistentVolumes",
	"PersistentVolumeClaims",
	"PodTemplates",
	"ResourceQuotas",
	"Secrets",
	"Services",
	"ServiceAccounts",
	"DaemonSets",
	"Deployments",
	"ReplicaSets",
	"StatefulSets",
}

// Getter the
// This should be go routine ready. Such that getter can be called via goroutines and over a channel the value can be passed to a switch type through with the respective printer can be called.
func Getter(namespace string, clientset *kubernetes.Clientset, kinds string, c chan interface{}) {
	defer close(c)
	var err error
	var list interface{}

	if kinds != "" {
		resources = strings.Split(kinds, ",")
	}

	for _, resource := range resources {
		switch resource {
		case "Pods":
			list, err = clientset.CoreV1().Pods(namespace).List(metav1.ListOptions{})
			handleError(err, resource)
		case "ComponentStatuses":
			list, err = clientset.CoreV1().ComponentStatuses().List(metav1.ListOptions{})
			handleError(err, resource)
		case "ConfigMaps":
			list, err = clientset.CoreV1().ConfigMaps(namespace).List(metav1.ListOptions{})
			handleError(err, resource)
		case "Endpoints":
			list, err = clientset.CoreV1().Endpoints(namespace).List(metav1.ListOptions{})
			handleError(err, resource)
		case "Events":
			list, err = clientset.CoreV1().Events(namespace).List(metav1.ListOptions{})
			handleError(err, resource)
		case "LimitRanges":
			list, err = clientset.CoreV1().LimitRanges(namespace).List(metav1.ListOptions{})
			handleError(err, resource)
		case "Namespaces":
			list, err = clientset.CoreV1().Namespaces().List(metav1.ListOptions{})
			handleError(err, resource)
		case "PersistentVolumes":
			list, err = clientset.CoreV1().PersistentVolumes().List(metav1.ListOptions{})
			handleError(err, resource)
		case "PersistentVolumeClaims":
			list, err = clientset.CoreV1().PersistentVolumeClaims(namespace).List(metav1.ListOptions{})
			handleError(err, resource)
		case "PodTemplates":
			list, err = clientset.CoreV1().PodTemplates(namespace).List(metav1.ListOptions{})
			handleError(err, resource)
		case "ResourceQuotas":
			list, err = clientset.CoreV1().ResourceQuotas(namespace).List(metav1.ListOptions{})
			handleError(err, resource)
		case "Secrets":
			list, err = clientset.CoreV1().Secrets(namespace).List(metav1.ListOptions{})
			handleError(err, resource)
		case "Services":
			list, err = clientset.CoreV1().Services(namespace).List(metav1.ListOptions{})
			handleError(err, resource)
		case "ServiceAccounts":
			list, err = clientset.CoreV1().ServiceAccounts(namespace).List(metav1.ListOptions{})
			handleError(err, resource)

		// these will be from the AppsV1
		case "DaemonSets":
			list, err = clientset.AppsV1().DaemonSets(namespace).List(metav1.ListOptions{})
			handleError(err, resource)
		case "Deployments":
			list, err = clientset.AppsV1().Deployments(namespace).List(metav1.ListOptions{})
			handleError(err, resource)
		case "ReplicaSets":
			list, err = clientset.AppsV1().ReplicaSets(namespace).List(metav1.ListOptions{})
			handleError(err, resource)
		case "StatefulSets":
			list, err = clientset.AppsV1().StatefulSets(namespace).List(metav1.ListOptions{})
			handleError(err, resource)
		default:
			log.Error("Given kind is not supported")
			return
		}

		c <- list
	}
}

func handleError(err error, r string) {
	if err != nil {
		log.Errorf("There was an error getting the %s from clientset", r)
	}
}
