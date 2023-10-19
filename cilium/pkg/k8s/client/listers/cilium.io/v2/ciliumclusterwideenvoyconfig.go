// SPDX-License-Identifier: Apache-2.0
// Copyright Authors of Cilium

// Code generated by lister-gen. DO NOT EDIT.

package v2

import (
	v2 "github.com/cilium/cilium/pkg/k8s/apis/cilium.io/v2"
	"k8s.io/apimachinery/pkg/api/errors"
	"k8s.io/apimachinery/pkg/labels"
	"k8s.io/client-go/tools/cache"
)

// CiliumClusterwideEnvoyConfigLister helps list CiliumClusterwideEnvoyConfigs.
// All objects returned here must be treated as read-only.
type CiliumClusterwideEnvoyConfigLister interface {
	// List lists all CiliumClusterwideEnvoyConfigs in the indexer.
	// Objects returned here must be treated as read-only.
	List(selector labels.Selector) (ret []*v2.CiliumClusterwideEnvoyConfig, err error)
	// Get retrieves the CiliumClusterwideEnvoyConfig from the index for a given name.
	// Objects returned here must be treated as read-only.
	Get(name string) (*v2.CiliumClusterwideEnvoyConfig, error)
	CiliumClusterwideEnvoyConfigListerExpansion
}

// ciliumClusterwideEnvoyConfigLister implements the CiliumClusterwideEnvoyConfigLister interface.
type ciliumClusterwideEnvoyConfigLister struct {
	indexer cache.Indexer
}

// NewCiliumClusterwideEnvoyConfigLister returns a new CiliumClusterwideEnvoyConfigLister.
func NewCiliumClusterwideEnvoyConfigLister(indexer cache.Indexer) CiliumClusterwideEnvoyConfigLister {
	return &ciliumClusterwideEnvoyConfigLister{indexer: indexer}
}

// List lists all CiliumClusterwideEnvoyConfigs in the indexer.
func (s *ciliumClusterwideEnvoyConfigLister) List(selector labels.Selector) (ret []*v2.CiliumClusterwideEnvoyConfig, err error) {
	err = cache.ListAll(s.indexer, selector, func(m interface{}) {
		ret = append(ret, m.(*v2.CiliumClusterwideEnvoyConfig))
	})
	return ret, err
}

// Get retrieves the CiliumClusterwideEnvoyConfig from the index for a given name.
func (s *ciliumClusterwideEnvoyConfigLister) Get(name string) (*v2.CiliumClusterwideEnvoyConfig, error) {
	obj, exists, err := s.indexer.GetByKey(name)
	if err != nil {
		return nil, err
	}
	if !exists {
		return nil, errors.NewNotFound(v2.Resource("ciliumclusterwideenvoyconfig"), name)
	}
	return obj.(*v2.CiliumClusterwideEnvoyConfig), nil
}
