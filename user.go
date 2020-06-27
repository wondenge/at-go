package atgo

import (
	"context"
	"fmt"
	at "github.com/wondenge/at-go/internal/pkg/gen/africastalking"
)

const (
	UserLiveURL    = "https://api.africastalking.com/version1/user"
	UserSandboxURL = "https://api.sandbox.africastalking.com/version1/user"
)

type (
	User interface {
		// Initiate an application data request.
		initiateAppData(ctx context.Context, p string) (res *at.UserResponse, err error)
	}
)

// Initiate an application data request.
func (c *Client) initiateAppData(ctx context.Context, p string) (res *at.UserResponse, err error) {

	req, err := c.NewRequest("GET", fmt.Sprintf("%s%s", "", ""), p)
	if err != nil {
		return nil, fmt.Errorf("could not make new http request: %w", err)
	}

	// Set Header Parameters.
	req.Header.Set("Accept", "application/json")
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")
	req.Header.Set("Apikey", "MyAppApiKey")

	res = &at.UserResponse{}
	if err := c.sendRequest(ctx, req, res); err != nil {
		return nil, err
	}

	return res, nil
}
