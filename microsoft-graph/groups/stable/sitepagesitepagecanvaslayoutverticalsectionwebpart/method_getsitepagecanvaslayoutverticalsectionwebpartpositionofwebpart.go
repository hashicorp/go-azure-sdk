package sitepagesitepagecanvaslayoutverticalsectionwebpart

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

type GetSitePageCanvasLayoutVerticalSectionWebpartPositionOfWebPartOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
	Model        *stable.WebPartPosition
}

type GetSitePageCanvasLayoutVerticalSectionWebpartPositionOfWebPartOperationOptions struct {
	Metadata  *odata.Metadata
	RetryFunc client.RequestRetryFunc
}

func DefaultGetSitePageCanvasLayoutVerticalSectionWebpartPositionOfWebPartOperationOptions() GetSitePageCanvasLayoutVerticalSectionWebpartPositionOfWebPartOperationOptions {
	return GetSitePageCanvasLayoutVerticalSectionWebpartPositionOfWebPartOperationOptions{}
}

func (o GetSitePageCanvasLayoutVerticalSectionWebpartPositionOfWebPartOperationOptions) ToHeaders() *client.Headers {
	out := client.Headers{}

	return &out
}

func (o GetSitePageCanvasLayoutVerticalSectionWebpartPositionOfWebPartOperationOptions) ToOData() *odata.Query {
	out := odata.Query{}
	if o.Metadata != nil {
		out.Metadata = *o.Metadata
	}
	return &out
}

func (o GetSitePageCanvasLayoutVerticalSectionWebpartPositionOfWebPartOperationOptions) ToQuery() *client.QueryParams {
	out := client.QueryParams{}

	return &out
}

// GetSitePageCanvasLayoutVerticalSectionWebpartPositionOfWebPart - Invoke action getPositionOfWebPart
func (c SitePageSitePageCanvasLayoutVerticalSectionWebpartClient) GetSitePageCanvasLayoutVerticalSectionWebpartPositionOfWebPart(ctx context.Context, id stable.GroupIdSiteIdPageIdSitePageCanvasLayoutVerticalSectionWebpartId, options GetSitePageCanvasLayoutVerticalSectionWebpartPositionOfWebPartOperationOptions) (result GetSitePageCanvasLayoutVerticalSectionWebpartPositionOfWebPartOperationResponse, err error) {
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

	var model stable.WebPartPosition
	result.Model = &model
	if err = resp.Unmarshal(result.Model); err != nil {
		return
	}

	return
}
