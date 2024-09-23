package teamprimarychanneltab

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

type CreateTeamPrimaryChannelTabOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
	Model        *beta.TeamsTab
}

type CreateTeamPrimaryChannelTabOperationOptions struct {
	Metadata  *odata.Metadata
	RetryFunc client.RequestRetryFunc
}

func DefaultCreateTeamPrimaryChannelTabOperationOptions() CreateTeamPrimaryChannelTabOperationOptions {
	return CreateTeamPrimaryChannelTabOperationOptions{}
}

func (o CreateTeamPrimaryChannelTabOperationOptions) ToHeaders() *client.Headers {
	out := client.Headers{}

	return &out
}

func (o CreateTeamPrimaryChannelTabOperationOptions) ToOData() *odata.Query {
	out := odata.Query{}
	if o.Metadata != nil {
		out.Metadata = *o.Metadata
	}
	return &out
}

func (o CreateTeamPrimaryChannelTabOperationOptions) ToQuery() *client.QueryParams {
	out := client.QueryParams{}

	return &out
}

// CreateTeamPrimaryChannelTab - Create new navigation property to tabs for groups
func (c TeamPrimaryChannelTabClient) CreateTeamPrimaryChannelTab(ctx context.Context, id beta.GroupId, input beta.TeamsTab, options CreateTeamPrimaryChannelTabOperationOptions) (result CreateTeamPrimaryChannelTabOperationResponse, err error) {
	opts := client.RequestOptions{
		ContentType: "application/json; charset=utf-8",
		ExpectedStatusCodes: []int{
			http.StatusCreated,
		},
		HttpMethod:    http.MethodPost,
		OptionsObject: options,
		Path:          fmt.Sprintf("%s/team/primaryChannel/tabs", id.ID()),
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

	var model beta.TeamsTab
	result.Model = &model
	if err = resp.Unmarshal(result.Model); err != nil {
		return
	}

	return
}
