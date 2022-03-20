/*
Copyright 2022.

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
// Code generated by main. DO NOT EDIT.

package fake

import (
	"context"

	welancontrollerv1 "github.com/weizhoublue/github-action-test/pkg/k8s/api/welancontroller/v1"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	labels "k8s.io/apimachinery/pkg/labels"
	schema "k8s.io/apimachinery/pkg/runtime/schema"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	testing "k8s.io/client-go/testing"
)

// FakeFoos implements FooInterface
type FakeFoos struct {
	Fake *FakeWelanV1
	ns   string
}

var foosResource = schema.GroupVersionResource{Group: "welan.k8s.io", Version: "v1", Resource: "foos"}

var foosKind = schema.GroupVersionKind{Group: "welan.k8s.io", Version: "v1", Kind: "Foo"}

// Get takes name of the foo, and returns the corresponding foo object, and an error if there is any.
func (c *FakeFoos) Get(ctx context.Context, name string, options v1.GetOptions) (result *welancontrollerv1.Foo, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewGetAction(foosResource, c.ns, name), &welancontrollerv1.Foo{})

	if obj == nil {
		return nil, err
	}
	return obj.(*welancontrollerv1.Foo), err
}

// List takes label and field selectors, and returns the list of Foos that match those selectors.
func (c *FakeFoos) List(ctx context.Context, opts v1.ListOptions) (result *welancontrollerv1.FooList, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewListAction(foosResource, foosKind, c.ns, opts), &welancontrollerv1.FooList{})

	if obj == nil {
		return nil, err
	}

	label, _, _ := testing.ExtractFromListOptions(opts)
	if label == nil {
		label = labels.Everything()
	}
	list := &welancontrollerv1.FooList{ListMeta: obj.(*welancontrollerv1.FooList).ListMeta}
	for _, item := range obj.(*welancontrollerv1.FooList).Items {
		if label.Matches(labels.Set(item.Labels)) {
			list.Items = append(list.Items, item)
		}
	}
	return list, err
}

// Watch returns a watch.Interface that watches the requested foos.
func (c *FakeFoos) Watch(ctx context.Context, opts v1.ListOptions) (watch.Interface, error) {
	return c.Fake.
		InvokesWatch(testing.NewWatchAction(foosResource, c.ns, opts))

}

// Create takes the representation of a foo and creates it.  Returns the server's representation of the foo, and an error, if there is any.
func (c *FakeFoos) Create(ctx context.Context, foo *welancontrollerv1.Foo, opts v1.CreateOptions) (result *welancontrollerv1.Foo, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewCreateAction(foosResource, c.ns, foo), &welancontrollerv1.Foo{})

	if obj == nil {
		return nil, err
	}
	return obj.(*welancontrollerv1.Foo), err
}

// Update takes the representation of a foo and updates it. Returns the server's representation of the foo, and an error, if there is any.
func (c *FakeFoos) Update(ctx context.Context, foo *welancontrollerv1.Foo, opts v1.UpdateOptions) (result *welancontrollerv1.Foo, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateAction(foosResource, c.ns, foo), &welancontrollerv1.Foo{})

	if obj == nil {
		return nil, err
	}
	return obj.(*welancontrollerv1.Foo), err
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().
func (c *FakeFoos) UpdateStatus(ctx context.Context, foo *welancontrollerv1.Foo, opts v1.UpdateOptions) (*welancontrollerv1.Foo, error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateSubresourceAction(foosResource, "status", c.ns, foo), &welancontrollerv1.Foo{})

	if obj == nil {
		return nil, err
	}
	return obj.(*welancontrollerv1.Foo), err
}

// Delete takes name of the foo and deletes it. Returns an error if one occurs.
func (c *FakeFoos) Delete(ctx context.Context, name string, opts v1.DeleteOptions) error {
	_, err := c.Fake.
		Invokes(testing.NewDeleteActionWithOptions(foosResource, c.ns, name, opts), &welancontrollerv1.Foo{})

	return err
}

// DeleteCollection deletes a collection of objects.
func (c *FakeFoos) DeleteCollection(ctx context.Context, opts v1.DeleteOptions, listOpts v1.ListOptions) error {
	action := testing.NewDeleteCollectionAction(foosResource, c.ns, listOpts)

	_, err := c.Fake.Invokes(action, &welancontrollerv1.FooList{})
	return err
}

// Patch applies the patch and returns the patched foo.
func (c *FakeFoos) Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts v1.PatchOptions, subresources ...string) (result *welancontrollerv1.Foo, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewPatchSubresourceAction(foosResource, c.ns, name, pt, data, subresources...), &welancontrollerv1.Foo{})

	if obj == nil {
		return nil, err
	}
	return obj.(*welancontrollerv1.Foo), err
}
