package exchangeonpremisespolicy

import (
	"context"
	"net/http"

	"github.com/hashicorp/go-azure-sdk/microsoft-graph/common-types/beta"
	"github.com/hashicorp/go-azure-sdk/sdk/client"
	"github.com/hashicorp/go-azure-sdk/sdk/odata"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type GetExchangeOnPremisesPolicyOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
	Model        *beta.DeviceManagementExchangeOnPremisesPolicy
}

type GetExchangeOnPremisesPolicyOperationOptions struct {
	Expand *odata.Expand
	Select *[]string
}

func DefaultGetExchangeOnPremisesPolicyOperationOptions() GetExchangeOnPremisesPolicyOperationOptions {
	return GetExchangeOnPremisesPolicyOperationOptions{}
}

func (o GetExchangeOnPremisesPolicyOperationOptions) ToHeaders() *client.Headers {
	out := client.Headers{}

	return &out
}

func (o GetExchangeOnPremisesPolicyOperationOptions) ToOData() *odata.Query {
	out := odata.Query{}
	if o.Expand != nil {
		out.Expand = *o.Expand
	}
	if o.Select != nil {
		out.Select = *o.Select
	}
	return &out
}

func (o GetExchangeOnPremisesPolicyOperationOptions) ToQuery() *client.QueryParams {
	out := client.QueryParams{}

	return &out
}

// GetExchangeOnPremisesPolicy - Get exchangeOnPremisesPolicy from deviceManagement. The policy which controls mobile
// device access to Exchange On Premises
func (c ExchangeOnPremisesPolicyClient) GetExchangeOnPremisesPolicy(ctx context.Context, options GetExchangeOnPremisesPolicyOperationOptions) (result GetExchangeOnPremisesPolicyOperationResponse, err error) {
	opts := client.RequestOptions{
		ContentType: "application/json; charset=utf-8",
		ExpectedStatusCodes: []int{
			http.StatusOK,
		},
		HttpMethod:    http.MethodGet,
		OptionsObject: options,
		Path:          "/deviceManagement/exchangeOnPremisesPolicy",
	}

	req, err := c.Client.NewRequest(ctx, opts)
	if err != nil {
		return
	}

	var resp *client.Response
	resp, err = req.Execute(ctx)
	if resp != nil {
		result.OData = resp.OData
		result.HttpResponse = resp.Response
	}
	if err != nil {
		return
	}

	var model beta.DeviceManagementExchangeOnPremisesPolicy
	result.Model = &model
	if err = resp.Unmarshal(result.Model); err != nil {
		return
	}

	return
}
