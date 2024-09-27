package partnerbillingusageunbilled

import (
	"context"
	"fmt"
	"net/http"

	"github.com/hashicorp/go-azure-sdk/sdk/client"
	"github.com/hashicorp/go-azure-sdk/sdk/odata"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type DeletePartnerBillingUsageUnbilledOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
}

type DeletePartnerBillingUsageUnbilledOperationOptions struct {
	IfMatch   *string
	Metadata  *odata.Metadata
	RetryFunc client.RequestRetryFunc
}

func DefaultDeletePartnerBillingUsageUnbilledOperationOptions() DeletePartnerBillingUsageUnbilledOperationOptions {
	return DeletePartnerBillingUsageUnbilledOperationOptions{}
}

func (o DeletePartnerBillingUsageUnbilledOperationOptions) ToHeaders() *client.Headers {
	out := client.Headers{}
	if o.IfMatch != nil {
		out.Append("If-Match", fmt.Sprintf("%v", *o.IfMatch))
	}
	return &out
}

func (o DeletePartnerBillingUsageUnbilledOperationOptions) ToOData() *odata.Query {
	out := odata.Query{}
	if o.Metadata != nil {
		out.Metadata = *o.Metadata
	}
	return &out
}

func (o DeletePartnerBillingUsageUnbilledOperationOptions) ToQuery() *client.QueryParams {
	out := client.QueryParams{}

	return &out
}

// DeletePartnerBillingUsageUnbilled - Delete navigation property unbilled for reports
func (c PartnerBillingUsageUnbilledClient) DeletePartnerBillingUsageUnbilled(ctx context.Context, options DeletePartnerBillingUsageUnbilledOperationOptions) (result DeletePartnerBillingUsageUnbilledOperationResponse, err error) {
	opts := client.RequestOptions{
		ContentType: "application/json; charset=utf-8",
		ExpectedStatusCodes: []int{
			http.StatusNoContent,
			http.StatusOK,
		},
		HttpMethod:    http.MethodDelete,
		OptionsObject: options,
		Path:          "/reports/partners/billing/usage/unbilled",
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

	return
}
