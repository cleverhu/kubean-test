// Code generated by informer-gen. DO NOT EDIT.

package v1alpha1

import (
	"context"
	time "time"

	kubeanclusterv1alpha1 "github.com/daocloud/kubean/pkg/apis/kubeancluster/v1alpha1"
	versioned "github.com/daocloud/kubean/pkg/generated/kubeancluster/clientset/versioned"
	internalinterfaces "github.com/daocloud/kubean/pkg/generated/kubeancluster/informers/externalversions/internalinterfaces"
	v1alpha1 "github.com/daocloud/kubean/pkg/generated/kubeancluster/listers/kubeancluster/v1alpha1"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	runtime "k8s.io/apimachinery/pkg/runtime"
	watch "k8s.io/apimachinery/pkg/watch"
	cache "k8s.io/client-go/tools/cache"
)

// KuBeanClusterInformer provides access to a shared informer and lister for
// KuBeanClusters.
type KuBeanClusterInformer interface {
	Informer() cache.SharedIndexInformer
	Lister() v1alpha1.KuBeanClusterLister
}

type kuBeanClusterInformer struct {
	factory          internalinterfaces.SharedInformerFactory
	tweakListOptions internalinterfaces.TweakListOptionsFunc
}

// NewKuBeanClusterInformer constructs a new informer for KuBeanCluster type.
// Always prefer using an informer factory to get a shared informer instead of getting an independent
// one. This reduces memory footprint and number of connections to the server.
func NewKuBeanClusterInformer(client versioned.Interface, resyncPeriod time.Duration, indexers cache.Indexers) cache.SharedIndexInformer {
	return NewFilteredKuBeanClusterInformer(client, resyncPeriod, indexers, nil)
}

// NewFilteredKuBeanClusterInformer constructs a new informer for KuBeanCluster type.
// Always prefer using an informer factory to get a shared informer instead of getting an independent
// one. This reduces memory footprint and number of connections to the server.
func NewFilteredKuBeanClusterInformer(client versioned.Interface, resyncPeriod time.Duration, indexers cache.Indexers, tweakListOptions internalinterfaces.TweakListOptionsFunc) cache.SharedIndexInformer {
	return cache.NewSharedIndexInformer(
		&cache.ListWatch{
			ListFunc: func(options v1.ListOptions) (runtime.Object, error) {
				if tweakListOptions != nil {
					tweakListOptions(&options)
				}
				return client.KubeanclusterV1alpha1().KuBeanClusters().List(context.TODO(), options)
			},
			WatchFunc: func(options v1.ListOptions) (watch.Interface, error) {
				if tweakListOptions != nil {
					tweakListOptions(&options)
				}
				return client.KubeanclusterV1alpha1().KuBeanClusters().Watch(context.TODO(), options)
			},
		},
		&kubeanclusterv1alpha1.KuBeanCluster{},
		resyncPeriod,
		indexers,
	)
}

func (f *kuBeanClusterInformer) defaultInformer(client versioned.Interface, resyncPeriod time.Duration) cache.SharedIndexInformer {
	return NewFilteredKuBeanClusterInformer(client, resyncPeriod, cache.Indexers{cache.NamespaceIndex: cache.MetaNamespaceIndexFunc}, f.tweakListOptions)
}

func (f *kuBeanClusterInformer) Informer() cache.SharedIndexInformer {
	return f.factory.InformerFor(&kubeanclusterv1alpha1.KuBeanCluster{}, f.defaultInformer)
}

func (f *kuBeanClusterInformer) Lister() v1alpha1.KuBeanClusterLister {
	return v1alpha1.NewKuBeanClusterLister(f.Informer().GetIndexer())
}