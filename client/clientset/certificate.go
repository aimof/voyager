package clientset

import (
	aci "github.com/appscode/voyager/api"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/watch"
	"k8s.io/client-go/rest"
)

type CertificateGetter interface {
	Certificates(namespace string) CertificateInterface
}

type CertificateInterface interface {
	List(opts metav1.ListOptions) (*aci.CertificateList, error)
	Get(name string) (*aci.Certificate, error)
	Create(certificate *aci.Certificate) (*aci.Certificate, error)
	Update(certificate *aci.Certificate) (*aci.Certificate, error)
	Delete(name string) error
	Watch(opts metav1.ListOptions) (watch.Interface, error)
	UpdateStatus(certificate *aci.Certificate) (*aci.Certificate, error)
}

type CertificateImpl struct {
	r  rest.Interface
	ns string
}

var _ CertificateInterface = &CertificateImpl{}

func newCertificate(c *ExtensionClient, namespace string) *CertificateImpl {
	return &CertificateImpl{c.restClient, namespace}
}

func (c *CertificateImpl) List(opts metav1.ListOptions) (result *aci.CertificateList, err error) {
	result = &aci.CertificateList{}
	err = c.r.Get().
		Namespace(c.ns).
		Resource("certificates").
		VersionedParams(&opts, ExtendedCodec).
		Do().
		Into(result)
	return
}

func (c *CertificateImpl) Get(name string) (result *aci.Certificate, err error) {
	result = &aci.Certificate{}
	err = c.r.Get().
		Namespace(c.ns).
		Resource("certificates").
		Name(name).
		Do().
		Into(result)
	return
}

func (c *CertificateImpl) Create(certificate *aci.Certificate) (result *aci.Certificate, err error) {
	result = &aci.Certificate{}
	err = c.r.Post().
		Namespace(c.ns).
		Resource("certificates").
		Body(certificate).
		Do().
		Into(result)
	return
}

func (c *CertificateImpl) Update(certificate *aci.Certificate) (result *aci.Certificate, err error) {
	result = &aci.Certificate{}
	err = c.r.Put().
		Namespace(c.ns).
		Resource("certificates").
		Name(certificate.Name).
		Body(certificate).
		Do().
		Into(result)
	return
}

func (c *CertificateImpl) Delete(name string) (err error) {
	return c.r.Delete().
		Namespace(c.ns).
		Resource("certificates").
		Name(name).
		Do().
		Error()
}

func (c *CertificateImpl) Watch(opts metav1.ListOptions) (watch.Interface, error) {
	return c.r.Get().
		Prefix("watch").
		Namespace(c.ns).
		Resource("certificates").
		VersionedParams(&opts, ExtendedCodec).
		Watch()
}

func (c *CertificateImpl) UpdateStatus(certificate *aci.Certificate) (result *aci.Certificate, err error) {
	result = &aci.Certificate{}
	err = c.r.Put().
		Namespace(c.ns).
		Resource("certificates").
		Name(certificate.Name).
		SubResource("status").
		Body(certificate).
		Do().
		Into(result)
	return
}
