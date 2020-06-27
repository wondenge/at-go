package atgo

import (
	"context"
	"fmt"
	"github.com/wondenge/at-go/internal/pkg/gen/africastalking"
)

const (
	CardCheckoutLiveURL    = "https://payments.africastalking.com/card/checkout/charge"
	CardCheckoutSandboxURL = "https://payments.sandbox.africastalking.com/card/checkout/charge"

	CardCheckoutValidateLiveURL    = "https://payments.africastalking.com/card/checkout/validate"
	CardCheckoutValidateSandboxURL = "https://payments.sandbox.africastalking.com/card/checkout/validate"
)

// Collect money into your Payment Wallet by initiating transactions that deduct
// money from a customers Debit or Credit Card.
// These APIs are currently only available in Nigeria on MasterCard and Verve cards.
func (c *Client) cardCheckout(ctx context.Context, p *africastalking.CardCheckoutPayload) (res *africastalking.CardCheckoutResponse, err error) {

	req, err := c.NewRequest("POST", "https://payments.sandbox.africastalking.com/card/checkout/charge", p)
	if err != nil {
		return nil, fmt.Errorf("could not make new http request: %w", err)
	}

	// Set Header Parameters.
	req.Header.Set("Accept", "application/json")
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Apikey", "MyAppAPIKey")

	res = &africastalking.CardCheckoutResponse{}
	if err := c.sendRequest(ctx, req, res); err != nil {
		return nil, err
	}

	return res, nil
}

// Allows your application to validate card checkout charge requests.
func (c *Client) cardCheckoutValidate(ctx context.Context, p *africastalking.CardCheckoutValidatePayload) (res *africastalking.CardCheckoutValidateResponse, err error) {

	req, err := c.NewRequest("POST", "https://payments.sandbox.africastalking.com/card/checkout/validate", p)
	if err != nil {
		return nil, fmt.Errorf("could not make new http request: %w", err)
	}

	// Set Header Parameters.
	req.Header.Set("Accept", "application/json")
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Apikey", "MyAppAPIKey")

	res = &africastalking.CardCheckoutValidateResponse{}
	if err := c.sendRequest(ctx, req, res); err != nil {
		return nil, err
	}

	return res, nil
}