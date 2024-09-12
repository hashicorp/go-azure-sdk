package teamworkassociatedteam

import (
	"context"
	"net/http"

	"github.com/hashicorp/go-azure-sdk/microsoft-graph/common-types/stable"
	"github.com/hashicorp/go-azure-sdk/sdk/client"
	"github.com/hashicorp/go-azure-sdk/sdk/odata"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type GetTeamworkAssociatedTeamOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
	Model        *stable.AssociatedTeamInfo
}

type GetTeamworkAssociatedTeamOperationOptions struct {
	Expand *odata.Expand
	Select *[]string
}

func DefaultGetTeamworkAssociatedTeamOperationOptions() GetTeamworkAssociatedTeamOperationOptions {
	return GetTeamworkAssociatedTeamOperationOptions{}
}

func (o GetTeamworkAssociatedTeamOperationOptions) ToHeaders() *client.Headers {
	out := client.Headers{}

	return &out
}

func (o GetTeamworkAssociatedTeamOperationOptions) ToOData() *odata.Query {
	out := odata.Query{}
	if o.Expand != nil {
		out.Expand = *o.Expand
	}
	if o.Select != nil {
		out.Select = *o.Select
	}
	return &out
}

func (o GetTeamworkAssociatedTeamOperationOptions) ToQuery() *client.QueryParams {
	out := client.QueryParams{}

	return &out
}

// GetTeamworkAssociatedTeam - Get associatedTeams from me. The list of associatedTeamInfo objects that a user is
// associated with.
func (c TeamworkAssociatedTeamClient) GetTeamworkAssociatedTeam(ctx context.Context, id stable.MeTeamworkAssociatedTeamId, options GetTeamworkAssociatedTeamOperationOptions) (result GetTeamworkAssociatedTeamOperationResponse, err error) {
	opts := client.RequestOptions{
		ContentType: "application/json; charset=utf-8",
		ExpectedStatusCodes: []int{
			http.StatusOK,
		},
		HttpMethod:    http.MethodGet,
		OptionsObject: options,
		Path:          id.ID(),
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

	var model stable.AssociatedTeamInfo
	result.Model = &model
	if err = resp.Unmarshal(result.Model); err != nil {
		return
	}

	return
}
