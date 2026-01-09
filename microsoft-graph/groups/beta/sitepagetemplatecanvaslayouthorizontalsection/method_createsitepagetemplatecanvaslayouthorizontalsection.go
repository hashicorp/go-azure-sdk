package sitepagetemplatecanvaslayouthorizontalsection

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

type CreateSitePageTemplateCanvasLayoutHorizontalSectionOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
	Model        *beta.HorizontalSection
}

type CreateSitePageTemplateCanvasLayoutHorizontalSectionOperationOptions struct {
	Metadata  *odata.Metadata
	RetryFunc client.RequestRetryFunc
}

func DefaultCreateSitePageTemplateCanvasLayoutHorizontalSectionOperationOptions() CreateSitePageTemplateCanvasLayoutHorizontalSectionOperationOptions {
	return CreateSitePageTemplateCanvasLayoutHorizontalSectionOperationOptions{}
}

func (o CreateSitePageTemplateCanvasLayoutHorizontalSectionOperationOptions) ToHeaders() *client.Headers {
	out := client.Headers{}

	return &out
}

func (o CreateSitePageTemplateCanvasLayoutHorizontalSectionOperationOptions) ToOData() *odata.Query {
	out := odata.Query{}
	if o.Metadata != nil {
		out.Metadata = *o.Metadata
	}
	return &out
}

func (o CreateSitePageTemplateCanvasLayoutHorizontalSectionOperationOptions) ToQuery() *client.QueryParams {
	out := client.QueryParams{}

	return &out
}

// CreateSitePageTemplateCanvasLayoutHorizontalSection - Create new navigation property to horizontalSections for groups
func (c SitePageTemplateCanvasLayoutHorizontalSectionClient) CreateSitePageTemplateCanvasLayoutHorizontalSection(ctx context.Context, id beta.GroupIdSiteIdPageTemplateId, input beta.HorizontalSection, options CreateSitePageTemplateCanvasLayoutHorizontalSectionOperationOptions) (result CreateSitePageTemplateCanvasLayoutHorizontalSectionOperationResponse, err error) {
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
		Path:          fmt.Sprintf("%s/canvasLayout/horizontalSections", id.ID()),
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

	var model beta.HorizontalSection
	result.Model = &model
	if err = resp.Unmarshal(result.Model); err != nil {
		return
	}

	return
}
