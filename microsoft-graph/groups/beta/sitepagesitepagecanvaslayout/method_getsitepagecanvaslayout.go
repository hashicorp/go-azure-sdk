package sitepagesitepagecanvaslayout

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

type GetSitePageCanvasLayoutOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
	Model        *beta.CanvasLayout
}

type GetSitePageCanvasLayoutOperationOptions struct {
	Expand    *odata.Expand
	Metadata  *odata.Metadata
	RetryFunc client.RequestRetryFunc
	Select    *[]string
}

func DefaultGetSitePageCanvasLayoutOperationOptions() GetSitePageCanvasLayoutOperationOptions {
	return GetSitePageCanvasLayoutOperationOptions{}
}

func (o GetSitePageCanvasLayoutOperationOptions) ToHeaders() *client.Headers {
	out := client.Headers{}

	return &out
}

func (o GetSitePageCanvasLayoutOperationOptions) ToOData() *odata.Query {
	out := odata.Query{}
	if o.Expand != nil {
		out.Expand = *o.Expand
	}
	if o.Metadata != nil {
		out.Metadata = *o.Metadata
	}
	if o.Select != nil {
		out.Select = *o.Select
	}
	return &out
}

func (o GetSitePageCanvasLayoutOperationOptions) ToQuery() *client.QueryParams {
	out := client.QueryParams{}

	return &out
}

// GetSitePageCanvasLayout - Get canvasLayout from groups. Indicates the layout of the content in a given SharePoint
// page, including horizontal sections and vertical sections.
func (c SitePageSitePageCanvasLayoutClient) GetSitePageCanvasLayout(ctx context.Context, id beta.GroupIdSiteIdPageId, options GetSitePageCanvasLayoutOperationOptions) (result GetSitePageCanvasLayoutOperationResponse, err error) {
	opts := client.RequestOptions{
		ContentType: "application/json; charset=utf-8",
		ExpectedStatusCodes: []int{
			http.StatusOK,
		},
		HttpMethod:    http.MethodGet,
		OptionsObject: options,
		Path:          fmt.Sprintf("%s/sitePage/canvasLayout", id.ID()),
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

	var model beta.CanvasLayout
	result.Model = &model
	if err = resp.Unmarshal(result.Model); err != nil {
		return
	}

	return
}
