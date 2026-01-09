package sitepagetemplatecanvaslayouthorizontalsectioncolumn

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

type CreateSitePageTemplateCanvasLayoutHorizontalSectionColumnOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
	Model        *beta.HorizontalSectionColumn
}

type CreateSitePageTemplateCanvasLayoutHorizontalSectionColumnOperationOptions struct {
	Metadata  *odata.Metadata
	RetryFunc client.RequestRetryFunc
}

func DefaultCreateSitePageTemplateCanvasLayoutHorizontalSectionColumnOperationOptions() CreateSitePageTemplateCanvasLayoutHorizontalSectionColumnOperationOptions {
	return CreateSitePageTemplateCanvasLayoutHorizontalSectionColumnOperationOptions{}
}

func (o CreateSitePageTemplateCanvasLayoutHorizontalSectionColumnOperationOptions) ToHeaders() *client.Headers {
	out := client.Headers{}

	return &out
}

func (o CreateSitePageTemplateCanvasLayoutHorizontalSectionColumnOperationOptions) ToOData() *odata.Query {
	out := odata.Query{}
	if o.Metadata != nil {
		out.Metadata = *o.Metadata
	}
	return &out
}

func (o CreateSitePageTemplateCanvasLayoutHorizontalSectionColumnOperationOptions) ToQuery() *client.QueryParams {
	out := client.QueryParams{}

	return &out
}

// CreateSitePageTemplateCanvasLayoutHorizontalSectionColumn - Create new navigation property to columns for groups
func (c SitePageTemplateCanvasLayoutHorizontalSectionColumnClient) CreateSitePageTemplateCanvasLayoutHorizontalSectionColumn(ctx context.Context, id beta.GroupIdSiteIdPageTemplateIdCanvasLayoutHorizontalSectionId, input beta.HorizontalSectionColumn, options CreateSitePageTemplateCanvasLayoutHorizontalSectionColumnOperationOptions) (result CreateSitePageTemplateCanvasLayoutHorizontalSectionColumnOperationResponse, err error) {
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
		Path:          fmt.Sprintf("%s/columns", id.ID()),
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

	var model beta.HorizontalSectionColumn
	result.Model = &model
	if err = resp.Unmarshal(result.Model); err != nil {
		return
	}

	return
}
