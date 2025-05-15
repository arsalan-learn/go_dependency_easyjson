package main

import (
	"fmt"

	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/kubernetes"
	ctrl "sigs.k8s.io/controller-runtime"
	"sigs.k8s.io/controller-runtime/pkg/client"
)

// This file exists to demonstrate all the dependencies for SCA scanning

// DemoResource is a demo resource type
type DemoResource struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              DemoSpec   `json:"spec,omitempty"`
	Status            DemoStatus `json:"status,omitempty"`
}

// DemoSpec represents the spec for DemoResource
type DemoSpec struct {
	Size int `json:"size"`
}

// DemoStatus represents the status for DemoResource
type DemoStatus struct {
	Ready bool `json:"ready"`
}

// DemoClient demonstrates client-go usage
func DemoClient() {
	// This function exists just to demonstrate dependencies
	fmt.Println("Demonstrating client-go usage")
	config, _ := ctrl.GetConfig()
	clientset, _ := kubernetes.NewForConfig(config)
	_ = clientset
}

// DemoController demonstrates controller-runtime usage
func DemoController() {
	// This function exists just to demonstrate dependencies
	fmt.Println("Demonstrating controller-runtime usage")
	c, _ := client.New(ctrl.GetConfigOrDie(), client.Options{})
	_ = c
}
