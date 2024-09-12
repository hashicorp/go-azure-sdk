package teamprimarychannelmessagereplyhostedcontent

import (
	"context"
	"net/http"

	"github.com/hashicorp/go-azure-sdk/microsoft-graph/common-types/beta"
	"github.com/hashicorp/go-azure-sdk/sdk/client"
	"github.com/hashicorp/go-azure-sdk/sdk/odata"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type GetTeamPrimaryChannelMessageReplyHostedContentOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
	Model        *beta.ChatMessageHostedContent
}

type GetTeamPrimaryChannelMessageReplyHostedContentOperationOptions struct {
	Expand *odata.Expand
	Select *[]string
}

func DefaultGetTeamPrimaryChannelMessageReplyHostedContentOperationOptions() GetTeamPrimaryChannelMessageReplyHostedContentOperationOptions {
	return GetTeamPrimaryChannelMessageReplyHostedContentOperationOptions{}
}

func (o GetTeamPrimaryChannelMessageReplyHostedContentOperationOptions) ToHeaders() *client.Headers {
	out := client.Headers{}

	return &out
}

func (o GetTeamPrimaryChannelMessageReplyHostedContentOperationOptions) ToOData() *odata.Query {
	out := odata.Query{}
	if o.Expand != nil {
		out.Expand = *o.Expand
	}
	if o.Select != nil {
		out.Select = *o.Select
	}
	return &out
}

func (o GetTeamPrimaryChannelMessageReplyHostedContentOperationOptions) ToQuery() *client.QueryParams {
	out := client.QueryParams{}

	return &out
}

// GetTeamPrimaryChannelMessageReplyHostedContent - Get hostedContents from groups. Content in a message hosted by
// Microsoft Teams - for example, images or code snippets.
func (c TeamPrimaryChannelMessageReplyHostedContentClient) GetTeamPrimaryChannelMessageReplyHostedContent(ctx context.Context, id beta.GroupIdTeamPrimaryChannelMessageIdReplyIdHostedContentId, options GetTeamPrimaryChannelMessageReplyHostedContentOperationOptions) (result GetTeamPrimaryChannelMessageReplyHostedContentOperationResponse, err error) {
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

	var model beta.ChatMessageHostedContent
	result.Model = &model
	if err = resp.Unmarshal(result.Model); err != nil {
		return
	}

	return
}
