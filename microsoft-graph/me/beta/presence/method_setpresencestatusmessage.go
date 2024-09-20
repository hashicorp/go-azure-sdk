package presence

import (
	"context"
	"net/http"

	"github.com/hashicorp/go-azure-sdk/sdk/client"
	"github.com/hashicorp/go-azure-sdk/sdk/odata"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type SetPresenceStatusMessageOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
}

type SetPresenceStatusMessageOperationOptions struct {
	Metadata *odata.Metadata
}

func DefaultSetPresenceStatusMessageOperationOptions() SetPresenceStatusMessageOperationOptions {
	return SetPresenceStatusMessageOperationOptions{}
}

func (o SetPresenceStatusMessageOperationOptions) ToHeaders() *client.Headers {
	out := client.Headers{}

	return &out
}

func (o SetPresenceStatusMessageOperationOptions) ToOData() *odata.Query {
	out := odata.Query{}
	if o.Metadata != nil {
		out.Metadata = *o.Metadata
	}
	return &out
}

func (o SetPresenceStatusMessageOperationOptions) ToQuery() *client.QueryParams {
	out := client.QueryParams{}

	return &out
}

// SetPresenceStatusMessage - Invoke action setStatusMessage. Set a presence status message for a user. An optional
// expiration date and time can be supplied.
func (c PresenceClient) SetPresenceStatusMessage(ctx context.Context, input SetPresenceStatusMessageRequest, options SetPresenceStatusMessageOperationOptions) (result SetPresenceStatusMessageOperationResponse, err error) {
	opts := client.RequestOptions{
		ContentType: "application/json; charset=utf-8",
		ExpectedStatusCodes: []int{
			http.StatusNoContent,
		},
		HttpMethod:    http.MethodPost,
		OptionsObject: options,
		Path:          "/me/presence/setStatusMessage",
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
