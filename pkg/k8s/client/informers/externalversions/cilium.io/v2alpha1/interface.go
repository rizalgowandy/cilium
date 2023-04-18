// SPDX-License-Identifier: Apache-2.0
// Copyright Authors of Cilium

// Code generated by informer-gen. DO NOT EDIT.

package v2alpha1

import (
	internalinterfaces "github.com/cilium/cilium/pkg/k8s/client/informers/externalversions/internalinterfaces"
)

// Interface provides access to all the informers in this group version.
type Interface interface {
	// CiliumBGPPeeringPolicies returns a CiliumBGPPeeringPolicyInformer.
	CiliumBGPPeeringPolicies() CiliumBGPPeeringPolicyInformer
	// CiliumCIDRGroups returns a CiliumCIDRGroupInformer.
	CiliumCIDRGroups() CiliumCIDRGroupInformer
	// CiliumEndpointSlices returns a CiliumEndpointSliceInformer.
	CiliumEndpointSlices() CiliumEndpointSliceInformer
	// CiliumLoadBalancerIPPools returns a CiliumLoadBalancerIPPoolInformer.
	CiliumLoadBalancerIPPools() CiliumLoadBalancerIPPoolInformer
	// CiliumNodeConfigs returns a CiliumNodeConfigInformer.
	CiliumNodeConfigs() CiliumNodeConfigInformer
}

type version struct {
	factory          internalinterfaces.SharedInformerFactory
	namespace        string
	tweakListOptions internalinterfaces.TweakListOptionsFunc
}

// New returns a new Interface.
func New(f internalinterfaces.SharedInformerFactory, namespace string, tweakListOptions internalinterfaces.TweakListOptionsFunc) Interface {
	return &version{factory: f, namespace: namespace, tweakListOptions: tweakListOptions}
}

// CiliumBGPPeeringPolicies returns a CiliumBGPPeeringPolicyInformer.
func (v *version) CiliumBGPPeeringPolicies() CiliumBGPPeeringPolicyInformer {
	return &ciliumBGPPeeringPolicyInformer{factory: v.factory, tweakListOptions: v.tweakListOptions}
}

// CiliumCIDRGroups returns a CiliumCIDRGroupInformer.
func (v *version) CiliumCIDRGroups() CiliumCIDRGroupInformer {
	return &ciliumCIDRGroupInformer{factory: v.factory, tweakListOptions: v.tweakListOptions}
}

// CiliumEndpointSlices returns a CiliumEndpointSliceInformer.
func (v *version) CiliumEndpointSlices() CiliumEndpointSliceInformer {
	return &ciliumEndpointSliceInformer{factory: v.factory, tweakListOptions: v.tweakListOptions}
}

// CiliumLoadBalancerIPPools returns a CiliumLoadBalancerIPPoolInformer.
func (v *version) CiliumLoadBalancerIPPools() CiliumLoadBalancerIPPoolInformer {
	return &ciliumLoadBalancerIPPoolInformer{factory: v.factory, tweakListOptions: v.tweakListOptions}
}

// CiliumNodeConfigs returns a CiliumNodeConfigInformer.
func (v *version) CiliumNodeConfigs() CiliumNodeConfigInformer {
	return &ciliumNodeConfigInformer{factory: v.factory, namespace: v.namespace, tweakListOptions: v.tweakListOptions}
}
