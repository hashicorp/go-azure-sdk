package joinedteamchannelmessagereply

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

type CreateJoinedTeamChannelMessageReplySoftDeleteOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
}

// CreateJoinedTeamChannelMessageReplySoftDelete - Invoke action softDelete. Delete a single chatMessage or a chat
// message reply in a channel or a chat.
func (c JoinedTeamChannelMessageReplyClient) CreateJoinedTeamChannelMessageReplySoftDelete(ctx context.Context, id stable.UserIdJoinedTeamIdChannelIdMessageIdReplyId) (result CreateJoinedTeamChannelMessageReplySoftDeleteOperationResponse, err error) {
	opts := client.RequestOptions{
		ContentType: "application/json; charset=utf-8",
		ExpectedStatusCodes: []int{
			http.StatusNoContent,
		},
		HttpMethod: http.MethodPost,
		Path:       fmt.Sprintf("%s/softDelete", id.ID()),
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
