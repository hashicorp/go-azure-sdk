package partnerbillingusagebilled

import (
	"context"
	"encoding/json"
	"net/http"

	"github.com/hashicorp/go-azure-sdk/microsoft-graph/common-types/stable"
	"github.com/hashicorp/go-azure-sdk/sdk/client"
	"github.com/hashicorp/go-azure-sdk/sdk/odata"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type CreatePartnerBillingUsageBilledExportOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
	Model        stable.PartnersBillingOperation
}

type CreatePartnerBillingUsageBilledExportOperationOptions struct {
	Metadata  *odata.Metadata
	RetryFunc client.RequestRetryFunc
}

func DefaultCreatePartnerBillingUsageBilledExportOperationOptions() CreatePartnerBillingUsageBilledExportOperationOptions {
	return CreatePartnerBillingUsageBilledExportOperationOptions{}
}

func (o CreatePartnerBillingUsageBilledExportOperationOptions) ToHeaders() *client.Headers {
	out := client.Headers{}

	return &out
}

func (o CreatePartnerBillingUsageBilledExportOperationOptions) ToOData() *odata.Query {
	out := odata.Query{}
	if o.Metadata != nil {
		out.Metadata = *o.Metadata
	}
	return &out
}

func (o CreatePartnerBillingUsageBilledExportOperationOptions) ToQuery() *client.QueryParams {
	out := client.QueryParams{}

	return &out
}

// CreatePartnerBillingUsageBilledExport - Invoke action export. Export the billed Azure usage data.
func (c PartnerBillingUsageBilledClient) CreatePartnerBillingUsageBilledExport(ctx context.Context, input CreatePartnerBillingUsageBilledExportRequest, options CreatePartnerBillingUsageBilledExportOperationOptions) (result CreatePartnerBillingUsageBilledExportOperationResponse, err error) {
	opts := client.RequestOptions{
		ContentType: "application/json; charset=utf-8",
		ExpectedStatusCodes: []int{
			http.StatusOK,
		},
		HttpMethod:    http.MethodPost,
		OptionsObject: options,
		Path:          "/reports/partners/billing/usage/billed/partners.billing.export",
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
	model, err := stable.UnmarshalPartnersBillingOperationImplementation(respObj)
	if err != nil {
		return
	}
	result.Model = model

	return
}
