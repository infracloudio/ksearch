package printers

import (
	"fmt"
	"os"
	"strings"
	"text/tabwriter"

	v1 "k8s.io/api/core/v1"

	appsv1 "k8s.io/api/apps/v1"
)

func PrintPodDetails(pods *v1.PodList, resName string) {
	fmt.Printf("\nPods\n----\n")
	w := tabwriter.NewWriter(os.Stdout, 0, 0, 1, ' ', 0)
	fmt.Fprintf(w, "%v\t%v\t%v\t%v\n", "NAME", "READY", "STATUS", "RESTARTS")

	for _, pod := range pods.Items {
		if resName != "" {
			if strings.Contains(pod.Name, resName) {
				fmt.Fprintf(w, "%v\t%v\t%v\t%v\n", pod.Name, "", pod.Status.Phase, "")
			}
		} else {
			fmt.Fprintf(w, "%v\t%v\t%v\t%v\n", pod.Name, "", pod.Status.Phase, "")
		}
	}
	w.Flush()
}
func PrintPodTemplates(podTemplates *v1.PodTemplateList, resName string) {
	fmt.Printf("\nPodTemplates\n--------------\n")
	w := tabwriter.NewWriter(os.Stdout, 0, 0, 1, ' ', 0)
	fmt.Fprintf(w, "%v\n", "NAME")
	for _, podTemplate := range podTemplates.Items {
		if resName != "" {
			if strings.Contains(podTemplate.Name, resName) {
				fmt.Fprintf(w, "%v\n", podTemplate.Name)
			}
		} else {
			fmt.Fprintf(w, "%v\n", podTemplate.Name)
		}
	}
	w.Flush()
}
func PrintComponentStatuses(componentStatuses *v1.ComponentStatusList, resName string) {
	fmt.Printf("\nComponentStatuses\n-------------\n")
	w := tabwriter.NewWriter(os.Stdout, 0, 0, 1, ' ', 0)
	fmt.Fprintf(w, "%v\t%v\t%v\t%v\n", "NAME", "STATUS", "MESSAGE", "ERROR")
	for _, componentStatus := range componentStatuses.Items {
		if resName != "" {
			if strings.Contains(componentStatus.Name, resName) {
				fmt.Fprintf(w, "%v\t%v\t%v\t%v\n", componentStatus.Name, componentStatus.Conditions[0].Type, componentStatus.Conditions[0].Message, componentStatus.Conditions[0].Error)
			}
		} else {
			fmt.Fprintf(w, "%v\t%v\t%v\t%v\n", componentStatus.Name, componentStatus.Conditions[0].Type, componentStatus.Conditions[0].Message, componentStatus.Conditions[0].Error)
		}
	}
	w.Flush()
}
func PrintConfigMaps(cms *v1.ConfigMapList, resName string) {
	fmt.Printf("\nConfigMaps\n--------------\n")
	w := tabwriter.NewWriter(os.Stdout, 0, 0, 1, ' ', 0)
	fmt.Fprintf(w, "%v\t%v\t%v\n", "NAME", "DATA", "AGE")
	for _, configMap := range cms.Items {
		if resName != "" {
			if strings.Contains(configMap.Name, resName) {
				fmt.Fprintf(w, "%v\t%v\t%v\n", configMap.Name, len(configMap.Data), "")
			}
		} else {
			fmt.Fprintf(w, "%v\t%v\t%v\n", configMap.Name, len(configMap.Data), "")
		}
	}
	w.Flush()
}
func PrintEndpoints(endPoints *v1.EndpointsList, resName string) {
	fmt.Printf("\nEndpoints\n--------------\n")
	w := tabwriter.NewWriter(os.Stdout, 0, 0, 1, ' ', 0)
	fmt.Fprintf(w, "%v\t%v\t%v\n", "NAME", "ENDPOINTS", "AGE")
	for _, endpoint := range endPoints.Items {
		if resName != "" {
			if strings.Contains(endpoint.Name, resName) {
				fmt.Fprintf(w, "%v\t%v\t%v\n", endpoint.Name, "", "")
			}
		} else {
			fmt.Fprintf(w, "%v\t%v\t%v\n", endpoint.Name, "", "")
		}
	}
	w.Flush()
}
func PrintEvents(events *v1.EventList, resName string) {
	fmt.Printf("\nEvents\n--------------\n")
	w := tabwriter.NewWriter(os.Stdout, 0, 0, 1, ' ', 0)
	fmt.Fprintf(w, "%v\t%v\t%v\t%v\t%v\t%v\n", "NAMESPACE", "LAST SEEN", "TYPE", "REASON", "OBJECT", "MESSAGE")
	for _, event := range events.Items {
		fmt.Fprintf(w, "%v\t%v\t%v\t%v\t%v\t%v\n", event.Namespace, "", event.Type, "", event.InvolvedObject.Kind+"/"+event.InvolvedObject.Name, event.Message)
	}
	w.Flush()
}
func PrintLimitRanges(limitRanges *v1.LimitRangeList, resName string) {
	fmt.Printf("\nLimitRanges\n--------------\n")
	w := tabwriter.NewWriter(os.Stdout, 0, 0, 1, ' ', 0)
	fmt.Fprintf(w, "%v\t%v\n", "NAME", "CREATED AT")
	for _, limitRange := range limitRanges.Items {
		if resName != "" {
			if strings.Contains(limitRange.Name, resName) {
				fmt.Fprintf(w, "%v\t%v\n", limitRange.Name, limitRange.CreationTimestamp)
			}
		} else {
			fmt.Fprintf(w, "%v\t%v\n", limitRange.Name, limitRange.CreationTimestamp)
		}
	}
	w.Flush()
}
func PrintNamespaces(namespaces *v1.NamespaceList, resName string) {
	fmt.Printf("\nNamespaces\n--------------\n")
	w := tabwriter.NewWriter(os.Stdout, 0, 0, 1, ' ', 0)
	fmt.Fprintf(w, "%v\t%v\t%v\n", "NAME", "STATUS", "AGE")
	for _, namespace := range namespaces.Items {
		if resName != "" {
			if strings.Contains(namespace.Name, resName) {
				fmt.Fprintf(w, "%v\t%v\t%v\n", namespace.Name, namespace.Status, "")
			}
		} else {
			fmt.Fprintf(w, "%v\t%v\t%v\n", namespace.Name, namespace.Status, "")
		}
	}
	w.Flush()
}
func PrintPVs(pvs *v1.PersistentVolumeList, resName string) {
	fmt.Printf("\nPersistentVolumes\n--------------\n")
	w := tabwriter.NewWriter(os.Stdout, 0, 0, 1, ' ', 0)
	fmt.Fprintf(w, "%v\t%v\t%v\t%v\t%v\t%v\t%v\t%v\t%v\n", "NAME", "CAPACITY", "ACCESS MODES", "RECLAIM POLICY", "STATUS", "CLAIM", "STORAGECLASS", "REASON", "AGE")

	for _, pv := range pvs.Items {
		if resName != "" {
			if strings.Contains(pv.Name, resName) {
				fmt.Fprintf(w, "%v\t%v\t%v\t%v\t%v\t%v\t%v\t%v\t%v\n", pv.Name, "", pv.Spec.AccessModes, pv.Spec.PersistentVolumeReclaimPolicy, pv.Status, pv.Spec.ClaimRef.Namespace+"/"+pv.Spec.ClaimRef.Name, pv.Spec.StorageClassName, pv.Status.Reason, "")
			}
		} else {
			fmt.Fprintf(w, "%v\t%v\t%v\t%v\t%v\t%v\t%v\t%v\t%v\n", pv.Name, "", pv.Spec.AccessModes, pv.Spec.PersistentVolumeReclaimPolicy, pv.Status, pv.Spec.ClaimRef.Namespace+"/"+pv.Spec.ClaimRef.Name, pv.Spec.StorageClassName, pv.Status.Reason, "")
		}
	}
	w.Flush()
}
func PrintPVCs(pvcs *v1.PersistentVolumeClaimList, resName string) {
	fmt.Printf("\nPersistentVolumeClaims\n--------------\n")
	w := tabwriter.NewWriter(os.Stdout, 0, 0, 1, ' ', 0)
	fmt.Fprintf(w, "%v\t%v\t%v\t%v\t%v\t%v\t%v\n", "NAME", "STATUS", "VOLUME", "CAPACITY", "ACCESS MODES", "STORAGECLASS", "AGE")
	for _, pvc := range pvcs.Items {
		if resName != "" {
			if strings.Contains(pvc.Name, resName) {
				fmt.Fprintf(w, "%v\t%v\t%v\t%v\t%v\t%v\t%v\n", pvc.Name, pvc.Status, "", pvc.Status.Capacity.Cpu(), pvc.Spec.AccessModes, pvc.Spec.StorageClassName, "")
			}
		} else {
			fmt.Fprintf(w, "%v\t%v\t%v\t%v\t%v\t%v\t%v\n", pvc.Name, pvc.Status, "", pvc.Status.Capacity.Cpu(), pvc.Spec.AccessModes, pvc.Spec.StorageClassName, "")
		}

	}
	w.Flush()
}

func PrintResourceQuotas(resQuotas *v1.ResourceQuotaList, resName string) {
	fmt.Printf("\nResourceQuotas\n--------------\n")
	w := tabwriter.NewWriter(os.Stdout, 0, 0, 1, ' ', 0)
	fmt.Fprintf(w, "%v\t%v\n", "NAME", "CREATED AT")
	for _, resQ := range resQuotas.Items {
		if resName != "" {
			if strings.Contains(resQ.Name, resName) {
				fmt.Fprintf(w, "%v\t%v\n", resQ.Name, resQ.CreationTimestamp)
			}
		} else {
			fmt.Fprintf(w, "%v\t%v\n", resQ.Name, resQ.CreationTimestamp)
		}
	}
	w.Flush()
}
func PrintSecrets(secrets *v1.SecretList, resName string) {
	fmt.Printf("\nSecrets\n--------------\n")
	w := tabwriter.NewWriter(os.Stdout, 0, 0, 1, ' ', 0)
	fmt.Fprintf(w, "%v\t%v\t%v\t%v\n", "NAME", "TYPE", "DATA", "AGE")
	for _, secret := range secrets.Items {
		if resName != "" {
			if strings.Contains(secret.Name, resName) {
				fmt.Fprintf(w, "%v\t%v\t%v\t%v\n", secret.Name, secret.Type, len(secret.Data), "")
			}
		} else {
			fmt.Fprintf(w, "%v\t%v\t%v\t%v\n", secret.Name, secret.Type, len(secret.Data), "")
		}
	}
	w.Flush()
}
func PrintServices(services *v1.ServiceList, resName string) {
	fmt.Printf("\nServices\n--------------\n")
	w := tabwriter.NewWriter(os.Stdout, 0, 0, 1, ' ', 0)
	fmt.Fprintf(w, "%v\t%v\t%v\t%v\t%v\t%v\n", "NAME", "TYPE", "CLUSTER-IP", "EXTERNAL-IP", "PORT(S)", "AGE")

	for _, service := range services.Items {
		if resName != "" {
			if strings.Contains(service.Name, resName) {
				fmt.Fprintf(w, "%v\t%v\t%v\t%v\t%v\t%v\n", service.Name, service.Spec.Type, service.Spec.ClusterIP, service.Spec.ExternalIPs, service.Spec.Ports, "")
			}
		} else {
			fmt.Fprintf(w, "%v\t%v\t%v\t%v\t%v\t%v\n", service.Name, service.Spec.Type, service.Spec.ClusterIP, service.Spec.ExternalIPs, service.Spec.Ports, "")
		}
	}
	w.Flush()
}
func PrintServiceAccounts(serviceAccs *v1.ServiceAccountList, resName string) {
	fmt.Printf("\nServiceAccounts\n--------------\n")
	w := tabwriter.NewWriter(os.Stdout, 0, 0, 1, ' ', 0)
	fmt.Fprintf(w, "%v\t%v\t%v\n", "NAME", "SECRETS", "AGE")
	for _, serviceAcc := range serviceAccs.Items {
		if resName != "" {
			if strings.Contains(serviceAcc.Name, resName) {
				fmt.Fprintf(w, "%v\t%v\t%v\n", serviceAcc.Name, len(serviceAcc.Secrets), "")
			}
		} else {
			fmt.Fprintf(w, "%v\t%v\t%v\n", serviceAcc.Name, len(serviceAcc.Secrets), "")
		}
	}
	w.Flush()
}
func PrintDaemonSets(daemonsets *appsv1.DaemonSetList, resName string) {
	fmt.Printf("\nDaemonSets\n--------------\n")
	w := tabwriter.NewWriter(os.Stdout, 0, 0, 1, ' ', 0)
	fmt.Fprintf(w, "%v\t%v\t%v\t%v\t%v\t%v\t%v\t%v\t%v\n", "NAMESPACE", "NAME", "DESIRED", "CURRENT", "READY", "UP-TO-DATE", "AVAILABLE", "NODE SELECTOR", "AGE")
	for _, ds := range daemonsets.Items {
		if resName != "" {
			if strings.Contains(ds.Name, resName) {
				fmt.Fprintf(w, "%v\t%v\t%v\t%v\t%v\t%v\t%v\t%v\t%v\n", ds.Namespace, ds.Name, ds.Status.DesiredNumberScheduled, ds.Status.CurrentNumberScheduled, ds.Status.NumberReady, "", ds.Status.NumberAvailable, ds.Spec.Template.Spec.NodeSelector, "")
			}
		} else {
			fmt.Fprintf(w, "%v\t%v\t%v\t%v\t%v\t%v\t%v\t%v\t%v\n", ds.Namespace, ds.Name, ds.Status.DesiredNumberScheduled, ds.Status.CurrentNumberScheduled, ds.Status.NumberReady, "", ds.Status.NumberAvailable, ds.Spec.Template.Spec.NodeSelector, "")
		}
	}
	w.Flush()
}
func PrintDeployments(deployments *appsv1.DeploymentList, resName string) {
	fmt.Printf("\nDeployments\n--------------\n")
	w := tabwriter.NewWriter(os.Stdout, 0, 0, 1, ' ', 0)
	fmt.Fprintf(w, "%v\t%v\t%v\t%v\t%v\n", "NAME", "READY", "UP-TO-DATE", "AVAILABLE", "AGE")
	for _, deployment := range deployments.Items {
		if resName != "" {
			if strings.Contains(deployment.Name, resName) {
				fmt.Fprintf(w, "%v\t%v\t%v\t%v\t%v\n", deployment.Name, deployment.Status.ReadyReplicas, "", deployment.Status.AvailableReplicas, "")
			}
		} else {
			fmt.Fprintf(w, "%v\t%v\t%v\t%v\t%v\n", deployment.Name, deployment.Status.ReadyReplicas, "", deployment.Status.AvailableReplicas, "")
		}
	}
	w.Flush()
}
func PrintReplicaSets(rsets *appsv1.ReplicaSetList, resName string) {
	fmt.Printf("\nReplicaSets\n--------------\n")
	w := tabwriter.NewWriter(os.Stdout, 0, 0, 1, ' ', 0)
	fmt.Fprintf(w, "%v\t%v\t%v\t%v\t%v\n", "NAME", "DESIRED", "CURRENT", "READY", "AGE")
	for _, rs := range rsets.Items {
		if resName != "" {
			if strings.Contains(rs.Name, resName) {
				fmt.Fprintf(w, "%v\t%v\t%v\t%v\t%v\n", rs.Name, "", "", "", "")
			}
		} else {
			fmt.Fprintf(w, "%v\t%v\t%v\t%v\t%v\n", rs.Name, "", "", "", "")
		}
	}
	w.Flush()
}
func PrintStateFulSets(ssets *appsv1.StatefulSetList, resName string) {
	fmt.Printf("\nStatefulSets\n--------------\n")
	w := tabwriter.NewWriter(os.Stdout, 0, 0, 1, ' ', 0)
	fmt.Fprintf(w, "%v\t%v\t%v\n", "NAME", "READY", "AGE")
	for _, sset := range ssets.Items {
		if resName != "" {
			if strings.Contains(sset.Name, resName) {
				fmt.Fprintf(w, "%v\t%v\t%v\n", sset.Name, sset.Status.ReadyReplicas, "")
			}
		} else {
			fmt.Fprintf(w, "%v\t%v\t%v\n", sset.Name, sset.Status.ReadyReplicas, "")
		}
	}
	w.Flush()
}
