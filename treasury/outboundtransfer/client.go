//
//
// File generated from our OpenAPI spec
//
//

// Package outboundtransfer provides the /treasury/outbound_transfers APIs
package outboundtransfer

import (
	"net/http"

	stripe "github.com/stripe/stripe-go/v76"
	"github.com/stripe/stripe-go/v76/form"
)

// Client is used to invoke /treasury/outbound_transfers APIs.
type Client struct {
	B   stripe.Backend
	Key string
}

// New creates a new treasury outbound transfer.
func New(params *stripe.TreasuryOutboundTransferParams) (*stripe.TreasuryOutboundTransfer, error) {
	return getC().New(params)
}

// New creates a new treasury outbound transfer.
func (c Client) New(params *stripe.TreasuryOutboundTransferParams) (*stripe.TreasuryOutboundTransfer, error) {
	outboundtransfer := &stripe.TreasuryOutboundTransfer{}
	var err error
	sr := stripe.NewStripeRequest(
		http.MethodPost,
		"/v1/treasury/outbound_transfers",
		c.Key,
	)
	err = sr.SetParams(params)
	if err != nil {
		return nil, err
	}
	err = c.B.Call(sr, outboundtransfer)
	return outboundtransfer, err
}

// Get returns the details of a treasury outbound transfer.
func Get(id string, params *stripe.TreasuryOutboundTransferParams) (*stripe.TreasuryOutboundTransfer, error) {
	return getC().Get(id, params)
}

// Get returns the details of a treasury outbound transfer.
func (c Client) Get(id string, params *stripe.TreasuryOutboundTransferParams) (*stripe.TreasuryOutboundTransfer, error) {
	path := stripe.FormatURLPath("/v1/treasury/outbound_transfers/%s", id)
	outboundtransfer := &stripe.TreasuryOutboundTransfer{}
	var err error
	sr := stripe.NewStripeRequest(http.MethodGet, path, c.Key)
	err = sr.SetParams(params)
	if err != nil {
		return nil, err
	}
	err = c.B.Call(sr, outboundtransfer)
	return outboundtransfer, err
}

// Cancel is the method for the `POST /v1/treasury/outbound_transfers/{outbound_transfer}/cancel` API.
func Cancel(id string, params *stripe.TreasuryOutboundTransferCancelParams) (*stripe.TreasuryOutboundTransfer, error) {
	return getC().Cancel(id, params)
}

// Cancel is the method for the `POST /v1/treasury/outbound_transfers/{outbound_transfer}/cancel` API.
func (c Client) Cancel(id string, params *stripe.TreasuryOutboundTransferCancelParams) (*stripe.TreasuryOutboundTransfer, error) {
	path := stripe.FormatURLPath("/v1/treasury/outbound_transfers/%s/cancel", id)
	outboundtransfer := &stripe.TreasuryOutboundTransfer{}
	var err error
	sr := stripe.NewStripeRequest(http.MethodPost, path, c.Key)
	err = sr.SetParams(params)
	if err != nil {
		return nil, err
	}
	err = c.B.Call(sr, outboundtransfer)
	return outboundtransfer, err
}

// List returns a list of treasury outbound transfers.
func List(params *stripe.TreasuryOutboundTransferListParams) *Iter {
	return getC().List(params)
}

// List returns a list of treasury outbound transfers.
func (c Client) List(listParams *stripe.TreasuryOutboundTransferListParams) *Iter {
	return &Iter{
		Iter: stripe.GetIter(listParams, func(p *stripe.Params, b *form.Values) ([]interface{}, stripe.ListContainer, error) {
			list := &stripe.TreasuryOutboundTransferList{}
			sr := stripe.NewStripeRequest(
				http.MethodGet,
				"/v1/treasury/outbound_transfers",
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

// Iter is an iterator for treasury outbound transfers.
type Iter struct {
	*stripe.Iter
}

// TreasuryOutboundTransfer returns the treasury outbound transfer which the iterator is currently pointing to.
func (i *Iter) TreasuryOutboundTransfer() *stripe.TreasuryOutboundTransfer {
	return i.Current().(*stripe.TreasuryOutboundTransfer)
}

// TreasuryOutboundTransferList returns the current list object which the iterator is
// currently using. List objects will change as new API calls are made to
// continue pagination.
func (i *Iter) TreasuryOutboundTransferList() *stripe.TreasuryOutboundTransferList {
	return i.List().(*stripe.TreasuryOutboundTransferList)
}

func getC() Client {
	return Client{stripe.GetBackend(stripe.APIBackend), stripe.Key}
}
