package teamprimarychannelmessagehostedcontent

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

type SetTeamPrimaryChannelMessageHostedContentValueOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
}

type SetTeamPrimaryChannelMessageHostedContentValueOperationOptions struct {
	Metadata *odata.Metadata
}

func DefaultSetTeamPrimaryChannelMessageHostedContentValueOperationOptions() SetTeamPrimaryChannelMessageHostedContentValueOperationOptions {
	return SetTeamPrimaryChannelMessageHostedContentValueOperationOptions{}
}

func (o SetTeamPrimaryChannelMessageHostedContentValueOperationOptions) ToHeaders() *client.Headers {
	out := client.Headers{}

	return &out
}

func (o SetTeamPrimaryChannelMessageHostedContentValueOperationOptions) ToOData() *odata.Query {
	out := odata.Query{}
	if o.Metadata != nil {
		out.Metadata = *o.Metadata
	}
	return &out
}

func (o SetTeamPrimaryChannelMessageHostedContentValueOperationOptions) ToQuery() *client.QueryParams {
	out := client.QueryParams{}

	return &out
}

// SetTeamPrimaryChannelMessageHostedContentValue - Update media content for the navigation property hostedContents in
// groups. The unique identifier for an entity. Read-only.
func (c TeamPrimaryChannelMessageHostedContentClient) SetTeamPrimaryChannelMessageHostedContentValue(ctx context.Context, id stable.GroupIdTeamPrimaryChannelMessageIdHostedContentId, input []byte, options SetTeamPrimaryChannelMessageHostedContentValueOperationOptions) (result SetTeamPrimaryChannelMessageHostedContentValueOperationResponse, err error) {
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
