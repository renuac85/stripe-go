//
//
// File generated from our OpenAPI spec
//
//

// Package taxrate provides the /tax_rates APIs
package taxrate

import (
	"net/http"

	stripe "github.com/stripe/stripe-go/v76"
	"github.com/stripe/stripe-go/v76/form"
)

// Client is used to invoke /tax_rates APIs.
type Client struct {
	B   stripe.Backend
	Key string
}

// New creates a new tax rate.
func New(params *stripe.TaxRateParams) (*stripe.TaxRate, error) {
	return getC().New(params)
}

// New creates a new tax rate.
func (c Client) New(params *stripe.TaxRateParams) (*stripe.TaxRate, error) {
	taxrate := &stripe.TaxRate{}
	var err error
	sr := stripe.NewStripeRequest(http.MethodPost, "/v1/tax_rates", c.Key)
	err = sr.SetParams(params)
	if err != nil {
		return nil, err
	}
	err = c.B.Call(sr, taxrate)
	return taxrate, err
}

// Get returns the details of a tax rate.
func Get(id string, params *stripe.TaxRateParams) (*stripe.TaxRate, error) {
	return getC().Get(id, params)
}

// Get returns the details of a tax rate.
func (c Client) Get(id string, params *stripe.TaxRateParams) (*stripe.TaxRate, error) {
	path := stripe.FormatURLPath("/v1/tax_rates/%s", id)
	taxrate := &stripe.TaxRate{}
	var err error
	sr := stripe.NewStripeRequest(http.MethodGet, path, c.Key)
	err = sr.SetParams(params)
	if err != nil {
		return nil, err
	}
	err = c.B.Call(sr, taxrate)
	return taxrate, err
}

// Update updates a tax rate's properties.
func Update(id string, params *stripe.TaxRateParams) (*stripe.TaxRate, error) {
	return getC().Update(id, params)
}

// Update updates a tax rate's properties.
func (c Client) Update(id string, params *stripe.TaxRateParams) (*stripe.TaxRate, error) {
	path := stripe.FormatURLPath("/v1/tax_rates/%s", id)
	taxrate := &stripe.TaxRate{}
	var err error
	sr := stripe.NewStripeRequest(http.MethodPost, path, c.Key)
	err = sr.SetParams(params)
	if err != nil {
		return nil, err
	}
	err = c.B.Call(sr, taxrate)
	return taxrate, err
}

// List returns a list of tax rates.
func List(params *stripe.TaxRateListParams) *Iter {
	return getC().List(params)
}

// List returns a list of tax rates.
func (c Client) List(listParams *stripe.TaxRateListParams) *Iter {
	return &Iter{
		Iter: stripe.GetIter(listParams, func(p *stripe.Params, b *form.Values) ([]interface{}, stripe.ListContainer, error) {
			list := &stripe.TaxRateList{}
			sr := stripe.NewStripeRequest(
				http.MethodGet,
				"/v1/tax_rates",
				c.Key,
			)
			err := sr.SetRawForm(p, b)
			if err != nil {
				return nil, list, err
			}
			err = c.B.Call(sr, list)

			ret := make([]interface{}, len(list.Data))
			for i, v := range list.Data {
				ret[i] = v
			}

			return ret, list, err
		}),
	}
}

// Iter is an iterator for tax rates.
type Iter struct {
	*stripe.Iter
}

// TaxRate returns the tax rate which the iterator is currently pointing to.
func (i *Iter) TaxRate() *stripe.TaxRate {
	return i.Current().(*stripe.TaxRate)
}

// TaxRateList returns the current list object which the iterator is
// currently using. List objects will change as new API calls are made to
// continue pagination.
func (i *Iter) TaxRateList() *stripe.TaxRateList {
	return i.List().(*stripe.TaxRateList)
}

func getC() Client {
	return Client{stripe.GetBackend(stripe.APIBackend), stripe.Key}
}
