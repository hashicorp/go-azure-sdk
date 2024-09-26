package onenotepage

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

type CreateOnenotePageOnenotePatchContentOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
}

type CreateOnenotePageOnenotePatchContentOperationOptions struct {
	Metadata  *odata.Metadata
	RetryFunc client.RequestRetryFunc
}

func DefaultCreateOnenotePageOnenotePatchContentOperationOptions() CreateOnenotePageOnenotePatchContentOperationOptions {
	return CreateOnenotePageOnenotePatchContentOperationOptions{}
}

func (o CreateOnenotePageOnenotePatchContentOperationOptions) ToHeaders() *client.Headers {
	out := client.Headers{}

	return &out
}

func (o CreateOnenotePageOnenotePatchContentOperationOptions) ToOData() *odata.Query {
	out := odata.Query{}
	if o.Metadata != nil {
		out.Metadata = *o.Metadata
	}
	return &out
}

func (o CreateOnenotePageOnenotePatchContentOperationOptions) ToQuery() *client.QueryParams {
	out := client.QueryParams{}

	return &out
}

// CreateOnenotePageOnenotePatchContent - Invoke action onenotePatchContent
func (c OnenotePageClient) CreateOnenotePageOnenotePatchContent(ctx context.Context, id stable.GroupIdOnenotePageId, input CreateOnenotePageOnenotePatchContentRequest, options CreateOnenotePageOnenotePatchContentOperationOptions) (result CreateOnenotePageOnenotePatchContentOperationResponse, err error) {
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
		Path:          fmt.Sprintf("%s/onenotePatchContent", id.ID()),
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

	return
}
