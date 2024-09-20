package joinedteamprimarychannelmessagehostedcontent

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

type GetJoinedTeamPrimaryChannelMessageHostedContentValueOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
	Model        *[]byte
}

type GetJoinedTeamPrimaryChannelMessageHostedContentValueOperationOptions struct {
	Metadata *odata.Metadata
}

func DefaultGetJoinedTeamPrimaryChannelMessageHostedContentValueOperationOptions() GetJoinedTeamPrimaryChannelMessageHostedContentValueOperationOptions {
	return GetJoinedTeamPrimaryChannelMessageHostedContentValueOperationOptions{}
}

func (o GetJoinedTeamPrimaryChannelMessageHostedContentValueOperationOptions) ToHeaders() *client.Headers {
	out := client.Headers{}

	return &out
}

func (o GetJoinedTeamPrimaryChannelMessageHostedContentValueOperationOptions) ToOData() *odata.Query {
	out := odata.Query{}
	if o.Metadata != nil {
		out.Metadata = *o.Metadata
	}
	return &out
}

func (o GetJoinedTeamPrimaryChannelMessageHostedContentValueOperationOptions) ToQuery() *client.QueryParams {
	out := client.QueryParams{}

	return &out
}

// GetJoinedTeamPrimaryChannelMessageHostedContentValue - Get media content for the navigation property hostedContents
// from me. The unique identifier for an entity. Read-only.
func (c JoinedTeamPrimaryChannelMessageHostedContentClient) GetJoinedTeamPrimaryChannelMessageHostedContentValue(ctx context.Context, id stable.MeJoinedTeamIdPrimaryChannelMessageIdHostedContentId, options GetJoinedTeamPrimaryChannelMessageHostedContentValueOperationOptions) (result GetJoinedTeamPrimaryChannelMessageHostedContentValueOperationResponse, err error) {
	opts := client.RequestOptions{
		ContentType: "application/octet-stream",
		ExpectedStatusCodes: []int{
			http.StatusOK,
		},
		HttpMethod:    http.MethodGet,
		OptionsObject: options,
		Path:          fmt.Sprintf("%s/$value", id.ID()),
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
