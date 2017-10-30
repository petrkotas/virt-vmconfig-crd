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

package v1alpha1

import (
	v1alpha1 "github.com/petrkotas/virt-vmconfig-crd/pkg/apis/persistentvm/v1alpha1"
	scheme "github.com/petrkotas/virt-vmconfig-crd/pkg/client/clientset/versioned/scheme"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	rest "k8s.io/client-go/rest"
)

// PVMsGetter has a method to return a PVMInterface.
// A group's client should implement this interface.
type PVMsGetter interface {
	PVMs(namespace string) PVMInterface
}

// PVMInterface has methods to work with PVM resources.
type PVMInterface interface {
	Create(*v1alpha1.PVM) (*v1alpha1.PVM, error)
	Update(*v1alpha1.PVM) (*v1alpha1.PVM, error)
	Delete(name string, options *v1.DeleteOptions) error
	DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error
	Get(name string, options v1.GetOptions) (*v1alpha1.PVM, error)
	List(opts v1.ListOptions) (*v1alpha1.PVMList, error)
	Watch(opts v1.ListOptions) (watch.Interface, error)
	Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v1alpha1.PVM, err error)
	PVMExpansion
}

// pVMs implements PVMInterface
type pVMs struct {
	client rest.Interface
	ns     string
}

// newPVMs returns a PVMs
func newPVMs(c *PvmV1alpha1Client, namespace string) *pVMs {
	return &pVMs{
		client: c.RESTClient(),
		ns:     namespace,
	}
}

// Get takes name of the pVM, and returns the corresponding pVM object, and an error if there is any.
func (c *pVMs) Get(name string, options v1.GetOptions) (result *v1alpha1.PVM, err error) {
	result = &v1alpha1.PVM{}
	err = c.client.Get().
		Namespace(c.ns).
		Resource("pvms").
		Name(name).
		VersionedParams(&options, scheme.ParameterCodec).
		Do().
		Into(result)
	return
}

// List takes label and field selectors, and returns the list of PVMs that match those selectors.
func (c *pVMs) List(opts v1.ListOptions) (result *v1alpha1.PVMList, err error) {
	result = &v1alpha1.PVMList{}
	err = c.client.Get().
		Namespace(c.ns).
		Resource("pvms").
		VersionedParams(&opts, scheme.ParameterCodec).
		Do().
		Into(result)
	return
}

// Watch returns a watch.Interface that watches the requested pVMs.
func (c *pVMs) Watch(opts v1.ListOptions) (watch.Interface, error) {
	opts.Watch = true
	return c.client.Get().
		Namespace(c.ns).
		Resource("pvms").
		VersionedParams(&opts, scheme.ParameterCodec).
		Watch()
}

// Create takes the representation of a pVM and creates it.  Returns the server's representation of the pVM, and an error, if there is any.
func (c *pVMs) Create(pVM *v1alpha1.PVM) (result *v1alpha1.PVM, err error) {
	result = &v1alpha1.PVM{}
	err = c.client.Post().
		Namespace(c.ns).
		Resource("pvms").
		Body(pVM).
		Do().
		Into(result)
	return
}

// Update takes the representation of a pVM and updates it. Returns the server's representation of the pVM, and an error, if there is any.
func (c *pVMs) Update(pVM *v1alpha1.PVM) (result *v1alpha1.PVM, err error) {
	result = &v1alpha1.PVM{}
	err = c.client.Put().
		Namespace(c.ns).
		Resource("pvms").
		Name(pVM.Name).
		Body(pVM).
		Do().
		Into(result)
	return
}

// Delete takes name of the pVM and deletes it. Returns an error if one occurs.
func (c *pVMs) Delete(name string, options *v1.DeleteOptions) error {
	return c.client.Delete().
		Namespace(c.ns).
		Resource("pvms").
		Name(name).
		Body(options).
		Do().
		Error()
}

// DeleteCollection deletes a collection of objects.
func (c *pVMs) DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error {
	return c.client.Delete().
		Namespace(c.ns).
		Resource("pvms").
		VersionedParams(&listOptions, scheme.ParameterCodec).
		Body(options).
		Do().
		Error()
}

// Patch applies the patch and returns the patched pVM.
func (c *pVMs) Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v1alpha1.PVM, err error) {
	result = &v1alpha1.PVM{}
	err = c.client.Patch(pt).
		Namespace(c.ns).
		Resource("pvms").
		SubResource(subresources...).
		Name(name).
		Body(data).
		Do().
		Into(result)
	return
}
