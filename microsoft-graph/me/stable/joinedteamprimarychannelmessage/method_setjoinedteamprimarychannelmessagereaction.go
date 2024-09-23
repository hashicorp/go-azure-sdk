package joinedteamprimarychannelmessage

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

type SetJoinedTeamPrimaryChannelMessageReactionOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
}

type SetJoinedTeamPrimaryChannelMessageReactionOperationOptions struct {
	Metadata  *odata.Metadata
	RetryFunc client.RequestRetryFunc
}

func DefaultSetJoinedTeamPrimaryChannelMessageReactionOperationOptions() SetJoinedTeamPrimaryChannelMessageReactionOperationOptions {
	return SetJoinedTeamPrimaryChannelMessageReactionOperationOptions{}
}

func (o SetJoinedTeamPrimaryChannelMessageReactionOperationOptions) ToHeaders() *client.Headers {
	out := client.Headers{}

	return &out
}

func (o SetJoinedTeamPrimaryChannelMessageReactionOperationOptions) ToOData() *odata.Query {
	out := odata.Query{}
	if o.Metadata != nil {
		out.Metadata = *o.Metadata
	}
	return &out
}

func (o SetJoinedTeamPrimaryChannelMessageReactionOperationOptions) ToQuery() *client.QueryParams {
	out := client.QueryParams{}

	return &out
}

// SetJoinedTeamPrimaryChannelMessageReaction - Invoke action setReaction
func (c JoinedTeamPrimaryChannelMessageClient) SetJoinedTeamPrimaryChannelMessageReaction(ctx context.Context, id stable.MeJoinedTeamIdPrimaryChannelMessageId, input SetJoinedTeamPrimaryChannelMessageReactionRequest, options SetJoinedTeamPrimaryChannelMessageReactionOperationOptions) (result SetJoinedTeamPrimaryChannelMessageReactionOperationResponse, err error) {
	opts := client.RequestOptions{
		ContentType: "application/json; charset=utf-8",
		ExpectedStatusCodes: []int{
			http.StatusNoContent,
		},
		HttpMethod:    http.MethodPost,
		OptionsObject: options,
		Path:          fmt.Sprintf("%s/setReaction", id.ID()),
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
