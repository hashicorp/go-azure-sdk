package joinedteamprimarychanneltabteamsapp

import (
	"context"
	"fmt"
	"net/http"

	"github.com/hashicorp/go-azure-sdk/microsoft-graph/common-types/stable"
	"github.com/hashicorp/go-azure-sdk/sdk/client"
	"github.com/hashicorp/go-azure-sdk/sdk/odata"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type GetJoinedTeamPrimaryChannelTabTeamsAppOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
	Model        *stable.TeamsApp
}

type GetJoinedTeamPrimaryChannelTabTeamsAppOperationOptions struct {
	Expand   *odata.Expand
	Metadata *odata.Metadata
	Select   *[]string
}

func DefaultGetJoinedTeamPrimaryChannelTabTeamsAppOperationOptions() GetJoinedTeamPrimaryChannelTabTeamsAppOperationOptions {
	return GetJoinedTeamPrimaryChannelTabTeamsAppOperationOptions{}
}

func (o GetJoinedTeamPrimaryChannelTabTeamsAppOperationOptions) ToHeaders() *client.Headers {
	out := client.Headers{}

	return &out
}

func (o GetJoinedTeamPrimaryChannelTabTeamsAppOperationOptions) ToOData() *odata.Query {
	out := odata.Query{}
	if o.Expand != nil {
		out.Expand = *o.Expand
	}
	if o.Metadata != nil {
		out.Metadata = *o.Metadata
	}
	if o.Select != nil {
		out.Select = *o.Select
	}
	return &out
}

func (o GetJoinedTeamPrimaryChannelTabTeamsAppOperationOptions) ToQuery() *client.QueryParams {
	out := client.QueryParams{}

	return &out
}

// GetJoinedTeamPrimaryChannelTabTeamsApp - Get teamsApp from me. The application that is linked to the tab. This can't
// be changed after tab creation.
func (c JoinedTeamPrimaryChannelTabTeamsAppClient) GetJoinedTeamPrimaryChannelTabTeamsApp(ctx context.Context, id stable.MeJoinedTeamIdPrimaryChannelTabId, options GetJoinedTeamPrimaryChannelTabTeamsAppOperationOptions) (result GetJoinedTeamPrimaryChannelTabTeamsAppOperationResponse, err error) {
	opts := client.RequestOptions{
		ContentType: "application/json; charset=utf-8",
		ExpectedStatusCodes: []int{
			http.StatusOK,
		},
		HttpMethod:    http.MethodGet,
		OptionsObject: options,
		Path:          fmt.Sprintf("%s/teamsApp", id.ID()),
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

	var model stable.TeamsApp
	result.Model = &model
	if err = resp.Unmarshal(result.Model); err != nil {
		return
	}

	return
}
