package permissionsanalyticazurefinding

import (
	"context"
	"net/http"

	"github.com/hashicorp/go-azure-sdk/sdk/client"
	"github.com/hashicorp/go-azure-sdk/sdk/odata"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type GetPermissionsAnalyticAzureFindingsCountOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
	Model        *[]byte
}

type GetPermissionsAnalyticAzureFindingsCountOperationOptions struct {
	Filter *string
	Search *string
}

func DefaultGetPermissionsAnalyticAzureFindingsCountOperationOptions() GetPermissionsAnalyticAzureFindingsCountOperationOptions {
	return GetPermissionsAnalyticAzureFindingsCountOperationOptions{}
}

func (o GetPermissionsAnalyticAzureFindingsCountOperationOptions) ToHeaders() *client.Headers {
	out := client.Headers{}

	return &out
}

func (o GetPermissionsAnalyticAzureFindingsCountOperationOptions) ToOData() *odata.Query {
	out := odata.Query{}
	if o.Filter != nil {
		out.Filter = *o.Filter
	}
	if o.Search != nil {
		out.Search = *o.Search
	}
	return &out
}

func (o GetPermissionsAnalyticAzureFindingsCountOperationOptions) ToQuery() *client.QueryParams {
	out := client.QueryParams{}

	return &out
}

// GetPermissionsAnalyticAzureFindingsCount - Get the number of the resource
func (c PermissionsAnalyticAzureFindingClient) GetPermissionsAnalyticAzureFindingsCount(ctx context.Context, options GetPermissionsAnalyticAzureFindingsCountOperationOptions) (result GetPermissionsAnalyticAzureFindingsCountOperationResponse, err error) {
	opts := client.RequestOptions{
		ContentType: "text/plain",
		ExpectedStatusCodes: []int{
			http.StatusOK,
		},
		HttpMethod:    http.MethodGet,
		OptionsObject: options,
		Path:          "/identityGovernance/permissionsAnalytics/azure/findings/$count",
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

	var model []byte
	result.Model = &model
	if err = resp.Unmarshal(result.Model); err != nil {
		return
	}

	return
}
