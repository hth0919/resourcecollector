package CRM

import (
	"k8s.io/client-go/kubernetes"
)

type ClientSet struct {
	ClientSet        *kubernetes.Clientset
}