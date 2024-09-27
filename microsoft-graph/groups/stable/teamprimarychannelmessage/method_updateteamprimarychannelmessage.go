package teamprimarychannelmessage

import (
	"context"
	"net/http"

	"github.com/hashicorp/go-azure-sdk/microsoft-graph/common-types/stable"
	"github.com/hashicorp/go-azure-sdk/sdk/client"
	"github.com/hashicorp/go-azure-sdk/sdk/odata"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type UpdateTeamPrimaryChannelMessageOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
}

type UpdateTeamPrimaryChannelMessageOperationOptions struct {
	Metadata  *odata.Metadata
	RetryFunc client.RequestRetryFunc
}

func DefaultUpdateTeamPrimaryChannelMessageOperationOptions() UpdateTeamPrimaryChannelMessageOperationOptions {
	return UpdateTeamPrimaryChannelMessageOperationOptions{}
}

func (o UpdateTeamPrimaryChannelMessageOperationOptions) ToHeaders() *client.Headers {
	out := client.Headers{}

	return &out
}

func (o UpdateTeamPrimaryChannelMessageOperationOptions) ToOData() *odata.Query {
	out := odata.Query{}
	if o.Metadata != nil {
		out.Metadata = *o.Metadata
	}
	return &out
}

func (o UpdateTeamPrimaryChannelMessageOperationOptions) ToQuery() *client.QueryParams {
	out := client.QueryParams{}

	return &out
}

// UpdateTeamPrimaryChannelMessage - Update the navigation property messages in groups
func (c TeamPrimaryChannelMessageClient) UpdateTeamPrimaryChannelMessage(ctx context.Context, id stable.GroupIdTeamPrimaryChannelMessageId, input stable.ChatMessage, options UpdateTeamPrimaryChannelMessageOperationOptions) (result UpdateTeamPrimaryChannelMessageOperationResponse, err error) {
	opts := client.RequestOptions{
		ContentType: "application/json; charset=utf-8",
		ExpectedStatusCodes: []int{
			http.StatusAccepted,
			http.StatusNoContent,
			http.StatusOK,
		},
		HttpMethod:    http.MethodPatch,
		OptionsObject: options,
		Path:          id.ID(),
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
