// Code generated by lister-gen. DO NOT EDIT.

package v1alpha1

import (
	v1alpha1 "github.com/daocloud/kubean/pkg/apis/kubeanclusterops/v1alpha1"
	"k8s.io/apimachinery/pkg/api/errors"
	"k8s.io/apimachinery/pkg/labels"
	"k8s.io/client-go/tools/cache"
)

// KuBeanClusterOpsLister helps list KuBeanClusterOpses.
// All objects returned here must be treated as read-only.
type KuBeanClusterOpsLister interface {
	// List lists all KuBeanClusterOpses in the indexer.
	// Objects returned here must be treated as read-only.
	List(selector labels.Selector) (ret []*v1alpha1.KuBeanClusterOps, err error)
	// Get retrieves the KuBeanClusterOps from the index for a given name.
	// Objects returned here must be treated as read-only.
	Get(name string) (*v1alpha1.KuBeanClusterOps, error)
	KuBeanClusterOpsListerExpansion
}

// kuBeanClusterOpsLister implements the KuBeanClusterOpsLister interface.
type kuBeanClusterOpsLister struct {
	indexer cache.Indexer
}

// NewKuBeanClusterOpsLister returns a new KuBeanClusterOpsLister.
func NewKuBeanClusterOpsLister(indexer cache.Indexer) KuBeanClusterOpsLister {
	return &kuBeanClusterOpsLister{indexer: indexer}
}

// List lists all KuBeanClusterOpses in the indexer.
func (s *kuBeanClusterOpsLister) List(selector labels.Selector) (ret []*v1alpha1.KuBeanClusterOps, err error) {
	err = cache.ListAll(s.indexer, selector, func(m interface{}) {
		ret = append(ret, m.(*v1alpha1.KuBeanClusterOps))
	})
	return ret, err
}

// Get retrieves the KuBeanClusterOps from the index for a given name.
func (s *kuBeanClusterOpsLister) Get(name string) (*v1alpha1.KuBeanClusterOps, error) {
	obj, exists, err := s.indexer.GetByKey(name)
	if err != nil {
		return nil, err
	}
	if !exists {
		return nil, errors.NewNotFound(v1alpha1.Resource("kubeanclusterops"), name)
	}
	return obj.(*v1alpha1.KuBeanClusterOps), nil
}
