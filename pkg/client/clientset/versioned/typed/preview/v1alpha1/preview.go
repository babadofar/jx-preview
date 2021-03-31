/*
Copyright 2020 The Jenkins X Authors.

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

package v1alpha1

import (
	"context"
	"time"

	v1alpha1 "github.com/jenkins-x-plugins/jx-preview/pkg/apis/preview/v1alpha1"
	scheme "github.com/jenkins-x-plugins/jx-preview/pkg/client/clientset/versioned/scheme"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	rest "k8s.io/client-go/rest"
)

// PreviewsGetter has a method to return a PreviewInterface.
// A group's client should implement this interface.
type PreviewsGetter interface {
	Previews(namespace string) PreviewInterface
}

// PreviewInterface has methods to work with Preview resources.
type PreviewInterface interface {
	Create(ctx context.Context, preview *v1alpha1.Preview, opts v1.CreateOptions) (*v1alpha1.Preview, error)
	Update(ctx context.Context, preview *v1alpha1.Preview, opts v1.UpdateOptions) (*v1alpha1.Preview, error)
	Delete(ctx context.Context, name string, opts v1.DeleteOptions) error
	DeleteCollection(ctx context.Context, opts v1.DeleteOptions, listOpts v1.ListOptions) error
	Get(ctx context.Context, name string, opts v1.GetOptions) (*v1alpha1.Preview, error)
	List(ctx context.Context, opts v1.ListOptions) (*v1alpha1.PreviewList, error)
	Watch(ctx context.Context, opts v1.ListOptions) (watch.Interface, error)
	Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts v1.PatchOptions, subresources ...string) (result *v1alpha1.Preview, err error)
	PreviewExpansion
}

// previews implements PreviewInterface
type previews struct {
	client rest.Interface
	ns     string
}

// newPreviews returns a Previews
func newPreviews(c *PreviewV1alpha1Client, namespace string) *previews {
	return &previews{
		client: c.RESTClient(),
		ns:     namespace,
	}
}

// Get takes name of the preview, and returns the corresponding preview object, and an error if there is any.
func (c *previews) Get(ctx context.Context, name string, options v1.GetOptions) (result *v1alpha1.Preview, err error) {
	result = &v1alpha1.Preview{}
	err = c.client.Get().
		Namespace(c.ns).
		Resource("previews").
		Name(name).
		VersionedParams(&options, scheme.ParameterCodec).
		Do(ctx).
		Into(result)
	return
}

// List takes label and field selectors, and returns the list of Previews that match those selectors.
func (c *previews) List(ctx context.Context, opts v1.ListOptions) (result *v1alpha1.PreviewList, err error) {
	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	result = &v1alpha1.PreviewList{}
	err = c.client.Get().
		Namespace(c.ns).
		Resource("previews").
		VersionedParams(&opts, scheme.ParameterCodec).
		Timeout(timeout).
		Do(ctx).
		Into(result)
	return
}

// Watch returns a watch.Interface that watches the requested previews.
func (c *previews) Watch(ctx context.Context, opts v1.ListOptions) (watch.Interface, error) {
	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	opts.Watch = true
	return c.client.Get().
		Namespace(c.ns).
		Resource("previews").
		VersionedParams(&opts, scheme.ParameterCodec).
		Timeout(timeout).
		Watch(ctx)
}

// Create takes the representation of a preview and creates it.  Returns the server's representation of the preview, and an error, if there is any.
func (c *previews) Create(ctx context.Context, preview *v1alpha1.Preview, opts v1.CreateOptions) (result *v1alpha1.Preview, err error) {
	result = &v1alpha1.Preview{}
	err = c.client.Post().
		Namespace(c.ns).
		Resource("previews").
		VersionedParams(&opts, scheme.ParameterCodec).
		Body(preview).
		Do(ctx).
		Into(result)
	return
}

// Update takes the representation of a preview and updates it. Returns the server's representation of the preview, and an error, if there is any.
func (c *previews) Update(ctx context.Context, preview *v1alpha1.Preview, opts v1.UpdateOptions) (result *v1alpha1.Preview, err error) {
	result = &v1alpha1.Preview{}
	err = c.client.Put().
		Namespace(c.ns).
		Resource("previews").
		Name(preview.Name).
		VersionedParams(&opts, scheme.ParameterCodec).
		Body(preview).
		Do(ctx).
		Into(result)
	return
}

// Delete takes name of the preview and deletes it. Returns an error if one occurs.
func (c *previews) Delete(ctx context.Context, name string, opts v1.DeleteOptions) error {
	return c.client.Delete().
		Namespace(c.ns).
		Resource("previews").
		Name(name).
		Body(&opts).
		Do(ctx).
		Error()
}

// DeleteCollection deletes a collection of objects.
func (c *previews) DeleteCollection(ctx context.Context, opts v1.DeleteOptions, listOpts v1.ListOptions) error {
	var timeout time.Duration
	if listOpts.TimeoutSeconds != nil {
		timeout = time.Duration(*listOpts.TimeoutSeconds) * time.Second
	}
	return c.client.Delete().
		Namespace(c.ns).
		Resource("previews").
		VersionedParams(&listOpts, scheme.ParameterCodec).
		Timeout(timeout).
		Body(&opts).
		Do(ctx).
		Error()
}

// Patch applies the patch and returns the patched preview.
func (c *previews) Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts v1.PatchOptions, subresources ...string) (result *v1alpha1.Preview, err error) {
	result = &v1alpha1.Preview{}
	err = c.client.Patch(pt).
		Namespace(c.ns).
		Resource("previews").
		Name(name).
		SubResource(subresources...).
		VersionedParams(&opts, scheme.ParameterCodec).
		Body(data).
		Do(ctx).
		Into(result)
	return
}
