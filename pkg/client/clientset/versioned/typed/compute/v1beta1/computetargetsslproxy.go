// Copyright 2020 Google LLC
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

// *** DISCLAIMER ***
// Config Connector's go-client for CRDs is currently in ALPHA, which means
// that future versions of the go-client may include breaking changes.
// Please try it out and give us feedback!

// Code generated by main. DO NOT EDIT.

package v1beta1

import (
	"context"
	"time"

	v1beta1 "github.com/GoogleCloudPlatform/k8s-config-connector/pkg/apis/compute/v1beta1"
	scheme "github.com/GoogleCloudPlatform/k8s-config-connector/pkg/client/clientset/versioned/scheme"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	rest "k8s.io/client-go/rest"
)

// ComputeTargetSSLProxiesGetter has a method to return a ComputeTargetSSLProxyInterface.
// A group's client should implement this interface.
type ComputeTargetSSLProxiesGetter interface {
	ComputeTargetSSLProxies(namespace string) ComputeTargetSSLProxyInterface
}

// ComputeTargetSSLProxyInterface has methods to work with ComputeTargetSSLProxy resources.
type ComputeTargetSSLProxyInterface interface {
	Create(ctx context.Context, computeTargetSSLProxy *v1beta1.ComputeTargetSSLProxy, opts v1.CreateOptions) (*v1beta1.ComputeTargetSSLProxy, error)
	Update(ctx context.Context, computeTargetSSLProxy *v1beta1.ComputeTargetSSLProxy, opts v1.UpdateOptions) (*v1beta1.ComputeTargetSSLProxy, error)
	UpdateStatus(ctx context.Context, computeTargetSSLProxy *v1beta1.ComputeTargetSSLProxy, opts v1.UpdateOptions) (*v1beta1.ComputeTargetSSLProxy, error)
	Delete(ctx context.Context, name string, opts v1.DeleteOptions) error
	DeleteCollection(ctx context.Context, opts v1.DeleteOptions, listOpts v1.ListOptions) error
	Get(ctx context.Context, name string, opts v1.GetOptions) (*v1beta1.ComputeTargetSSLProxy, error)
	List(ctx context.Context, opts v1.ListOptions) (*v1beta1.ComputeTargetSSLProxyList, error)
	Watch(ctx context.Context, opts v1.ListOptions) (watch.Interface, error)
	Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts v1.PatchOptions, subresources ...string) (result *v1beta1.ComputeTargetSSLProxy, err error)
	ComputeTargetSSLProxyExpansion
}

// computeTargetSSLProxies implements ComputeTargetSSLProxyInterface
type computeTargetSSLProxies struct {
	client rest.Interface
	ns     string
}

// newComputeTargetSSLProxies returns a ComputeTargetSSLProxies
func newComputeTargetSSLProxies(c *ComputeV1beta1Client, namespace string) *computeTargetSSLProxies {
	return &computeTargetSSLProxies{
		client: c.RESTClient(),
		ns:     namespace,
	}
}

// Get takes name of the computeTargetSSLProxy, and returns the corresponding computeTargetSSLProxy object, and an error if there is any.
func (c *computeTargetSSLProxies) Get(ctx context.Context, name string, options v1.GetOptions) (result *v1beta1.ComputeTargetSSLProxy, err error) {
	result = &v1beta1.ComputeTargetSSLProxy{}
	err = c.client.Get().
		Namespace(c.ns).
		Resource("computetargetsslproxies").
		Name(name).
		VersionedParams(&options, scheme.ParameterCodec).
		Do(ctx).
		Into(result)
	return
}

// List takes label and field selectors, and returns the list of ComputeTargetSSLProxies that match those selectors.
func (c *computeTargetSSLProxies) List(ctx context.Context, opts v1.ListOptions) (result *v1beta1.ComputeTargetSSLProxyList, err error) {
	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	result = &v1beta1.ComputeTargetSSLProxyList{}
	err = c.client.Get().
		Namespace(c.ns).
		Resource("computetargetsslproxies").
		VersionedParams(&opts, scheme.ParameterCodec).
		Timeout(timeout).
		Do(ctx).
		Into(result)
	return
}

// Watch returns a watch.Interface that watches the requested computeTargetSSLProxies.
func (c *computeTargetSSLProxies) Watch(ctx context.Context, opts v1.ListOptions) (watch.Interface, error) {
	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	opts.Watch = true
	return c.client.Get().
		Namespace(c.ns).
		Resource("computetargetsslproxies").
		VersionedParams(&opts, scheme.ParameterCodec).
		Timeout(timeout).
		Watch(ctx)
}

// Create takes the representation of a computeTargetSSLProxy and creates it.  Returns the server's representation of the computeTargetSSLProxy, and an error, if there is any.
func (c *computeTargetSSLProxies) Create(ctx context.Context, computeTargetSSLProxy *v1beta1.ComputeTargetSSLProxy, opts v1.CreateOptions) (result *v1beta1.ComputeTargetSSLProxy, err error) {
	result = &v1beta1.ComputeTargetSSLProxy{}
	err = c.client.Post().
		Namespace(c.ns).
		Resource("computetargetsslproxies").
		VersionedParams(&opts, scheme.ParameterCodec).
		Body(computeTargetSSLProxy).
		Do(ctx).
		Into(result)
	return
}

// Update takes the representation of a computeTargetSSLProxy and updates it. Returns the server's representation of the computeTargetSSLProxy, and an error, if there is any.
func (c *computeTargetSSLProxies) Update(ctx context.Context, computeTargetSSLProxy *v1beta1.ComputeTargetSSLProxy, opts v1.UpdateOptions) (result *v1beta1.ComputeTargetSSLProxy, err error) {
	result = &v1beta1.ComputeTargetSSLProxy{}
	err = c.client.Put().
		Namespace(c.ns).
		Resource("computetargetsslproxies").
		Name(computeTargetSSLProxy.Name).
		VersionedParams(&opts, scheme.ParameterCodec).
		Body(computeTargetSSLProxy).
		Do(ctx).
		Into(result)
	return
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().
func (c *computeTargetSSLProxies) UpdateStatus(ctx context.Context, computeTargetSSLProxy *v1beta1.ComputeTargetSSLProxy, opts v1.UpdateOptions) (result *v1beta1.ComputeTargetSSLProxy, err error) {
	result = &v1beta1.ComputeTargetSSLProxy{}
	err = c.client.Put().
		Namespace(c.ns).
		Resource("computetargetsslproxies").
		Name(computeTargetSSLProxy.Name).
		SubResource("status").
		VersionedParams(&opts, scheme.ParameterCodec).
		Body(computeTargetSSLProxy).
		Do(ctx).
		Into(result)
	return
}

// Delete takes name of the computeTargetSSLProxy and deletes it. Returns an error if one occurs.
func (c *computeTargetSSLProxies) Delete(ctx context.Context, name string, opts v1.DeleteOptions) error {
	return c.client.Delete().
		Namespace(c.ns).
		Resource("computetargetsslproxies").
		Name(name).
		Body(&opts).
		Do(ctx).
		Error()
}

// DeleteCollection deletes a collection of objects.
func (c *computeTargetSSLProxies) DeleteCollection(ctx context.Context, opts v1.DeleteOptions, listOpts v1.ListOptions) error {
	var timeout time.Duration
	if listOpts.TimeoutSeconds != nil {
		timeout = time.Duration(*listOpts.TimeoutSeconds) * time.Second
	}
	return c.client.Delete().
		Namespace(c.ns).
		Resource("computetargetsslproxies").
		VersionedParams(&listOpts, scheme.ParameterCodec).
		Timeout(timeout).
		Body(&opts).
		Do(ctx).
		Error()
}

// Patch applies the patch and returns the patched computeTargetSSLProxy.
func (c *computeTargetSSLProxies) Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts v1.PatchOptions, subresources ...string) (result *v1beta1.ComputeTargetSSLProxy, err error) {
	result = &v1beta1.ComputeTargetSSLProxy{}
	err = c.client.Patch(pt).
		Namespace(c.ns).
		Resource("computetargetsslproxies").
		Name(name).
		SubResource(subresources...).
		VersionedParams(&opts, scheme.ParameterCodec).
		Body(data).
		Do(ctx).
		Into(result)
	return
}