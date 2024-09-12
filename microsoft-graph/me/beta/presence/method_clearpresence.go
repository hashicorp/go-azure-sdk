package presence

import (
	"context"
	"net/http"

	"github.com/hashicorp/go-azure-sdk/sdk/client"
	"github.com/hashicorp/go-azure-sdk/sdk/odata"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type ClearPresenceOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
}

// ClearPresence - Invoke action clearPresence. Clear a presence session of an application for a user. If it is the
// user's only presence session, a successful clearPresence changes the user's presence to Offline/Offline. Read more
// about presence sessions and their time-out and expiration.
func (c PresenceClient) ClearPresence(ctx context.Context, input ClearPresenceRequest) (result ClearPresenceOperationResponse, err error) {
	opts := client.RequestOptions{
		ContentType: "application/json; charset=utf-8",
		ExpectedStatusCodes: []int{
			http.StatusNoContent,
		},
		HttpMethod: http.MethodPost,
		Path:       "/me/presence/clearPresence",
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
