package teamprimarychannelsharedwithteam

import (
	"context"
	"net/http"

	"github.com/hashicorp/go-azure-sdk/microsoft-graph/common-types/beta"
	"github.com/hashicorp/go-azure-sdk/sdk/client"
	"github.com/hashicorp/go-azure-sdk/sdk/odata"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type UpdateTeamPrimaryChannelSharedWithTeamOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
}

type UpdateTeamPrimaryChannelSharedWithTeamOperationOptions struct {
	Metadata *odata.Metadata
}

func DefaultUpdateTeamPrimaryChannelSharedWithTeamOperationOptions() UpdateTeamPrimaryChannelSharedWithTeamOperationOptions {
	return UpdateTeamPrimaryChannelSharedWithTeamOperationOptions{}
}

func (o UpdateTeamPrimaryChannelSharedWithTeamOperationOptions) ToHeaders() *client.Headers {
	out := client.Headers{}

	return &out
}

func (o UpdateTeamPrimaryChannelSharedWithTeamOperationOptions) ToOData() *odata.Query {
	out := odata.Query{}
	if o.Metadata != nil {
		out.Metadata = *o.Metadata
	}
	return &out
}

func (o UpdateTeamPrimaryChannelSharedWithTeamOperationOptions) ToQuery() *client.QueryParams {
	out := client.QueryParams{}

	return &out
}

// UpdateTeamPrimaryChannelSharedWithTeam - Update the navigation property sharedWithTeams in groups
func (c TeamPrimaryChannelSharedWithTeamClient) UpdateTeamPrimaryChannelSharedWithTeam(ctx context.Context, id beta.GroupIdTeamPrimaryChannelSharedWithTeamId, input beta.SharedWithChannelTeamInfo, options UpdateTeamPrimaryChannelSharedWithTeamOperationOptions) (result UpdateTeamPrimaryChannelSharedWithTeamOperationResponse, err error) {
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
