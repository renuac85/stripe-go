//
//
// File generated from our OpenAPI spec
//
//

// Package transaction provides the /issuing/transactions APIs
package transaction

import (
	"net/http"

	stripe "github.com/stripe/stripe-go/v76"
	"github.com/stripe/stripe-go/v76/form"
)

// Client is used to invoke /issuing/transactions APIs.
type Client struct {
	B   stripe.Backend
	Key string
}

// Get returns the details of an issuing transaction.
func Get(id string, params *stripe.IssuingTransactionParams) (*stripe.IssuingTransaction, error) {
	return getC().Get(id, params)
}

// Get returns the details of an issuing transaction.
func (c Client) Get(id string, params *stripe.IssuingTransactionParams) (*stripe.IssuingTransaction, error) {
	path := stripe.FormatURLPath("/v1/issuing/transactions/%s", id)
	transaction := &stripe.IssuingTransaction{}
	var err error
	sr := stripe.NewStripeRequest(http.MethodGet, path, c.Key)
	err = sr.SetParams(params)
	if err != nil {
		return nil, err
	}
	err = c.B.Call(sr, transaction)
	return transaction, err
}

// Update updates an issuing transaction's properties.
func Update(id string, params *stripe.IssuingTransactionParams) (*stripe.IssuingTransaction, error) {
	return getC().Update(id, params)
}

// Update updates an issuing transaction's properties.
func (c Client) Update(id string, params *stripe.IssuingTransactionParams) (*stripe.IssuingTransaction, error) {
	path := stripe.FormatURLPath("/v1/issuing/transactions/%s", id)
	transaction := &stripe.IssuingTransaction{}
	var err error
	sr := stripe.NewStripeRequest(http.MethodPost, path, c.Key)
	err = sr.SetParams(params)
	if err != nil {
		return nil, err
	}
	err = c.B.Call(sr, transaction)
	return transaction, err
}

// List returns a list of issuing transactions.
func List(params *stripe.IssuingTransactionListParams) *Iter {
	return getC().List(params)
}

// List returns a list of issuing transactions.
func (c Client) List(listParams *stripe.IssuingTransactionListParams) *Iter {
	return &Iter{
		Iter: stripe.GetIter(listParams, func(p *stripe.Params, b *form.Values) ([]interface{}, stripe.ListContainer, error) {
			list := &stripe.IssuingTransactionList{}
			sr := stripe.NewStripeRequest(
				http.MethodGet,
				"/v1/issuing/transactions",
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

// Iter is an iterator for issuing transactions.
type Iter struct {
	*stripe.Iter
}

// IssuingTransaction returns the issuing transaction which the iterator is currently pointing to.
func (i *Iter) IssuingTransaction() *stripe.IssuingTransaction {
	return i.Current().(*stripe.IssuingTransaction)
}

// IssuingTransactionList returns the current list object which the iterator is
// currently using. List objects will change as new API calls are made to
// continue pagination.
func (i *Iter) IssuingTransactionList() *stripe.IssuingTransactionList {
	return i.List().(*stripe.IssuingTransactionList)
}

func getC() Client {
	return Client{stripe.GetBackend(stripe.APIBackend), stripe.Key}
}
