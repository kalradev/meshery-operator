package istio

import (
	"log"

	networkingV1alpha3 "istio.io/client-go/pkg/apis/networking/v1alpha3"
	networkingV1beta1 "istio.io/client-go/pkg/apis/networking/v1beta1"
	securityV1beta1 "istio.io/client-go/pkg/apis/security/v1beta1"

	broker "github.com/layer5io/meshery-operator/pkg/broker"
	discovery "github.com/layer5io/meshery-operator/pkg/discovery"
	inf "github.com/layer5io/meshery-operator/pkg/informers"
	informers "github.com/layer5io/meshery-operator/pkg/meshsync/meshes/istio/informers"
	pipeline "github.com/layer5io/meshery-operator/pkg/meshsync/meshes/istio/pipeline"
)

type Resources struct {
	VirtualServices        []networkingV1beta1.VirtualService      `json:"virtualservices,omitempty"`
	Sidecars               []networkingV1beta1.Sidecar             `json:"sidecars,omitempty"`
	WorkloadEntries        []networkingV1beta1.WorkloadEntry       `json:"workloadentries,omitempty"`
	AuthorizationPolicies  []securityV1beta1.AuthorizationPolicy   `json:"authorizationpolicies,omitempty"`
	DestinationRules       []networkingV1beta1.DestinationRule     `json:"destinationrules,omitempty"`
	EnvoyFilters           []networkingV1alpha3.EnvoyFilter        `json:"envoyfilters,omitempty"`
	Gateways               []networkingV1beta1.Gateway             `json:"gateways,omitempty"`
	PeerAuthentications    []securityV1beta1.PeerAuthentication    `json:"peerauthenticatons,omitempty"`
	RequestAuthentications []securityV1beta1.RequestAuthentication `json:"requestauthentications,omitempty"`
	ServiceEntries         []networkingV1beta1.ServiceEntry        `json:"serviceentries,omitempty"`
	WorkloadGroups         []networkingV1alpha3.WorkloadGroup      `json:"workloadgroups,omitempty"`
}

func StartDiscovery(dclient *discovery.Client, broker broker.Broker) error {
	// Get pipeline instance
	pl := pipeline.Initialize(dclient, broker)
	// run pipelines
	result := pl.Run()
	if result.Error != nil {
		log.Printf("Error executing cluster pipeline: %s", result.Error)
		return result.Error
	}
	return nil
}

func StartInformer(iclient *inf.Client, broker broker.Broker) {
	informers.Initialize(iclient, broker)
}
