package sitepagetemplatecanvaslayouthorizontalsectioncolumn

import (
	"context"
	"net/http"

	"github.com/hashicorp/go-azure-sdk/microsoft-graph/common-types/beta"
	"github.com/hashicorp/go-azure-sdk/sdk/client"
	"github.com/hashicorp/go-azure-sdk/sdk/odata"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type GetSitePageTemplateCanvasLayoutHorizontalSectionColumnOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
	Model        *beta.HorizontalSectionColumn
}

type GetSitePageTemplateCanvasLayoutHorizontalSectionColumnOperationOptions struct {
	Expand    *odata.Expand
	Metadata  *odata.Metadata
	RetryFunc client.RequestRetryFunc
	Select    *[]string
}

func DefaultGetSitePageTemplateCanvasLayoutHorizontalSectionColumnOperationOptions() GetSitePageTemplateCanvasLayoutHorizontalSectionColumnOperationOptions {
	return GetSitePageTemplateCanvasLayoutHorizontalSectionColumnOperationOptions{}
}

func (o GetSitePageTemplateCanvasLayoutHorizontalSectionColumnOperationOptions) ToHeaders() *client.Headers {
	out := client.Headers{}

	return &out
}

func (o GetSitePageTemplateCanvasLayoutHorizontalSectionColumnOperationOptions) ToOData() *odata.Query {
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

func (o GetSitePageTemplateCanvasLayoutHorizontalSectionColumnOperationOptions) ToQuery() *client.QueryParams {
	out := client.QueryParams{}

	return &out
}

// GetSitePageTemplateCanvasLayoutHorizontalSectionColumn - Get columns from groups. The set of vertical columns in this
// section.
func (c SitePageTemplateCanvasLayoutHorizontalSectionColumnClient) GetSitePageTemplateCanvasLayoutHorizontalSectionColumn(ctx context.Context, id beta.GroupIdSiteIdPageTemplateIdCanvasLayoutHorizontalSectionIdColumnId, options GetSitePageTemplateCanvasLayoutHorizontalSectionColumnOperationOptions) (result GetSitePageTemplateCanvasLayoutHorizontalSectionColumnOperationResponse, err error) {
	opts := client.RequestOptions{
		ContentType: "application/json; charset=utf-8",
		ExpectedStatusCodes: []int{
			http.StatusOK,
		},
		HttpMethod:    http.MethodGet,
		OptionsObject: options,
		Path:          id.ID(),
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

	var model beta.HorizontalSectionColumn
	result.Model = &model
	if err = resp.Unmarshal(result.Model); err != nil {
		return
	}

	return
}
