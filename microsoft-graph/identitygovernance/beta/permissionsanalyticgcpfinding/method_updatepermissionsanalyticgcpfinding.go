package permissionsanalyticgcpfinding

import (
	"context"
	"net/http"

	"github.com/hashicorp/go-azure-sdk/microsoft-graph/common-types/beta"
	"github.com/hashicorp/go-azure-sdk/sdk/client"
	"github.com/hashicorp/go-azure-sdk/sdk/odata"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type UpdatePermissionsAnalyticGcpFindingOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
}

type UpdatePermissionsAnalyticGcpFindingOperationOptions struct {
	Metadata *odata.Metadata
}

func DefaultUpdatePermissionsAnalyticGcpFindingOperationOptions() UpdatePermissionsAnalyticGcpFindingOperationOptions {
	return UpdatePermissionsAnalyticGcpFindingOperationOptions{}
}

func (o UpdatePermissionsAnalyticGcpFindingOperationOptions) ToHeaders() *client.Headers {
	out := client.Headers{}

	return &out
}

func (o UpdatePermissionsAnalyticGcpFindingOperationOptions) ToOData() *odata.Query {
	out := odata.Query{}
	if o.Metadata != nil {
		out.Metadata = *o.Metadata
	}
	return &out
}

func (o UpdatePermissionsAnalyticGcpFindingOperationOptions) ToQuery() *client.QueryParams {
	out := client.QueryParams{}

	return &out
}

// UpdatePermissionsAnalyticGcpFinding - Update the navigation property findings in identityGovernance
func (c PermissionsAnalyticGcpFindingClient) UpdatePermissionsAnalyticGcpFinding(ctx context.Context, id beta.IdentityGovernancePermissionsAnalyticGcpFindingId, input beta.Finding, options UpdatePermissionsAnalyticGcpFindingOperationOptions) (result UpdatePermissionsAnalyticGcpFindingOperationResponse, err error) {
	opts := client.RequestOptions{
		ContentType: "application/json; charset=utf-8",
		ExpectedStatusCodes: []int{
			http.StatusNoContent,
		},
		HttpMethod:    http.MethodPatch,
		OptionsObject: options,
		Path:          id.ID(),
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
