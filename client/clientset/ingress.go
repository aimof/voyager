package clientset

import (
	aci "github.com/appscode/voyager/api"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/watch"
	"k8s.io/client-go/rest"
)

// IngressesGetter has methods to work with ExtendedIngress resources in a namespace
type IngressesGetter interface {
	Ingresses(namespace string) IngressInterface
}

// ExtendedIngressInterface exposes methods to work on ExtendedIngress resources.
type IngressInterface interface {
	List(opts metav1.ListOptions) (*aci.IngressList, error)
	Get(name string) (*aci.Ingress, error)
	Create(ExtendedIngress *aci.Ingress) (*aci.Ingress, error)
	Update(ExtendedIngress *aci.Ingress) (*aci.Ingress, error)
	Delete(name string) error
	Watch(opts metav1.ListOptions) (watch.Interface, error)
	UpdateStatus(ExtendedIngress *aci.Ingress) (*aci.Ingress, error)
}

// IngressImpl implements IngressesGetter interface
type IngressImpl struct {
	r  rest.Interface
	ns string
}

var _ IngressInterface = &IngressImpl{}

// newExtendedIngress returns a ExtendedIngress
func newExtendedIngress(c *ExtensionClient, namespace string) *IngressImpl {
	return &IngressImpl{c.restClient, namespace}
}

// List returns a list of ExtendedIngress that match the label and field selectors.
func (c *IngressImpl) List(opts metav1.ListOptions) (result *aci.IngressList, err error) {
	result = &aci.IngressList{}
	err = c.r.Get().
		Namespace(c.ns).
		Resource("ingresses").
		VersionedParams(&opts, ExtendedCodec).
		Do().
		Into(result)
	return
}

// Get returns information about a particular ExtendedIngress.
func (c *IngressImpl) Get(name string) (result *aci.Ingress, err error) {
	result = &aci.Ingress{}
	err = c.r.Get().
		Namespace(c.ns).
		Resource("ingresses").
		Name(name).
		Do().
		Into(result)
	return
}

// Create creates a new ExtendedIngress.
func (c *IngressImpl) Create(extendedIngress *aci.Ingress) (result *aci.Ingress, err error) {
	result = &aci.Ingress{}
	err = c.r.Post().
		Namespace(c.ns).
		Resource("ingresses").
		Body(extendedIngress).
		Do().
		Into(result)
	return
}

// Update updates an existing ExtendedIngress.
func (c *IngressImpl) Update(extendedIngress *aci.Ingress) (result *aci.Ingress, err error) {
	result = &aci.Ingress{}
	err = c.r.Put().
		Namespace(c.ns).
		Resource("ingresses").
		Name(extendedIngress.Name).
		Body(extendedIngress).
		Do().
		Into(result)
	return
}

// Delete deletes a ExtendedIngress, returns error if one occurs.
func (c *IngressImpl) Delete(name string) (err error) {
	return c.r.Delete().
		Namespace(c.ns).
		Resource("ingresses").
		Name(name).
		Do().
		Error()
}

// Watch returns a watch.Interface that watches the requested ExtendedIngress.
func (c *IngressImpl) Watch(opts metav1.ListOptions) (watch.Interface, error) {
	return c.r.Get().
		Prefix("watch").
		Namespace(c.ns).
		Resource("ingresses").
		VersionedParams(&opts, ExtendedCodec).
		Watch()
}

// UpdateStatus takes the name of the ExtendedIngress and the new status.  Returns the server's representation of the ExtendedIngress, and an error, if it occurs.
func (c *IngressImpl) UpdateStatus(extendedIngress *aci.Ingress) (result *aci.Ingress, err error) {
	result = &aci.Ingress{}
	err = c.r.Put().
		Namespace(c.ns).
		Resource("ingresses").
		Name(extendedIngress.Name).
		SubResource("status").
		Body(extendedIngress).
		Do().
		Into(result)
	return
}
