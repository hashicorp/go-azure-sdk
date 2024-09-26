package teamchannelmessagereply

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

type CreateTeamChannelMessageReplyUndoSoftDeleteOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
}

type CreateTeamChannelMessageReplyUndoSoftDeleteOperationOptions struct {
	Metadata  *odata.Metadata
	RetryFunc client.RequestRetryFunc
}

func DefaultCreateTeamChannelMessageReplyUndoSoftDeleteOperationOptions() CreateTeamChannelMessageReplyUndoSoftDeleteOperationOptions {
	return CreateTeamChannelMessageReplyUndoSoftDeleteOperationOptions{}
}

func (o CreateTeamChannelMessageReplyUndoSoftDeleteOperationOptions) ToHeaders() *client.Headers {
	out := client.Headers{}

	return &out
}

func (o CreateTeamChannelMessageReplyUndoSoftDeleteOperationOptions) ToOData() *odata.Query {
	out := odata.Query{}
	if o.Metadata != nil {
		out.Metadata = *o.Metadata
	}
	return &out
}

func (o CreateTeamChannelMessageReplyUndoSoftDeleteOperationOptions) ToQuery() *client.QueryParams {
	out := client.QueryParams{}

	return &out
}

// CreateTeamChannelMessageReplyUndoSoftDelete - Invoke action undoSoftDelete. Undo soft deletion of a single
// chatMessage or a chat message reply in a channel or a chat.
func (c TeamChannelMessageReplyClient) CreateTeamChannelMessageReplyUndoSoftDelete(ctx context.Context, id beta.GroupIdTeamChannelIdMessageIdReplyId, options CreateTeamChannelMessageReplyUndoSoftDeleteOperationOptions) (result CreateTeamChannelMessageReplyUndoSoftDeleteOperationResponse, err error) {
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
		Path:          fmt.Sprintf("%s/undoSoftDelete", id.ID()),
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
