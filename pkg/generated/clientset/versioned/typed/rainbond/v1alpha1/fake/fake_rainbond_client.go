// RAINBOND, Application Management Platform
// Copyright (C) 2014-2017 Goodrain Co., Ltd.

// This program is free software: you can redistribute it and/or modify
// it under the terms of the GNU General Public License as published by
// the Free Software Foundation, either version 3 of the License, or
// (at your option) any later version. For any non-GPL usage of Rainbond,
// one or multiple Commercial Licenses authorized by Goodrain Co., Ltd.
// must be obtained first.

// This program is distributed in the hope that it will be useful,
// but WITHOUT ANY WARRANTY; without even the implied warranty of
// MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE. See the
// GNU General Public License for more details.

// You should have received a copy of the GNU General Public License
// along with this program. If not, see <http://www.gnu.org/licenses/>.

// Code generated by client-gen. DO NOT EDIT.

package fake

import (
	v1alpha1 "github.com/goodrain/rainbond-operator/pkg/generated/clientset/versioned/typed/rainbond/v1alpha1"
	rest "k8s.io/client-go/rest"
	testing "k8s.io/client-go/testing"
)

type FakeRainbondV1alpha1 struct {
	*testing.Fake
}

func (c *FakeRainbondV1alpha1) RainbondClusters(namespace string) v1alpha1.RainbondClusterInterface {
	return &FakeRainbondClusters{c, namespace}
}

func (c *FakeRainbondV1alpha1) RainbondPackages(namespace string) v1alpha1.RainbondPackageInterface {
	return &FakeRainbondPackages{c, namespace}
}

func (c *FakeRainbondV1alpha1) RbdComponents(namespace string) v1alpha1.RbdComponentInterface {
	return &FakeRbdComponents{c, namespace}
}

// RESTClient returns a RESTClient that is used to communicate
// with API server by this client implementation.
func (c *FakeRainbondV1alpha1) RESTClient() rest.Interface {
	var ret *rest.RESTClient
	return ret
}
