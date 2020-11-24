package cache

import (
	"github.com/layer5io/meshery-operator/pkg/meshsync/cluster"
	"github.com/layer5io/meshery-operator/pkg/meshsync/meshes/istio"
)

type Resources struct {
	Cluster cluster.Resources `json:"cluster,omitempty"`
	Istio   istio.Resources   `json:"istio,omitempty"`
}

type Cache interface {
	Read(result *Resources) error
	Write(result *Resources) error
}
