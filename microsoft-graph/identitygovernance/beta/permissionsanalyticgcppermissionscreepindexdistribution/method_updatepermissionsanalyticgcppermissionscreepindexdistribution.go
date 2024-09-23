package permissionsanalyticgcppermissionscreepindexdistribution

import (
	"context"
	"net/http"

	"github.com/hashicorp/go-azure-sdk/microsoft-graph/common-types/beta"
	"github.com/hashicorp/go-azure-sdk/sdk/client"
	"github.com/hashicorp/go-azure-sdk/sdk/odata"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type UpdatePermissionsAnalyticGcpPermissionsCreepIndexDistributionOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
}

type UpdatePermissionsAnalyticGcpPermissionsCreepIndexDistributionOperationOptions struct {
	Metadata  *odata.Metadata
	RetryFunc client.RequestRetryFunc
}

func DefaultUpdatePermissionsAnalyticGcpPermissionsCreepIndexDistributionOperationOptions() UpdatePermissionsAnalyticGcpPermissionsCreepIndexDistributionOperationOptions {
	return UpdatePermissionsAnalyticGcpPermissionsCreepIndexDistributionOperationOptions{}
}

func (o UpdatePermissionsAnalyticGcpPermissionsCreepIndexDistributionOperationOptions) ToHeaders() *client.Headers {
	out := client.Headers{}

	return &out
}

func (o UpdatePermissionsAnalyticGcpPermissionsCreepIndexDistributionOperationOptions) ToOData() *odata.Query {
	out := odata.Query{}
	if o.Metadata != nil {
		out.Metadata = *o.Metadata
	}
	return &out
}

func (o UpdatePermissionsAnalyticGcpPermissionsCreepIndexDistributionOperationOptions) ToQuery() *client.QueryParams {
	out := client.QueryParams{}

	return &out
}

// UpdatePermissionsAnalyticGcpPermissionsCreepIndexDistribution - Update the navigation property
// permissionsCreepIndexDistributions in identityGovernance
func (c PermissionsAnalyticGcpPermissionsCreepIndexDistributionClient) UpdatePermissionsAnalyticGcpPermissionsCreepIndexDistribution(ctx context.Context, id beta.IdentityGovernancePermissionsAnalyticGcpPermissionsCreepIndexDistributionId, input beta.PermissionsCreepIndexDistribution, options UpdatePermissionsAnalyticGcpPermissionsCreepIndexDistributionOperationOptions) (result UpdatePermissionsAnalyticGcpPermissionsCreepIndexDistributionOperationResponse, err error) {
	opts := client.RequestOptions{
		ContentType: "application/json; charset=utf-8",
		ExpectedStatusCodes: []int{
			http.StatusNoContent,
		},
		HttpMethod:    http.MethodPatch,
		OptionsObject: options,
		Path:          id.ID(),
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

	return
}
