package presence

import (
	"context"
	"net/http"

	"github.com/hashicorp/go-azure-sdk/sdk/client"
	"github.com/hashicorp/go-azure-sdk/sdk/odata"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type ClearPresencePresenceOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
}

type ClearPresencePresenceOperationOptions struct {
	Metadata  *odata.Metadata
	RetryFunc client.RequestRetryFunc
}

func DefaultClearPresencePresenceOperationOptions() ClearPresencePresenceOperationOptions {
	return ClearPresencePresenceOperationOptions{}
}

func (o ClearPresencePresenceOperationOptions) ToHeaders() *client.Headers {
	out := client.Headers{}

	return &out
}

func (o ClearPresencePresenceOperationOptions) ToOData() *odata.Query {
	out := odata.Query{}
	if o.Metadata != nil {
		out.Metadata = *o.Metadata
	}
	return &out
}

func (o ClearPresencePresenceOperationOptions) ToQuery() *client.QueryParams {
	out := client.QueryParams{}

	return &out
}

// ClearPresencePresence - Invoke action clearPresence. Clear a presence session of an application for a user. If it is
// the user's only presence session, a successful clearPresence changes the user's presence to Offline/Offline. Read
// more about presence sessions and their time-out and expiration.
func (c PresenceClient) ClearPresencePresence(ctx context.Context, input ClearPresencePresenceRequest, options ClearPresencePresenceOperationOptions) (result ClearPresencePresenceOperationResponse, err error) {
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
		Path:          "/me/presence/clearPresence",
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
