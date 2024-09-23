package joinedteamprimarychannel

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

type CreateJoinedTeamPrimaryChannelArchiveOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
}

type CreateJoinedTeamPrimaryChannelArchiveOperationOptions struct {
	Metadata  *odata.Metadata
	RetryFunc client.RequestRetryFunc
}

func DefaultCreateJoinedTeamPrimaryChannelArchiveOperationOptions() CreateJoinedTeamPrimaryChannelArchiveOperationOptions {
	return CreateJoinedTeamPrimaryChannelArchiveOperationOptions{}
}

func (o CreateJoinedTeamPrimaryChannelArchiveOperationOptions) ToHeaders() *client.Headers {
	out := client.Headers{}

	return &out
}

func (o CreateJoinedTeamPrimaryChannelArchiveOperationOptions) ToOData() *odata.Query {
	out := odata.Query{}
	if o.Metadata != nil {
		out.Metadata = *o.Metadata
	}
	return &out
}

func (o CreateJoinedTeamPrimaryChannelArchiveOperationOptions) ToQuery() *client.QueryParams {
	out := client.QueryParams{}

	return &out
}

// CreateJoinedTeamPrimaryChannelArchive - Invoke action archive. Archive a channel in a team. When a channel is
// archived, users can't send new messages or react to existing messages in the channel, edit the channel settings, or
// make other changes to the channel. You can delete an archived channel or add and remove members from it. If you
// archive a team, its channels are also archived. Archiving is an asynchronous operation; a channel is archived after
// the asynchronous archiving operation completes successfully, which might occur after the response returns. A channel
// without an owner or that belongs to a group that has no owner, can't be archived. To restore a channel from its
// archived state, use the channel: unarchive method. A channel can’t be archived or unarchived if its team is
// archived.
func (c JoinedTeamPrimaryChannelClient) CreateJoinedTeamPrimaryChannelArchive(ctx context.Context, id stable.UserIdJoinedTeamId, input CreateJoinedTeamPrimaryChannelArchiveRequest, options CreateJoinedTeamPrimaryChannelArchiveOperationOptions) (result CreateJoinedTeamPrimaryChannelArchiveOperationResponse, err error) {
	opts := client.RequestOptions{
		ContentType: "application/json; charset=utf-8",
		ExpectedStatusCodes: []int{
			http.StatusNoContent,
		},
		HttpMethod:    http.MethodPost,
		OptionsObject: options,
		Path:          fmt.Sprintf("%s/primaryChannel/archive", id.ID()),
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
