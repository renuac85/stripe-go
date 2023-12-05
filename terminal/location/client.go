//
//
// File generated from our OpenAPI spec
//
//

// Package location provides the /terminal/locations APIs
package location

import (
	"net/http"

	stripe "github.com/stripe/stripe-go/v76"
	"github.com/stripe/stripe-go/v76/form"
)

// Client is used to invoke /terminal/locations APIs.
type Client struct {
	B   stripe.Backend
	Key string
}

// New creates a new terminal location.
func New(params *stripe.TerminalLocationParams) (*stripe.TerminalLocation, error) {
	return getC().New(params)
}

// New creates a new terminal location.
func (c Client) New(params *stripe.TerminalLocationParams) (*stripe.TerminalLocation, error) {
	location := &stripe.TerminalLocation{}
	var err error
	sr := stripe.NewStripeRequest(
		http.MethodPost,
		"/v1/terminal/locations",
		c.Key,
	)
	err = sr.SetParams(params)
	if err != nil {
		return nil, err
	}
	err = c.B.Call(sr, location)
	return location, err
}

// Get returns the details of a terminal location.
func Get(id string, params *stripe.TerminalLocationParams) (*stripe.TerminalLocation, error) {
	return getC().Get(id, params)
}

// Get returns the details of a terminal location.
func (c Client) Get(id string, params *stripe.TerminalLocationParams) (*stripe.TerminalLocation, error) {
	path := stripe.FormatURLPath("/v1/terminal/locations/%s", id)
	location := &stripe.TerminalLocation{}
	var err error
	sr := stripe.NewStripeRequest(http.MethodGet, path, c.Key)
	err = sr.SetParams(params)
	if err != nil {
		return nil, err
	}
	err = c.B.Call(sr, location)
	return location, err
}

// Update updates a terminal location's properties.
func Update(id string, params *stripe.TerminalLocationParams) (*stripe.TerminalLocation, error) {
	return getC().Update(id, params)
}

// Update updates a terminal location's properties.
func (c Client) Update(id string, params *stripe.TerminalLocationParams) (*stripe.TerminalLocation, error) {
	path := stripe.FormatURLPath("/v1/terminal/locations/%s", id)
	location := &stripe.TerminalLocation{}
	var err error
	sr := stripe.NewStripeRequest(http.MethodPost, path, c.Key)
	err = sr.SetParams(params)
	if err != nil {
		return nil, err
	}
	err = c.B.Call(sr, location)
	return location, err
}

// Del removes a terminal location.
func Del(id string, params *stripe.TerminalLocationParams) (*stripe.TerminalLocation, error) {
	return getC().Del(id, params)
}

// Del removes a terminal location.
func (c Client) Del(id string, params *stripe.TerminalLocationParams) (*stripe.TerminalLocation, error) {
	path := stripe.FormatURLPath("/v1/terminal/locations/%s", id)
	location := &stripe.TerminalLocation{}
	var err error
	sr := stripe.NewStripeRequest(http.MethodDelete, path, c.Key)
	err = sr.SetParams(params)
	if err != nil {
		return nil, err
	}
	err = c.B.Call(sr, location)
	return location, err
}

// List returns a list of terminal locations.
func List(params *stripe.TerminalLocationListParams) *Iter {
	return getC().List(params)
}

// List returns a list of terminal locations.
func (c Client) List(listParams *stripe.TerminalLocationListParams) *Iter {
	return &Iter{
		Iter: stripe.GetIter(listParams, func(p *stripe.Params, b *form.Values) ([]interface{}, stripe.ListContainer, error) {
			list := &stripe.TerminalLocationList{}
			sr := stripe.NewStripeRequest(
				http.MethodGet,
				"/v1/terminal/locations",
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

// Iter is an iterator for terminal locations.
type Iter struct {
	*stripe.Iter
}

// TerminalLocation returns the terminal location which the iterator is currently pointing to.
func (i *Iter) TerminalLocation() *stripe.TerminalLocation {
	return i.Current().(*stripe.TerminalLocation)
}

// TerminalLocationList returns the current list object which the iterator is
// currently using. List objects will change as new API calls are made to
// continue pagination.
func (i *Iter) TerminalLocationList() *stripe.TerminalLocationList {
	return i.List().(*stripe.TerminalLocationList)
}

func getC() Client {
	return Client{stripe.GetBackend(stripe.APIBackend), stripe.Key}
}
