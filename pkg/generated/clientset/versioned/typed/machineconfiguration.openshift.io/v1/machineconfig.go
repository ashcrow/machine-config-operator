// Code generated by client-gen. DO NOT EDIT.

package v1

import (
	v1 "github.com/openshift/machine-config-operator/pkg/apis/machineconfiguration.openshift.io/v1"
	scheme "github.com/openshift/machine-config-operator/pkg/generated/clientset/versioned/scheme"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	rest "k8s.io/client-go/rest"
)

// MachineConfigsGetter has a method to return a MachineConfigInterface.
// A group's client should implement this interface.
type MachineConfigsGetter interface {
	MachineConfigs(namespace string) MachineConfigInterface
}

// MachineConfigInterface has methods to work with MachineConfig resources.
type MachineConfigInterface interface {
	Create(*v1.MachineConfig) (*v1.MachineConfig, error)
	Update(*v1.MachineConfig) (*v1.MachineConfig, error)
	Delete(name string, options *metav1.DeleteOptions) error
	DeleteCollection(options *metav1.DeleteOptions, listOptions metav1.ListOptions) error
	Get(name string, options metav1.GetOptions) (*v1.MachineConfig, error)
	List(opts metav1.ListOptions) (*v1.MachineConfigList, error)
	Watch(opts metav1.ListOptions) (watch.Interface, error)
	Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v1.MachineConfig, err error)
	MachineConfigExpansion
}

// machineConfigs implements MachineConfigInterface
type machineConfigs struct {
	client rest.Interface
	ns     string
}

// newMachineConfigs returns a MachineConfigs
func newMachineConfigs(c *MachineconfigurationV1Client, namespace string) *machineConfigs {
	return &machineConfigs{
		client: c.RESTClient(),
		ns:     namespace,
	}
}

// Get takes name of the machineConfig, and returns the corresponding machineConfig object, and an error if there is any.
func (c *machineConfigs) Get(name string, options metav1.GetOptions) (result *v1.MachineConfig, err error) {
	result = &v1.MachineConfig{}
	err = c.client.Get().
		Namespace(c.ns).
		Resource("machineconfigs").
		Name(name).
		VersionedParams(&options, scheme.ParameterCodec).
		Do().
		Into(result)
	return
}

// List takes label and field selectors, and returns the list of MachineConfigs that match those selectors.
func (c *machineConfigs) List(opts metav1.ListOptions) (result *v1.MachineConfigList, err error) {
	result = &v1.MachineConfigList{}
	err = c.client.Get().
		Namespace(c.ns).
		Resource("machineconfigs").
		VersionedParams(&opts, scheme.ParameterCodec).
		Do().
		Into(result)
	return
}

// Watch returns a watch.Interface that watches the requested machineConfigs.
func (c *machineConfigs) Watch(opts metav1.ListOptions) (watch.Interface, error) {
	opts.Watch = true
	return c.client.Get().
		Namespace(c.ns).
		Resource("machineconfigs").
		VersionedParams(&opts, scheme.ParameterCodec).
		Watch()
}

// Create takes the representation of a machineConfig and creates it.  Returns the server's representation of the machineConfig, and an error, if there is any.
func (c *machineConfigs) Create(machineConfig *v1.MachineConfig) (result *v1.MachineConfig, err error) {
	result = &v1.MachineConfig{}
	err = c.client.Post().
		Namespace(c.ns).
		Resource("machineconfigs").
		Body(machineConfig).
		Do().
		Into(result)
	return
}

// Update takes the representation of a machineConfig and updates it. Returns the server's representation of the machineConfig, and an error, if there is any.
func (c *machineConfigs) Update(machineConfig *v1.MachineConfig) (result *v1.MachineConfig, err error) {
	result = &v1.MachineConfig{}
	err = c.client.Put().
		Namespace(c.ns).
		Resource("machineconfigs").
		Name(machineConfig.Name).
		Body(machineConfig).
		Do().
		Into(result)
	return
}

// Delete takes name of the machineConfig and deletes it. Returns an error if one occurs.
func (c *machineConfigs) Delete(name string, options *metav1.DeleteOptions) error {
	return c.client.Delete().
		Namespace(c.ns).
		Resource("machineconfigs").
		Name(name).
		Body(options).
		Do().
		Error()
}

// DeleteCollection deletes a collection of objects.
func (c *machineConfigs) DeleteCollection(options *metav1.DeleteOptions, listOptions metav1.ListOptions) error {
	return c.client.Delete().
		Namespace(c.ns).
		Resource("machineconfigs").
		VersionedParams(&listOptions, scheme.ParameterCodec).
		Body(options).
		Do().
		Error()
}

// Patch applies the patch and returns the patched machineConfig.
func (c *machineConfigs) Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v1.MachineConfig, err error) {
	result = &v1.MachineConfig{}
	err = c.client.Patch(pt).
		Namespace(c.ns).
		Resource("machineconfigs").
		SubResource(subresources...).
		Name(name).
		Body(data).
		Do().
		Into(result)
	return
}
