package sitepagetemplatewebpart

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

type GetSitePageTemplateWebPartPositionOfWebPartOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
	Model        *beta.WebPartPosition
}

type GetSitePageTemplateWebPartPositionOfWebPartOperationOptions struct {
	Metadata  *odata.Metadata
	RetryFunc client.RequestRetryFunc
}

func DefaultGetSitePageTemplateWebPartPositionOfWebPartOperationOptions() GetSitePageTemplateWebPartPositionOfWebPartOperationOptions {
	return GetSitePageTemplateWebPartPositionOfWebPartOperationOptions{}
}

func (o GetSitePageTemplateWebPartPositionOfWebPartOperationOptions) ToHeaders() *client.Headers {
	out := client.Headers{}

	return &out
}

func (o GetSitePageTemplateWebPartPositionOfWebPartOperationOptions) ToOData() *odata.Query {
	out := odata.Query{}
	if o.Metadata != nil {
		out.Metadata = *o.Metadata
	}
	return &out
}

func (o GetSitePageTemplateWebPartPositionOfWebPartOperationOptions) ToQuery() *client.QueryParams {
	out := client.QueryParams{}

	return &out
}

// GetSitePageTemplateWebPartPositionOfWebPart - Invoke action getPositionOfWebPart
func (c SitePageTemplateWebPartClient) GetSitePageTemplateWebPartPositionOfWebPart(ctx context.Context, id beta.GroupIdSiteIdPageTemplateIdWebPartId, options GetSitePageTemplateWebPartPositionOfWebPartOperationOptions) (result GetSitePageTemplateWebPartPositionOfWebPartOperationResponse, err error) {
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
