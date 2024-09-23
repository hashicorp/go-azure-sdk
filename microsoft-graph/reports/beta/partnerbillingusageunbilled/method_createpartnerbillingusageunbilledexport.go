package partnerbillingusageunbilled

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

type CreatePartnerBillingUsageUnbilledExportOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
	Model        beta.PartnersBillingOperation
}

type CreatePartnerBillingUsageUnbilledExportOperationOptions struct {
	Metadata  *odata.Metadata
	RetryFunc client.RequestRetryFunc
}

func DefaultCreatePartnerBillingUsageUnbilledExportOperationOptions() CreatePartnerBillingUsageUnbilledExportOperationOptions {
	return CreatePartnerBillingUsageUnbilledExportOperationOptions{}
}

func (o CreatePartnerBillingUsageUnbilledExportOperationOptions) ToHeaders() *client.Headers {
	out := client.Headers{}

	return &out
}

func (o CreatePartnerBillingUsageUnbilledExportOperationOptions) ToOData() *odata.Query {
	out := odata.Query{}
	if o.Metadata != nil {
		out.Metadata = *o.Metadata
	}
	return &out
}

func (o CreatePartnerBillingUsageUnbilledExportOperationOptions) ToQuery() *client.QueryParams {
	out := client.QueryParams{}

	return &out
}

// CreatePartnerBillingUsageUnbilledExport - Invoke action export. Export unbilled Azure usage data for a specific
// billing period and currency.
func (c PartnerBillingUsageUnbilledClient) CreatePartnerBillingUsageUnbilledExport(ctx context.Context, input CreatePartnerBillingUsageUnbilledExportRequest, options CreatePartnerBillingUsageUnbilledExportOperationOptions) (result CreatePartnerBillingUsageUnbilledExportOperationResponse, err error) {
	opts := client.RequestOptions{
		ContentType: "application/json; charset=utf-8",
		ExpectedStatusCodes: []int{
			http.StatusOK,
		},
		HttpMethod:    http.MethodPost,
		OptionsObject: options,
		Path:          "/reports/partners/billing/usage/unbilled/partners.billing.export",
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
