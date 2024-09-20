package teamtagmember

import (
	"context"
	"net/http"

	"github.com/hashicorp/go-azure-sdk/microsoft-graph/common-types/beta"
	"github.com/hashicorp/go-azure-sdk/sdk/client"
	"github.com/hashicorp/go-azure-sdk/sdk/odata"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type UpdateTeamTagMemberOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
}

type UpdateTeamTagMemberOperationOptions struct {
	Metadata *odata.Metadata
}

func DefaultUpdateTeamTagMemberOperationOptions() UpdateTeamTagMemberOperationOptions {
	return UpdateTeamTagMemberOperationOptions{}
}

func (o UpdateTeamTagMemberOperationOptions) ToHeaders() *client.Headers {
	out := client.Headers{}

	return &out
}

func (o UpdateTeamTagMemberOperationOptions) ToOData() *odata.Query {
	out := odata.Query{}
	if o.Metadata != nil {
		out.Metadata = *o.Metadata
	}
	return &out
}

func (o UpdateTeamTagMemberOperationOptions) ToQuery() *client.QueryParams {
	out := client.QueryParams{}

	return &out
}

// UpdateTeamTagMember - Update the navigation property members in groups
func (c TeamTagMemberClient) UpdateTeamTagMember(ctx context.Context, id beta.GroupIdTeamTagIdMemberId, input beta.TeamworkTagMember, options UpdateTeamTagMemberOperationOptions) (result UpdateTeamTagMemberOperationResponse, err error) {
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
