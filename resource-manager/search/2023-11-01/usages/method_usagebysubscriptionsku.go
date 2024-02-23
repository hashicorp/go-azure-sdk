package usages

import (
	"context"
	"fmt"
	"net/http"

	"github.com/hashicorp/go-azure-sdk/sdk/client"
	"github.com/hashicorp/go-azure-sdk/sdk/odata"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type UsageBySubscriptionSkuOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
	Model        *QuotaUsageResult
}

type UsageBySubscriptionSkuOperationOptions struct {
	XMsClientRequestId *string
}

func DefaultUsageBySubscriptionSkuOperationOptions() UsageBySubscriptionSkuOperationOptions {
	return UsageBySubscriptionSkuOperationOptions{}
}

func (o UsageBySubscriptionSkuOperationOptions) ToHeaders() *client.Headers {
	out := client.Headers{}
	if o.XMsClientRequestId != nil {
		out.Append("x-ms-client-request-id", fmt.Sprintf("%v", *o.XMsClientRequestId))
	}
	return &out
}

func (o UsageBySubscriptionSkuOperationOptions) ToOData() *odata.Query {
	out := odata.Query{}
	return &out
}

func (o UsageBySubscriptionSkuOperationOptions) ToQuery() *client.QueryParams {
	out := client.QueryParams{}

	return &out
}

// UsageBySubscriptionSku ...
func (c UsagesClient) UsageBySubscriptionSku(ctx context.Context, id UsageId, options UsageBySubscriptionSkuOperationOptions) (result UsageBySubscriptionSkuOperationResponse, err error) {
	opts := client.RequestOptions{
		ContentType: "application/json; charset=utf-8",
		ExpectedStatusCodes: []int{
			http.StatusOK,
		},
		HttpMethod:    http.MethodGet,
		Path:          id.ID(),
		OptionsObject: options,
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

	var model QuotaUsageResult
	result.Model = &model

	if err = resp.Unmarshal(result.Model); err != nil {
		return
	}

	return
}
