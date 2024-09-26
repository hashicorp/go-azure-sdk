package joinedteam

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

type CreateJoinedTeamUnarchiveOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
}

type CreateJoinedTeamUnarchiveOperationOptions struct {
	Metadata  *odata.Metadata
	RetryFunc client.RequestRetryFunc
}

func DefaultCreateJoinedTeamUnarchiveOperationOptions() CreateJoinedTeamUnarchiveOperationOptions {
	return CreateJoinedTeamUnarchiveOperationOptions{}
}

func (o CreateJoinedTeamUnarchiveOperationOptions) ToHeaders() *client.Headers {
	out := client.Headers{}

	return &out
}

func (o CreateJoinedTeamUnarchiveOperationOptions) ToOData() *odata.Query {
	out := odata.Query{}
	if o.Metadata != nil {
		out.Metadata = *o.Metadata
	}
	return &out
}

func (o CreateJoinedTeamUnarchiveOperationOptions) ToQuery() *client.QueryParams {
	out := client.QueryParams{}

	return &out
}

// CreateJoinedTeamUnarchive - Invoke action unarchive. Restore an archived team. This API restores users' ability to
// send messages and edit the team, abiding by tenant and team settings. A Team is archived using the archive API.
// Unarchiving is an async operation. A team is unarchived once the async operation completes successfully, which might
// occur subsequent to a response from this API.
func (c JoinedTeamClient) CreateJoinedTeamUnarchive(ctx context.Context, id stable.UserIdJoinedTeamId, options CreateJoinedTeamUnarchiveOperationOptions) (result CreateJoinedTeamUnarchiveOperationResponse, err error) {
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
		Path:          fmt.Sprintf("%s/unarchive", id.ID()),
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
