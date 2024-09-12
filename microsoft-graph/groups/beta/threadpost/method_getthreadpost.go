package threadpost

import (
	"context"
	"net/http"

	"github.com/hashicorp/go-azure-sdk/microsoft-graph/common-types/beta"
	"github.com/hashicorp/go-azure-sdk/sdk/client"
	"github.com/hashicorp/go-azure-sdk/sdk/odata"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type GetThreadPostOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
	Model        *beta.Post
}

type GetThreadPostOperationOptions struct {
	Expand *odata.Expand
	Select *[]string
}

func DefaultGetThreadPostOperationOptions() GetThreadPostOperationOptions {
	return GetThreadPostOperationOptions{}
}

func (o GetThreadPostOperationOptions) ToHeaders() *client.Headers {
	out := client.Headers{}

	return &out
}

func (o GetThreadPostOperationOptions) ToOData() *odata.Query {
	out := odata.Query{}
	if o.Expand != nil {
		out.Expand = *o.Expand
	}
	if o.Select != nil {
		out.Select = *o.Select
	}
	return &out
}

func (o GetThreadPostOperationOptions) ToQuery() *client.QueryParams {
	out := client.QueryParams{}

	return &out
}

// GetThreadPost - Get post. Get the properties and relationships of a post in a specified thread. You can specify both
// the parent conversation and the thread, or, you can specify the thread without referencing the parent conversation.
// Since the post resource supports extensions, you can also use the GET operation to get custom properties and
// extension data in a post instance.
func (c ThreadPostClient) GetThreadPost(ctx context.Context, id beta.GroupIdThreadIdPostId, options GetThreadPostOperationOptions) (result GetThreadPostOperationResponse, err error) {
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

	var model beta.Post
	result.Model = &model
	if err = resp.Unmarshal(result.Model); err != nil {
		return
	}

	return
}
