// Code generated by informer-gen. DO NOT EDIT.

package v1alpha1

import (
	internalinterfaces "github.com/daocloud/kubean/pkg/generated/kubeanclusterops/informers/externalversions/internalinterfaces"
)

// Interface provides access to all the informers in this group version.
type Interface interface {
	// KuBeanClusterOpses returns a KuBeanClusterOpsInformer.
	KuBeanClusterOpses() KuBeanClusterOpsInformer
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

// KuBeanClusterOpses returns a KuBeanClusterOpsInformer.
func (v *version) KuBeanClusterOpses() KuBeanClusterOpsInformer {
	return &kuBeanClusterOpsInformer{factory: v.factory, tweakListOptions: v.tweakListOptions}
}
