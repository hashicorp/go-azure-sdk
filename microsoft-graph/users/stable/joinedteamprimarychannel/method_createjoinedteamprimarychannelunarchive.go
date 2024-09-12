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

type CreateJoinedTeamPrimaryChannelUnarchiveOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
}

// CreateJoinedTeamPrimaryChannelUnarchive - Invoke action unarchive. Restore an archived channel. Unarchiving restores
// the ability for users to send messages and edit the channel. Channels are archived via the channel: archive method.
// Unarchiving is an asynchronous operation; a channel is unarchived when the asynchronous unarchiving operation
// completes successfully, which might occur after this method responds.
func (c JoinedTeamPrimaryChannelClient) CreateJoinedTeamPrimaryChannelUnarchive(ctx context.Context, id stable.UserIdJoinedTeamId) (result CreateJoinedTeamPrimaryChannelUnarchiveOperationResponse, err error) {
	opts := client.RequestOptions{
		ContentType: "application/json; charset=utf-8",
		ExpectedStatusCodes: []int{
			http.StatusNoContent,
		},
		HttpMethod: http.MethodPost,
		Path:       fmt.Sprintf("%s/primaryChannel/unarchive", id.ID()),
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
