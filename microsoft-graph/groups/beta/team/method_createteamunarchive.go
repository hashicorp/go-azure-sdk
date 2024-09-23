package team

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

type CreateTeamUnarchiveOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
}

type CreateTeamUnarchiveOperationOptions struct {
	Metadata  *odata.Metadata
	RetryFunc client.RequestRetryFunc
}

func DefaultCreateTeamUnarchiveOperationOptions() CreateTeamUnarchiveOperationOptions {
	return CreateTeamUnarchiveOperationOptions{}
}

func (o CreateTeamUnarchiveOperationOptions) ToHeaders() *client.Headers {
	out := client.Headers{}

	return &out
}

func (o CreateTeamUnarchiveOperationOptions) ToOData() *odata.Query {
	out := odata.Query{}
	if o.Metadata != nil {
		out.Metadata = *o.Metadata
	}
	return &out
}

func (o CreateTeamUnarchiveOperationOptions) ToQuery() *client.QueryParams {
	out := client.QueryParams{}

	return &out
}

// CreateTeamUnarchive - Invoke action unarchive. Restore an archived team and restores users' ability to send messages
// and edit the team, abiding by tenant and team settings. Teams are archived using the archive API. Unarchiving is an
// async operation. A team is unarchived once the async operation completes successfully, which might occur subsequent
// to a response from this API.
func (c TeamClient) CreateTeamUnarchive(ctx context.Context, id beta.GroupId, options CreateTeamUnarchiveOperationOptions) (result CreateTeamUnarchiveOperationResponse, err error) {
	opts := client.RequestOptions{
		ContentType: "application/json; charset=utf-8",
		ExpectedStatusCodes: []int{
			http.StatusNoContent,
		},
		HttpMethod:    http.MethodPost,
		OptionsObject: options,
		Path:          fmt.Sprintf("%s/team/unarchive", id.ID()),
		RetryFunc:     options.RetryFunc,
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

	return
}
