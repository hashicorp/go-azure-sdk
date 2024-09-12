package teaminstalledappteamsappdefinition

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

type GetTeamInstalledAppTeamsAppDefinitionOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
	Model        *stable.TeamsAppDefinition
}

type GetTeamInstalledAppTeamsAppDefinitionOperationOptions struct {
	Expand *odata.Expand
	Select *[]string
}

func DefaultGetTeamInstalledAppTeamsAppDefinitionOperationOptions() GetTeamInstalledAppTeamsAppDefinitionOperationOptions {
	return GetTeamInstalledAppTeamsAppDefinitionOperationOptions{}
}

func (o GetTeamInstalledAppTeamsAppDefinitionOperationOptions) ToHeaders() *client.Headers {
	out := client.Headers{}

	return &out
}

func (o GetTeamInstalledAppTeamsAppDefinitionOperationOptions) ToOData() *odata.Query {
	out := odata.Query{}
	if o.Expand != nil {
		out.Expand = *o.Expand
	}
	if o.Select != nil {
		out.Select = *o.Select
	}
	return &out
}

func (o GetTeamInstalledAppTeamsAppDefinitionOperationOptions) ToQuery() *client.QueryParams {
	out := client.QueryParams{}

	return &out
}

// GetTeamInstalledAppTeamsAppDefinition - Get teamsAppDefinition from groups. The details of this version of the app.
func (c TeamInstalledAppTeamsAppDefinitionClient) GetTeamInstalledAppTeamsAppDefinition(ctx context.Context, id stable.GroupIdTeamInstalledAppId, options GetTeamInstalledAppTeamsAppDefinitionOperationOptions) (result GetTeamInstalledAppTeamsAppDefinitionOperationResponse, err error) {
	opts := client.RequestOptions{
		ContentType: "application/json; charset=utf-8",
		ExpectedStatusCodes: []int{
			http.StatusOK,
		},
		HttpMethod:    http.MethodGet,
		OptionsObject: options,
		Path:          fmt.Sprintf("%s/teamsAppDefinition", id.ID()),
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

	var model stable.TeamsAppDefinition
	result.Model = &model
	if err = resp.Unmarshal(result.Model); err != nil {
		return
	}

	return
}
