package sitepagetemplatecanvaslayoutverticalsectionwebpart

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

type GetSitePageTemplateCanvasLayoutVerticalSectionWebpartPositionOfWebPartOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
	Model        *beta.WebPartPosition
}

type GetSitePageTemplateCanvasLayoutVerticalSectionWebpartPositionOfWebPartOperationOptions struct {
	Metadata  *odata.Metadata
	RetryFunc client.RequestRetryFunc
}

func DefaultGetSitePageTemplateCanvasLayoutVerticalSectionWebpartPositionOfWebPartOperationOptions() GetSitePageTemplateCanvasLayoutVerticalSectionWebpartPositionOfWebPartOperationOptions {
	return GetSitePageTemplateCanvasLayoutVerticalSectionWebpartPositionOfWebPartOperationOptions{}
}

func (o GetSitePageTemplateCanvasLayoutVerticalSectionWebpartPositionOfWebPartOperationOptions) ToHeaders() *client.Headers {
	out := client.Headers{}

	return &out
}

func (o GetSitePageTemplateCanvasLayoutVerticalSectionWebpartPositionOfWebPartOperationOptions) ToOData() *odata.Query {
	out := odata.Query{}
	if o.Metadata != nil {
		out.Metadata = *o.Metadata
	}
	return &out
}

func (o GetSitePageTemplateCanvasLayoutVerticalSectionWebpartPositionOfWebPartOperationOptions) ToQuery() *client.QueryParams {
	out := client.QueryParams{}

	return &out
}

// GetSitePageTemplateCanvasLayoutVerticalSectionWebpartPositionOfWebPart - Invoke action getPositionOfWebPart
func (c SitePageTemplateCanvasLayoutVerticalSectionWebpartClient) GetSitePageTemplateCanvasLayoutVerticalSectionWebpartPositionOfWebPart(ctx context.Context, id beta.GroupIdSiteIdPageTemplateIdCanvasLayoutVerticalSectionWebpartId, options GetSitePageTemplateCanvasLayoutVerticalSectionWebpartPositionOfWebPartOperationOptions) (result GetSitePageTemplateCanvasLayoutVerticalSectionWebpartPositionOfWebPartOperationResponse, err error) {
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
