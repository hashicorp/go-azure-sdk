package teaminstalledappteamsapp

import (
	"context"
	"fmt"
	"net/http"

	"github.com/hashicorp/go-azure-sdk/microsoft-graph/common-types/beta"
	"github.com/hashicorp/go-azure-sdk/sdk/client"
	"github.com/hashicorp/go-azure-sdk/sdk/odata"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type GetTeamInstalledAppTeamsAppOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
	Model        *beta.TeamsApp
}

type GetTeamInstalledAppTeamsAppOperationOptions struct {
	Expand *odata.Expand
	Select *[]string
}

func DefaultGetTeamInstalledAppTeamsAppOperationOptions() GetTeamInstalledAppTeamsAppOperationOptions {
	return GetTeamInstalledAppTeamsAppOperationOptions{}
}

func (o GetTeamInstalledAppTeamsAppOperationOptions) ToHeaders() *client.Headers {
	out := client.Headers{}

	return &out
}

func (o GetTeamInstalledAppTeamsAppOperationOptions) ToOData() *odata.Query {
	out := odata.Query{}
	if o.Expand != nil {
		out.Expand = *o.Expand
	}
	if o.Select != nil {
		out.Select = *o.Select
	}
	return &out
}

func (o GetTeamInstalledAppTeamsAppOperationOptions) ToQuery() *client.QueryParams {
	out := client.QueryParams{}

	return &out
}

// GetTeamInstalledAppTeamsApp - Get teamsApp from groups. The app that is installed.
func (c TeamInstalledAppTeamsAppClient) GetTeamInstalledAppTeamsApp(ctx context.Context, id beta.GroupIdTeamInstalledAppId, options GetTeamInstalledAppTeamsAppOperationOptions) (result GetTeamInstalledAppTeamsAppOperationResponse, err error) {
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

	var model beta.TeamsApp
	result.Model = &model
	if err = resp.Unmarshal(result.Model); err != nil {
		return
	}

	return
}
