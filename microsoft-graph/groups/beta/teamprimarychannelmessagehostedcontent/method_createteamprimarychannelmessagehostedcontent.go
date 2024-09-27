package teamprimarychannelmessagehostedcontent

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

type CreateTeamPrimaryChannelMessageHostedContentOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
	Model        *beta.ChatMessageHostedContent
}

type CreateTeamPrimaryChannelMessageHostedContentOperationOptions struct {
	Metadata  *odata.Metadata
	RetryFunc client.RequestRetryFunc
}

func DefaultCreateTeamPrimaryChannelMessageHostedContentOperationOptions() CreateTeamPrimaryChannelMessageHostedContentOperationOptions {
	return CreateTeamPrimaryChannelMessageHostedContentOperationOptions{}
}

func (o CreateTeamPrimaryChannelMessageHostedContentOperationOptions) ToHeaders() *client.Headers {
	out := client.Headers{}

	return &out
}

func (o CreateTeamPrimaryChannelMessageHostedContentOperationOptions) ToOData() *odata.Query {
	out := odata.Query{}
	if o.Metadata != nil {
		out.Metadata = *o.Metadata
	}
	return &out
}

func (o CreateTeamPrimaryChannelMessageHostedContentOperationOptions) ToQuery() *client.QueryParams {
	out := client.QueryParams{}

	return &out
}

// CreateTeamPrimaryChannelMessageHostedContent - Create new navigation property to hostedContents for groups
func (c TeamPrimaryChannelMessageHostedContentClient) CreateTeamPrimaryChannelMessageHostedContent(ctx context.Context, id beta.GroupIdTeamPrimaryChannelMessageId, input beta.ChatMessageHostedContent, options CreateTeamPrimaryChannelMessageHostedContentOperationOptions) (result CreateTeamPrimaryChannelMessageHostedContentOperationResponse, err error) {
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
		Path:          fmt.Sprintf("%s/hostedContents", id.ID()),
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

	var model beta.ChatMessageHostedContent
	result.Model = &model
	if err = resp.Unmarshal(result.Model); err != nil {
		return
	}

	return
}
