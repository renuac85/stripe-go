//
//
// File generated from our OpenAPI spec
//
//

// Package confirmationtoken provides the /confirmation_tokens APIs
package confirmationtoken

import (
	"net/http"

	stripe "github.com/stripe/stripe-go/v76"
)

// Client is used to invoke /confirmation_tokens APIs.
type Client struct {
	B   stripe.Backend
	Key string
}

// Get returns the details of a confirmation token.
func Get(id string, params *stripe.ConfirmationTokenParams) (*stripe.ConfirmationToken, error) {
	return getC().Get(id, params)
}

// Get returns the details of a confirmation token.
func (c Client) Get(id string, params *stripe.ConfirmationTokenParams) (*stripe.ConfirmationToken, error) {
	path := stripe.FormatURLPath("/v1/confirmation_tokens/%s", id)
	confirmationtoken := &stripe.ConfirmationToken{}
	var err error
	sr := stripe.NewStripeRequest(http.MethodGet, path, c.Key)
	err = sr.SetParams(params)
	if err != nil {
		return nil, err
	}
	err = c.B.Call(sr, confirmationtoken)
	return confirmationtoken, err
}

func getC() Client {
	return Client{stripe.GetBackend(stripe.APIBackend), stripe.Key}
}
