package teamworkinstalledapp

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

type CreateTeamworkInstalledAppOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
	Model        *stable.UserScopeTeamsAppInstallation
}

type CreateTeamworkInstalledAppOperationOptions struct {
	Metadata *odata.Metadata
}

func DefaultCreateTeamworkInstalledAppOperationOptions() CreateTeamworkInstalledAppOperationOptions {
	return CreateTeamworkInstalledAppOperationOptions{}
}

func (o CreateTeamworkInstalledAppOperationOptions) ToHeaders() *client.Headers {
	out := client.Headers{}

	return &out
}

func (o CreateTeamworkInstalledAppOperationOptions) ToOData() *odata.Query {
	out := odata.Query{}
	if o.Metadata != nil {
		out.Metadata = *o.Metadata
	}
	return &out
}

func (o CreateTeamworkInstalledAppOperationOptions) ToQuery() *client.QueryParams {
	out := client.QueryParams{}

	return &out
}

// CreateTeamworkInstalledApp - Install app for user. Install an app in the personal scope of the specified user.
func (c TeamworkInstalledAppClient) CreateTeamworkInstalledApp(ctx context.Context, id stable.UserId, input stable.UserScopeTeamsAppInstallation, options CreateTeamworkInstalledAppOperationOptions) (result CreateTeamworkInstalledAppOperationResponse, err error) {
	opts := client.RequestOptions{
		ContentType: "application/json; charset=utf-8",
		ExpectedStatusCodes: []int{
			http.StatusCreated,
		},
		HttpMethod:    http.MethodPost,
		OptionsObject: options,
		Path:          fmt.Sprintf("%s/teamwork/installedApps", id.ID()),
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

	var model stable.UserScopeTeamsAppInstallation
	result.Model = &model
	if err = resp.Unmarshal(result.Model); err != nil {
		return
	}

	return
}
