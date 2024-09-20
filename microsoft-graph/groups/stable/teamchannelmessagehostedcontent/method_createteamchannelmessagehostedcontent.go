package teamchannelmessagehostedcontent

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

type CreateTeamChannelMessageHostedContentOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
	Model        *stable.ChatMessageHostedContent
}

type CreateTeamChannelMessageHostedContentOperationOptions struct {
	Metadata *odata.Metadata
}

func DefaultCreateTeamChannelMessageHostedContentOperationOptions() CreateTeamChannelMessageHostedContentOperationOptions {
	return CreateTeamChannelMessageHostedContentOperationOptions{}
}

func (o CreateTeamChannelMessageHostedContentOperationOptions) ToHeaders() *client.Headers {
	out := client.Headers{}

	return &out
}

func (o CreateTeamChannelMessageHostedContentOperationOptions) ToOData() *odata.Query {
	out := odata.Query{}
	if o.Metadata != nil {
		out.Metadata = *o.Metadata
	}
	return &out
}

func (o CreateTeamChannelMessageHostedContentOperationOptions) ToQuery() *client.QueryParams {
	out := client.QueryParams{}

	return &out
}

// CreateTeamChannelMessageHostedContent - Create new navigation property to hostedContents for groups
func (c TeamChannelMessageHostedContentClient) CreateTeamChannelMessageHostedContent(ctx context.Context, id stable.GroupIdTeamChannelIdMessageId, input stable.ChatMessageHostedContent, options CreateTeamChannelMessageHostedContentOperationOptions) (result CreateTeamChannelMessageHostedContentOperationResponse, err error) {
	opts := client.RequestOptions{
		ContentType: "application/json; charset=utf-8",
		ExpectedStatusCodes: []int{
			http.StatusCreated,
		},
		HttpMethod:    http.MethodPost,
		OptionsObject: options,
		Path:          fmt.Sprintf("%s/hostedContents", id.ID()),
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

	var model stable.ChatMessageHostedContent
	result.Model = &model
	if err = resp.Unmarshal(result.Model); err != nil {
		return
	}

	return
}
