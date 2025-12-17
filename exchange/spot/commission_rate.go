package spot

import (
	"context"

	"github.com/UnipayFI/go-aster/spot"
)

func (c *Client) GetCommissionRate(symbol string) (*CommissionRate, error) {
	rate, err := c.NewSpotClient().NewCommissionRateService(symbol).Do(context.Background())
	if err != nil {
		return nil, err
	}
	return &CommissionRate{CommissionRateResponse: *rate}, nil
}

type CommissionRate struct {
	spot.CommissionRateResponse
}
