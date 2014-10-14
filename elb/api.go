package elb

import (
	"fmt"
	"net/http"
	"github.com/stevencorona/elastic-haproxy/haproxy"
)

type ElbApi struct {
	Cluster haproxy.Cluster
}

func InitApiHandlers() {

	api := new(ElbApi)

	http.HandleFunc("/", api.ELBHandler)
	http.ListenAndServe(":8080", nil)
}

func (api *ElbApi) ELBHandler(w http.ResponseWriter, r *http.Request) {
	action := r.FormValue("Action")
	version := r.FormValue("Version")

	fmt.Println(action, version)

	// This is going to be nasty, but we need to redispatch the
	// request to the correct handler.
	switch action {
	case "CreateLoadBalancer":
		api.CreateLoadBalancerHandler(w, r)
	case "CreateLoadBalancerListeners":
		api.CreateLoadBalancerListenersHandler(w, r)
	case "CreateLoadBalancerPolicy":
		api.CreateLoadBalancerPolicyHandler(w, r)
	case "RegisterInstancesWithLoadBalancer":
		api.RegisterInstancesWithLoadBalancerHandler(w, r)
	}
}

func (api *ElbAPi) CreateLoadBalancerHandler(w http.ResponseWriter, r *http.Request) {

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

func (api *ElbAPi) CreateLoadBalancerListenersHandler(w http.ResponseWriter, r *http.Request) {
	loadBalancerName := r.FormValue("LoadBalancerName")
	fmt.Println(loadBalancerName)
}

func (api *ElbAPi) CreateLoadBalancerPolicyHandler(w http.ResponseWriter, r *http.Request) {
	loadBalancerName := r.FormValue("LoadBalancerName")
	fmt.Println(loadBalancerName)
}

func (api *ElbAPi) RegisterInstancesWithLoadBalancerHandler(w http.ResponseWriter, r *http.Request) {
	loadBalancerName := r.FormValue("LoadBalancerName")
	fmt.Println(loadBalancerName)
}

func (api *ElbAPi) parseMembersFromInput() {

}
