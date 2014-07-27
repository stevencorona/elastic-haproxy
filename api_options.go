package main

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
	PolicyAttributes []PolicyAttributeOptions
	PolicyName       string
	PolicyTypeName   string
}

type RegisterInstancesWithLoadBalancerOptions struct {
	Instances        []InstanceOptions
	LoadBalancerName string
}

type ListenerOptions struct {
	LoadBalancerPort string
	Protocol         string
	InstancePort     string
	InstanceProtocol string
	SSLCertificateId string
	SSLCertificate   string
}

type InstanceOptions struct {
	InstanceId string
	IpAddress  string
}

type PolicyAttributeOptions struct {
	AttributeName  string
	AttributeValue string
}
