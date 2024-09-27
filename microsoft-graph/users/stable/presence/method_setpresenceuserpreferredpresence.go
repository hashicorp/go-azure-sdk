package presence

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

type SetPresenceUserPreferredPresenceOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
}

type SetPresenceUserPreferredPresenceOperationOptions struct {
	Metadata  *odata.Metadata
	RetryFunc client.RequestRetryFunc
}

func DefaultSetPresenceUserPreferredPresenceOperationOptions() SetPresenceUserPreferredPresenceOperationOptions {
	return SetPresenceUserPreferredPresenceOperationOptions{}
}

func (o SetPresenceUserPreferredPresenceOperationOptions) ToHeaders() *client.Headers {
	out := client.Headers{}

	return &out
}

func (o SetPresenceUserPreferredPresenceOperationOptions) ToOData() *odata.Query {
	out := odata.Query{}
	if o.Metadata != nil {
		out.Metadata = *o.Metadata
	}
	return &out
}

func (o SetPresenceUserPreferredPresenceOperationOptions) ToQuery() *client.QueryParams {
	out := client.QueryParams{}

	return &out
}

// SetPresenceUserPreferredPresence - Invoke action setUserPreferredPresence. Set the preferred availability and
// activity status for a user. If the preferred presence of a user is set, the user's presence shows as the preferred
// status. Preferred presence takes effect only when at least one presence session exists for the user. Otherwise, the
// user's presence shows as Offline. A presence session is created as a result of a successful setPresence operation, or
// if the user is signed in on a Microsoft Teams client. For more details, see presence sessions and time-out and
// expiration.
func (c PresenceClient) SetPresenceUserPreferredPresence(ctx context.Context, id stable.UserId, input SetPresenceUserPreferredPresenceRequest, options SetPresenceUserPreferredPresenceOperationOptions) (result SetPresenceUserPreferredPresenceOperationResponse, err error) {
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
		Path:          fmt.Sprintf("%s/presence/setUserPreferredPresence", id.ID()),
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
