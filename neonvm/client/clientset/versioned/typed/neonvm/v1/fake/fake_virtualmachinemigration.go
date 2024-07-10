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
// Code generated by client-gen. DO NOT EDIT.

package fake

import (
	"context"

	v1 "github.com/neondatabase/autoscaling/neonvm/apis/neonvm/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	labels "k8s.io/apimachinery/pkg/labels"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	testing "k8s.io/client-go/testing"
)

// FakeVirtualMachineMigrations implements VirtualMachineMigrationInterface
type FakeVirtualMachineMigrations struct {
	Fake *FakeNeonvmV1
	ns   string
}

var virtualmachinemigrationsResource = v1.SchemeGroupVersion.WithResource("virtualmachinemigrations")

var virtualmachinemigrationsKind = v1.SchemeGroupVersion.WithKind("VirtualMachineMigration")

// Get takes name of the virtualMachineMigration, and returns the corresponding virtualMachineMigration object, and an error if there is any.
func (c *FakeVirtualMachineMigrations) Get(ctx context.Context, name string, options metav1.GetOptions) (result *v1.VirtualMachineMigration, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewGetAction(virtualmachinemigrationsResource, c.ns, name), &v1.VirtualMachineMigration{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1.VirtualMachineMigration), err
}

// List takes label and field selectors, and returns the list of VirtualMachineMigrations that match those selectors.
func (c *FakeVirtualMachineMigrations) List(ctx context.Context, opts metav1.ListOptions) (result *v1.VirtualMachineMigrationList, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewListAction(virtualmachinemigrationsResource, virtualmachinemigrationsKind, c.ns, opts), &v1.VirtualMachineMigrationList{})

	if obj == nil {
		return nil, err
	}

	label, _, _ := testing.ExtractFromListOptions(opts)
	if label == nil {
		label = labels.Everything()
	}
	list := &v1.VirtualMachineMigrationList{ListMeta: obj.(*v1.VirtualMachineMigrationList).ListMeta}
	for _, item := range obj.(*v1.VirtualMachineMigrationList).Items {
		if label.Matches(labels.Set(item.Labels)) {
			list.Items = append(list.Items, item)
		}
	}
	return list, err
}

// Watch returns a watch.Interface that watches the requested virtualMachineMigrations.
func (c *FakeVirtualMachineMigrations) Watch(ctx context.Context, opts metav1.ListOptions) (watch.Interface, error) {
	return c.Fake.
		InvokesWatch(testing.NewWatchAction(virtualmachinemigrationsResource, c.ns, opts))

}

// Create takes the representation of a virtualMachineMigration and creates it.  Returns the server's representation of the virtualMachineMigration, and an error, if there is any.
func (c *FakeVirtualMachineMigrations) Create(ctx context.Context, virtualMachineMigration *v1.VirtualMachineMigration, opts metav1.CreateOptions) (result *v1.VirtualMachineMigration, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewCreateAction(virtualmachinemigrationsResource, c.ns, virtualMachineMigration), &v1.VirtualMachineMigration{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1.VirtualMachineMigration), err
}

// Update takes the representation of a virtualMachineMigration and updates it. Returns the server's representation of the virtualMachineMigration, and an error, if there is any.
func (c *FakeVirtualMachineMigrations) Update(ctx context.Context, virtualMachineMigration *v1.VirtualMachineMigration, opts metav1.UpdateOptions) (result *v1.VirtualMachineMigration, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateAction(virtualmachinemigrationsResource, c.ns, virtualMachineMigration), &v1.VirtualMachineMigration{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1.VirtualMachineMigration), err
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().
func (c *FakeVirtualMachineMigrations) UpdateStatus(ctx context.Context, virtualMachineMigration *v1.VirtualMachineMigration, opts metav1.UpdateOptions) (*v1.VirtualMachineMigration, error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateSubresourceAction(virtualmachinemigrationsResource, "status", c.ns, virtualMachineMigration), &v1.VirtualMachineMigration{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1.VirtualMachineMigration), err
}

// Delete takes name of the virtualMachineMigration and deletes it. Returns an error if one occurs.
func (c *FakeVirtualMachineMigrations) Delete(ctx context.Context, name string, opts metav1.DeleteOptions) error {
	_, err := c.Fake.
		Invokes(testing.NewDeleteActionWithOptions(virtualmachinemigrationsResource, c.ns, name, opts), &v1.VirtualMachineMigration{})

	return err
}

// DeleteCollection deletes a collection of objects.
func (c *FakeVirtualMachineMigrations) DeleteCollection(ctx context.Context, opts metav1.DeleteOptions, listOpts metav1.ListOptions) error {
	action := testing.NewDeleteCollectionAction(virtualmachinemigrationsResource, c.ns, listOpts)

	_, err := c.Fake.Invokes(action, &v1.VirtualMachineMigrationList{})
	return err
}

// Patch applies the patch and returns the patched virtualMachineMigration.
func (c *FakeVirtualMachineMigrations) Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts metav1.PatchOptions, subresources ...string) (result *v1.VirtualMachineMigration, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewPatchSubresourceAction(virtualmachinemigrationsResource, c.ns, name, pt, data, subresources...), &v1.VirtualMachineMigration{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1.VirtualMachineMigration), err
}
