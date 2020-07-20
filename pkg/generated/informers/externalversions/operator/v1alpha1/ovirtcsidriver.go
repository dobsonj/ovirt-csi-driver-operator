// Code generated by informer-gen. DO NOT EDIT.

package v1alpha1

import (
	"context"
	time "time"

	operatorv1alpha1 "github.com/ovirt/csi-driver-operator/pkg/apis/operator/v1alpha1"
	versioned "github.com/ovirt/csi-driver-operator/pkg/generated/clientset/versioned"
	internalinterfaces "github.com/ovirt/csi-driver-operator/pkg/generated/informers/externalversions/internalinterfaces"
	v1alpha1 "github.com/ovirt/csi-driver-operator/pkg/generated/listers/operator/v1alpha1"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	runtime "k8s.io/apimachinery/pkg/runtime"
	watch "k8s.io/apimachinery/pkg/watch"
	cache "k8s.io/client-go/tools/cache"
)

// OvirtCSIDriverInformer provides access to a shared informer and lister for
// OvirtCSIDrivers.
type OvirtCSIDriverInformer interface {
	Informer() cache.SharedIndexInformer
	Lister() v1alpha1.OvirtCSIDriverLister
}

type ovirtCSIDriverInformer struct {
	factory          internalinterfaces.SharedInformerFactory
	tweakListOptions internalinterfaces.TweakListOptionsFunc
}

// NewOvirtCSIDriverInformer constructs a new informer for OvirtCSIDriver type.
// Always prefer using an informer factory to get a shared informer instead of getting an independent
// one. This reduces memory footprint and number of connections to the server.
func NewOvirtCSIDriverInformer(client versioned.Interface, resyncPeriod time.Duration, indexers cache.Indexers) cache.SharedIndexInformer {
	return NewFilteredOvirtCSIDriverInformer(client, resyncPeriod, indexers, nil)
}

// NewFilteredOvirtCSIDriverInformer constructs a new informer for OvirtCSIDriver type.
// Always prefer using an informer factory to get a shared informer instead of getting an independent
// one. This reduces memory footprint and number of connections to the server.
func NewFilteredOvirtCSIDriverInformer(client versioned.Interface, resyncPeriod time.Duration, indexers cache.Indexers, tweakListOptions internalinterfaces.TweakListOptionsFunc) cache.SharedIndexInformer {
	return cache.NewSharedIndexInformer(
		&cache.ListWatch{
			ListFunc: func(options v1.ListOptions) (runtime.Object, error) {
				if tweakListOptions != nil {
					tweakListOptions(&options)
				}
				return client.CsiV1alpha1().OvirtCSIDrivers().List(context.TODO(), options)
			},
			WatchFunc: func(options v1.ListOptions) (watch.Interface, error) {
				if tweakListOptions != nil {
					tweakListOptions(&options)
				}
				return client.CsiV1alpha1().OvirtCSIDrivers().Watch(context.TODO(), options)
			},
		},
		&operatorv1alpha1.OvirtCSIDriver{},
		resyncPeriod,
		indexers,
	)
}

func (f *ovirtCSIDriverInformer) defaultInformer(client versioned.Interface, resyncPeriod time.Duration) cache.SharedIndexInformer {
	return NewFilteredOvirtCSIDriverInformer(client, resyncPeriod, cache.Indexers{cache.NamespaceIndex: cache.MetaNamespaceIndexFunc}, f.tweakListOptions)
}

func (f *ovirtCSIDriverInformer) Informer() cache.SharedIndexInformer {
	return f.factory.InformerFor(&operatorv1alpha1.OvirtCSIDriver{}, f.defaultInformer)
}

func (f *ovirtCSIDriverInformer) Lister() v1alpha1.OvirtCSIDriverLister {
	return v1alpha1.NewOvirtCSIDriverLister(f.Informer().GetIndexer())
}