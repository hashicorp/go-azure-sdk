package teaminstalledapp

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

type CreateTeamInstalledAppUpgradeOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
}

type CreateTeamInstalledAppUpgradeOperationOptions struct {
	Metadata  *odata.Metadata
	RetryFunc client.RequestRetryFunc
}

func DefaultCreateTeamInstalledAppUpgradeOperationOptions() CreateTeamInstalledAppUpgradeOperationOptions {
	return CreateTeamInstalledAppUpgradeOperationOptions{}
}

func (o CreateTeamInstalledAppUpgradeOperationOptions) ToHeaders() *client.Headers {
	out := client.Headers{}

	return &out
}

func (o CreateTeamInstalledAppUpgradeOperationOptions) ToOData() *odata.Query {
	out := odata.Query{}
	if o.Metadata != nil {
		out.Metadata = *o.Metadata
	}
	return &out
}

func (o CreateTeamInstalledAppUpgradeOperationOptions) ToQuery() *client.QueryParams {
	out := client.QueryParams{}

	return &out
}

// CreateTeamInstalledAppUpgrade - Invoke action upgrade. Upgrade an app installation within a chat.
func (c TeamInstalledAppClient) CreateTeamInstalledAppUpgrade(ctx context.Context, id beta.GroupIdTeamInstalledAppId, input CreateTeamInstalledAppUpgradeRequest, options CreateTeamInstalledAppUpgradeOperationOptions) (result CreateTeamInstalledAppUpgradeOperationResponse, err error) {
	opts := client.RequestOptions{
		ContentType: "application/json; charset=utf-8",
		ExpectedStatusCodes: []int{
			http.StatusAccepted,
			http.StatusCreated,
			http.StatusNoContent,
			http.StatusOK,
		},
		HttpMethod:    http.MethodPost,
		OptionsObject: options,
		Path:          fmt.Sprintf("%s/upgrade", id.ID()),
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
