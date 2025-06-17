package sitepagesitepagecanvaslayoutverticalsection

import (
	"context"
	"fmt"
	"net/http"

	"github.com/hashicorp/go-azure-sdk/microsoft-graph/common-types/stable"
	"github.com/hashicorp/go-azure-sdk/sdk/client"
	"github.com/hashicorp/go-azure-sdk/sdk/odata"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type UpdateSitePageCanvasLayoutVerticalSectionOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
}

type UpdateSitePageCanvasLayoutVerticalSectionOperationOptions struct {
	Metadata  *odata.Metadata
	RetryFunc client.RequestRetryFunc
}

func DefaultUpdateSitePageCanvasLayoutVerticalSectionOperationOptions() UpdateSitePageCanvasLayoutVerticalSectionOperationOptions {
	return UpdateSitePageCanvasLayoutVerticalSectionOperationOptions{}
}

func (o UpdateSitePageCanvasLayoutVerticalSectionOperationOptions) ToHeaders() *client.Headers {
	out := client.Headers{}

	return &out
}

func (o UpdateSitePageCanvasLayoutVerticalSectionOperationOptions) ToOData() *odata.Query {
	out := odata.Query{}
	if o.Metadata != nil {
		out.Metadata = *o.Metadata
	}
	return &out
}

func (o UpdateSitePageCanvasLayoutVerticalSectionOperationOptions) ToQuery() *client.QueryParams {
	out := client.QueryParams{}

	return &out
}

// UpdateSitePageCanvasLayoutVerticalSection - Update the navigation property verticalSection in groups
func (c SitePageSitePageCanvasLayoutVerticalSectionClient) UpdateSitePageCanvasLayoutVerticalSection(ctx context.Context, id stable.GroupIdSiteIdPageId, input stable.VerticalSection, options UpdateSitePageCanvasLayoutVerticalSectionOperationOptions) (result UpdateSitePageCanvasLayoutVerticalSectionOperationResponse, err error) {
	opts := client.RequestOptions{
		ContentType: "application/json; charset=utf-8",
		ExpectedStatusCodes: []int{
			http.StatusAccepted,
			http.StatusNoContent,
			http.StatusOK,
		},
		HttpMethod:    http.MethodPatch,
		OptionsObject: options,
		Path:          fmt.Sprintf("%s/sitePage/canvasLayout/verticalSection", id.ID()),
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
