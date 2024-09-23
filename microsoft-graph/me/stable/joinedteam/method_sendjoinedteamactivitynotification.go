package joinedteam

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

type SendJoinedTeamActivityNotificationOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
}

type SendJoinedTeamActivityNotificationOperationOptions struct {
	Metadata  *odata.Metadata
	RetryFunc client.RequestRetryFunc
}

func DefaultSendJoinedTeamActivityNotificationOperationOptions() SendJoinedTeamActivityNotificationOperationOptions {
	return SendJoinedTeamActivityNotificationOperationOptions{}
}

func (o SendJoinedTeamActivityNotificationOperationOptions) ToHeaders() *client.Headers {
	out := client.Headers{}

	return &out
}

func (o SendJoinedTeamActivityNotificationOperationOptions) ToOData() *odata.Query {
	out := odata.Query{}
	if o.Metadata != nil {
		out.Metadata = *o.Metadata
	}
	return &out
}

func (o SendJoinedTeamActivityNotificationOperationOptions) ToQuery() *client.QueryParams {
	out := client.QueryParams{}

	return &out
}

// SendJoinedTeamActivityNotification - Invoke action sendActivityNotification. Send an activity feed notification in
// the scope of a team. For more information about sending notifications and the requirements for doing so, see sending
// Teams activity notifications.
func (c JoinedTeamClient) SendJoinedTeamActivityNotification(ctx context.Context, id stable.MeJoinedTeamId, input SendJoinedTeamActivityNotificationRequest, options SendJoinedTeamActivityNotificationOperationOptions) (result SendJoinedTeamActivityNotificationOperationResponse, err error) {
	opts := client.RequestOptions{
		ContentType: "application/json; charset=utf-8",
		ExpectedStatusCodes: []int{
			http.StatusNoContent,
		},
		HttpMethod:    http.MethodPost,
		OptionsObject: options,
		Path:          fmt.Sprintf("%s/sendActivityNotification", id.ID()),
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
