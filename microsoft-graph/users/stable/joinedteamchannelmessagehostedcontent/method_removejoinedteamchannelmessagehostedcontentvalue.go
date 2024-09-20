package joinedteamchannelmessagehostedcontent

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

type RemoveJoinedTeamChannelMessageHostedContentValueOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
}

type RemoveJoinedTeamChannelMessageHostedContentValueOperationOptions struct {
	IfMatch  *string
	Metadata *odata.Metadata
}

func DefaultRemoveJoinedTeamChannelMessageHostedContentValueOperationOptions() RemoveJoinedTeamChannelMessageHostedContentValueOperationOptions {
	return RemoveJoinedTeamChannelMessageHostedContentValueOperationOptions{}
}

func (o RemoveJoinedTeamChannelMessageHostedContentValueOperationOptions) ToHeaders() *client.Headers {
	out := client.Headers{}
	if o.IfMatch != nil {
		out.Append("If-Match", fmt.Sprintf("%v", *o.IfMatch))
	}
	return &out
}

func (o RemoveJoinedTeamChannelMessageHostedContentValueOperationOptions) ToOData() *odata.Query {
	out := odata.Query{}
	if o.Metadata != nil {
		out.Metadata = *o.Metadata
	}
	return &out
}

func (o RemoveJoinedTeamChannelMessageHostedContentValueOperationOptions) ToQuery() *client.QueryParams {
	out := client.QueryParams{}

	return &out
}

// RemoveJoinedTeamChannelMessageHostedContentValue - Delete media content for the navigation property hostedContents in
// users. The unique identifier for an entity. Read-only.
func (c JoinedTeamChannelMessageHostedContentClient) RemoveJoinedTeamChannelMessageHostedContentValue(ctx context.Context, id stable.UserIdJoinedTeamIdChannelIdMessageIdHostedContentId, options RemoveJoinedTeamChannelMessageHostedContentValueOperationOptions) (result RemoveJoinedTeamChannelMessageHostedContentValueOperationResponse, err error) {
	opts := client.RequestOptions{
		ContentType: "application/json; charset=utf-8",
		ExpectedStatusCodes: []int{
			http.StatusNoContent,
		},
		HttpMethod:    http.MethodDelete,
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

	return
}
