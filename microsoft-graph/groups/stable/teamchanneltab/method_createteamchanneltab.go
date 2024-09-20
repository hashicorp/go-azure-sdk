package teamchanneltab

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

type CreateTeamChannelTabOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
	Model        *stable.TeamsTab
}

type CreateTeamChannelTabOperationOptions struct {
	Metadata *odata.Metadata
}

func DefaultCreateTeamChannelTabOperationOptions() CreateTeamChannelTabOperationOptions {
	return CreateTeamChannelTabOperationOptions{}
}

func (o CreateTeamChannelTabOperationOptions) ToHeaders() *client.Headers {
	out := client.Headers{}

	return &out
}

func (o CreateTeamChannelTabOperationOptions) ToOData() *odata.Query {
	out := odata.Query{}
	if o.Metadata != nil {
		out.Metadata = *o.Metadata
	}
	return &out
}

func (o CreateTeamChannelTabOperationOptions) ToQuery() *client.QueryParams {
	out := client.QueryParams{}

	return &out
}

// CreateTeamChannelTab - Create new navigation property to tabs for groups
func (c TeamChannelTabClient) CreateTeamChannelTab(ctx context.Context, id stable.GroupIdTeamChannelId, input stable.TeamsTab, options CreateTeamChannelTabOperationOptions) (result CreateTeamChannelTabOperationResponse, err error) {
	opts := client.RequestOptions{
		ContentType: "application/json; charset=utf-8",
		ExpectedStatusCodes: []int{
			http.StatusCreated,
		},
		HttpMethod:    http.MethodPost,
		OptionsObject: options,
		Path:          fmt.Sprintf("%s/tabs", id.ID()),
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

	var model stable.TeamsTab
	result.Model = &model
	if err = resp.Unmarshal(result.Model); err != nil {
		return
	}

	return
}
