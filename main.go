package main

import (
	"flag"
	"fmt"
	"os"
	"strings"
	"text/tabwriter"

	log "github.com/sirupsen/logrus"
	appsv1 "k8s.io/api/apps/v1"
	v1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/tools/clientcmd"
)

func main() {
	var kubeconfig string
	resName := flag.String("name", "", "Name of the pod that you want to get.")
	namespace := flag.String("n", "", "Namespace you want that resource to be searched in.")
	kinds := flag.String("kinds", "", "List all the kinds that you want to be displayed.")

	flag.Parse()

	if envVar := os.Getenv("KUBECONFIG"); len(envVar) > 0 {
		kubeconfig = envVar
	} else {
		log.Error("KUBECONFIG env variable is not set, please set the variable.")
	}

	config, err := clientcmd.BuildConfigFromFlags("", kubeconfig)
	if err != nil {
		log.Info("There was an error getting the config from kubeconfig.")
	}

	clientset, err := kubernetes.NewForConfig(config)
	if err != nil {
		log.Info("There was an error getting clientset from config")
	}

	pods, err := clientset.CoreV1().Pods(*namespace).List(metav1.ListOptions{})
	if err != nil {
		log.Info("There was an error getting the pod from clientset", err)
	}
	printPodDetails(pods, resName)

	componentStatuses, err := clientset.CoreV1().ComponentStatuses().List(metav1.ListOptions{})
	if err != nil {
		log.Info("There was an error getting components statusses from clientset", err)
	}
	if result := checkKinds("componentstatus", *kinds); result == true {
		printComponentStatuses(componentStatuses, resName)
	}

	cms, err := clientset.CoreV1().ConfigMaps(*namespace).List(metav1.ListOptions{})
	if err != nil {
		log.Info("There was an error getting configmaps from cleintset", err)
	}
	if result := checkKinds("configmap", *kinds); result == true {
		printConfigMaps(cms, resName)
	}

	endPoints, err := clientset.CoreV1().Endpoints(*namespace).List(metav1.ListOptions{})
	if err != nil {
		log.Info("There was an error getting endpoints from clientset", err)
	}
	if result := checkKinds("endpoint", *kinds); result == true {
		printEndpoints(endPoints, resName)
	}

	events, err := clientset.CoreV1().Events(*namespace).List(metav1.ListOptions{})
	if err != nil {
		log.Info("There was an error getting the events from clienesat", err)
	}
	if result := checkKinds("event", *kinds); result == true {
		printEvents(events, resName)
	}

	limitRanges, err := clientset.CoreV1().LimitRanges(*namespace).List(metav1.ListOptions{})
	if err != nil {
		log.Info("There was an error getting limitrange from clientset", err)
	}
	if result := checkKinds("limitrange", *kinds); result == true {
		printLimitRanges(limitRanges, resName)
	}

	namespaces, err := clientset.CoreV1().Namespaces().List(metav1.ListOptions{})
	if err != nil {
		log.Info("There was an error getting namespaces from clientset", err)
	}
	if result := checkKinds("namespace", *kinds); result == true {
		printNamespaces(namespaces, resName)
	}

	pvs, err := clientset.CoreV1().PersistentVolumes().List(metav1.ListOptions{})
	if err != nil {
		log.Info("There was an error getting PVs throuhg CLIentset", err)
	}
	if result := checkKinds("persistentvolume", *kinds); result == true {
		printPVs(pvs, resName)
	}

	pvcs, err := clientset.CoreV1().PersistentVolumeClaims(*namespace).List(metav1.ListOptions{})
	if err != nil {
		log.Info("There was an error getting PVCs throuhg CLIentset", err)
	}
	if result := checkKinds("persistentvolumeclaim", *kinds); result == true {
		printPVCs(pvcs, resName)
	}

	podTemplates, err := clientset.CoreV1().PodTemplates(*namespace).List(metav1.ListOptions{})
	if err != nil {
		log.Info("There was an error getting podTemplates throuhg CLIentset", err)
	}
	if result := checkKinds("podtemplate", *kinds); result == true {
		printPodTemplates(podTemplates, resName)
	}

	resQuotas, err := clientset.CoreV1().ResourceQuotas(*namespace).List(metav1.ListOptions{})
	if err != nil {
		log.Info("There was an error getting resourceQuots throuhg CLIentset", err)
	}
	if result := checkKinds("resourcequota", *kinds); result == true {
		printResourceQuotas(resQuotas, resName)
	}

	secrets, err := clientset.CoreV1().Secrets(*namespace).List(metav1.ListOptions{})
	if err != nil {
		log.Info("There was an error getting secrets throuhg CLIentset", err)
	}
	if result := checkKinds("secret", *kinds); result == true {
		printSecrets(secrets, resName)
	}

	services, err := clientset.CoreV1().Services(*namespace).List(metav1.ListOptions{})
	if err != nil {
		log.Info("There was an error getting services  throuhg CLIentset", err)
	}
	printServices(services, resName)

	serviceAccs, err := clientset.CoreV1().ServiceAccounts(*namespace).List(metav1.ListOptions{})
	if err != nil {
		log.Info("There was an error getting serviceacc throuhg CLIentset", err)
	}
	if result := checkKinds("serviceaccount", *kinds); result == true {
		printServiceAccounts(serviceAccs, resName)
	}

	// these will be from the appsV1

	daemonsets, err := clientset.AppsV1().DaemonSets(*namespace).List(metav1.ListOptions{})
	if err != nil {
		log.Info("There was an error getting ds from clientset", err)
	}
	printDaemonSets(daemonsets, resName)

	deployments, err := clientset.AppsV1().Deployments(*namespace).List(metav1.ListOptions{})
	if err != nil {
		log.Info("There was an error getting deployment from clientset", err)
	}
	printDeployments(deployments, resName)

	rsets, err := clientset.AppsV1().ReplicaSets(*namespace).List(metav1.ListOptions{})
	if err != nil {
		log.Info("There was an error getting replicasets from clientset ", err)
	}
	printReplicaSets(rsets, resName)

	ssets, err := clientset.AppsV1().StatefulSets(*namespace).List(metav1.ListOptions{})
	if err != nil {
		log.Info("There was an error getting statefulset from clientset", err)
	}
	if result := checkKinds("statefulset", *kinds); result == true {
		printStateFulSets(ssets, resName)
	}
}

func printPodDetails(pods *v1.PodList, resName *string) {
	fmt.Printf("\nPods\n----\n")
	w := tabwriter.NewWriter(os.Stdout, 0, 0, 1, ' ', 0)
	fmt.Fprintf(w, "%v\t%v\t%v\t%v\n", "NAME", "READY", "STATUS", "RESTARTS")

	for _, pod := range pods.Items {
		if *resName != "" {
			if strings.Contains(pod.Name, *resName) {
				fmt.Fprintf(w, "%v\t%v\t%v\t%v\n", pod.Name, "", pod.Status.Phase, "")
			}
		} else {
			fmt.Fprintf(w, "%v\t%v\t%v\t%v\n", pod.Name, "", pod.Status.Phase, "")
		}
	}
	w.Flush()
}

func printComponentStatuses(componentStatuses *v1.ComponentStatusList, resName *string) {
	fmt.Printf("\nComponentStatuses\n-------------\n")
	w := tabwriter.NewWriter(os.Stdout, 0, 0, 1, ' ', 0)
	fmt.Fprintf(w, "%v\t%v\t%v\t%v\n", "NAME", "STATUS", "MESSAGE", "ERROR")
	for _, componentStatus := range componentStatuses.Items {
		if *resName != "" {
			if strings.Contains(componentStatus.Name, *resName) {
				fmt.Fprintf(w, "%v\t%v\t%v\t%v\n", componentStatus.Name, componentStatus.Conditions[0].Type, componentStatus.Conditions[0].Message, componentStatus.Conditions[0].Error)
			}
		} else {
			fmt.Fprintf(w, "%v\t%v\t%v\t%v\n", componentStatus.Name, componentStatus.Conditions[0].Type, componentStatus.Conditions[0].Message, componentStatus.Conditions[0].Error)
		}
	}
	w.Flush()
}

func printConfigMaps(cms *v1.ConfigMapList, resName *string) {
	fmt.Printf("\nConfigMaps\n--------------\n")
	w := tabwriter.NewWriter(os.Stdout, 0, 0, 1, ' ', 0)
	fmt.Fprintf(w, "%v\t%v\t%v\n", "NAME", "DATA", "AGE")
	for _, configMap := range cms.Items {
		if *resName != "" {
			if strings.Contains(configMap.Name, *resName) {
				fmt.Fprintf(w, "%v\t%v\t%v\n", configMap.Name, len(configMap.Data), "")
			}
		} else {
			fmt.Fprintf(w, "%v\t%v\t%v\n", configMap.Name, len(configMap.Data), "")
		}
	}
	w.Flush()
}

func printEndpoints(endPoints *v1.EndpointsList, resName *string) {
	fmt.Printf("\nEndpoints\n--------------\n")
	w := tabwriter.NewWriter(os.Stdout, 0, 0, 1, ' ', 0)
	fmt.Fprintf(w, "%v\t%v\t%v\n", "NAME", "ENDPOINTS", "AGE")
	for _, endpoint := range endPoints.Items {
		if *resName != "" {
			if strings.Contains(endpoint.Name, *resName) {
				fmt.Fprintf(w, "%v\t%v\t%v\n", endpoint.Name, "", "")
			}
		} else {
			fmt.Fprintf(w, "%v\t%v\t%v\n", endpoint.Name, "", "")
		}
	}
	w.Flush()
}

func printEvents(events *v1.EventList, resName *string) {
	fmt.Printf("\nEvents\n--------------\n")
	w := tabwriter.NewWriter(os.Stdout, 0, 0, 1, ' ', 0)
	fmt.Fprintf(w, "%v\t%v\t%v\t%v\t%v\t%v\n", "NAMESPACE", "LAST SEEN", "TYPE", "REASON", "OBJECT", "MESSAGE")
	for _, event := range events.Items {
		fmt.Fprintf(w, "%v\t%v\t%v\t%v\t%v\t%v\n", event.Namespace, "", event.Type, "", event.InvolvedObject.Kind+"/"+event.InvolvedObject.Name, event.Message)
	}
	w.Flush()
}

func printLimitRanges(limitRanges *v1.LimitRangeList, resName *string) {
	fmt.Printf("\nLimitRanges\n--------------\n")
	w := tabwriter.NewWriter(os.Stdout, 0, 0, 1, ' ', 0)
	fmt.Fprintf(w, "%v\t%v\n", "NAME", "CREATED AT")
	for _, limitRange := range limitRanges.Items {
		if *resName != "" {
			if strings.Contains(limitRange.Name, *resName) {
				fmt.Fprintf(w, "%v\t%v\n", limitRange.Name, limitRange.CreationTimestamp)
			}
		} else {
			fmt.Fprintf(w, "%v\t%v\n", limitRange.Name, limitRange.CreationTimestamp)
		}
	}
	w.Flush()
}

func printNamespaces(namespaces *v1.NamespaceList, resName *string) {
	fmt.Printf("\nNamespaces\n--------------\n")
	w := tabwriter.NewWriter(os.Stdout, 0, 0, 1, ' ', 0)
	fmt.Fprintf(w, "%v\t%v\t%v\n", "NAME", "STATUS", "AGE")
	for _, namespace := range namespaces.Items {
		if *resName != "" {
			if strings.Contains(namespace.Name, *resName) {
				fmt.Fprintf(w, "%v\t%v\t%v\n", namespace.Name, namespace.Status, "")
			}
		} else {
			fmt.Fprintf(w, "%v\t%v\t%v\n", namespace.Name, namespace.Status, "")
		}
	}
	w.Flush()
}

func printPVs(pvs *v1.PersistentVolumeList, resName *string) {
	fmt.Printf("\nPersistentVolumes\n--------------\n")
	w := tabwriter.NewWriter(os.Stdout, 0, 0, 1, ' ', 0)
	fmt.Fprintf(w, "%v\t%v\t%v\t%v\t%v\t%v\t%v\t%v\t%v\n", "NAME", "CAPACITY", "ACCESS MODES", "RECLAIM POLICY", "STATUS", "CLAIM", "STORAGECLASS", "REASON", "AGE")

	for _, pv := range pvs.Items {
		if *resName != "" {
			if strings.Contains(pv.Name, *resName) {
				fmt.Fprintf(w, "%v\t%v\t%v\t%v\t%v\t%v\t%v\t%v\t%v\n", pv.Name, "", pv.Spec.AccessModes, pv.Spec.PersistentVolumeReclaimPolicy, pv.Status, pv.Spec.ClaimRef.Namespace+"/"+pv.Spec.ClaimRef.Name, pv.Spec.StorageClassName, pv.Status.Reason, "")
			}
		} else {
			fmt.Fprintf(w, "%v\t%v\t%v\t%v\t%v\t%v\t%v\t%v\t%v\n", pv.Name, "", pv.Spec.AccessModes, pv.Spec.PersistentVolumeReclaimPolicy, pv.Status, pv.Spec.ClaimRef.Namespace+"/"+pv.Spec.ClaimRef.Name, pv.Spec.StorageClassName, pv.Status.Reason, "")
		}
	}
	w.Flush()
}

func printPVCs(pvcs *v1.PersistentVolumeClaimList, resName *string) {
	fmt.Printf("\nPersistentVolumeClaims\n--------------\n")
	w := tabwriter.NewWriter(os.Stdout, 0, 0, 1, ' ', 0)
	fmt.Fprintf(w, "%v\t%v\t%v\t%v\t%v\t%v\t%v\n", "NAME", "STATUS", "VOLUME", "CAPACITY", "ACCESS MODES", "STORAGECLASS", "AGE")
	for _, pvc := range pvcs.Items {
		if *resName != "" {
			if strings.Contains(pvc.Name, *resName) {
				fmt.Fprintf(w, "%v\t%v\t%v\t%v\t%v\t%v\t%v\n", pvc.Name, pvc.Status, "", pvc.Status.Capacity.Cpu(), pvc.Spec.AccessModes, pvc.Spec.StorageClassName, "")
			}
		} else {
			fmt.Fprintf(w, "%v\t%v\t%v\t%v\t%v\t%v\t%v\n", pvc.Name, pvc.Status, "", pvc.Status.Capacity.Cpu(), pvc.Spec.AccessModes, pvc.Spec.StorageClassName, "")
		}

	}
	w.Flush()
}

func printPodTemplates(podTemplates *v1.PodTemplateList, resName *string) {
	fmt.Printf("\nPodTemplates\n--------------\n")
	w := tabwriter.NewWriter(os.Stdout, 0, 0, 1, ' ', 0)
	fmt.Fprintf(w, "%v\n", "NAME")
	for _, podTemplate := range podTemplates.Items {
		if *resName != "" {
			if strings.Contains(podTemplate.Name, *resName) {
				fmt.Fprintf(w, "%v\n", podTemplate.Name)
			}
		} else {
			fmt.Fprintf(w, "%v\n", podTemplate.Name)
		}
	}
	w.Flush()
}

func printResourceQuotas(resQuotas *v1.ResourceQuotaList, resName *string) {
	fmt.Printf("\nResourceQuotas\n--------------\n")
	w := tabwriter.NewWriter(os.Stdout, 0, 0, 1, ' ', 0)
	fmt.Fprintf(w, "%v\t%v\n", "NAME", "CREATED AT")
	for _, resQ := range resQuotas.Items {
		if *resName != "" {
			if strings.Contains(resQ.Name, *resName) {
				fmt.Fprintf(w, "%v\t%v\n", resQ.Name, resQ.CreationTimestamp)
			}
		} else {
			fmt.Fprintf(w, "%v\t%v\n", resQ.Name, resQ.CreationTimestamp)
		}
	}
	w.Flush()
}

func printSecrets(secrets *v1.SecretList, resName *string) {
	fmt.Printf("\nSecrets\n--------------\n")
	w := tabwriter.NewWriter(os.Stdout, 0, 0, 1, ' ', 0)
	fmt.Fprintf(w, "%v\t%v\t%v\t%v\n", "NAME", "TYPE", "DATA", "AGE")
	for _, secret := range secrets.Items {
		if *resName != "" {
			if strings.Contains(secret.Name, *resName) {
				fmt.Fprintf(w, "%v\t%v\t%v\t%v\n", secret.Name, secret.Type, len(secret.Data), "")
			}
		} else {
			fmt.Fprintf(w, "%v\t%v\t%v\t%v\n", secret.Name, secret.Type, len(secret.Data), "")
		}
	}
	w.Flush()
}

func printServices(services *v1.ServiceList, resName *string) {
	fmt.Printf("\nServices\n--------------\n")
	w := tabwriter.NewWriter(os.Stdout, 0, 0, 1, ' ', 0)
	fmt.Fprintf(w, "%v\t%v\t%v\t%v\t%v\t%v\n", "NAME", "TYPE", "CLUSTER-IP", "EXTERNAL-IP", "PORT(S)", "AGE")

	for _, service := range services.Items {
		if *resName != "" {
			if strings.Contains(service.Name, *resName) {
				fmt.Fprintf(w, "%v\t%v\t%v\t%v\t%v\t%v\n", service.Name, service.Spec.Type, service.Spec.ClusterIP, service.Spec.ExternalIPs, service.Spec.Ports, "")
			}
		} else {
			fmt.Fprintf(w, "%v\t%v\t%v\t%v\t%v\t%v\n", service.Name, service.Spec.Type, service.Spec.ClusterIP, service.Spec.ExternalIPs, service.Spec.Ports, "")
		}
	}
	w.Flush()
}

func printServiceAccounts(serviceAccs *v1.ServiceAccountList, resName *string) {
	fmt.Printf("\nServiceAccounts\n--------------\n")
	w := tabwriter.NewWriter(os.Stdout, 0, 0, 1, ' ', 0)
	fmt.Fprintf(w, "%v\t%v\t%v\n", "NAME", "SECRETS", "AGE")
	for _, serviceAcc := range serviceAccs.Items {
		if *resName != "" {
			if strings.Contains(serviceAcc.Name, *resName) {
				fmt.Fprintf(w, "%v\t%v\t%v\n", serviceAcc.Name, len(serviceAcc.Secrets), "")
			}
		} else {
			fmt.Fprintf(w, "%v\t%v\t%v\n", serviceAcc.Name, len(serviceAcc.Secrets), "")
		}
	}
	w.Flush()
}

func printDaemonSets(daemonsets *appsv1.DaemonSetList, resName *string) {
	fmt.Printf("\nDaemonSets\n--------------\n")
	w := tabwriter.NewWriter(os.Stdout, 0, 0, 1, ' ', 0)
	fmt.Fprintf(w, "%v\t%v\t%v\t%v\t%v\t%v\t%v\t%v\t%v\n", "NAMESPACE", "NAME", "DESIRED", "CURRENT", "READY", "UP-TO-DATE", "AVAILABLE", "NODE SELECTOR", "AGE")
	for _, ds := range daemonsets.Items {
		if *resName != "" {
			if strings.Contains(ds.Name, *resName) {
				fmt.Fprintf(w, "%v\t%v\t%v\t%v\t%v\t%v\t%v\t%v\t%v\n", ds.Namespace, ds.Name, ds.Status.DesiredNumberScheduled, ds.Status.CurrentNumberScheduled, ds.Status.NumberReady, "", ds.Status.NumberAvailable, ds.Spec.Template.Spec.NodeSelector, "")
			}
		} else {
			fmt.Fprintf(w, "%v\t%v\t%v\t%v\t%v\t%v\t%v\t%v\t%v\n", ds.Namespace, ds.Name, ds.Status.DesiredNumberScheduled, ds.Status.CurrentNumberScheduled, ds.Status.NumberReady, "", ds.Status.NumberAvailable, ds.Spec.Template.Spec.NodeSelector, "")
		}
	}
	w.Flush()
}

func printDeployments(deployments *appsv1.DeploymentList, resName *string) {
	fmt.Printf("\nDeployments\n--------------\n")
	w := tabwriter.NewWriter(os.Stdout, 0, 0, 1, ' ', 0)
	fmt.Fprintf(w, "%v\t%v\t%v\t%v\t%v\n", "NAME", "READY", "UP-TO-DATE", "AVAILABLE", "AGE")
	for _, deployment := range deployments.Items {
		if *resName != "" {
			if strings.Contains(deployment.Name, *resName) {
				fmt.Fprintf(w, "%v\t%v\t%v\t%v\t%v\n", deployment.Name, deployment.Status.ReadyReplicas, "", deployment.Status.AvailableReplicas, "")
			}
		} else {
			fmt.Fprintf(w, "%v\t%v\t%v\t%v\t%v\n", deployment.Name, deployment.Status.ReadyReplicas, "", deployment.Status.AvailableReplicas, "")
		}
	}
	w.Flush()
}

func printReplicaSets(rsets *appsv1.ReplicaSetList, resName *string) {
	fmt.Printf("\nReplicaSets\n--------------\n")
	w := tabwriter.NewWriter(os.Stdout, 0, 0, 1, ' ', 0)
	fmt.Fprintf(w, "%v\t%v\t%v\t%v\t%v\n", "NAME", "DESIRED", "CURRENT", "READY", "AGE")
	for _, rs := range rsets.Items {
		if *resName != "" {
			if strings.Contains(rs.Name, *resName) {
				fmt.Fprintf(w, "%v\t%v\t%v\t%v\t%v\n", rs.Name, "", "", "", "")
			}
		} else {
			fmt.Fprintf(w, "%v\t%v\t%v\t%v\t%v\n", rs.Name, "", "", "", "")
		}
	}
	w.Flush()
}

func printStateFulSets(ssets *appsv1.StatefulSetList, resName *string) {
	fmt.Printf("\nStatefulSets\n--------------\n")
	w := tabwriter.NewWriter(os.Stdout, 0, 0, 1, ' ', 0)
	fmt.Fprintf(w, "%v\t%v\t%v\n", "NAME", "READY", "AGE")
	for _, sset := range ssets.Items {
		if *resName != "" {
			if strings.Contains(sset.Name, *resName) {
				fmt.Fprintf(w, "%v\t%v\t%v\n", sset.Name, sset.Status.ReadyReplicas, "")
			}
		} else {
			fmt.Fprintf(w, "%v\t%v\t%v\n", sset.Name, sset.Status.ReadyReplicas, "")
		}
	}
	w.Flush()
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
