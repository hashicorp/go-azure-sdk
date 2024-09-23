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

type CreateJoinedTeamArchiveOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
}

type CreateJoinedTeamArchiveOperationOptions struct {
	Metadata  *odata.Metadata
	RetryFunc client.RequestRetryFunc
}

func DefaultCreateJoinedTeamArchiveOperationOptions() CreateJoinedTeamArchiveOperationOptions {
	return CreateJoinedTeamArchiveOperationOptions{}
}

func (o CreateJoinedTeamArchiveOperationOptions) ToHeaders() *client.Headers {
	out := client.Headers{}

	return &out
}

func (o CreateJoinedTeamArchiveOperationOptions) ToOData() *odata.Query {
	out := odata.Query{}
	if o.Metadata != nil {
		out.Metadata = *o.Metadata
	}
	return &out
}

func (o CreateJoinedTeamArchiveOperationOptions) ToQuery() *client.QueryParams {
	out := client.QueryParams{}

	return &out
}

// CreateJoinedTeamArchive - Invoke action archive. Archive the specified team. When a team is archived, users can no
// longer make most changes to the team. For example, users can no longer: send or like messages on any channel in the
// team; edit the team's name or description; nor edit other settings. However, membership changes to the team are still
// allowed. Archiving is an async operation. A team is archived once the async operation completes successfully, which
// might occur subsequent to a response from this API. To archive a team, the team and group must have an owner. To
// restore a team from its archived state, use the API to unarchive.
func (c JoinedTeamClient) CreateJoinedTeamArchive(ctx context.Context, id stable.MeJoinedTeamId, input CreateJoinedTeamArchiveRequest, options CreateJoinedTeamArchiveOperationOptions) (result CreateJoinedTeamArchiveOperationResponse, err error) {
	opts := client.RequestOptions{
		ContentType: "application/json; charset=utf-8",
		ExpectedStatusCodes: []int{
			http.StatusNoContent,
		},
		HttpMethod:    http.MethodPost,
		OptionsObject: options,
		Path:          fmt.Sprintf("%s/archive", id.ID()),
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
