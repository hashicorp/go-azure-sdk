package partnerbillingreconciliationunbilled

import (
	"context"
	"encoding/json"
	"net/http"

	"github.com/hashicorp/go-azure-sdk/microsoft-graph/common-types/beta"
	"github.com/hashicorp/go-azure-sdk/sdk/client"
	"github.com/hashicorp/go-azure-sdk/sdk/odata"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type CreatePartnerBillingReconciliationUnbilledExportOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
	Model        beta.PartnersBillingOperation
}

type CreatePartnerBillingReconciliationUnbilledExportOperationOptions struct {
	Metadata  *odata.Metadata
	RetryFunc client.RequestRetryFunc
}

func DefaultCreatePartnerBillingReconciliationUnbilledExportOperationOptions() CreatePartnerBillingReconciliationUnbilledExportOperationOptions {
	return CreatePartnerBillingReconciliationUnbilledExportOperationOptions{}
}

func (o CreatePartnerBillingReconciliationUnbilledExportOperationOptions) ToHeaders() *client.Headers {
	out := client.Headers{}

	return &out
}

func (o CreatePartnerBillingReconciliationUnbilledExportOperationOptions) ToOData() *odata.Query {
	out := odata.Query{}
	if o.Metadata != nil {
		out.Metadata = *o.Metadata
	}
	return &out
}

func (o CreatePartnerBillingReconciliationUnbilledExportOperationOptions) ToQuery() *client.QueryParams {
	out := client.QueryParams{}

	return &out
}

// CreatePartnerBillingReconciliationUnbilledExport - Invoke action export. Export the unbilled invoice reconciliation
// data for a specific billing period and a given currency.
func (c PartnerBillingReconciliationUnbilledClient) CreatePartnerBillingReconciliationUnbilledExport(ctx context.Context, input CreatePartnerBillingReconciliationUnbilledExportRequest, options CreatePartnerBillingReconciliationUnbilledExportOperationOptions) (result CreatePartnerBillingReconciliationUnbilledExportOperationResponse, err error) {
	opts := client.RequestOptions{
		ContentType: "application/json; charset=utf-8",
		ExpectedStatusCodes: []int{
			http.StatusAccepted,
			http.StatusCreated,
			http.StatusNoContent,
			http.StatusOK,
		},
		HttpMethod:    http.MethodPost,
		OptionsObject: options,
		Path:          "/reports/partners/billing/reconciliation/unbilled/partners.billing.export",
		RetryFunc:     options.RetryFunc,
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

	var respObj json.RawMessage
	if err = resp.Unmarshal(&respObj); err != nil {
		return
	}
	model, err := beta.UnmarshalPartnersBillingOperationImplementation(respObj)
	if err != nil {
		return
	}
	result.Model = model

	return
}
