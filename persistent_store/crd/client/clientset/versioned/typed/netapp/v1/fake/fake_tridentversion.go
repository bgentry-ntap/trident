// Copyright 2019 NetApp, Inc. All Rights Reserved.

// Code generated by client-gen. DO NOT EDIT.

package fake

import (
	netappv1 "github.com/netapp/trident/persistent_store/crd/apis/netapp/v1"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	labels "k8s.io/apimachinery/pkg/labels"
	schema "k8s.io/apimachinery/pkg/runtime/schema"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	testing "k8s.io/client-go/testing"
)

// FakeTridentVersions implements TridentVersionInterface
type FakeTridentVersions struct {
	Fake *FakeTridentV1
	ns   string
}

var tridentversionsResource = schema.GroupVersionResource{Group: "trident.netapp.io", Version: "v1", Resource: "tridentversions"}

var tridentversionsKind = schema.GroupVersionKind{Group: "trident.netapp.io", Version: "v1", Kind: "TridentVersion"}

// Get takes name of the tridentVersion, and returns the corresponding tridentVersion object, and an error if there is any.
func (c *FakeTridentVersions) Get(name string, options v1.GetOptions) (result *netappv1.TridentVersion, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewGetAction(tridentversionsResource, c.ns, name), &netappv1.TridentVersion{})

	if obj == nil {
		return nil, err
	}
	return obj.(*netappv1.TridentVersion), err
}

// List takes label and field selectors, and returns the list of TridentVersions that match those selectors.
func (c *FakeTridentVersions) List(opts v1.ListOptions) (result *netappv1.TridentVersionList, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewListAction(tridentversionsResource, tridentversionsKind, c.ns, opts), &netappv1.TridentVersionList{})

	if obj == nil {
		return nil, err
	}

	label, _, _ := testing.ExtractFromListOptions(opts)
	if label == nil {
		label = labels.Everything()
	}
	list := &netappv1.TridentVersionList{ListMeta: obj.(*netappv1.TridentVersionList).ListMeta}
	for _, item := range obj.(*netappv1.TridentVersionList).Items {
		if label.Matches(labels.Set(item.Labels)) {
			list.Items = append(list.Items, item)
		}
	}
	return list, err
}

// Watch returns a watch.Interface that watches the requested tridentVersions.
func (c *FakeTridentVersions) Watch(opts v1.ListOptions) (watch.Interface, error) {
	return c.Fake.
		InvokesWatch(testing.NewWatchAction(tridentversionsResource, c.ns, opts))

}

// Create takes the representation of a tridentVersion and creates it.  Returns the server's representation of the tridentVersion, and an error, if there is any.
func (c *FakeTridentVersions) Create(tridentVersion *netappv1.TridentVersion) (result *netappv1.TridentVersion, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewCreateAction(tridentversionsResource, c.ns, tridentVersion), &netappv1.TridentVersion{})

	if obj == nil {
		return nil, err
	}
	return obj.(*netappv1.TridentVersion), err
}

// Update takes the representation of a tridentVersion and updates it. Returns the server's representation of the tridentVersion, and an error, if there is any.
func (c *FakeTridentVersions) Update(tridentVersion *netappv1.TridentVersion) (result *netappv1.TridentVersion, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateAction(tridentversionsResource, c.ns, tridentVersion), &netappv1.TridentVersion{})

	if obj == nil {
		return nil, err
	}
	return obj.(*netappv1.TridentVersion), err
}

// Delete takes name of the tridentVersion and deletes it. Returns an error if one occurs.
func (c *FakeTridentVersions) Delete(name string, options *v1.DeleteOptions) error {
	_, err := c.Fake.
		Invokes(testing.NewDeleteAction(tridentversionsResource, c.ns, name), &netappv1.TridentVersion{})

	return err
}

// DeleteCollection deletes a collection of objects.
func (c *FakeTridentVersions) DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error {
	action := testing.NewDeleteCollectionAction(tridentversionsResource, c.ns, listOptions)

	_, err := c.Fake.Invokes(action, &netappv1.TridentVersionList{})
	return err
}

// Patch applies the patch and returns the patched tridentVersion.
func (c *FakeTridentVersions) Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *netappv1.TridentVersion, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewPatchSubresourceAction(tridentversionsResource, c.ns, name, pt, data, subresources...), &netappv1.TridentVersion{})

	if obj == nil {
		return nil, err
	}
	return obj.(*netappv1.TridentVersion), err
}
