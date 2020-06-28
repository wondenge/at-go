package atgo

import (
	"context"
	"fmt"
	at "github.com/wondenge/at-go/internal/pkg/gen/africastalking"
)

type (
	Card interface {
		// Collect money into your Payment Wallet by initiating transactions that deduct money from a customers Debit or Credit Card.
		// These APIs are currently only available in Nigeria on MasterCard and Verve cards.
		cardCheckout(ctx context.Context, p *at.CardCheckoutPayload) (res *at.CardCheckoutResponse, err error)

		// Allows your application to validate card checkout charge requests.
		cardCheckoutValidate(ctx context.Context, p *at.CardCheckoutValidatePayload) (res *at.CardCheckoutValidateResponse, err error)
	}
)

// Collect money into your Payment Wallet by initiating transactions that deduct
// money from a customers Debit or Credit Card.
// These APIs are currently only available in Nigeria on MasterCard and Verve cards.
func (c *ATClient) cardCheckout(ctx context.Context, p *at.CardCheckoutPayload) (res *at.CardCheckoutResponse, err error) {

	req, err := c.newRequest("POST", fmt.Sprintf("%s%s", c.PaymentEndpoint, "/card/checkout/charge"), p)
	if err != nil {
		return nil, fmt.Errorf("could not make new http request: %w", err)
	}

	// Set Header Parameters.
	req.Header.Set("Accept", "application/json")
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Apikey", c.APIKey)

	res = &at.CardCheckoutResponse{}
	if err := c.sendRequest(ctx, req, res); err != nil {
		return nil, err
	}

	return res, nil
}

// Allows your application to validate card checkout charge requests.
func (c *ATClient) cardCheckoutValidate(ctx context.Context, p *at.CardCheckoutValidatePayload) (res *at.CardCheckoutValidateResponse, err error) {

	req, err := c.newRequest("POST", fmt.Sprintf("%s%s", c.PaymentEndpoint, "/card/checkout/validate"), p)
	if err != nil {
		return nil, fmt.Errorf("could not make new http request: %w", err)
	}

	// Set Header Parameters.
	req.Header.Set("Accept", "application/json")
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Apikey", c.APIKey)

	res = &at.CardCheckoutValidateResponse{}
	if err := c.sendRequest(ctx, req, res); err != nil {
		return nil, err
	}

	return res, nil
}
