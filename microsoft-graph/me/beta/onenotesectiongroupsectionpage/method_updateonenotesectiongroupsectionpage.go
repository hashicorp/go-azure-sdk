package onenotesectiongroupsectionpage

import (
	"context"
	"net/http"

	"github.com/hashicorp/go-azure-sdk/microsoft-graph/common-types/beta"
	"github.com/hashicorp/go-azure-sdk/sdk/client"
	"github.com/hashicorp/go-azure-sdk/sdk/odata"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type UpdateOnenoteSectionGroupSectionPageOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
}

type UpdateOnenoteSectionGroupSectionPageOperationOptions struct {
	Metadata  *odata.Metadata
	RetryFunc client.RequestRetryFunc
}

func DefaultUpdateOnenoteSectionGroupSectionPageOperationOptions() UpdateOnenoteSectionGroupSectionPageOperationOptions {
	return UpdateOnenoteSectionGroupSectionPageOperationOptions{}
}

func (o UpdateOnenoteSectionGroupSectionPageOperationOptions) ToHeaders() *client.Headers {
	out := client.Headers{}

	return &out
}

func (o UpdateOnenoteSectionGroupSectionPageOperationOptions) ToOData() *odata.Query {
	out := odata.Query{}
	if o.Metadata != nil {
		out.Metadata = *o.Metadata
	}
	return &out
}

func (o UpdateOnenoteSectionGroupSectionPageOperationOptions) ToQuery() *client.QueryParams {
	out := client.QueryParams{}

	return &out
}

// UpdateOnenoteSectionGroupSectionPage - Update the navigation property pages in me
func (c OnenoteSectionGroupSectionPageClient) UpdateOnenoteSectionGroupSectionPage(ctx context.Context, id beta.MeOnenoteSectionGroupIdSectionIdPageId, input beta.OnenotePage, options UpdateOnenoteSectionGroupSectionPageOperationOptions) (result UpdateOnenoteSectionGroupSectionPageOperationResponse, err error) {
	opts := client.RequestOptions{
		ContentType: "application/json; charset=utf-8",
		ExpectedStatusCodes: []int{
			http.StatusNoContent,
		},
		HttpMethod:    http.MethodPatch,
		OptionsObject: options,
		Path:          id.ID(),
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
