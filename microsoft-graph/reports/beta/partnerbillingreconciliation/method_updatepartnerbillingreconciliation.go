package partnerbillingreconciliation

import (
	"context"
	"net/http"

	"github.com/hashicorp/go-azure-sdk/microsoft-graph/common-types/beta"
	"github.com/hashicorp/go-azure-sdk/sdk/client"
	"github.com/hashicorp/go-azure-sdk/sdk/odata"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type UpdatePartnerBillingReconciliationOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
}

type UpdatePartnerBillingReconciliationOperationOptions struct {
	Metadata *odata.Metadata
}

func DefaultUpdatePartnerBillingReconciliationOperationOptions() UpdatePartnerBillingReconciliationOperationOptions {
	return UpdatePartnerBillingReconciliationOperationOptions{}
}

func (o UpdatePartnerBillingReconciliationOperationOptions) ToHeaders() *client.Headers {
	out := client.Headers{}

	return &out
}

func (o UpdatePartnerBillingReconciliationOperationOptions) ToOData() *odata.Query {
	out := odata.Query{}
	if o.Metadata != nil {
		out.Metadata = *o.Metadata
	}
	return &out
}

func (o UpdatePartnerBillingReconciliationOperationOptions) ToQuery() *client.QueryParams {
	out := client.QueryParams{}

	return &out
}

// UpdatePartnerBillingReconciliation - Update the navigation property reconciliation in reports
func (c PartnerBillingReconciliationClient) UpdatePartnerBillingReconciliation(ctx context.Context, input beta.PartnersBillingBillingReconciliation, options UpdatePartnerBillingReconciliationOperationOptions) (result UpdatePartnerBillingReconciliationOperationResponse, err error) {
	opts := client.RequestOptions{
		ContentType: "application/json; charset=utf-8",
		ExpectedStatusCodes: []int{
			http.StatusNoContent,
		},
		HttpMethod:    http.MethodPatch,
		OptionsObject: options,
		Path:          "/reports/partners/billing/reconciliation",
	}

	req, err := c.Client.NewRequest(ctx, opts)
	if err != nil {
		return
	}

	if err = req.Marshal(input); err != nil {
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

	return
}
