package dimensions

import (
	"context"
	"fmt"
	"net/http"

	"github.com/hashicorp/go-azure-sdk/sdk/client"
	"github.com/hashicorp/go-azure-sdk/sdk/odata"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type ByExternalCloudProviderTypeOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
	Model        *DimensionsListResult
}

type ByExternalCloudProviderTypeOperationOptions struct {
	Expand *string
	Filter *string
	Top    *int64
}

func DefaultByExternalCloudProviderTypeOperationOptions() ByExternalCloudProviderTypeOperationOptions {
	return ByExternalCloudProviderTypeOperationOptions{}
}

func (o ByExternalCloudProviderTypeOperationOptions) ToHeaders() *client.Headers {
	out := client.Headers{}

	return &out
}

func (o ByExternalCloudProviderTypeOperationOptions) ToOData() *odata.Query {
	out := odata.Query{}

	return &out
}

func (o ByExternalCloudProviderTypeOperationOptions) ToQuery() *client.QueryParams {
	out := client.QueryParams{}
	if o.Expand != nil {
		out.Append("$expand", fmt.Sprintf("%v", *o.Expand))
	}
	if o.Filter != nil {
		out.Append("$filter", fmt.Sprintf("%v", *o.Filter))
	}
	if o.Top != nil {
		out.Append("$top", fmt.Sprintf("%v", *o.Top))
	}
	return &out
}

// ByExternalCloudProviderType ...
func (c DimensionsClient) ByExternalCloudProviderType(ctx context.Context, id ExternalCloudProviderTypeId, options ByExternalCloudProviderTypeOperationOptions) (result ByExternalCloudProviderTypeOperationResponse, err error) {
	opts := client.RequestOptions{
		ContentType: "application/json; charset=utf-8",
		ExpectedStatusCodes: []int{
			http.StatusOK,
		},
		HttpMethod:    http.MethodGet,
		OptionsObject: options,
		Path:          fmt.Sprintf("%s/dimensions", id.ID()),
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

	var model DimensionsListResult
	result.Model = &model
	if err = resp.Unmarshal(result.Model); err != nil {
		return
	}

	return
}
