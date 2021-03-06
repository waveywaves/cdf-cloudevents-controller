/*
Copyright 2020 waveywaves

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

package v1alpha1

import (
	v1alpha1 "github.com/waveywaves/cloudevents-controller/pkg/apis/samples/v1alpha1"
	"k8s.io/apimachinery/pkg/api/errors"
	"k8s.io/apimachinery/pkg/labels"
	"k8s.io/client-go/tools/cache"
)

// CloudeventSinkLister helps list CloudeventSinks.
// All objects returned here must be treated as read-only.
type CloudeventSinkLister interface {
	// List lists all CloudeventSinks in the indexer.
	// Objects returned here must be treated as read-only.
	List(selector labels.Selector) (ret []*v1alpha1.CloudeventSink, err error)
	// CloudeventSinks returns an object that can list and get CloudeventSinks.
	CloudeventSinks(namespace string) CloudeventSinkNamespaceLister
	CloudeventSinkListerExpansion
}

// cloudeventSinkLister implements the CloudeventSinkLister interface.
type cloudeventSinkLister struct {
	indexer cache.Indexer
}

// NewCloudeventSinkLister returns a new CloudeventSinkLister.
func NewCloudeventSinkLister(indexer cache.Indexer) CloudeventSinkLister {
	return &cloudeventSinkLister{indexer: indexer}
}

// List lists all CloudeventSinks in the indexer.
func (s *cloudeventSinkLister) List(selector labels.Selector) (ret []*v1alpha1.CloudeventSink, err error) {
	err = cache.ListAll(s.indexer, selector, func(m interface{}) {
		ret = append(ret, m.(*v1alpha1.CloudeventSink))
	})
	return ret, err
}

// CloudeventSinks returns an object that can list and get CloudeventSinks.
func (s *cloudeventSinkLister) CloudeventSinks(namespace string) CloudeventSinkNamespaceLister {
	return cloudeventSinkNamespaceLister{indexer: s.indexer, namespace: namespace}
}

// CloudeventSinkNamespaceLister helps list and get CloudeventSinks.
// All objects returned here must be treated as read-only.
type CloudeventSinkNamespaceLister interface {
	// List lists all CloudeventSinks in the indexer for a given namespace.
	// Objects returned here must be treated as read-only.
	List(selector labels.Selector) (ret []*v1alpha1.CloudeventSink, err error)
	// Get retrieves the CloudeventSink from the indexer for a given namespace and name.
	// Objects returned here must be treated as read-only.
	Get(name string) (*v1alpha1.CloudeventSink, error)
	CloudeventSinkNamespaceListerExpansion
}

// cloudeventSinkNamespaceLister implements the CloudeventSinkNamespaceLister
// interface.
type cloudeventSinkNamespaceLister struct {
	indexer   cache.Indexer
	namespace string
}

// List lists all CloudeventSinks in the indexer for a given namespace.
func (s cloudeventSinkNamespaceLister) List(selector labels.Selector) (ret []*v1alpha1.CloudeventSink, err error) {
	err = cache.ListAllByNamespace(s.indexer, s.namespace, selector, func(m interface{}) {
		ret = append(ret, m.(*v1alpha1.CloudeventSink))
	})
	return ret, err
}

// Get retrieves the CloudeventSink from the indexer for a given namespace and name.
func (s cloudeventSinkNamespaceLister) Get(name string) (*v1alpha1.CloudeventSink, error) {
	obj, exists, err := s.indexer.GetByKey(s.namespace + "/" + name)
	if err != nil {
		return nil, err
	}
	if !exists {
		return nil, errors.NewNotFound(v1alpha1.Resource("cloudeventsink"), name)
	}
	return obj.(*v1alpha1.CloudeventSink), nil
}
