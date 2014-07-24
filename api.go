package main

import (
	"fmt"
	"net/http"
)

type CreateLoadBalancerOptions struct {
	AvailabilityZones string
	Listeners         []ListenerOptions
	LoadBalancerName  string
	Scheme            string
	SecurityGroups    []string
	Subnets           []string
}

type CreateLoadBalancerListenersOptions struct {
	Listeners        []ListenerOptions
	LoadBalancerName string
}

type CreateLoadBalancerPolicyOptions struct {
	PolicyAttributes []string
	PolicyName       string
	PolicyTypeName   string
}

type RegisterInstancesWithLoadBalancerOptions struct {
	Instances        []string
	LoadBalancerName string
}

type ListenerOptions struct {
	LoadBalancerPort int
	InstancePort     int
	InstanceProtocol string
	SSLCertificateId string
}

func SetupApiHandlers() {
	http.HandleFunc("/", ELBHandler)
	http.ListenAndServe(":8080", nil)
}

func ELBHandler(w http.ResponseWriter, r *http.Request) {
	action := r.FormValue("Action")
	version := r.FormValue("Version")

	fmt.Println(action, version)

	// This is going to be nasty, but we need to redispatch the
	// request to the correct handler.
	switch action {
	case "CreateLoadBalancer":
		CreateLoadBalancerHandler(w, r)
	case "CreateLoadBalancerListeners":
		CreateLoadBalancerListenersHandler(w, r)
	case "CreateLoadBalancerPolicy":
		CreateLoadBalancerPolicyHandler(w, r)
	case "RegisterInstancesWithLoadBalancer":
		RegisterInstancesWithLoadBalancerHandler(w, r)
	}
}

func CreateLoadBalancerHandler(w http.ResponseWriter, r *http.Request) {
	loadBalancerName := r.FormValue("LoadBalancerName")
	fmt.Println(loadBalancerName)
}

func CreateLoadBalancerListenersHandler(w http.ResponseWriter, r *http.Request) {
	loadBalancerName := r.FormValue("LoadBalancerName")
	fmt.Println(loadBalancerName)
}

func CreateLoadBalancerPolicyHandler(w http.ResponseWriter, r *http.Request) {
	loadBalancerName := r.FormValue("LoadBalancerName")
	fmt.Println(loadBalancerName)
}

func RegisterInstancesWithLoadBalancerHandler(w http.ResponseWriter, r *http.Request) {
	loadBalancerName := r.FormValue("LoadBalancerName")
	fmt.Println(loadBalancerName)
}
