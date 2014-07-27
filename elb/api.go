package elb

import (
	"fmt"
	"net/http"
)

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

	optionSet := new(CreateLoadBalancerOptions)

	optionSet.LoadBalancerName = r.FormValue("LoadBalancerName")
	optionSet.Scheme = r.FormValue("Scheme")

	listenerSet := ListenerOptions{}
	listenerSet.LoadBalancerPort = r.FormValue("Listeners.member.1.LoadBalancerPort")
	listenerSet.InstancePort = r.FormValue("Listeners.member.1.InstancePort")
	listenerSet.Protocol = r.FormValue("Listeners.member.1.Protocol")
	listenerSet.InstanceProtocol = r.FormValue("Listeners.member.1.InstanceProtocol")

	optionSet.Listeners = append(optionSet.Listeners, listenerSet)

	fmt.Println(optionSet)
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

func parseMembersFromInput() {

}