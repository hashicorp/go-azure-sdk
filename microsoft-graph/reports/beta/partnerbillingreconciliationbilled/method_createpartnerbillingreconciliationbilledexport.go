package partnerbillingreconciliationbilled

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

type CreatePartnerBillingReconciliationBilledExportOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
	Model        beta.PartnersBillingOperation
}

type CreatePartnerBillingReconciliationBilledExportOperationOptions struct {
	Metadata  *odata.Metadata
	RetryFunc client.RequestRetryFunc
}

func DefaultCreatePartnerBillingReconciliationBilledExportOperationOptions() CreatePartnerBillingReconciliationBilledExportOperationOptions {
	return CreatePartnerBillingReconciliationBilledExportOperationOptions{}
}

func (o CreatePartnerBillingReconciliationBilledExportOperationOptions) ToHeaders() *client.Headers {
	out := client.Headers{}

	return &out
}

func (o CreatePartnerBillingReconciliationBilledExportOperationOptions) ToOData() *odata.Query {
	out := odata.Query{}
	if o.Metadata != nil {
		out.Metadata = *o.Metadata
	}
	return &out
}

func (o CreatePartnerBillingReconciliationBilledExportOperationOptions) ToQuery() *client.QueryParams {
	out := client.QueryParams{}

	return &out
}

// CreatePartnerBillingReconciliationBilledExport - Invoke action export. Export the billed invoice reconciliation data.
func (c PartnerBillingReconciliationBilledClient) CreatePartnerBillingReconciliationBilledExport(ctx context.Context, input CreatePartnerBillingReconciliationBilledExportRequest, options CreatePartnerBillingReconciliationBilledExportOperationOptions) (result CreatePartnerBillingReconciliationBilledExportOperationResponse, err error) {
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
		Path:          "/reports/partners/billing/reconciliation/billed/partners.billing.export",
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
