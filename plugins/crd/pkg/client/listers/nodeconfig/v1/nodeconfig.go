// Copyright (c) 2018 Cisco and/or its affiliates.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at:
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

// Code generated by lister-gen. DO NOT EDIT.

package v1

import (
	v1 "github.com/contiv/vpp/plugins/crd/pkg/apis/nodeconfig/v1"
	"k8s.io/apimachinery/pkg/api/errors"
	"k8s.io/apimachinery/pkg/labels"
	"k8s.io/client-go/tools/cache"
)

// NodeConfigLister helps list NodeConfigs.
type NodeConfigLister interface {
	// List lists all NodeConfigs in the indexer.
	List(selector labels.Selector) (ret []*v1.NodeConfig, err error)
	// NodeConfigs returns an object that can list and get NodeConfigs.
	NodeConfigs(namespace string) NodeConfigNamespaceLister
	NodeConfigListerExpansion
}

// nodeConfigLister implements the NodeConfigLister interface.
type nodeConfigLister struct {
	indexer cache.Indexer
}

// NewNodeConfigLister returns a new NodeConfigLister.
func NewNodeConfigLister(indexer cache.Indexer) NodeConfigLister {
	return &nodeConfigLister{indexer: indexer}
}

// List lists all NodeConfigs in the indexer.
func (s *nodeConfigLister) List(selector labels.Selector) (ret []*v1.NodeConfig, err error) {
	err = cache.ListAll(s.indexer, selector, func(m interface{}) {
		ret = append(ret, m.(*v1.NodeConfig))
	})
	return ret, err
}

// NodeConfigs returns an object that can list and get NodeConfigs.
func (s *nodeConfigLister) NodeConfigs(namespace string) NodeConfigNamespaceLister {
	return nodeConfigNamespaceLister{indexer: s.indexer, namespace: namespace}
}

// NodeConfigNamespaceLister helps list and get NodeConfigs.
type NodeConfigNamespaceLister interface {
	// List lists all NodeConfigs in the indexer for a given namespace.
	List(selector labels.Selector) (ret []*v1.NodeConfig, err error)
	// Get retrieves the NodeConfig from the indexer for a given namespace and name.
	Get(name string) (*v1.NodeConfig, error)
	NodeConfigNamespaceListerExpansion
}

// nodeConfigNamespaceLister implements the NodeConfigNamespaceLister
// interface.
type nodeConfigNamespaceLister struct {
	indexer   cache.Indexer
	namespace string
}

// List lists all NodeConfigs in the indexer for a given namespace.
func (s nodeConfigNamespaceLister) List(selector labels.Selector) (ret []*v1.NodeConfig, err error) {
	err = cache.ListAllByNamespace(s.indexer, s.namespace, selector, func(m interface{}) {
		ret = append(ret, m.(*v1.NodeConfig))
	})
	return ret, err
}

// Get retrieves the NodeConfig from the indexer for a given namespace and name.
func (s nodeConfigNamespaceLister) Get(name string) (*v1.NodeConfig, error) {
	obj, exists, err := s.indexer.GetByKey(s.namespace + "/" + name)
	if err != nil {
		return nil, err
	}
	if !exists {
		return nil, errors.NewNotFound(v1.Resource("nodeconfig"), name)
	}
	return obj.(*v1.NodeConfig), nil
}