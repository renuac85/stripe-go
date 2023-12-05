//
//
// File generated from our OpenAPI spec
//
//

// Package customercashbalancetransaction provides the /customers/{customer}/cash_balance_transactions APIs
package customercashbalancetransaction

import (
	"net/http"

	stripe "github.com/stripe/stripe-go/v76"
	"github.com/stripe/stripe-go/v76/form"
)

// Client is used to invoke /customers/{customer}/cash_balance_transactions APIs.
type Client struct {
	B   stripe.Backend
	Key string
}

// Get returns the details of a customer cash balance transaction.
func Get(id string, params *stripe.CustomerCashBalanceTransactionParams) (*stripe.CustomerCashBalanceTransaction, error) {
	return getC().Get(id, params)
}

// Get returns the details of a customer cash balance transaction.
func (c Client) Get(id string, params *stripe.CustomerCashBalanceTransactionParams) (*stripe.CustomerCashBalanceTransaction, error) {
	path := stripe.FormatURLPath(
		"/v1/customers/%s/cash_balance_transactions/%s",
		stripe.StringValue(params.Customer),
		id,
	)
	customercashbalancetransaction := &stripe.CustomerCashBalanceTransaction{}
	var err error
	sr := stripe.NewStripeRequest(http.MethodGet, path, c.Key)
	err = sr.SetParams(params)
	if err != nil {
		return nil, err
	}
	err = c.B.Call(sr, customercashbalancetransaction)
	return customercashbalancetransaction, err
}

// List returns a list of customer cash balance transactions.
func List(params *stripe.CustomerCashBalanceTransactionListParams) *Iter {
	return getC().List(params)
}

// List returns a list of customer cash balance transactions.
func (c Client) List(listParams *stripe.CustomerCashBalanceTransactionListParams) *Iter {
	path := stripe.FormatURLPath(
		"/v1/customers/%s/cash_balance_transactions",
		stripe.StringValue(listParams.Customer),
	)
	return &Iter{
		Iter: stripe.GetIter(listParams, func(p *stripe.Params, b *form.Values) ([]interface{}, stripe.ListContainer, error) {
			list := &stripe.CustomerCashBalanceTransactionList{}
			sr := stripe.NewStripeRequest(
				http.MethodGet,
				path,
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

// Iter is an iterator for customer cash balance transactions.
type Iter struct {
	*stripe.Iter
}

// CustomerCashBalanceTransaction returns the customer cash balance transaction which the iterator is currently pointing to.
func (i *Iter) CustomerCashBalanceTransaction() *stripe.CustomerCashBalanceTransaction {
	return i.Current().(*stripe.CustomerCashBalanceTransaction)
}

// CustomerCashBalanceTransactionList returns the current list object which the iterator is
// currently using. List objects will change as new API calls are made to
// continue pagination.
func (i *Iter) CustomerCashBalanceTransactionList() *stripe.CustomerCashBalanceTransactionList {
	return i.List().(*stripe.CustomerCashBalanceTransactionList)
}

func getC() Client {
	return Client{stripe.GetBackend(stripe.APIBackend), stripe.Key}
}
