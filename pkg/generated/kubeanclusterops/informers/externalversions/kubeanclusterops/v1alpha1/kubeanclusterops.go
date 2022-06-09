// Code generated by informer-gen. DO NOT EDIT.

package v1alpha1

import (
	"context"
	time "time"

	kubeanclusteropsv1alpha1 "github.com/daocloud/kubean/pkg/apis/kubeanclusterops/v1alpha1"
	versioned "github.com/daocloud/kubean/pkg/generated/kubeanclusterops/clientset/versioned"
	internalinterfaces "github.com/daocloud/kubean/pkg/generated/kubeanclusterops/informers/externalversions/internalinterfaces"
	v1alpha1 "github.com/daocloud/kubean/pkg/generated/kubeanclusterops/listers/kubeanclusterops/v1alpha1"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	runtime "k8s.io/apimachinery/pkg/runtime"
	watch "k8s.io/apimachinery/pkg/watch"
	cache "k8s.io/client-go/tools/cache"
)

// KuBeanClusterOpsInformer provides access to a shared informer and lister for
// KuBeanClusterOpses.
type KuBeanClusterOpsInformer interface {
	Informer() cache.SharedIndexInformer
	Lister() v1alpha1.KuBeanClusterOpsLister
}

type kuBeanClusterOpsInformer struct {
	factory          internalinterfaces.SharedInformerFactory
	tweakListOptions internalinterfaces.TweakListOptionsFunc
}

// NewKuBeanClusterOpsInformer constructs a new informer for KuBeanClusterOps type.
// Always prefer using an informer factory to get a shared informer instead of getting an independent
// one. This reduces memory footprint and number of connections to the server.
func NewKuBeanClusterOpsInformer(client versioned.Interface, resyncPeriod time.Duration, indexers cache.Indexers) cache.SharedIndexInformer {
	return NewFilteredKuBeanClusterOpsInformer(client, resyncPeriod, indexers, nil)
}

// NewFilteredKuBeanClusterOpsInformer constructs a new informer for KuBeanClusterOps type.
// Always prefer using an informer factory to get a shared informer instead of getting an independent
// one. This reduces memory footprint and number of connections to the server.
func NewFilteredKuBeanClusterOpsInformer(client versioned.Interface, resyncPeriod time.Duration, indexers cache.Indexers, tweakListOptions internalinterfaces.TweakListOptionsFunc) cache.SharedIndexInformer {
	return cache.NewSharedIndexInformer(
		&cache.ListWatch{
			ListFunc: func(options v1.ListOptions) (runtime.Object, error) {
				if tweakListOptions != nil {
					tweakListOptions(&options)
				}
				return client.KubeanclusteropsV1alpha1().KuBeanClusterOpses().List(context.TODO(), options)
			},
			WatchFunc: func(options v1.ListOptions) (watch.Interface, error) {
				if tweakListOptions != nil {
					tweakListOptions(&options)
				}
				return client.KubeanclusteropsV1alpha1().KuBeanClusterOpses().Watch(context.TODO(), options)
			},
		},
		&kubeanclusteropsv1alpha1.KuBeanClusterOps{},
		resyncPeriod,
		indexers,
	)
}

func (f *kuBeanClusterOpsInformer) defaultInformer(client versioned.Interface, resyncPeriod time.Duration) cache.SharedIndexInformer {
	return NewFilteredKuBeanClusterOpsInformer(client, resyncPeriod, cache.Indexers{cache.NamespaceIndex: cache.MetaNamespaceIndexFunc}, f.tweakListOptions)
}

func (f *kuBeanClusterOpsInformer) Informer() cache.SharedIndexInformer {
	return f.factory.InformerFor(&kubeanclusteropsv1alpha1.KuBeanClusterOps{}, f.defaultInformer)
}

func (f *kuBeanClusterOpsInformer) Lister() v1alpha1.KuBeanClusterOpsLister {
	return v1alpha1.NewKuBeanClusterOpsLister(f.Informer().GetIndexer())
}