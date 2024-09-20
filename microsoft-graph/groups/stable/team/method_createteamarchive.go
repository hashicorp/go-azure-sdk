package team

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

type CreateTeamArchiveOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
}

type CreateTeamArchiveOperationOptions struct {
	Metadata *odata.Metadata
}

func DefaultCreateTeamArchiveOperationOptions() CreateTeamArchiveOperationOptions {
	return CreateTeamArchiveOperationOptions{}
}

func (o CreateTeamArchiveOperationOptions) ToHeaders() *client.Headers {
	out := client.Headers{}

	return &out
}

func (o CreateTeamArchiveOperationOptions) ToOData() *odata.Query {
	out := odata.Query{}
	if o.Metadata != nil {
		out.Metadata = *o.Metadata
	}
	return &out
}

func (o CreateTeamArchiveOperationOptions) ToQuery() *client.QueryParams {
	out := client.QueryParams{}

	return &out
}

// CreateTeamArchive - Invoke action archive. Archive the specified team. When a team is archived, users can no longer
// make most changes to the team. For example, users can no longer: send or like messages on any channel in the team;
// edit the team's name or description; nor edit other settings. However, membership changes to the team are still
// allowed. Archiving is an async operation. A team is archived once the async operation completes successfully, which
// might occur subsequent to a response from this API. To archive a team, the team and group must have an owner. To
// restore a team from its archived state, use the API to unarchive.
func (c TeamClient) CreateTeamArchive(ctx context.Context, id stable.GroupId, input CreateTeamArchiveRequest, options CreateTeamArchiveOperationOptions) (result CreateTeamArchiveOperationResponse, err error) {
	opts := client.RequestOptions{
		ContentType: "application/json; charset=utf-8",
		ExpectedStatusCodes: []int{
			http.StatusNoContent,
		},
		HttpMethod:    http.MethodPost,
		OptionsObject: options,
		Path:          fmt.Sprintf("%s/team/archive", id.ID()),
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
