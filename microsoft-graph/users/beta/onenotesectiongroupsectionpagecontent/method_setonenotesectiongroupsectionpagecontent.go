package onenotesectiongroupsectionpagecontent

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

type SetOnenoteSectionGroupSectionPageContentOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
}

type SetOnenoteSectionGroupSectionPageContentOperationOptions struct {
	Metadata *odata.Metadata
}

func DefaultSetOnenoteSectionGroupSectionPageContentOperationOptions() SetOnenoteSectionGroupSectionPageContentOperationOptions {
	return SetOnenoteSectionGroupSectionPageContentOperationOptions{}
}

func (o SetOnenoteSectionGroupSectionPageContentOperationOptions) ToHeaders() *client.Headers {
	out := client.Headers{}

	return &out
}

func (o SetOnenoteSectionGroupSectionPageContentOperationOptions) ToOData() *odata.Query {
	out := odata.Query{}
	if o.Metadata != nil {
		out.Metadata = *o.Metadata
	}
	return &out
}

func (o SetOnenoteSectionGroupSectionPageContentOperationOptions) ToQuery() *client.QueryParams {
	out := client.QueryParams{}

	return &out
}

// SetOnenoteSectionGroupSectionPageContent - Update content for the navigation property pages in users. The page's HTML
// content.
func (c OnenoteSectionGroupSectionPageContentClient) SetOnenoteSectionGroupSectionPageContent(ctx context.Context, id beta.UserIdOnenoteSectionGroupIdSectionIdPageId, input []byte, options SetOnenoteSectionGroupSectionPageContentOperationOptions) (result SetOnenoteSectionGroupSectionPageContentOperationResponse, err error) {
	opts := client.RequestOptions{
		ContentType: "application/json; charset=utf-8",
		ExpectedStatusCodes: []int{
			http.StatusNoContent,
		},
		HttpMethod:    http.MethodPut,
		OptionsObject: options,
		Path:          fmt.Sprintf("%s/content", id.ID()),
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
