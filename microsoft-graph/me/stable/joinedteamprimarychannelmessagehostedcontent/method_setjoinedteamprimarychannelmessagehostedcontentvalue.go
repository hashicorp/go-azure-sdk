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

type SetJoinedTeamPrimaryChannelMessageHostedContentValueOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
}

type SetJoinedTeamPrimaryChannelMessageHostedContentValueOperationOptions struct {
	Metadata *odata.Metadata
}

func DefaultSetJoinedTeamPrimaryChannelMessageHostedContentValueOperationOptions() SetJoinedTeamPrimaryChannelMessageHostedContentValueOperationOptions {
	return SetJoinedTeamPrimaryChannelMessageHostedContentValueOperationOptions{}
}

func (o SetJoinedTeamPrimaryChannelMessageHostedContentValueOperationOptions) ToHeaders() *client.Headers {
	out := client.Headers{}

	return &out
}

func (o SetJoinedTeamPrimaryChannelMessageHostedContentValueOperationOptions) ToOData() *odata.Query {
	out := odata.Query{}
	if o.Metadata != nil {
		out.Metadata = *o.Metadata
	}
	return &out
}

func (o SetJoinedTeamPrimaryChannelMessageHostedContentValueOperationOptions) ToQuery() *client.QueryParams {
	out := client.QueryParams{}

	return &out
}

// SetJoinedTeamPrimaryChannelMessageHostedContentValue - Update media content for the navigation property
// hostedContents in me. The unique identifier for an entity. Read-only.
func (c JoinedTeamPrimaryChannelMessageHostedContentClient) SetJoinedTeamPrimaryChannelMessageHostedContentValue(ctx context.Context, id stable.MeJoinedTeamIdPrimaryChannelMessageIdHostedContentId, input []byte, options SetJoinedTeamPrimaryChannelMessageHostedContentValueOperationOptions) (result SetJoinedTeamPrimaryChannelMessageHostedContentValueOperationResponse, err error) {
	opts := client.RequestOptions{
		ContentType: "application/json; charset=utf-8",
		ExpectedStatusCodes: []int{
			http.StatusNoContent,
		},
		HttpMethod:    http.MethodPut,
		OptionsObject: options,
		Path:          fmt.Sprintf("%s/$value", id.ID()),
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
