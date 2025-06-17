package sitepagetemplatecanvaslayouthorizontalsectioncolumnwebpart

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

type GetSitePageTemplateCanvasLayoutHorizontalSectionColumnWebpartPositionOfWebPartOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
	Model        *beta.WebPartPosition
}

type GetSitePageTemplateCanvasLayoutHorizontalSectionColumnWebpartPositionOfWebPartOperationOptions struct {
	Metadata  *odata.Metadata
	RetryFunc client.RequestRetryFunc
}

func DefaultGetSitePageTemplateCanvasLayoutHorizontalSectionColumnWebpartPositionOfWebPartOperationOptions() GetSitePageTemplateCanvasLayoutHorizontalSectionColumnWebpartPositionOfWebPartOperationOptions {
	return GetSitePageTemplateCanvasLayoutHorizontalSectionColumnWebpartPositionOfWebPartOperationOptions{}
}

func (o GetSitePageTemplateCanvasLayoutHorizontalSectionColumnWebpartPositionOfWebPartOperationOptions) ToHeaders() *client.Headers {
	out := client.Headers{}

	return &out
}

func (o GetSitePageTemplateCanvasLayoutHorizontalSectionColumnWebpartPositionOfWebPartOperationOptions) ToOData() *odata.Query {
	out := odata.Query{}
	if o.Metadata != nil {
		out.Metadata = *o.Metadata
	}
	return &out
}

func (o GetSitePageTemplateCanvasLayoutHorizontalSectionColumnWebpartPositionOfWebPartOperationOptions) ToQuery() *client.QueryParams {
	out := client.QueryParams{}

	return &out
}

// GetSitePageTemplateCanvasLayoutHorizontalSectionColumnWebpartPositionOfWebPart - Invoke action getPositionOfWebPart
func (c SitePageTemplateCanvasLayoutHorizontalSectionColumnWebpartClient) GetSitePageTemplateCanvasLayoutHorizontalSectionColumnWebpartPositionOfWebPart(ctx context.Context, id beta.GroupIdSiteIdPageTemplateIdCanvasLayoutHorizontalSectionIdColumnIdWebpartId, options GetSitePageTemplateCanvasLayoutHorizontalSectionColumnWebpartPositionOfWebPartOperationOptions) (result GetSitePageTemplateCanvasLayoutHorizontalSectionColumnWebpartPositionOfWebPartOperationResponse, err error) {
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
		Path:          fmt.Sprintf("%s/getPositionOfWebPart", id.ID()),
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

	var model beta.WebPartPosition
	result.Model = &model
	if err = resp.Unmarshal(result.Model); err != nil {
		return
	}

	return
}
