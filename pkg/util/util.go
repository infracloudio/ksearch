package util

import (
	"strings"

	"github.com/infracloudio/ksearch/pkg/printers"
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
func Getter(namespace string, clientset *kubernetes.Clientset, kinds string, resName string) {

	for _, resource := range resources {

		switch resource {
		case "Pods":
			pods, err := clientset.CoreV1().Pods(namespace).List(metav1.ListOptions{})
			handleError(err, resource)
			if result := checkKinds("pod", kinds); result == true {
				printers.PrintPodDetails(pods, resName)
			}
		case "ComponentStatuses":
			componentStatuses, err := clientset.CoreV1().ComponentStatuses().List(metav1.ListOptions{})
			handleError(err, resource)
			if result := checkKinds("componentstatus", kinds); result == true {
				printers.PrintComponentStatuses(componentStatuses, resName)
			}
		case "ConfigMaps":
			cms, err := clientset.CoreV1().ConfigMaps(namespace).List(metav1.ListOptions{})
			handleError(err, resource)
			if result := checkKinds("configmap", kinds); result == true {
				printers.PrintConfigMaps(cms, resName)
			}
		case "Endpoints":
			endPoints, err := clientset.CoreV1().Endpoints(namespace).List(metav1.ListOptions{})
			handleError(err, resource)
			if result := checkKinds("endpoint", kinds); result == true {
				printers.PrintEndpoints(endPoints, resName)
			}
		case "Events":
			events, err := clientset.CoreV1().Events(namespace).List(metav1.ListOptions{})
			handleError(err, resource)
			if result := checkKinds("event", kinds); result == true {
				printers.PrintEvents(events, resName)
			}
		case "LimitRanges":
			limitRanges, err := clientset.CoreV1().LimitRanges(namespace).List(metav1.ListOptions{})
			handleError(err, resource)
			if result := checkKinds("limitrange", kinds); result == true {
				printers.PrintLimitRanges(limitRanges, resName)
			}
		case "Namespaces":
			namespaces, err := clientset.CoreV1().Namespaces().List(metav1.ListOptions{})
			handleError(err, resource)
			if result := checkKinds("namespace", kinds); result == true {
				printers.PrintNamespaces(namespaces, resName)
			}
		case "PersistentVolumes":
			pvs, err := clientset.CoreV1().PersistentVolumes().List(metav1.ListOptions{})
			handleError(err, resource)
			if result := checkKinds("persistentvolume", kinds); result == true {
				printers.PrintPVs(pvs, resName)
			}
		case "PersistentVolumeClaims":
			pvcs, err := clientset.CoreV1().PersistentVolumeClaims(namespace).List(metav1.ListOptions{})
			handleError(err, resource)
			if result := checkKinds("persistentvolumeclaim", kinds); result == true {
				printers.PrintPVCs(pvcs, resName)
			}
		case "PodTemplates":
			podTemplates, err := clientset.CoreV1().PodTemplates(namespace).List(metav1.ListOptions{})
			handleError(err, resource)
			if result := checkKinds("podtemplate", kinds); result == true {
				printers.PrintPodTemplates(podTemplates, resName)
			}
		case "ResourceQuotas":
			resQuotas, err := clientset.CoreV1().ResourceQuotas(namespace).List(metav1.ListOptions{})
			handleError(err, resource)
			if result := checkKinds("resourcequota", kinds); result == true {
				printers.PrintResourceQuotas(resQuotas, resName)
			}
		case "Secrets":
			secrets, err := clientset.CoreV1().Secrets(namespace).List(metav1.ListOptions{})
			handleError(err, resource)
			if result := checkKinds("secret", kinds); result == true {
				printers.PrintSecrets(secrets, resName)
			}
		case "Services":
			services, err := clientset.CoreV1().Services(namespace).List(metav1.ListOptions{})
			handleError(err, resource)
			if result := checkKinds("service", kinds); result == true {
				printers.PrintServices(services, resName)
			}
		case "ServiceAccounts":
			serviceAccs, err := clientset.CoreV1().ServiceAccounts(namespace).List(metav1.ListOptions{})
			handleError(err, resource)
			if result := checkKinds("serviceaccount", kinds); result == true {
				printers.PrintServiceAccounts(serviceAccs, resName)
			}

		// these will be from the AppsV1
		case "DaemonSets":
			daemonsets, err := clientset.AppsV1().DaemonSets(namespace).List(metav1.ListOptions{})
			handleError(err, resource)
			if result := checkKinds("daemonset", kinds); result == true {
				printers.PrintDaemonSets(daemonsets, resName)
			}
		case "Deployments":
			deployments, err := clientset.AppsV1().Deployments(namespace).List(metav1.ListOptions{})
			handleError(err, resource)
			if result := checkKinds("deployment", kinds); result == true {
				printers.PrintDeployments(deployments, resName)
			}
		case "ReplicaSets":
			rsets, err := clientset.AppsV1().ReplicaSets(namespace).List(metav1.ListOptions{})
			handleError(err, resource)
			if result := checkKinds("replicaset", kinds); result == true {
				printers.PrintReplicaSets(rsets, resName)
			}
		case "StatefulSets":
			ssets, err := clientset.AppsV1().StatefulSets(namespace).List(metav1.ListOptions{})
			handleError(err, resource)
			if result := checkKinds("statefulset", kinds); result == true {
				printers.PrintStateFulSets(ssets, resName)
			}
		}

	}

}

func checkKinds(kind string, providedKinds string) bool {
	if providedKinds == "" {
		return true
	}
	stringSlice := strings.Split(providedKinds, ",")
	for _, val := range stringSlice {
		if kind == val {
			return true
		}
	}
	return false
}

func handleError(err error, r string) {
	if err != nil {
		log.Errorf("There was an error getting the %s from clientset", r)
	}
}
