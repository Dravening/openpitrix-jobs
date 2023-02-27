/*
Copyright 2020 The D3os Authors.

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
	v1alpha1 "d3os.io/openpitrix-jobs/pkg/apis/application/v1alpha1"
	"k8s.io/apimachinery/pkg/api/errors"
	"k8s.io/apimachinery/pkg/labels"
	"k8s.io/client-go/tools/cache"
)

// HelmApplicationVersionLister helps list HelmApplicationVersions.
type HelmApplicationVersionLister interface {
	// List lists all HelmApplicationVersions in the indexer.
	List(selector labels.Selector) (ret []*v1alpha1.HelmApplicationVersion, err error)
	// Get retrieves the HelmApplicationVersion from the index for a given name.
	Get(name string) (*v1alpha1.HelmApplicationVersion, error)
	HelmApplicationVersionListerExpansion
}

// helmApplicationVersionLister implements the HelmApplicationVersionLister interface.
type helmApplicationVersionLister struct {
	indexer cache.Indexer
}

// NewHelmApplicationVersionLister returns a new HelmApplicationVersionLister.
func NewHelmApplicationVersionLister(indexer cache.Indexer) HelmApplicationVersionLister {
	return &helmApplicationVersionLister{indexer: indexer}
}

// List lists all HelmApplicationVersions in the indexer.
func (s *helmApplicationVersionLister) List(selector labels.Selector) (ret []*v1alpha1.HelmApplicationVersion, err error) {
	err = cache.ListAll(s.indexer, selector, func(m interface{}) {
		ret = append(ret, m.(*v1alpha1.HelmApplicationVersion))
	})
	return ret, err
}

// Get retrieves the HelmApplicationVersion from the index for a given name.
func (s *helmApplicationVersionLister) Get(name string) (*v1alpha1.HelmApplicationVersion, error) {
	obj, exists, err := s.indexer.GetByKey(name)
	if err != nil {
		return nil, err
	}
	if !exists {
		return nil, errors.NewNotFound(v1alpha1.Resource("helmapplicationversion"), name)
	}
	return obj.(*v1alpha1.HelmApplicationVersion), nil
}
