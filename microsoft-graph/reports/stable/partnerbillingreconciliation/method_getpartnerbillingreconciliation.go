package partnerbillingreconciliation

import (
	"context"
	"net/http"

	"github.com/hashicorp/go-azure-sdk/microsoft-graph/common-types/stable"
	"github.com/hashicorp/go-azure-sdk/sdk/client"
	"github.com/hashicorp/go-azure-sdk/sdk/odata"
)

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type GetPartnerBillingReconciliationOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
	Model        *stable.PartnersBillingBillingReconciliation
}

type GetPartnerBillingReconciliationOperationOptions struct {
	Expand    *odata.Expand
	Metadata  *odata.Metadata
	RetryFunc client.RequestRetryFunc
	Select    *[]string
}

func DefaultGetPartnerBillingReconciliationOperationOptions() GetPartnerBillingReconciliationOperationOptions {
	return GetPartnerBillingReconciliationOperationOptions{}
}

func (o GetPartnerBillingReconciliationOperationOptions) ToHeaders() *client.Headers {
	out := client.Headers{}

	return &out
}

func (o GetPartnerBillingReconciliationOperationOptions) ToOData() *odata.Query {
	out := odata.Query{}
	if o.Expand != nil {
		out.Expand = *o.Expand
	}
	if o.Metadata != nil {
		out.Metadata = *o.Metadata
	}
	if o.Select != nil {
		out.Select = *o.Select
	}
	return &out
}

func (o GetPartnerBillingReconciliationOperationOptions) ToQuery() *client.QueryParams {
	out := client.QueryParams{}

	return &out
}

// GetPartnerBillingReconciliation - Get reconciliation from reports. Represents details for billed invoice
// reconciliation data.
func (c PartnerBillingReconciliationClient) GetPartnerBillingReconciliation(ctx context.Context, options GetPartnerBillingReconciliationOperationOptions) (result GetPartnerBillingReconciliationOperationResponse, err error) {
	opts := client.RequestOptions{
		ContentType: "application/json; charset=utf-8",
		ExpectedStatusCodes: []int{
			http.StatusOK,
		},
		HttpMethod:    http.MethodGet,
		OptionsObject: options,
		Path:          "/reports/partners/billing/reconciliation",
		RetryFunc:     options.RetryFunc,
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

	var model stable.PartnersBillingBillingReconciliation
	result.Model = &model
	if err = resp.Unmarshal(result.Model); err != nil {
		return
	}

	return
}
