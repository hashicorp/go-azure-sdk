package siteonenotenotebooksectiongroupsectionpagecontent

import (
	"context"
	"fmt"
	"net/http"

	"github.com/hashicorp/go-azure-sdk/microsoft-graph/common-types/beta"
	"github.com/hashicorp/go-azure-sdk/sdk/client"
	"github.com/hashicorp/go-azure-sdk/sdk/odata"
)

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type GetSiteOnenoteNotebookSectionGroupSectionPageContentOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
	Model        *[]byte
}

type GetSiteOnenoteNotebookSectionGroupSectionPageContentOperationOptions struct {
	Metadata  *odata.Metadata
	RetryFunc client.RequestRetryFunc
}

func DefaultGetSiteOnenoteNotebookSectionGroupSectionPageContentOperationOptions() GetSiteOnenoteNotebookSectionGroupSectionPageContentOperationOptions {
	return GetSiteOnenoteNotebookSectionGroupSectionPageContentOperationOptions{}
}

func (o GetSiteOnenoteNotebookSectionGroupSectionPageContentOperationOptions) ToHeaders() *client.Headers {
	out := client.Headers{}

	return &out
}

func (o GetSiteOnenoteNotebookSectionGroupSectionPageContentOperationOptions) ToOData() *odata.Query {
	out := odata.Query{}
	if o.Metadata != nil {
		out.Metadata = *o.Metadata
	}
	return &out
}

func (o GetSiteOnenoteNotebookSectionGroupSectionPageContentOperationOptions) ToQuery() *client.QueryParams {
	out := client.QueryParams{}

	return &out
}

// GetSiteOnenoteNotebookSectionGroupSectionPageContent - Get content for the navigation property pages from groups. The
// page's HTML content.
func (c SiteOnenoteNotebookSectionGroupSectionPageContentClient) GetSiteOnenoteNotebookSectionGroupSectionPageContent(ctx context.Context, id beta.GroupIdSiteIdOnenoteNotebookIdSectionGroupIdSectionIdPageId, options GetSiteOnenoteNotebookSectionGroupSectionPageContentOperationOptions) (result GetSiteOnenoteNotebookSectionGroupSectionPageContentOperationResponse, err error) {
	opts := client.RequestOptions{
		ContentType: "application/octet-stream",
		ExpectedStatusCodes: []int{
			http.StatusOK,
		},
		HttpMethod:    http.MethodGet,
		OptionsObject: options,
		Path:          fmt.Sprintf("%s/content", id.ID()),
		RetryFunc:     options.RetryFunc,
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

	var model []byte
	result.Model = &model
	if err = resp.Unmarshal(result.Model); err != nil {
		return
	}

	return
}
