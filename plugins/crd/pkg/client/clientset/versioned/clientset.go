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

// Code generated by client-gen. DO NOT EDIT.

package versioned

import (
	contivv1 "github.com/contiv/vpp/plugins/crd/pkg/client/clientset/versioned/typed/contivtelemetry/v1"
	contivv1 "github.com/contiv/vpp/plugins/crd/pkg/client/clientset/versioned/typed/nodeconfig/v1"
	discovery "k8s.io/client-go/discovery"
	rest "k8s.io/client-go/rest"
	flowcontrol "k8s.io/client-go/util/flowcontrol"
)

type Interface interface {
	Discovery() discovery.DiscoveryInterface
	ContivV1() contivv1.ContivV1Interface
	// Deprecated: please explicitly pick a version if possible.
	Contiv() contivv1.ContivV1Interface
	ContivV1() contivv1.ContivV1Interface
	// Deprecated: please explicitly pick a version if possible.
	Contiv() contivv1.ContivV1Interface
}

// Clientset contains the clients for groups. Each group has exactly one
// version included in a Clientset.
type Clientset struct {
	*discovery.DiscoveryClient
	contivV1 *contivv1.ContivV1Client
	contivV1 *contivv1.ContivV1Client
}

// ContivV1 retrieves the ContivV1Client
func (c *Clientset) ContivV1() contivv1.ContivV1Interface {
	return c.contivV1
}

// Deprecated: Contiv retrieves the default version of ContivClient.
// Please explicitly pick a version.
func (c *Clientset) Contiv() contivv1.ContivV1Interface {
	return c.contivV1
}

// ContivV1 retrieves the ContivV1Client
func (c *Clientset) ContivV1() contivv1.ContivV1Interface {
	return c.contivV1
}

// Deprecated: Contiv retrieves the default version of ContivClient.
// Please explicitly pick a version.
func (c *Clientset) Contiv() contivv1.ContivV1Interface {
	return c.contivV1
}

// Discovery retrieves the DiscoveryClient
func (c *Clientset) Discovery() discovery.DiscoveryInterface {
	if c == nil {
		return nil
	}
	return c.DiscoveryClient
}

// NewForConfig creates a new Clientset for the given config.
func NewForConfig(c *rest.Config) (*Clientset, error) {
	configShallowCopy := *c
	if configShallowCopy.RateLimiter == nil && configShallowCopy.QPS > 0 {
		configShallowCopy.RateLimiter = flowcontrol.NewTokenBucketRateLimiter(configShallowCopy.QPS, configShallowCopy.Burst)
	}
	var cs Clientset
	var err error
	cs.contivV1, err = contivv1.NewForConfig(&configShallowCopy)
	if err != nil {
		return nil, err
	}
	cs.contivV1, err = contivv1.NewForConfig(&configShallowCopy)
	if err != nil {
		return nil, err
	}

	cs.DiscoveryClient, err = discovery.NewDiscoveryClientForConfig(&configShallowCopy)
	if err != nil {
		return nil, err
	}
	return &cs, nil
}

// NewForConfigOrDie creates a new Clientset for the given config and
// panics if there is an error in the config.
func NewForConfigOrDie(c *rest.Config) *Clientset {
	var cs Clientset
	cs.contivV1 = contivv1.NewForConfigOrDie(c)
	cs.contivV1 = contivv1.NewForConfigOrDie(c)

	cs.DiscoveryClient = discovery.NewDiscoveryClientForConfigOrDie(c)
	return &cs
}

// New creates a new Clientset for the given RESTClient.
func New(c rest.Interface) *Clientset {
	var cs Clientset
	cs.contivV1 = contivv1.New(c)
	cs.contivV1 = contivv1.New(c)

	cs.DiscoveryClient = discovery.NewDiscoveryClient(c)
	return &cs
}
