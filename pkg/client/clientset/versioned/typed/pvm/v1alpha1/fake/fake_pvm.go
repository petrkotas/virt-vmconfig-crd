/*
Copyright 2017 The Kubernetes Authors.

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

package fake

import (
	v1alpha1 "github.com/petrkotas/virt-vmconfig-crd/pkg/apis/persistentvm/v1alpha1"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	labels "k8s.io/apimachinery/pkg/labels"
	schema "k8s.io/apimachinery/pkg/runtime/schema"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	testing "k8s.io/client-go/testing"
)

// FakePVMs implements PVMInterface
type FakePVMs struct {
	Fake *FakePvmV1alpha1
	ns   string
}

var pvmsResource = schema.GroupVersionResource{Group: "pvm.kubevirt.io", Version: "v1alpha1", Resource: "pvms"}

var pvmsKind = schema.GroupVersionKind{Group: "pvm.kubevirt.io", Version: "v1alpha1", Kind: "PVM"}

// Get takes name of the pVM, and returns the corresponding pVM object, and an error if there is any.
func (c *FakePVMs) Get(name string, options v1.GetOptions) (result *v1alpha1.PVM, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewGetAction(pvmsResource, c.ns, name), &v1alpha1.PVM{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.PVM), err
}

// List takes label and field selectors, and returns the list of PVMs that match those selectors.
func (c *FakePVMs) List(opts v1.ListOptions) (result *v1alpha1.PVMList, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewListAction(pvmsResource, pvmsKind, c.ns, opts), &v1alpha1.PVMList{})

	if obj == nil {
		return nil, err
	}

	label, _, _ := testing.ExtractFromListOptions(opts)
	if label == nil {
		label = labels.Everything()
	}
	list := &v1alpha1.PVMList{}
	for _, item := range obj.(*v1alpha1.PVMList).Items {
		if label.Matches(labels.Set(item.Labels)) {
			list.Items = append(list.Items, item)
		}
	}
	return list, err
}

// Watch returns a watch.Interface that watches the requested pVMs.
func (c *FakePVMs) Watch(opts v1.ListOptions) (watch.Interface, error) {
	return c.Fake.
		InvokesWatch(testing.NewWatchAction(pvmsResource, c.ns, opts))

}

// Create takes the representation of a pVM and creates it.  Returns the server's representation of the pVM, and an error, if there is any.
func (c *FakePVMs) Create(pVM *v1alpha1.PVM) (result *v1alpha1.PVM, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewCreateAction(pvmsResource, c.ns, pVM), &v1alpha1.PVM{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.PVM), err
}

// Update takes the representation of a pVM and updates it. Returns the server's representation of the pVM, and an error, if there is any.
func (c *FakePVMs) Update(pVM *v1alpha1.PVM) (result *v1alpha1.PVM, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateAction(pvmsResource, c.ns, pVM), &v1alpha1.PVM{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.PVM), err
}

// Delete takes name of the pVM and deletes it. Returns an error if one occurs.
func (c *FakePVMs) Delete(name string, options *v1.DeleteOptions) error {
	_, err := c.Fake.
		Invokes(testing.NewDeleteAction(pvmsResource, c.ns, name), &v1alpha1.PVM{})

	return err
}

// DeleteCollection deletes a collection of objects.
func (c *FakePVMs) DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error {
	action := testing.NewDeleteCollectionAction(pvmsResource, c.ns, listOptions)

	_, err := c.Fake.Invokes(action, &v1alpha1.PVMList{})
	return err
}

// Patch applies the patch and returns the patched pVM.
func (c *FakePVMs) Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v1alpha1.PVM, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewPatchSubresourceAction(pvmsResource, c.ns, name, data, subresources...), &v1alpha1.PVM{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.PVM), err
}
