package sitepagetemplatecanvaslayouthorizontalsectioncolumn

import (
	"context"
	"net/http"

	"github.com/hashicorp/go-azure-sdk/microsoft-graph/common-types/beta"
	"github.com/hashicorp/go-azure-sdk/sdk/client"
	"github.com/hashicorp/go-azure-sdk/sdk/odata"
)

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type UpdateSitePageTemplateCanvasLayoutHorizontalSectionColumnOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
}

type UpdateSitePageTemplateCanvasLayoutHorizontalSectionColumnOperationOptions struct {
	Metadata  *odata.Metadata
	RetryFunc client.RequestRetryFunc
}

func DefaultUpdateSitePageTemplateCanvasLayoutHorizontalSectionColumnOperationOptions() UpdateSitePageTemplateCanvasLayoutHorizontalSectionColumnOperationOptions {
	return UpdateSitePageTemplateCanvasLayoutHorizontalSectionColumnOperationOptions{}
}

func (o UpdateSitePageTemplateCanvasLayoutHorizontalSectionColumnOperationOptions) ToHeaders() *client.Headers {
	out := client.Headers{}

	return &out
}

func (o UpdateSitePageTemplateCanvasLayoutHorizontalSectionColumnOperationOptions) ToOData() *odata.Query {
	out := odata.Query{}
	if o.Metadata != nil {
		out.Metadata = *o.Metadata
	}
	return &out
}

func (o UpdateSitePageTemplateCanvasLayoutHorizontalSectionColumnOperationOptions) ToQuery() *client.QueryParams {
	out := client.QueryParams{}

	return &out
}

// UpdateSitePageTemplateCanvasLayoutHorizontalSectionColumn - Update the navigation property columns in groups
func (c SitePageTemplateCanvasLayoutHorizontalSectionColumnClient) UpdateSitePageTemplateCanvasLayoutHorizontalSectionColumn(ctx context.Context, id beta.GroupIdSiteIdPageTemplateIdCanvasLayoutHorizontalSectionIdColumnId, input beta.HorizontalSectionColumn, options UpdateSitePageTemplateCanvasLayoutHorizontalSectionColumnOperationOptions) (result UpdateSitePageTemplateCanvasLayoutHorizontalSectionColumnOperationResponse, err error) {
	opts := client.RequestOptions{
		ContentType: "application/json; charset=utf-8",
		ExpectedStatusCodes: []int{
			http.StatusAccepted,
			http.StatusNoContent,
			http.StatusOK,
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
