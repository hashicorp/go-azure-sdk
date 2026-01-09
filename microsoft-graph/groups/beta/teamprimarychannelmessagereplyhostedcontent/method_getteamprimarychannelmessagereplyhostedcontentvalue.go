package teamprimarychannelmessagereplyhostedcontent

import (
	"context"
	"fmt"
	"net/http"

	"github.com/hashicorp/go-azure-sdk/microsoft-graph/common-types/beta"
	"github.com/hashicorp/go-azure-sdk/sdk/client"
	"github.com/hashicorp/go-azure-sdk/sdk/odata"
)

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type GetTeamPrimaryChannelMessageReplyHostedContentValueOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
	Model        *[]byte
}

type GetTeamPrimaryChannelMessageReplyHostedContentValueOperationOptions struct {
	Metadata  *odata.Metadata
	RetryFunc client.RequestRetryFunc
}

func DefaultGetTeamPrimaryChannelMessageReplyHostedContentValueOperationOptions() GetTeamPrimaryChannelMessageReplyHostedContentValueOperationOptions {
	return GetTeamPrimaryChannelMessageReplyHostedContentValueOperationOptions{}
}

func (o GetTeamPrimaryChannelMessageReplyHostedContentValueOperationOptions) ToHeaders() *client.Headers {
	out := client.Headers{}

	return &out
}

func (o GetTeamPrimaryChannelMessageReplyHostedContentValueOperationOptions) ToOData() *odata.Query {
	out := odata.Query{}
	if o.Metadata != nil {
		out.Metadata = *o.Metadata
	}
	return &out
}

func (o GetTeamPrimaryChannelMessageReplyHostedContentValueOperationOptions) ToQuery() *client.QueryParams {
	out := client.QueryParams{}

	return &out
}

// GetTeamPrimaryChannelMessageReplyHostedContentValue - Get media content for the navigation property hostedContents
// from groups. The unique identifier for an entity. Read-only.
func (c TeamPrimaryChannelMessageReplyHostedContentClient) GetTeamPrimaryChannelMessageReplyHostedContentValue(ctx context.Context, id beta.GroupIdTeamPrimaryChannelMessageIdReplyIdHostedContentId, options GetTeamPrimaryChannelMessageReplyHostedContentValueOperationOptions) (result GetTeamPrimaryChannelMessageReplyHostedContentValueOperationResponse, err error) {
	opts := client.RequestOptions{
		ContentType: "application/octet-stream",
		ExpectedStatusCodes: []int{
			http.StatusOK,
		},
		HttpMethod:    http.MethodGet,
		OptionsObject: options,
		Path:          fmt.Sprintf("%s/$value", id.ID()),
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

	var model []byte
	result.Model = &model
	if err = resp.Unmarshal(result.Model); err != nil {
		return
	}

	return
}
