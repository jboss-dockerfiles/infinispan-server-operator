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

// Code generated by client-gen. DO NOT EDIT.

package fake

import (
	infinispanv1 "github.com/jboss-dockerfiles/infinispan-server-operator/pkg/apis/infinispan/v1"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	labels "k8s.io/apimachinery/pkg/labels"
	schema "k8s.io/apimachinery/pkg/runtime/schema"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	testing "k8s.io/client-go/testing"
)

// FakeInfinispans implements InfinispanInterface
type FakeInfinispans struct {
	Fake *FakeInfinispanV1
	ns   string
}

var infinispansResource = schema.GroupVersionResource{Group: "infinispan.org", Version: "v1", Resource: "infinispans"}

var infinispansKind = schema.GroupVersionKind{Group: "infinispan.org", Version: "v1", Kind: "Infinispan"}

// Get takes name of the infinispan, and returns the corresponding infinispan object, and an error if there is any.
func (c *FakeInfinispans) Get(name string, options v1.GetOptions) (result *infinispanv1.Infinispan, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewGetAction(infinispansResource, c.ns, name), &infinispanv1.Infinispan{})

	if obj == nil {
		return nil, err
	}
	return obj.(*infinispanv1.Infinispan), err
}

// List takes label and field selectors, and returns the list of Infinispans that match those selectors.
func (c *FakeInfinispans) List(opts v1.ListOptions) (result *infinispanv1.InfinispanList, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewListAction(infinispansResource, infinispansKind, c.ns, opts), &infinispanv1.InfinispanList{})

	if obj == nil {
		return nil, err
	}

	label, _, _ := testing.ExtractFromListOptions(opts)
	if label == nil {
		label = labels.Everything()
	}
	list := &infinispanv1.InfinispanList{ListMeta: obj.(*infinispanv1.InfinispanList).ListMeta}
	for _, item := range obj.(*infinispanv1.InfinispanList).Items {
		if label.Matches(labels.Set(item.Labels)) {
			list.Items = append(list.Items, item)
		}
	}
	return list, err
}

// Watch returns a watch.Interface that watches the requested infinispans.
func (c *FakeInfinispans) Watch(opts v1.ListOptions) (watch.Interface, error) {
	return c.Fake.
		InvokesWatch(testing.NewWatchAction(infinispansResource, c.ns, opts))

}

// Create takes the representation of a infinispan and creates it.  Returns the server's representation of the infinispan, and an error, if there is any.
func (c *FakeInfinispans) Create(infinispan *infinispanv1.Infinispan) (result *infinispanv1.Infinispan, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewCreateAction(infinispansResource, c.ns, infinispan), &infinispanv1.Infinispan{})

	if obj == nil {
		return nil, err
	}
	return obj.(*infinispanv1.Infinispan), err
}

// Update takes the representation of a infinispan and updates it. Returns the server's representation of the infinispan, and an error, if there is any.
func (c *FakeInfinispans) Update(infinispan *infinispanv1.Infinispan) (result *infinispanv1.Infinispan, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateAction(infinispansResource, c.ns, infinispan), &infinispanv1.Infinispan{})

	if obj == nil {
		return nil, err
	}
	return obj.(*infinispanv1.Infinispan), err
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().
func (c *FakeInfinispans) UpdateStatus(infinispan *infinispanv1.Infinispan) (*infinispanv1.Infinispan, error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateSubresourceAction(infinispansResource, "status", c.ns, infinispan), &infinispanv1.Infinispan{})

	if obj == nil {
		return nil, err
	}
	return obj.(*infinispanv1.Infinispan), err
}

// Delete takes name of the infinispan and deletes it. Returns an error if one occurs.
func (c *FakeInfinispans) Delete(name string, options *v1.DeleteOptions) error {
	_, err := c.Fake.
		Invokes(testing.NewDeleteAction(infinispansResource, c.ns, name), &infinispanv1.Infinispan{})

	return err
}

// DeleteCollection deletes a collection of objects.
func (c *FakeInfinispans) DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error {
	action := testing.NewDeleteCollectionAction(infinispansResource, c.ns, listOptions)

	_, err := c.Fake.Invokes(action, &infinispanv1.InfinispanList{})
	return err
}

// Patch applies the patch and returns the patched infinispan.
func (c *FakeInfinispans) Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *infinispanv1.Infinispan, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewPatchSubresourceAction(infinispansResource, c.ns, name, data, subresources...), &infinispanv1.Infinispan{})

	if obj == nil {
		return nil, err
	}
	return obj.(*infinispanv1.Infinispan), err
}
