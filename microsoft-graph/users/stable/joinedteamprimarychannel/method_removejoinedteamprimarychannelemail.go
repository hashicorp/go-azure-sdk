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

type RemoveJoinedTeamPrimaryChannelEmailOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
}

type RemoveJoinedTeamPrimaryChannelEmailOperationOptions struct {
	Metadata  *odata.Metadata
	RetryFunc client.RequestRetryFunc
}

func DefaultRemoveJoinedTeamPrimaryChannelEmailOperationOptions() RemoveJoinedTeamPrimaryChannelEmailOperationOptions {
	return RemoveJoinedTeamPrimaryChannelEmailOperationOptions{}
}

func (o RemoveJoinedTeamPrimaryChannelEmailOperationOptions) ToHeaders() *client.Headers {
	out := client.Headers{}

	return &out
}

func (o RemoveJoinedTeamPrimaryChannelEmailOperationOptions) ToOData() *odata.Query {
	out := odata.Query{}
	if o.Metadata != nil {
		out.Metadata = *o.Metadata
	}
	return &out
}

func (o RemoveJoinedTeamPrimaryChannelEmailOperationOptions) ToQuery() *client.QueryParams {
	out := client.QueryParams{}

	return &out
}

// RemoveJoinedTeamPrimaryChannelEmail - Invoke action removeEmail. Remove the email address of a channel. You can
// remove an email address only if it was provisioned using the provisionEmail method or through the Microsoft Teams
// client.
func (c JoinedTeamPrimaryChannelClient) RemoveJoinedTeamPrimaryChannelEmail(ctx context.Context, id stable.UserIdJoinedTeamId, options RemoveJoinedTeamPrimaryChannelEmailOperationOptions) (result RemoveJoinedTeamPrimaryChannelEmailOperationResponse, err error) {
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
		Path:          fmt.Sprintf("%s/primaryChannel/removeEmail", id.ID()),
		RetryFunc:     options.RetryFunc,
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
