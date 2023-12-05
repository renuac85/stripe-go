//
//
// File generated from our OpenAPI spec
//
//

// Package financialaccount provides the /treasury/financial_accounts APIs
package financialaccount

import (
	"net/http"

	stripe "github.com/stripe/stripe-go/v76"
	"github.com/stripe/stripe-go/v76/form"
)

// Client is used to invoke /treasury/financial_accounts APIs.
type Client struct {
	B   stripe.Backend
	Key string
}

// New creates a new treasury financial account.
func New(params *stripe.TreasuryFinancialAccountParams) (*stripe.TreasuryFinancialAccount, error) {
	return getC().New(params)
}

// New creates a new treasury financial account.
func (c Client) New(params *stripe.TreasuryFinancialAccountParams) (*stripe.TreasuryFinancialAccount, error) {
	financialaccount := &stripe.TreasuryFinancialAccount{}
	var err error
	sr := stripe.NewStripeRequest(
		http.MethodPost,
		"/v1/treasury/financial_accounts",
		c.Key,
	)
	err = sr.SetParams(params)
	if err != nil {
		return nil, err
	}
	err = c.B.Call(sr, financialaccount)
	return financialaccount, err
}

// Get returns the details of a treasury financial account.
func Get(id string, params *stripe.TreasuryFinancialAccountParams) (*stripe.TreasuryFinancialAccount, error) {
	return getC().Get(id, params)
}

// Get returns the details of a treasury financial account.
func (c Client) Get(id string, params *stripe.TreasuryFinancialAccountParams) (*stripe.TreasuryFinancialAccount, error) {
	path := stripe.FormatURLPath("/v1/treasury/financial_accounts/%s", id)
	financialaccount := &stripe.TreasuryFinancialAccount{}
	var err error
	sr := stripe.NewStripeRequest(http.MethodGet, path, c.Key)
	err = sr.SetParams(params)
	if err != nil {
		return nil, err
	}
	err = c.B.Call(sr, financialaccount)
	return financialaccount, err
}

// Update updates a treasury financial account's properties.
func Update(id string, params *stripe.TreasuryFinancialAccountParams) (*stripe.TreasuryFinancialAccount, error) {
	return getC().Update(id, params)
}

// Update updates a treasury financial account's properties.
func (c Client) Update(id string, params *stripe.TreasuryFinancialAccountParams) (*stripe.TreasuryFinancialAccount, error) {
	path := stripe.FormatURLPath("/v1/treasury/financial_accounts/%s", id)
	financialaccount := &stripe.TreasuryFinancialAccount{}
	var err error
	sr := stripe.NewStripeRequest(http.MethodPost, path, c.Key)
	err = sr.SetParams(params)
	if err != nil {
		return nil, err
	}
	err = c.B.Call(sr, financialaccount)
	return financialaccount, err
}

// RetrieveFeatures is the method for the `GET /v1/treasury/financial_accounts/{financial_account}/features` API.
func RetrieveFeatures(id string, params *stripe.TreasuryFinancialAccountRetrieveFeaturesParams) (*stripe.TreasuryFinancialAccountFeatures, error) {
	return getC().RetrieveFeatures(id, params)
}

// RetrieveFeatures is the method for the `GET /v1/treasury/financial_accounts/{financial_account}/features` API.
func (c Client) RetrieveFeatures(id string, params *stripe.TreasuryFinancialAccountRetrieveFeaturesParams) (*stripe.TreasuryFinancialAccountFeatures, error) {
	path := stripe.FormatURLPath(
		"/v1/treasury/financial_accounts/%s/features",
		id,
	)
	financialaccountfeatures := &stripe.TreasuryFinancialAccountFeatures{}
	var err error
	sr := stripe.NewStripeRequest(http.MethodGet, path, c.Key)
	err = sr.SetParams(params)
	if err != nil {
		return nil, err
	}
	err = c.B.Call(sr, financialaccountfeatures)
	return financialaccountfeatures, err
}

// UpdateFeatures is the method for the `POST /v1/treasury/financial_accounts/{financial_account}/features` API.
func UpdateFeatures(id string, params *stripe.TreasuryFinancialAccountUpdateFeaturesParams) (*stripe.TreasuryFinancialAccountFeatures, error) {
	return getC().UpdateFeatures(id, params)
}

// UpdateFeatures is the method for the `POST /v1/treasury/financial_accounts/{financial_account}/features` API.
func (c Client) UpdateFeatures(id string, params *stripe.TreasuryFinancialAccountUpdateFeaturesParams) (*stripe.TreasuryFinancialAccountFeatures, error) {
	path := stripe.FormatURLPath(
		"/v1/treasury/financial_accounts/%s/features",
		id,
	)
	financialaccountfeatures := &stripe.TreasuryFinancialAccountFeatures{}
	var err error
	sr := stripe.NewStripeRequest(http.MethodPost, path, c.Key)
	err = sr.SetParams(params)
	if err != nil {
		return nil, err
	}
	err = c.B.Call(sr, financialaccountfeatures)
	return financialaccountfeatures, err
}

// List returns a list of treasury financial accounts.
func List(params *stripe.TreasuryFinancialAccountListParams) *Iter {
	return getC().List(params)
}

// List returns a list of treasury financial accounts.
func (c Client) List(listParams *stripe.TreasuryFinancialAccountListParams) *Iter {
	return &Iter{
		Iter: stripe.GetIter(listParams, func(p *stripe.Params, b *form.Values) ([]interface{}, stripe.ListContainer, error) {
			list := &stripe.TreasuryFinancialAccountList{}
			sr := stripe.NewStripeRequest(
				http.MethodGet,
				"/v1/treasury/financial_accounts",
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

// Iter is an iterator for treasury financial accounts.
type Iter struct {
	*stripe.Iter
}

// TreasuryFinancialAccount returns the treasury financial account which the iterator is currently pointing to.
func (i *Iter) TreasuryFinancialAccount() *stripe.TreasuryFinancialAccount {
	return i.Current().(*stripe.TreasuryFinancialAccount)
}

// TreasuryFinancialAccountList returns the current list object which the iterator is
// currently using. List objects will change as new API calls are made to
// continue pagination.
func (i *Iter) TreasuryFinancialAccountList() *stripe.TreasuryFinancialAccountList {
	return i.List().(*stripe.TreasuryFinancialAccountList)
}

func getC() Client {
	return Client{stripe.GetBackend(stripe.APIBackend), stripe.Key}
}
