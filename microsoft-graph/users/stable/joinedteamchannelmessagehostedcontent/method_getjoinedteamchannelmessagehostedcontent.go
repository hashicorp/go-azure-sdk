package joinedteamchannelmessagehostedcontent

import (
	"context"
	"net/http"

	"github.com/hashicorp/go-azure-sdk/microsoft-graph/common-types/stable"
	"github.com/hashicorp/go-azure-sdk/sdk/client"
	"github.com/hashicorp/go-azure-sdk/sdk/odata"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type GetJoinedTeamChannelMessageHostedContentOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
	Model        *stable.ChatMessageHostedContent
}

type GetJoinedTeamChannelMessageHostedContentOperationOptions struct {
	Expand   *odata.Expand
	Metadata *odata.Metadata
	Select   *[]string
}

func DefaultGetJoinedTeamChannelMessageHostedContentOperationOptions() GetJoinedTeamChannelMessageHostedContentOperationOptions {
	return GetJoinedTeamChannelMessageHostedContentOperationOptions{}
}

func (o GetJoinedTeamChannelMessageHostedContentOperationOptions) ToHeaders() *client.Headers {
	out := client.Headers{}

	return &out
}

func (o GetJoinedTeamChannelMessageHostedContentOperationOptions) ToOData() *odata.Query {
	out := odata.Query{}
	if o.Expand != nil {
		out.Expand = *o.Expand
	}
	if o.Metadata != nil {
		out.Metadata = *o.Metadata
	}
	if o.Select != nil {
		out.Select = *o.Select
	}
	return &out
}

func (o GetJoinedTeamChannelMessageHostedContentOperationOptions) ToQuery() *client.QueryParams {
	out := client.QueryParams{}

	return &out
}

// GetJoinedTeamChannelMessageHostedContent - Get hostedContents from users. Content in a message hosted by Microsoft
// Teams - for example, images or code snippets.
func (c JoinedTeamChannelMessageHostedContentClient) GetJoinedTeamChannelMessageHostedContent(ctx context.Context, id stable.UserIdJoinedTeamIdChannelIdMessageIdHostedContentId, options GetJoinedTeamChannelMessageHostedContentOperationOptions) (result GetJoinedTeamChannelMessageHostedContentOperationResponse, err error) {
	opts := client.RequestOptions{
		ContentType: "application/json; charset=utf-8",
		ExpectedStatusCodes: []int{
			http.StatusOK,
		},
		HttpMethod:    http.MethodGet,
		OptionsObject: options,
		Path:          id.ID(),
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

	var model stable.ChatMessageHostedContent
	result.Model = &model
	if err = resp.Unmarshal(result.Model); err != nil {
		return
	}

	return
}
