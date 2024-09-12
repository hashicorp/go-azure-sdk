package permissionsanalyticawpermissionscreepindexdistribution

import (
	"context"
	"net/http"

	"github.com/hashicorp/go-azure-sdk/sdk/client"
	"github.com/hashicorp/go-azure-sdk/sdk/odata"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type GetPermissionsAnalyticAwPermissionsCreepIndexDistributionsCountOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
	Model        *[]byte
}

type GetPermissionsAnalyticAwPermissionsCreepIndexDistributionsCountOperationOptions struct {
	Filter *string
	Search *string
}

func DefaultGetPermissionsAnalyticAwPermissionsCreepIndexDistributionsCountOperationOptions() GetPermissionsAnalyticAwPermissionsCreepIndexDistributionsCountOperationOptions {
	return GetPermissionsAnalyticAwPermissionsCreepIndexDistributionsCountOperationOptions{}
}

func (o GetPermissionsAnalyticAwPermissionsCreepIndexDistributionsCountOperationOptions) ToHeaders() *client.Headers {
	out := client.Headers{}

	return &out
}

func (o GetPermissionsAnalyticAwPermissionsCreepIndexDistributionsCountOperationOptions) ToOData() *odata.Query {
	out := odata.Query{}
	if o.Filter != nil {
		out.Filter = *o.Filter
	}
	if o.Search != nil {
		out.Search = *o.Search
	}
	return &out
}

func (o GetPermissionsAnalyticAwPermissionsCreepIndexDistributionsCountOperationOptions) ToQuery() *client.QueryParams {
	out := client.QueryParams{}

	return &out
}

// GetPermissionsAnalyticAwPermissionsCreepIndexDistributionsCount - Get the number of the resource
func (c PermissionsAnalyticAwPermissionsCreepIndexDistributionClient) GetPermissionsAnalyticAwPermissionsCreepIndexDistributionsCount(ctx context.Context, options GetPermissionsAnalyticAwPermissionsCreepIndexDistributionsCountOperationOptions) (result GetPermissionsAnalyticAwPermissionsCreepIndexDistributionsCountOperationResponse, err error) {
	opts := client.RequestOptions{
		ContentType: "text/plain",
		ExpectedStatusCodes: []int{
			http.StatusOK,
		},
		HttpMethod:    http.MethodGet,
		OptionsObject: options,
		Path:          "/identityGovernance/permissionsAnalytics/aws/permissionsCreepIndexDistributions/$count",
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
