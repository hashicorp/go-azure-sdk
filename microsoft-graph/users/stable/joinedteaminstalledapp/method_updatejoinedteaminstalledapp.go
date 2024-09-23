package joinedteaminstalledapp

import (
	"context"
	"net/http"

	"github.com/hashicorp/go-azure-sdk/microsoft-graph/common-types/stable"
	"github.com/hashicorp/go-azure-sdk/sdk/client"
	"github.com/hashicorp/go-azure-sdk/sdk/odata"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type UpdateJoinedTeamInstalledAppOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
}

type UpdateJoinedTeamInstalledAppOperationOptions struct {
	Metadata  *odata.Metadata
	RetryFunc client.RequestRetryFunc
}

func DefaultUpdateJoinedTeamInstalledAppOperationOptions() UpdateJoinedTeamInstalledAppOperationOptions {
	return UpdateJoinedTeamInstalledAppOperationOptions{}
}

func (o UpdateJoinedTeamInstalledAppOperationOptions) ToHeaders() *client.Headers {
	out := client.Headers{}

	return &out
}

func (o UpdateJoinedTeamInstalledAppOperationOptions) ToOData() *odata.Query {
	out := odata.Query{}
	if o.Metadata != nil {
		out.Metadata = *o.Metadata
	}
	return &out
}

func (o UpdateJoinedTeamInstalledAppOperationOptions) ToQuery() *client.QueryParams {
	out := client.QueryParams{}

	return &out
}

// UpdateJoinedTeamInstalledApp - Update the navigation property installedApps in users
func (c JoinedTeamInstalledAppClient) UpdateJoinedTeamInstalledApp(ctx context.Context, id stable.UserIdJoinedTeamIdInstalledAppId, input stable.TeamsAppInstallation, options UpdateJoinedTeamInstalledAppOperationOptions) (result UpdateJoinedTeamInstalledAppOperationResponse, err error) {
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
