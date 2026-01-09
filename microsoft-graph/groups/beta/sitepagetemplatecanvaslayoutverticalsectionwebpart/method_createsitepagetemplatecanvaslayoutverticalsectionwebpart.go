package sitepagetemplatecanvaslayoutverticalsectionwebpart

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/hashicorp/go-azure-sdk/microsoft-graph/common-types/beta"
	"github.com/hashicorp/go-azure-sdk/sdk/client"
	"github.com/hashicorp/go-azure-sdk/sdk/odata"
)

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type CreateSitePageTemplateCanvasLayoutVerticalSectionWebpartOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
	Model        beta.WebPart
}

type CreateSitePageTemplateCanvasLayoutVerticalSectionWebpartOperationOptions struct {
	Metadata  *odata.Metadata
	RetryFunc client.RequestRetryFunc
}

func DefaultCreateSitePageTemplateCanvasLayoutVerticalSectionWebpartOperationOptions() CreateSitePageTemplateCanvasLayoutVerticalSectionWebpartOperationOptions {
	return CreateSitePageTemplateCanvasLayoutVerticalSectionWebpartOperationOptions{}
}

func (o CreateSitePageTemplateCanvasLayoutVerticalSectionWebpartOperationOptions) ToHeaders() *client.Headers {
	out := client.Headers{}

	return &out
}

func (o CreateSitePageTemplateCanvasLayoutVerticalSectionWebpartOperationOptions) ToOData() *odata.Query {
	out := odata.Query{}
	if o.Metadata != nil {
		out.Metadata = *o.Metadata
	}
	return &out
}

func (o CreateSitePageTemplateCanvasLayoutVerticalSectionWebpartOperationOptions) ToQuery() *client.QueryParams {
	out := client.QueryParams{}

	return &out
}

// CreateSitePageTemplateCanvasLayoutVerticalSectionWebpart - Create new navigation property to webparts for groups
func (c SitePageTemplateCanvasLayoutVerticalSectionWebpartClient) CreateSitePageTemplateCanvasLayoutVerticalSectionWebpart(ctx context.Context, id beta.GroupIdSiteIdPageTemplateId, input beta.WebPart, options CreateSitePageTemplateCanvasLayoutVerticalSectionWebpartOperationOptions) (result CreateSitePageTemplateCanvasLayoutVerticalSectionWebpartOperationResponse, err error) {
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
		Path:          fmt.Sprintf("%s/canvasLayout/verticalSection/webparts", id.ID()),
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

	var respObj json.RawMessage
	if err = resp.Unmarshal(&respObj); err != nil {
		return
	}
	model, err := beta.UnmarshalWebPartImplementation(respObj)
	if err != nil {
		return
	}
	result.Model = model

	return
}
