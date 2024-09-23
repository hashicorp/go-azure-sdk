package onenotesectionpage

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

type CreateOnenoteSectionPageOnenotePatchContentOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
}

type CreateOnenoteSectionPageOnenotePatchContentOperationOptions struct {
	Metadata  *odata.Metadata
	RetryFunc client.RequestRetryFunc
}

func DefaultCreateOnenoteSectionPageOnenotePatchContentOperationOptions() CreateOnenoteSectionPageOnenotePatchContentOperationOptions {
	return CreateOnenoteSectionPageOnenotePatchContentOperationOptions{}
}

func (o CreateOnenoteSectionPageOnenotePatchContentOperationOptions) ToHeaders() *client.Headers {
	out := client.Headers{}

	return &out
}

func (o CreateOnenoteSectionPageOnenotePatchContentOperationOptions) ToOData() *odata.Query {
	out := odata.Query{}
	if o.Metadata != nil {
		out.Metadata = *o.Metadata
	}
	return &out
}

func (o CreateOnenoteSectionPageOnenotePatchContentOperationOptions) ToQuery() *client.QueryParams {
	out := client.QueryParams{}

	return &out
}

// CreateOnenoteSectionPageOnenotePatchContent - Invoke action onenotePatchContent
func (c OnenoteSectionPageClient) CreateOnenoteSectionPageOnenotePatchContent(ctx context.Context, id stable.GroupIdOnenoteSectionIdPageId, input CreateOnenoteSectionPageOnenotePatchContentRequest, options CreateOnenoteSectionPageOnenotePatchContentOperationOptions) (result CreateOnenoteSectionPageOnenotePatchContentOperationResponse, err error) {
	opts := client.RequestOptions{
		ContentType: "application/json; charset=utf-8",
		ExpectedStatusCodes: []int{
			http.StatusNoContent,
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
