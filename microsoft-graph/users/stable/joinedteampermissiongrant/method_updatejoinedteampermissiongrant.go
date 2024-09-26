package joinedteampermissiongrant

import (
	"context"
	"net/http"

	"github.com/hashicorp/go-azure-sdk/microsoft-graph/common-types/stable"
	"github.com/hashicorp/go-azure-sdk/sdk/client"
	"github.com/hashicorp/go-azure-sdk/sdk/odata"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type UpdateJoinedTeamPermissionGrantOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
}

type UpdateJoinedTeamPermissionGrantOperationOptions struct {
	Metadata  *odata.Metadata
	RetryFunc client.RequestRetryFunc
}

func DefaultUpdateJoinedTeamPermissionGrantOperationOptions() UpdateJoinedTeamPermissionGrantOperationOptions {
	return UpdateJoinedTeamPermissionGrantOperationOptions{}
}

func (o UpdateJoinedTeamPermissionGrantOperationOptions) ToHeaders() *client.Headers {
	out := client.Headers{}

	return &out
}

func (o UpdateJoinedTeamPermissionGrantOperationOptions) ToOData() *odata.Query {
	out := odata.Query{}
	if o.Metadata != nil {
		out.Metadata = *o.Metadata
	}
	return &out
}

func (o UpdateJoinedTeamPermissionGrantOperationOptions) ToQuery() *client.QueryParams {
	out := client.QueryParams{}

	return &out
}

// UpdateJoinedTeamPermissionGrant - Update the navigation property permissionGrants in users
func (c JoinedTeamPermissionGrantClient) UpdateJoinedTeamPermissionGrant(ctx context.Context, id stable.UserIdJoinedTeamIdPermissionGrantId, input stable.ResourceSpecificPermissionGrant, options UpdateJoinedTeamPermissionGrantOperationOptions) (result UpdateJoinedTeamPermissionGrantOperationResponse, err error) {
	opts := client.RequestOptions{
		ContentType: "application/json; charset=utf-8",
		ExpectedStatusCodes: []int{
			http.StatusAccepted,
			http.StatusNoContent,
			http.StatusOK,
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
