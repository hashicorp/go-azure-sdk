package onenotesectiongroupsectionpage

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

type CreateOnenoteSectionGroupSectionPageOnenotePatchContentOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
}

type CreateOnenoteSectionGroupSectionPageOnenotePatchContentOperationOptions struct {
	Metadata  *odata.Metadata
	RetryFunc client.RequestRetryFunc
}

func DefaultCreateOnenoteSectionGroupSectionPageOnenotePatchContentOperationOptions() CreateOnenoteSectionGroupSectionPageOnenotePatchContentOperationOptions {
	return CreateOnenoteSectionGroupSectionPageOnenotePatchContentOperationOptions{}
}

func (o CreateOnenoteSectionGroupSectionPageOnenotePatchContentOperationOptions) ToHeaders() *client.Headers {
	out := client.Headers{}

	return &out
}

func (o CreateOnenoteSectionGroupSectionPageOnenotePatchContentOperationOptions) ToOData() *odata.Query {
	out := odata.Query{}
	if o.Metadata != nil {
		out.Metadata = *o.Metadata
	}
	return &out
}

func (o CreateOnenoteSectionGroupSectionPageOnenotePatchContentOperationOptions) ToQuery() *client.QueryParams {
	out := client.QueryParams{}

	return &out
}

// CreateOnenoteSectionGroupSectionPageOnenotePatchContent - Invoke action onenotePatchContent
func (c OnenoteSectionGroupSectionPageClient) CreateOnenoteSectionGroupSectionPageOnenotePatchContent(ctx context.Context, id beta.GroupIdOnenoteSectionGroupIdSectionIdPageId, input CreateOnenoteSectionGroupSectionPageOnenotePatchContentRequest, options CreateOnenoteSectionGroupSectionPageOnenotePatchContentOperationOptions) (result CreateOnenoteSectionGroupSectionPageOnenotePatchContentOperationResponse, err error) {
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
