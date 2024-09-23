package teamwork

import (
	"context"
	"net/http"

	"github.com/hashicorp/go-azure-sdk/sdk/client"
	"github.com/hashicorp/go-azure-sdk/sdk/odata"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type SendTeamworkActivityNotificationOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
}

type SendTeamworkActivityNotificationOperationOptions struct {
	Metadata  *odata.Metadata
	RetryFunc client.RequestRetryFunc
}

func DefaultSendTeamworkActivityNotificationOperationOptions() SendTeamworkActivityNotificationOperationOptions {
	return SendTeamworkActivityNotificationOperationOptions{}
}

func (o SendTeamworkActivityNotificationOperationOptions) ToHeaders() *client.Headers {
	out := client.Headers{}

	return &out
}

func (o SendTeamworkActivityNotificationOperationOptions) ToOData() *odata.Query {
	out := odata.Query{}
	if o.Metadata != nil {
		out.Metadata = *o.Metadata
	}
	return &out
}

func (o SendTeamworkActivityNotificationOperationOptions) ToQuery() *client.QueryParams {
	out := client.QueryParams{}

	return &out
}

// SendTeamworkActivityNotification - Invoke action sendActivityNotification. Send an activity feed notification to a
// user. For more information, see sending Teams activity notifications.
func (c TeamworkClient) SendTeamworkActivityNotification(ctx context.Context, input SendTeamworkActivityNotificationRequest, options SendTeamworkActivityNotificationOperationOptions) (result SendTeamworkActivityNotificationOperationResponse, err error) {
	opts := client.RequestOptions{
		ContentType: "application/json; charset=utf-8",
		ExpectedStatusCodes: []int{
			http.StatusNoContent,
		},
		HttpMethod:    http.MethodPost,
		OptionsObject: options,
		Path:          "/me/teamwork/sendActivityNotification",
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
