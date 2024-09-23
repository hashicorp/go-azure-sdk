package teamchannel

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

type RemoveTeamChannelEmailOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
}

type RemoveTeamChannelEmailOperationOptions struct {
	Metadata  *odata.Metadata
	RetryFunc client.RequestRetryFunc
}

func DefaultRemoveTeamChannelEmailOperationOptions() RemoveTeamChannelEmailOperationOptions {
	return RemoveTeamChannelEmailOperationOptions{}
}

func (o RemoveTeamChannelEmailOperationOptions) ToHeaders() *client.Headers {
	out := client.Headers{}

	return &out
}

func (o RemoveTeamChannelEmailOperationOptions) ToOData() *odata.Query {
	out := odata.Query{}
	if o.Metadata != nil {
		out.Metadata = *o.Metadata
	}
	return &out
}

func (o RemoveTeamChannelEmailOperationOptions) ToQuery() *client.QueryParams {
	out := client.QueryParams{}

	return &out
}

// RemoveTeamChannelEmail - Invoke action removeEmail. Remove the email address of a channel. You can remove an email
// address only if it was provisioned using the provisionEmail method or through the Microsoft Teams client.
func (c TeamChannelClient) RemoveTeamChannelEmail(ctx context.Context, id stable.GroupIdTeamChannelId, options RemoveTeamChannelEmailOperationOptions) (result RemoveTeamChannelEmailOperationResponse, err error) {
	opts := client.RequestOptions{
		ContentType: "application/json; charset=utf-8",
		ExpectedStatusCodes: []int{
			http.StatusNoContent,
		},
		HttpMethod:    http.MethodPost,
		OptionsObject: options,
		Path:          fmt.Sprintf("%s/removeEmail", id.ID()),
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
