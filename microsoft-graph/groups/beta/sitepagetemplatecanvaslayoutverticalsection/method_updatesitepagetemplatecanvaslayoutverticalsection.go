package sitepagetemplatecanvaslayoutverticalsection

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

type UpdateSitePageTemplateCanvasLayoutVerticalSectionOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
}

type UpdateSitePageTemplateCanvasLayoutVerticalSectionOperationOptions struct {
	Metadata  *odata.Metadata
	RetryFunc client.RequestRetryFunc
}

func DefaultUpdateSitePageTemplateCanvasLayoutVerticalSectionOperationOptions() UpdateSitePageTemplateCanvasLayoutVerticalSectionOperationOptions {
	return UpdateSitePageTemplateCanvasLayoutVerticalSectionOperationOptions{}
}

func (o UpdateSitePageTemplateCanvasLayoutVerticalSectionOperationOptions) ToHeaders() *client.Headers {
	out := client.Headers{}

	return &out
}

func (o UpdateSitePageTemplateCanvasLayoutVerticalSectionOperationOptions) ToOData() *odata.Query {
	out := odata.Query{}
	if o.Metadata != nil {
		out.Metadata = *o.Metadata
	}
	return &out
}

func (o UpdateSitePageTemplateCanvasLayoutVerticalSectionOperationOptions) ToQuery() *client.QueryParams {
	out := client.QueryParams{}

	return &out
}

// UpdateSitePageTemplateCanvasLayoutVerticalSection - Update the navigation property verticalSection in groups
func (c SitePageTemplateCanvasLayoutVerticalSectionClient) UpdateSitePageTemplateCanvasLayoutVerticalSection(ctx context.Context, id beta.GroupIdSiteIdPageTemplateId, input beta.VerticalSection, options UpdateSitePageTemplateCanvasLayoutVerticalSectionOperationOptions) (result UpdateSitePageTemplateCanvasLayoutVerticalSectionOperationResponse, err error) {
	opts := client.RequestOptions{
		ContentType: "application/json; charset=utf-8",
		ExpectedStatusCodes: []int{
			http.StatusAccepted,
			http.StatusNoContent,
			http.StatusOK,
		},
		HttpMethod:    http.MethodPatch,
		OptionsObject: options,
		Path:          fmt.Sprintf("%s/canvasLayout/verticalSection", id.ID()),
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
