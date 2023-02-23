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
	"context"
	json "encoding/json"
	"fmt"

	v1alpha1 "github.com/nidhey27/respaldo/pkg/apis/nyctonid.dev/v1alpha1"
	nyctoniddevv1alpha1 "github.com/nidhey27/respaldo/pkg/client/applyconfiguration/nyctonid.dev/v1alpha1"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	labels "k8s.io/apimachinery/pkg/labels"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	testing "k8s.io/client-go/testing"
)

// FakeRespaldos implements RespaldoInterface
type FakeRespaldos struct {
	Fake *FakeNyctonidV1alpha1
	ns   string
}

var respaldosResource = v1alpha1.SchemeGroupVersion.WithResource("respaldos")

var respaldosKind = v1alpha1.SchemeGroupVersion.WithKind("Respaldo")

// Get takes name of the respaldo, and returns the corresponding respaldo object, and an error if there is any.
func (c *FakeRespaldos) Get(ctx context.Context, name string, options v1.GetOptions) (result *v1alpha1.Respaldo, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewGetAction(respaldosResource, c.ns, name), &v1alpha1.Respaldo{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.Respaldo), err
}

// List takes label and field selectors, and returns the list of Respaldos that match those selectors.
func (c *FakeRespaldos) List(ctx context.Context, opts v1.ListOptions) (result *v1alpha1.RespaldoList, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewListAction(respaldosResource, respaldosKind, c.ns, opts), &v1alpha1.RespaldoList{})

	if obj == nil {
		return nil, err
	}

	label, _, _ := testing.ExtractFromListOptions(opts)
	if label == nil {
		label = labels.Everything()
	}
	list := &v1alpha1.RespaldoList{ListMeta: obj.(*v1alpha1.RespaldoList).ListMeta}
	for _, item := range obj.(*v1alpha1.RespaldoList).Items {
		if label.Matches(labels.Set(item.Labels)) {
			list.Items = append(list.Items, item)
		}
	}
	return list, err
}

// Watch returns a watch.Interface that watches the requested respaldos.
func (c *FakeRespaldos) Watch(ctx context.Context, opts v1.ListOptions) (watch.Interface, error) {
	return c.Fake.
		InvokesWatch(testing.NewWatchAction(respaldosResource, c.ns, opts))

}

// Create takes the representation of a respaldo and creates it.  Returns the server's representation of the respaldo, and an error, if there is any.
func (c *FakeRespaldos) Create(ctx context.Context, respaldo *v1alpha1.Respaldo, opts v1.CreateOptions) (result *v1alpha1.Respaldo, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewCreateAction(respaldosResource, c.ns, respaldo), &v1alpha1.Respaldo{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.Respaldo), err
}

// Update takes the representation of a respaldo and updates it. Returns the server's representation of the respaldo, and an error, if there is any.
func (c *FakeRespaldos) Update(ctx context.Context, respaldo *v1alpha1.Respaldo, opts v1.UpdateOptions) (result *v1alpha1.Respaldo, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateAction(respaldosResource, c.ns, respaldo), &v1alpha1.Respaldo{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.Respaldo), err
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().
func (c *FakeRespaldos) UpdateStatus(ctx context.Context, respaldo *v1alpha1.Respaldo, opts v1.UpdateOptions) (*v1alpha1.Respaldo, error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateSubresourceAction(respaldosResource, "status", c.ns, respaldo), &v1alpha1.Respaldo{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.Respaldo), err
}

// Delete takes name of the respaldo and deletes it. Returns an error if one occurs.
func (c *FakeRespaldos) Delete(ctx context.Context, name string, opts v1.DeleteOptions) error {
	_, err := c.Fake.
		Invokes(testing.NewDeleteActionWithOptions(respaldosResource, c.ns, name, opts), &v1alpha1.Respaldo{})

	return err
}

// DeleteCollection deletes a collection of objects.
func (c *FakeRespaldos) DeleteCollection(ctx context.Context, opts v1.DeleteOptions, listOpts v1.ListOptions) error {
	action := testing.NewDeleteCollectionAction(respaldosResource, c.ns, listOpts)

	_, err := c.Fake.Invokes(action, &v1alpha1.RespaldoList{})
	return err
}

// Patch applies the patch and returns the patched respaldo.
func (c *FakeRespaldos) Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts v1.PatchOptions, subresources ...string) (result *v1alpha1.Respaldo, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewPatchSubresourceAction(respaldosResource, c.ns, name, pt, data, subresources...), &v1alpha1.Respaldo{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.Respaldo), err
}

// Apply takes the given apply declarative configuration, applies it and returns the applied respaldo.
func (c *FakeRespaldos) Apply(ctx context.Context, respaldo *nyctoniddevv1alpha1.RespaldoApplyConfiguration, opts v1.ApplyOptions) (result *v1alpha1.Respaldo, err error) {
	if respaldo == nil {
		return nil, fmt.Errorf("respaldo provided to Apply must not be nil")
	}
	data, err := json.Marshal(respaldo)
	if err != nil {
		return nil, err
	}
	name := respaldo.Name
	if name == nil {
		return nil, fmt.Errorf("respaldo.Name must be provided to Apply")
	}
	obj, err := c.Fake.
		Invokes(testing.NewPatchSubresourceAction(respaldosResource, c.ns, *name, types.ApplyPatchType, data), &v1alpha1.Respaldo{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.Respaldo), err
}

// ApplyStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating ApplyStatus().
func (c *FakeRespaldos) ApplyStatus(ctx context.Context, respaldo *nyctoniddevv1alpha1.RespaldoApplyConfiguration, opts v1.ApplyOptions) (result *v1alpha1.Respaldo, err error) {
	if respaldo == nil {
		return nil, fmt.Errorf("respaldo provided to Apply must not be nil")
	}
	data, err := json.Marshal(respaldo)
	if err != nil {
		return nil, err
	}
	name := respaldo.Name
	if name == nil {
		return nil, fmt.Errorf("respaldo.Name must be provided to Apply")
	}
	obj, err := c.Fake.
		Invokes(testing.NewPatchSubresourceAction(respaldosResource, c.ns, *name, types.ApplyPatchType, data, "status"), &v1alpha1.Respaldo{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.Respaldo), err
}
