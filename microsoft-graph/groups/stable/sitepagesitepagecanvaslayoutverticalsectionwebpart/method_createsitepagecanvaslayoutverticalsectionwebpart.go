package sitepagesitepagecanvaslayoutverticalsectionwebpart

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/hashicorp/go-azure-sdk/microsoft-graph/common-types/stable"
	"github.com/hashicorp/go-azure-sdk/sdk/client"
	"github.com/hashicorp/go-azure-sdk/sdk/odata"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type CreateSitePageCanvasLayoutVerticalSectionWebpartOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
	Model        stable.WebPart
}

type CreateSitePageCanvasLayoutVerticalSectionWebpartOperationOptions struct {
	Metadata  *odata.Metadata
	RetryFunc client.RequestRetryFunc
}

func DefaultCreateSitePageCanvasLayoutVerticalSectionWebpartOperationOptions() CreateSitePageCanvasLayoutVerticalSectionWebpartOperationOptions {
	return CreateSitePageCanvasLayoutVerticalSectionWebpartOperationOptions{}
}

func (o CreateSitePageCanvasLayoutVerticalSectionWebpartOperationOptions) ToHeaders() *client.Headers {
	out := client.Headers{}

	return &out
}

func (o CreateSitePageCanvasLayoutVerticalSectionWebpartOperationOptions) ToOData() *odata.Query {
	out := odata.Query{}
	if o.Metadata != nil {
		out.Metadata = *o.Metadata
	}
	return &out
}

func (o CreateSitePageCanvasLayoutVerticalSectionWebpartOperationOptions) ToQuery() *client.QueryParams {
	out := client.QueryParams{}

	return &out
}

// CreateSitePageCanvasLayoutVerticalSectionWebpart - Create new navigation property to webparts for groups
func (c SitePageSitePageCanvasLayoutVerticalSectionWebpartClient) CreateSitePageCanvasLayoutVerticalSectionWebpart(ctx context.Context, id stable.GroupIdSiteIdPageId, input stable.WebPart, options CreateSitePageCanvasLayoutVerticalSectionWebpartOperationOptions) (result CreateSitePageCanvasLayoutVerticalSectionWebpartOperationResponse, err error) {
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
		Path:          fmt.Sprintf("%s/sitePage/canvasLayout/verticalSection/webparts", id.ID()),
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
	model, err := stable.UnmarshalWebPartImplementation(respObj)
	if err != nil {
		return
	}
	result.Model = model

	return
}
