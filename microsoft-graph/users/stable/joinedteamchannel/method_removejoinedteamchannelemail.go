package joinedteamchannel

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

type RemoveJoinedTeamChannelEmailOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
}

// RemoveJoinedTeamChannelEmail - Invoke action removeEmail. Remove the email address of a channel. You can remove an
// email address only if it was provisioned using the provisionEmail method or through the Microsoft Teams client.
func (c JoinedTeamChannelClient) RemoveJoinedTeamChannelEmail(ctx context.Context, id stable.UserIdJoinedTeamIdChannelId) (result RemoveJoinedTeamChannelEmailOperationResponse, err error) {
	opts := client.RequestOptions{
		ContentType: "application/json; charset=utf-8",
		ExpectedStatusCodes: []int{
			http.StatusNoContent,
		},
		HttpMethod: http.MethodPost,
		Path:       fmt.Sprintf("%s/removeEmail", id.ID()),
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
