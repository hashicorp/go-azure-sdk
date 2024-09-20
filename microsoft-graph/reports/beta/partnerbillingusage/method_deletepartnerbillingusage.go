package partnerbillingusage

import (
	"context"
	"fmt"
	"net/http"

	"github.com/hashicorp/go-azure-sdk/sdk/client"
	"github.com/hashicorp/go-azure-sdk/sdk/odata"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type DeletePartnerBillingUsageOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
}

type DeletePartnerBillingUsageOperationOptions struct {
	IfMatch  *string
	Metadata *odata.Metadata
}

func DefaultDeletePartnerBillingUsageOperationOptions() DeletePartnerBillingUsageOperationOptions {
	return DeletePartnerBillingUsageOperationOptions{}
}

func (o DeletePartnerBillingUsageOperationOptions) ToHeaders() *client.Headers {
	out := client.Headers{}
	if o.IfMatch != nil {
		out.Append("If-Match", fmt.Sprintf("%v", *o.IfMatch))
	}
	return &out
}

func (o DeletePartnerBillingUsageOperationOptions) ToOData() *odata.Query {
	out := odata.Query{}
	if o.Metadata != nil {
		out.Metadata = *o.Metadata
	}
	return &out
}

func (o DeletePartnerBillingUsageOperationOptions) ToQuery() *client.QueryParams {
	out := client.QueryParams{}

	return &out
}

// DeletePartnerBillingUsage - Delete navigation property usage for reports
func (c PartnerBillingUsageClient) DeletePartnerBillingUsage(ctx context.Context, options DeletePartnerBillingUsageOperationOptions) (result DeletePartnerBillingUsageOperationResponse, err error) {
	opts := client.RequestOptions{
		ContentType: "application/json; charset=utf-8",
		ExpectedStatusCodes: []int{
			http.StatusNoContent,
		},
		HttpMethod:    http.MethodDelete,
		OptionsObject: options,
		Path:          "/reports/partners/billing/usage",
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
