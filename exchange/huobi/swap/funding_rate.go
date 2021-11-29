package swap

import (
	"context"
	"net/http"

	"github.com/NadiaSama/ccexgo/exchange"
	"github.com/pkg/errors"
	"github.com/shopspring/decimal"
)

type (
	FundingRateResp struct {
		EstimmatedRate  decimal.Decimal `json:"estimated_rate"`
		FundingRate     decimal.Decimal `json:"funding_rate"`
		ContractCode    string          `json:"contract_code"`
		Symbol          string          `json:"symbol"`
		FeeAsset        string          `json:"fee_asset"`
		FundingTime     string          `json:"funding_time"`
		NextFundingTime string          `json:"next_funding_time"`
	}

	FundingRateReq struct {
		*exchange.RestReq
	}
)

const (
	FundingRateEndPoint = "/swap-api/v1/swap_funding_rate"
)

func NewFundingRateReq(cc string) *FundingRateReq {
	r := exchange.NewRestReq()
	return &FundingRateReq{
		RestReq: r.AddFields("contract_code", cc),
	}
}

func (rc *RestClient) SwapFundingRate(ctx context.Context, req *FundingRateReq) (*FundingRateResp, error) {
	var resp FundingRateResp

	param, err := req.Values()
	if err != nil {
		return nil, errors.WithMessage(err, "build values fail")
	}
	if err := rc.Request(ctx, http.MethodGet, FundingRateEndPoint, param, nil, false, &resp); err != nil {
		return nil, errors.WithMessage(err, "request funding fail")
	}

	return &resp, nil
}
