//
//
// File generated from our OpenAPI spec
//
//

// Package file provides the /files APIs
package file

import (
	"fmt"
	"net/http"

	stripe "github.com/stripe/stripe-go/v76"
	"github.com/stripe/stripe-go/v76/form"
)

// Client is used to invoke /files APIs.
type Client struct {
	B        stripe.Backend
	BUploads stripe.Backend
	Key      string
}

// New creates a new file.
func New(params *stripe.FileParams) (*stripe.File, error) {
	return getC().New(params)
}

// New creates a new file.
func (c Client) New(params *stripe.FileParams) (*stripe.File, error) {
	if params == nil {
		return nil, fmt.Errorf(
			"params cannot be nil, and params.Purpose and params.File must be set",
		)
	}

	bodyBuffer, boundary, err := params.GetBody()
	if err != nil {
		return nil, err
	}

	file := &stripe.File{}
	sr := stripe.NewStripeRequest(http.MethodPost, "/v1/files", c.Key)
	err = sr.SetMultipart(&params.Params, boundary, bodyBuffer)
	if err != nil {
		return nil, err
	}
	err = c.BUploads.Call(sr, file)

	return file, err
}

// Get returns the details of a file.
func Get(id string, params *stripe.FileParams) (*stripe.File, error) {
	return getC().Get(id, params)
}

// Get returns the details of a file.
func (c Client) Get(id string, params *stripe.FileParams) (*stripe.File, error) {
	path := stripe.FormatURLPath("/v1/files/%s", id)
	file := &stripe.File{}
	var err error
	sr := stripe.NewStripeRequest(http.MethodGet, path, c.Key)
	err = sr.SetParams(params)
	if err != nil {
		return nil, err
	}
	err = c.B.Call(sr, file)
	return file, err
}

// List returns a list of files.
func List(params *stripe.FileListParams) *Iter {
	return getC().List(params)
}

// List returns a list of files.
func (c Client) List(listParams *stripe.FileListParams) *Iter {
	return &Iter{
		Iter: stripe.GetIter(listParams, func(p *stripe.Params, b *form.Values) ([]interface{}, stripe.ListContainer, error) {
			list := &stripe.FileList{}
			sr := stripe.NewStripeRequest(
				http.MethodGet,
				"/v1/files",
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

// Iter is an iterator for files.
type Iter struct {
	*stripe.Iter
}

// File returns the file which the iterator is currently pointing to.
func (i *Iter) File() *stripe.File {
	return i.Current().(*stripe.File)
}

// FileList returns the current list object which the iterator is
// currently using. List objects will change as new API calls are made to
// continue pagination.
func (i *Iter) FileList() *stripe.FileList {
	return i.List().(*stripe.FileList)
}

func getC() Client {
	return Client{stripe.GetBackend(stripe.APIBackend), stripe.GetBackend(stripe.UploadsBackend), stripe.Key}
}
