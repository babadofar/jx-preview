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

package fake

import (
	"context"

	v1alpha1 "github.com/jenkins-x-plugins/jx-preview/pkg/apis/preview/v1alpha1"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	labels "k8s.io/apimachinery/pkg/labels"
	schema "k8s.io/apimachinery/pkg/runtime/schema"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	testing "k8s.io/client-go/testing"
)

// FakePreviews implements PreviewInterface
type FakePreviews struct {
	Fake *FakePreviewV1alpha1
	ns   string
}

var previewsResource = schema.GroupVersionResource{Group: "preview.jenkins.io", Version: "v1alpha1", Resource: "previews"}

var previewsKind = schema.GroupVersionKind{Group: "preview.jenkins.io", Version: "v1alpha1", Kind: "Preview"}

// Get takes name of the preview, and returns the corresponding preview object, and an error if there is any.
func (c *FakePreviews) Get(ctx context.Context, name string, options v1.GetOptions) (result *v1alpha1.Preview, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewGetAction(previewsResource, c.ns, name), &v1alpha1.Preview{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.Preview), err
}

// List takes label and field selectors, and returns the list of Previews that match those selectors.
func (c *FakePreviews) List(ctx context.Context, opts v1.ListOptions) (result *v1alpha1.PreviewList, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewListAction(previewsResource, previewsKind, c.ns, opts), &v1alpha1.PreviewList{})

	if obj == nil {
		return nil, err
	}

	label, _, _ := testing.ExtractFromListOptions(opts)
	if label == nil {
		label = labels.Everything()
	}
	list := &v1alpha1.PreviewList{ListMeta: obj.(*v1alpha1.PreviewList).ListMeta}
	for _, item := range obj.(*v1alpha1.PreviewList).Items {
		if label.Matches(labels.Set(item.Labels)) {
			list.Items = append(list.Items, item)
		}
	}
	return list, err
}

// Watch returns a watch.Interface that watches the requested previews.
func (c *FakePreviews) Watch(ctx context.Context, opts v1.ListOptions) (watch.Interface, error) {
	return c.Fake.
		InvokesWatch(testing.NewWatchAction(previewsResource, c.ns, opts))

}

// Create takes the representation of a preview and creates it.  Returns the server's representation of the preview, and an error, if there is any.
func (c *FakePreviews) Create(ctx context.Context, preview *v1alpha1.Preview, opts v1.CreateOptions) (result *v1alpha1.Preview, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewCreateAction(previewsResource, c.ns, preview), &v1alpha1.Preview{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.Preview), err
}

// Update takes the representation of a preview and updates it. Returns the server's representation of the preview, and an error, if there is any.
func (c *FakePreviews) Update(ctx context.Context, preview *v1alpha1.Preview, opts v1.UpdateOptions) (result *v1alpha1.Preview, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateAction(previewsResource, c.ns, preview), &v1alpha1.Preview{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.Preview), err
}

// Delete takes name of the preview and deletes it. Returns an error if one occurs.
func (c *FakePreviews) Delete(ctx context.Context, name string, opts v1.DeleteOptions) error {
	_, err := c.Fake.
		Invokes(testing.NewDeleteAction(previewsResource, c.ns, name), &v1alpha1.Preview{})

	return err
}

// DeleteCollection deletes a collection of objects.
func (c *FakePreviews) DeleteCollection(ctx context.Context, opts v1.DeleteOptions, listOpts v1.ListOptions) error {
	action := testing.NewDeleteCollectionAction(previewsResource, c.ns, listOpts)

	_, err := c.Fake.Invokes(action, &v1alpha1.PreviewList{})
	return err
}

// Patch applies the patch and returns the patched preview.
func (c *FakePreviews) Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts v1.PatchOptions, subresources ...string) (result *v1alpha1.Preview, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewPatchSubresourceAction(previewsResource, c.ns, name, pt, data, subresources...), &v1alpha1.Preview{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.Preview), err
}
