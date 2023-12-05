//
//
// File generated from our OpenAPI spec
//
//

// Package accountsession provides the /account_sessions APIs
package accountsession

import (
	"net/http"

	stripe "github.com/stripe/stripe-go/v76"
)

// Client is used to invoke /account_sessions APIs.
type Client struct {
	B   stripe.Backend
	Key string
}

// New creates a new account session.
func New(params *stripe.AccountSessionParams) (*stripe.AccountSession, error) {
	return getC().New(params)
}

// New creates a new account session.
func (c Client) New(params *stripe.AccountSessionParams) (*stripe.AccountSession, error) {
	accountsession := &stripe.AccountSession{}
	var err error
	sr := stripe.NewStripeRequest(http.MethodPost, "/v1/account_sessions", c.Key)
	err = sr.SetParams(params)
	if err != nil {
		return nil, err
	}
	err = c.B.Call(sr, accountsession)
	return accountsession, err
}

func getC() Client {
	return Client{stripe.GetBackend(stripe.APIBackend), stripe.Key}
}
