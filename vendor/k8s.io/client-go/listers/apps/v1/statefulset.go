/*
Copyright The Kubernetes Authors.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

// Code generated by lister-gen. DO NOT EDIT.

package v1

import (
	v1 "k8s.io/api/apps/v1"
	"k8s.io/apimachinery/pkg/labels"
	"k8s.io/client-go/listers"
	"k8s.io/client-go/tools/cache"
)

// StatefulSetLister helps list StatefulSets.
// All objects returned here must be treated as read-only.
type StatefulSetLister interface {
	// List lists all StatefulSets in the indexer.
	// Objects returned here must be treated as read-only.
	List(selector labels.Selector) (ret []*v1.StatefulSet, err error)
	// StatefulSets returns an object that can list and get StatefulSets.
	StatefulSets(namespace string) StatefulSetNamespaceLister
	StatefulSetListerExpansion
}

// statefulSetLister implements the StatefulSetLister interface.
type statefulSetLister struct {
	listers.ResourceIndexer[*v1.StatefulSet]
}

// NewStatefulSetLister returns a new StatefulSetLister.
func NewStatefulSetLister(indexer cache.Indexer) StatefulSetLister {
	return &statefulSetLister{listers.New[*v1.StatefulSet](indexer, v1.Resource("statefulset"))}
}

// StatefulSets returns an object that can list and get StatefulSets.
func (s *statefulSetLister) StatefulSets(namespace string) StatefulSetNamespaceLister {
	return statefulSetNamespaceLister{listers.NewNamespaced[*v1.StatefulSet](s.ResourceIndexer, namespace)}
}

// StatefulSetNamespaceLister helps list and get StatefulSets.
// All objects returned here must be treated as read-only.
type StatefulSetNamespaceLister interface {
	// List lists all StatefulSets in the indexer for a given namespace.
	// Objects returned here must be treated as read-only.
	List(selector labels.Selector) (ret []*v1.StatefulSet, err error)
	// Get retrieves the StatefulSet from the indexer for a given namespace and name.
	// Objects returned here must be treated as read-only.
	Get(name string) (*v1.StatefulSet, error)
	StatefulSetNamespaceListerExpansion
}

// statefulSetNamespaceLister implements the StatefulSetNamespaceLister
// interface.
type statefulSetNamespaceLister struct {
	listers.ResourceIndexer[*v1.StatefulSet]
}
