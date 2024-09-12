package partnerbillingusagebilled

import (
	"context"
	"net/http"

	"github.com/hashicorp/go-azure-sdk/microsoft-graph/common-types/beta"
	"github.com/hashicorp/go-azure-sdk/sdk/client"
	"github.com/hashicorp/go-azure-sdk/sdk/odata"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type GetPartnerBillingUsageBilledOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
	Model        *beta.PartnersBillingBilledUsage
}

type GetPartnerBillingUsageBilledOperationOptions struct {
	Expand *odata.Expand
	Select *[]string
}

func DefaultGetPartnerBillingUsageBilledOperationOptions() GetPartnerBillingUsageBilledOperationOptions {
	return GetPartnerBillingUsageBilledOperationOptions{}
}

func (o GetPartnerBillingUsageBilledOperationOptions) ToHeaders() *client.Headers {
	out := client.Headers{}

	return &out
}

func (o GetPartnerBillingUsageBilledOperationOptions) ToOData() *odata.Query {
	out := odata.Query{}
	if o.Expand != nil {
		out.Expand = *o.Expand
	}
	if o.Select != nil {
		out.Select = *o.Select
	}
	return &out
}

func (o GetPartnerBillingUsageBilledOperationOptions) ToQuery() *client.QueryParams {
	out := client.QueryParams{}

	return &out
}

// GetPartnerBillingUsageBilled - Get billed from reports. Represents details for billed Azure usage data.
func (c PartnerBillingUsageBilledClient) GetPartnerBillingUsageBilled(ctx context.Context, options GetPartnerBillingUsageBilledOperationOptions) (result GetPartnerBillingUsageBilledOperationResponse, err error) {
	opts := client.RequestOptions{
		ContentType: "application/json; charset=utf-8",
		ExpectedStatusCodes: []int{
			http.StatusOK,
		},
		HttpMethod:    http.MethodGet,
		OptionsObject: options,
		Path:          "/reports/partners/billing/usage/billed",
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

	var model beta.PartnersBillingBilledUsage
	result.Model = &model
	if err = resp.Unmarshal(result.Model); err != nil {
		return
	}

	return
}
