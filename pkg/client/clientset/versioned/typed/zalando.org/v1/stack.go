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

package v1

import (
	"context"
	"time"

	v1 "github.com/zalando-incubator/stackset-controller/pkg/apis/zalando.org/v1"
	scheme "github.com/zalando-incubator/stackset-controller/pkg/client/clientset/versioned/scheme"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	rest "k8s.io/client-go/rest"
)

// StacksGetter has a method to return a StackInterface.
// A group's client should implement this interface.
type StacksGetter interface {
	Stacks(namespace string) StackInterface
}

// StackInterface has methods to work with Stack resources.
type StackInterface interface {
	Create(ctx context.Context, stack *v1.Stack, opts metav1.CreateOptions) (*v1.Stack, error)
	Update(ctx context.Context, stack *v1.Stack, opts metav1.UpdateOptions) (*v1.Stack, error)
	UpdateStatus(ctx context.Context, stack *v1.Stack, opts metav1.UpdateOptions) (*v1.Stack, error)
	Delete(ctx context.Context, name string, opts metav1.DeleteOptions) error
	DeleteCollection(ctx context.Context, opts metav1.DeleteOptions, listOpts metav1.ListOptions) error
	Get(ctx context.Context, name string, opts metav1.GetOptions) (*v1.Stack, error)
	List(ctx context.Context, opts metav1.ListOptions) (*v1.StackList, error)
	Watch(ctx context.Context, opts metav1.ListOptions) (watch.Interface, error)
	Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts metav1.PatchOptions, subresources ...string) (result *v1.Stack, err error)
	StackExpansion
}

// stacks implements StackInterface
type stacks struct {
	client rest.Interface
	ns     string
}

// newStacks returns a Stacks
func newStacks(c *ZalandoV1Client, namespace string) *stacks {
	return &stacks{
		client: c.RESTClient(),
		ns:     namespace,
	}
}

// Get takes name of the stack, and returns the corresponding stack object, and an error if there is any.
func (c *stacks) Get(ctx context.Context, name string, options metav1.GetOptions) (result *v1.Stack, err error) {
	result = &v1.Stack{}
	err = c.client.Get().
		Namespace(c.ns).
		Resource("stacks").
		Name(name).
		VersionedParams(&options, scheme.ParameterCodec).
		Do(ctx).
		Into(result)
	return
}

// List takes label and field selectors, and returns the list of Stacks that match those selectors.
func (c *stacks) List(ctx context.Context, opts metav1.ListOptions) (result *v1.StackList, err error) {
	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	result = &v1.StackList{}
	err = c.client.Get().
		Namespace(c.ns).
		Resource("stacks").
		VersionedParams(&opts, scheme.ParameterCodec).
		Timeout(timeout).
		Do(ctx).
		Into(result)
	return
}

// Watch returns a watch.Interface that watches the requested stacks.
func (c *stacks) Watch(ctx context.Context, opts metav1.ListOptions) (watch.Interface, error) {
	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	opts.Watch = true
	return c.client.Get().
		Namespace(c.ns).
		Resource("stacks").
		VersionedParams(&opts, scheme.ParameterCodec).
		Timeout(timeout).
		Watch(ctx)
}

// Create takes the representation of a stack and creates it.  Returns the server's representation of the stack, and an error, if there is any.
func (c *stacks) Create(ctx context.Context, stack *v1.Stack, opts metav1.CreateOptions) (result *v1.Stack, err error) {
	result = &v1.Stack{}
	err = c.client.Post().
		Namespace(c.ns).
		Resource("stacks").
		VersionedParams(&opts, scheme.ParameterCodec).
		Body(stack).
		Do(ctx).
		Into(result)
	return
}

// Update takes the representation of a stack and updates it. Returns the server's representation of the stack, and an error, if there is any.
func (c *stacks) Update(ctx context.Context, stack *v1.Stack, opts metav1.UpdateOptions) (result *v1.Stack, err error) {
	result = &v1.Stack{}
	err = c.client.Put().
		Namespace(c.ns).
		Resource("stacks").
		Name(stack.Name).
		VersionedParams(&opts, scheme.ParameterCodec).
		Body(stack).
		Do(ctx).
		Into(result)
	return
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().
func (c *stacks) UpdateStatus(ctx context.Context, stack *v1.Stack, opts metav1.UpdateOptions) (result *v1.Stack, err error) {
	result = &v1.Stack{}
	err = c.client.Put().
		Namespace(c.ns).
		Resource("stacks").
		Name(stack.Name).
		SubResource("status").
		VersionedParams(&opts, scheme.ParameterCodec).
		Body(stack).
		Do(ctx).
		Into(result)
	return
}

// Delete takes name of the stack and deletes it. Returns an error if one occurs.
func (c *stacks) Delete(ctx context.Context, name string, opts metav1.DeleteOptions) error {
	return c.client.Delete().
		Namespace(c.ns).
		Resource("stacks").
		Name(name).
		Body(&opts).
		Do(ctx).
		Error()
}

// DeleteCollection deletes a collection of objects.
func (c *stacks) DeleteCollection(ctx context.Context, opts metav1.DeleteOptions, listOpts metav1.ListOptions) error {
	var timeout time.Duration
	if listOpts.TimeoutSeconds != nil {
		timeout = time.Duration(*listOpts.TimeoutSeconds) * time.Second
	}
	return c.client.Delete().
		Namespace(c.ns).
		Resource("stacks").
		VersionedParams(&listOpts, scheme.ParameterCodec).
		Timeout(timeout).
		Body(&opts).
		Do(ctx).
		Error()
}

// Patch applies the patch and returns the patched stack.
func (c *stacks) Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts metav1.PatchOptions, subresources ...string) (result *v1.Stack, err error) {
	result = &v1.Stack{}
	err = c.client.Patch(pt).
		Namespace(c.ns).
		Resource("stacks").
		Name(name).
		SubResource(subresources...).
		VersionedParams(&opts, scheme.ParameterCodec).
		Body(data).
		Do(ctx).
		Into(result)
	return
}
