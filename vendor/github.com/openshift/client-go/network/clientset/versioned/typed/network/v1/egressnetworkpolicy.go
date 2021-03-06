package v1

import (
	v1 "github.com/openshift/api/network/v1"
	scheme "github.com/openshift/client-go/network/clientset/versioned/scheme"
	meta_v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	rest "k8s.io/client-go/rest"
)

// EgressNetworkPoliciesGetter has a method to return a EgressNetworkPolicyInterface.
// A group's client should implement this interface.
type EgressNetworkPoliciesGetter interface {
	EgressNetworkPolicies(namespace string) EgressNetworkPolicyInterface
}

// EgressNetworkPolicyInterface has methods to work with EgressNetworkPolicy resources.
type EgressNetworkPolicyInterface interface {
	Create(*v1.EgressNetworkPolicy) (*v1.EgressNetworkPolicy, error)
	Update(*v1.EgressNetworkPolicy) (*v1.EgressNetworkPolicy, error)
	Delete(name string, options *meta_v1.DeleteOptions) error
	DeleteCollection(options *meta_v1.DeleteOptions, listOptions meta_v1.ListOptions) error
	Get(name string, options meta_v1.GetOptions) (*v1.EgressNetworkPolicy, error)
	List(opts meta_v1.ListOptions) (*v1.EgressNetworkPolicyList, error)
	Watch(opts meta_v1.ListOptions) (watch.Interface, error)
	Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v1.EgressNetworkPolicy, err error)
	EgressNetworkPolicyExpansion
}

// egressNetworkPolicies implements EgressNetworkPolicyInterface
type egressNetworkPolicies struct {
	client rest.Interface
	ns     string
}

// newEgressNetworkPolicies returns a EgressNetworkPolicies
func newEgressNetworkPolicies(c *NetworkV1Client, namespace string) *egressNetworkPolicies {
	return &egressNetworkPolicies{
		client: c.RESTClient(),
		ns:     namespace,
	}
}

// Get takes name of the egressNetworkPolicy, and returns the corresponding egressNetworkPolicy object, and an error if there is any.
func (c *egressNetworkPolicies) Get(name string, options meta_v1.GetOptions) (result *v1.EgressNetworkPolicy, err error) {
	result = &v1.EgressNetworkPolicy{}
	err = c.client.Get().
		Namespace(c.ns).
		Resource("egressnetworkpolicies").
		Name(name).
		VersionedParams(&options, scheme.ParameterCodec).
		Do().
		Into(result)
	return
}

// List takes label and field selectors, and returns the list of EgressNetworkPolicies that match those selectors.
func (c *egressNetworkPolicies) List(opts meta_v1.ListOptions) (result *v1.EgressNetworkPolicyList, err error) {
	result = &v1.EgressNetworkPolicyList{}
	err = c.client.Get().
		Namespace(c.ns).
		Resource("egressnetworkpolicies").
		VersionedParams(&opts, scheme.ParameterCodec).
		Do().
		Into(result)
	return
}

// Watch returns a watch.Interface that watches the requested egressNetworkPolicies.
func (c *egressNetworkPolicies) Watch(opts meta_v1.ListOptions) (watch.Interface, error) {
	opts.Watch = true
	return c.client.Get().
		Namespace(c.ns).
		Resource("egressnetworkpolicies").
		VersionedParams(&opts, scheme.ParameterCodec).
		Watch()
}

// Create takes the representation of a egressNetworkPolicy and creates it.  Returns the server's representation of the egressNetworkPolicy, and an error, if there is any.
func (c *egressNetworkPolicies) Create(egressNetworkPolicy *v1.EgressNetworkPolicy) (result *v1.EgressNetworkPolicy, err error) {
	result = &v1.EgressNetworkPolicy{}
	err = c.client.Post().
		Namespace(c.ns).
		Resource("egressnetworkpolicies").
		Body(egressNetworkPolicy).
		Do().
		Into(result)
	return
}

// Update takes the representation of a egressNetworkPolicy and updates it. Returns the server's representation of the egressNetworkPolicy, and an error, if there is any.
func (c *egressNetworkPolicies) Update(egressNetworkPolicy *v1.EgressNetworkPolicy) (result *v1.EgressNetworkPolicy, err error) {
	result = &v1.EgressNetworkPolicy{}
	err = c.client.Put().
		Namespace(c.ns).
		Resource("egressnetworkpolicies").
		Name(egressNetworkPolicy.Name).
		Body(egressNetworkPolicy).
		Do().
		Into(result)
	return
}

// Delete takes name of the egressNetworkPolicy and deletes it. Returns an error if one occurs.
func (c *egressNetworkPolicies) Delete(name string, options *meta_v1.DeleteOptions) error {
	return c.client.Delete().
		Namespace(c.ns).
		Resource("egressnetworkpolicies").
		Name(name).
		Body(options).
		Do().
		Error()
}

// DeleteCollection deletes a collection of objects.
func (c *egressNetworkPolicies) DeleteCollection(options *meta_v1.DeleteOptions, listOptions meta_v1.ListOptions) error {
	return c.client.Delete().
		Namespace(c.ns).
		Resource("egressnetworkpolicies").
		VersionedParams(&listOptions, scheme.ParameterCodec).
		Body(options).
		Do().
		Error()
}

// Patch applies the patch and returns the patched egressNetworkPolicy.
func (c *egressNetworkPolicies) Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v1.EgressNetworkPolicy, err error) {
	result = &v1.EgressNetworkPolicy{}
	err = c.client.Patch(pt).
		Namespace(c.ns).
		Resource("egressnetworkpolicies").
		SubResource(subresources...).
		Name(name).
		Body(data).
		Do().
		Into(result)
	return
}
