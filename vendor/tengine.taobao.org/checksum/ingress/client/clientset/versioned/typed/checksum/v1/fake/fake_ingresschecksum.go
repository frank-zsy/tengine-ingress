/*
Copyright 2023 The Kubernetes Authors.

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

// Code generated by client-gen. DO NOT EDIT.

package fake

import (
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	labels "k8s.io/apimachinery/pkg/labels"
	schema "k8s.io/apimachinery/pkg/runtime/schema"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	testing "k8s.io/client-go/testing"
	checksumv1 "tengine.taobao.org/checksum/ingress/apis/checksum/v1"
)

// FakeIngressCheckSums implements IngressCheckSumInterface
type FakeIngressCheckSums struct {
	Fake *FakeTengineV1
	ns   string
}

var ingresschecksumsResource = schema.GroupVersionResource{Group: "tengine.taobao.org", Version: "v1", Resource: "ingresschecksums"}

var ingresschecksumsKind = schema.GroupVersionKind{Group: "tengine.taobao.org", Version: "v1", Kind: "IngressCheckSum"}

// Get takes name of the ingressCheckSum, and returns the corresponding ingressCheckSum object, and an error if there is any.
func (c *FakeIngressCheckSums) Get(name string, options v1.GetOptions) (result *checksumv1.IngressCheckSum, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewGetAction(ingresschecksumsResource, c.ns, name), &checksumv1.IngressCheckSum{})

	if obj == nil {
		return nil, err
	}
	return obj.(*checksumv1.IngressCheckSum), err
}

// List takes label and field selectors, and returns the list of IngressCheckSums that match those selectors.
func (c *FakeIngressCheckSums) List(opts v1.ListOptions) (result *checksumv1.IngressCheckSumList, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewListAction(ingresschecksumsResource, ingresschecksumsKind, c.ns, opts), &checksumv1.IngressCheckSumList{})

	if obj == nil {
		return nil, err
	}

	label, _, _ := testing.ExtractFromListOptions(opts)
	if label == nil {
		label = labels.Everything()
	}
	list := &checksumv1.IngressCheckSumList{ListMeta: obj.(*checksumv1.IngressCheckSumList).ListMeta}
	for _, item := range obj.(*checksumv1.IngressCheckSumList).Items {
		if label.Matches(labels.Set(item.Labels)) {
			list.Items = append(list.Items, item)
		}
	}
	return list, err
}

// Watch returns a watch.Interface that watches the requested ingressCheckSums.
func (c *FakeIngressCheckSums) Watch(opts v1.ListOptions) (watch.Interface, error) {
	return c.Fake.
		InvokesWatch(testing.NewWatchAction(ingresschecksumsResource, c.ns, opts))

}

// Create takes the representation of a ingressCheckSum and creates it.  Returns the server's representation of the ingressCheckSum, and an error, if there is any.
func (c *FakeIngressCheckSums) Create(ingressCheckSum *checksumv1.IngressCheckSum) (result *checksumv1.IngressCheckSum, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewCreateAction(ingresschecksumsResource, c.ns, ingressCheckSum), &checksumv1.IngressCheckSum{})

	if obj == nil {
		return nil, err
	}
	return obj.(*checksumv1.IngressCheckSum), err
}

// Update takes the representation of a ingressCheckSum and updates it. Returns the server's representation of the ingressCheckSum, and an error, if there is any.
func (c *FakeIngressCheckSums) Update(ingressCheckSum *checksumv1.IngressCheckSum) (result *checksumv1.IngressCheckSum, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateAction(ingresschecksumsResource, c.ns, ingressCheckSum), &checksumv1.IngressCheckSum{})

	if obj == nil {
		return nil, err
	}
	return obj.(*checksumv1.IngressCheckSum), err
}

// Delete takes name of the ingressCheckSum and deletes it. Returns an error if one occurs.
func (c *FakeIngressCheckSums) Delete(name string, options *v1.DeleteOptions) error {
	_, err := c.Fake.
		Invokes(testing.NewDeleteAction(ingresschecksumsResource, c.ns, name), &checksumv1.IngressCheckSum{})

	return err
}

// DeleteCollection deletes a collection of objects.
func (c *FakeIngressCheckSums) DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error {
	action := testing.NewDeleteCollectionAction(ingresschecksumsResource, c.ns, listOptions)

	_, err := c.Fake.Invokes(action, &checksumv1.IngressCheckSumList{})
	return err
}

// Patch applies the patch and returns the patched ingressCheckSum.
func (c *FakeIngressCheckSums) Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *checksumv1.IngressCheckSum, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewPatchSubresourceAction(ingresschecksumsResource, c.ns, name, pt, data, subresources...), &checksumv1.IngressCheckSum{})

	if obj == nil {
		return nil, err
	}
	return obj.(*checksumv1.IngressCheckSum), err
}