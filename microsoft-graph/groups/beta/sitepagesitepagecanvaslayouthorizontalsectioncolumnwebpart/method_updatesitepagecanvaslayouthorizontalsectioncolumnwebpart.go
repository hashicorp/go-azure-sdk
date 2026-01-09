package sitepagesitepagecanvaslayouthorizontalsectioncolumnwebpart

import (
	"context"
	"net/http"

	"github.com/hashicorp/go-azure-sdk/microsoft-graph/common-types/beta"
	"github.com/hashicorp/go-azure-sdk/sdk/client"
	"github.com/hashicorp/go-azure-sdk/sdk/odata"
)

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type UpdateSitePageCanvasLayoutHorizontalSectionColumnWebpartOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
}

type UpdateSitePageCanvasLayoutHorizontalSectionColumnWebpartOperationOptions struct {
	Metadata  *odata.Metadata
	RetryFunc client.RequestRetryFunc
}

func DefaultUpdateSitePageCanvasLayoutHorizontalSectionColumnWebpartOperationOptions() UpdateSitePageCanvasLayoutHorizontalSectionColumnWebpartOperationOptions {
	return UpdateSitePageCanvasLayoutHorizontalSectionColumnWebpartOperationOptions{}
}

func (o UpdateSitePageCanvasLayoutHorizontalSectionColumnWebpartOperationOptions) ToHeaders() *client.Headers {
	out := client.Headers{}

	return &out
}

func (o UpdateSitePageCanvasLayoutHorizontalSectionColumnWebpartOperationOptions) ToOData() *odata.Query {
	out := odata.Query{}
	if o.Metadata != nil {
		out.Metadata = *o.Metadata
	}
	return &out
}

func (o UpdateSitePageCanvasLayoutHorizontalSectionColumnWebpartOperationOptions) ToQuery() *client.QueryParams {
	out := client.QueryParams{}

	return &out
}

// UpdateSitePageCanvasLayoutHorizontalSectionColumnWebpart - Update the navigation property webparts in groups
func (c SitePageSitePageCanvasLayoutHorizontalSectionColumnWebpartClient) UpdateSitePageCanvasLayoutHorizontalSectionColumnWebpart(ctx context.Context, id beta.GroupIdSiteIdPageIdSitePageCanvasLayoutHorizontalSectionIdColumnIdWebpartId, input beta.WebPart, options UpdateSitePageCanvasLayoutHorizontalSectionColumnWebpartOperationOptions) (result UpdateSitePageCanvasLayoutHorizontalSectionColumnWebpartOperationResponse, err error) {
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
