// This file was automatically generated by informer-gen

package v1

import (
	internalinterfaces "github.com/openshift/origin/pkg/sdn/generated/informers/externalversions/internalinterfaces"
)

// Interface provides access to all the informers in this group version.
type Interface interface {
	// ClusterNetworks returns a ClusterNetworkInformer.
	ClusterNetworks() ClusterNetworkInformer
}

type version struct {
	internalinterfaces.SharedInformerFactory
}

// New returns a new Interface.
func New(f internalinterfaces.SharedInformerFactory) Interface {
	return &version{f}
}

// ClusterNetworks returns a ClusterNetworkInformer.
func (v *version) ClusterNetworks() ClusterNetworkInformer {
	return &clusterNetworkInformer{factory: v.SharedInformerFactory}
}
