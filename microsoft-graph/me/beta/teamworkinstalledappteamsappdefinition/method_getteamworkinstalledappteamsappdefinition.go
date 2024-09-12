package teamworkinstalledappteamsappdefinition

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

type GetTeamworkInstalledAppTeamsAppDefinitionOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
	Model        *beta.TeamsAppDefinition
}

type GetTeamworkInstalledAppTeamsAppDefinitionOperationOptions struct {
	Expand *odata.Expand
	Select *[]string
}

func DefaultGetTeamworkInstalledAppTeamsAppDefinitionOperationOptions() GetTeamworkInstalledAppTeamsAppDefinitionOperationOptions {
	return GetTeamworkInstalledAppTeamsAppDefinitionOperationOptions{}
}

func (o GetTeamworkInstalledAppTeamsAppDefinitionOperationOptions) ToHeaders() *client.Headers {
	out := client.Headers{}

	return &out
}

func (o GetTeamworkInstalledAppTeamsAppDefinitionOperationOptions) ToOData() *odata.Query {
	out := odata.Query{}
	if o.Expand != nil {
		out.Expand = *o.Expand
	}
	if o.Select != nil {
		out.Select = *o.Select
	}
	return &out
}

func (o GetTeamworkInstalledAppTeamsAppDefinitionOperationOptions) ToQuery() *client.QueryParams {
	out := client.QueryParams{}

	return &out
}

// GetTeamworkInstalledAppTeamsAppDefinition - Get teamsAppDefinition from me. The details of this version of the app.
func (c TeamworkInstalledAppTeamsAppDefinitionClient) GetTeamworkInstalledAppTeamsAppDefinition(ctx context.Context, id beta.MeTeamworkInstalledAppId, options GetTeamworkInstalledAppTeamsAppDefinitionOperationOptions) (result GetTeamworkInstalledAppTeamsAppDefinitionOperationResponse, err error) {
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

	var model beta.TeamsAppDefinition
	result.Model = &model
	if err = resp.Unmarshal(result.Model); err != nil {
		return
	}

	return
}
