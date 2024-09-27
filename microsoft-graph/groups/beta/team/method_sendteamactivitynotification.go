package team

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

type SendTeamActivityNotificationOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
}

type SendTeamActivityNotificationOperationOptions struct {
	Metadata  *odata.Metadata
	RetryFunc client.RequestRetryFunc
}

func DefaultSendTeamActivityNotificationOperationOptions() SendTeamActivityNotificationOperationOptions {
	return SendTeamActivityNotificationOperationOptions{}
}

func (o SendTeamActivityNotificationOperationOptions) ToHeaders() *client.Headers {
	out := client.Headers{}

	return &out
}

func (o SendTeamActivityNotificationOperationOptions) ToOData() *odata.Query {
	out := odata.Query{}
	if o.Metadata != nil {
		out.Metadata = *o.Metadata
	}
	return &out
}

func (o SendTeamActivityNotificationOperationOptions) ToQuery() *client.QueryParams {
	out := client.QueryParams{}

	return &out
}

// SendTeamActivityNotification - Invoke action sendActivityNotification. Send an activity feed notification in the
// scope of a team. For more information, see sending Teams activity notifications.
func (c TeamClient) SendTeamActivityNotification(ctx context.Context, id beta.GroupId, input SendTeamActivityNotificationRequest, options SendTeamActivityNotificationOperationOptions) (result SendTeamActivityNotificationOperationResponse, err error) {
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
		Path:          fmt.Sprintf("%s/team/sendActivityNotification", id.ID()),
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
