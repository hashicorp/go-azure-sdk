package siteonenotesection

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

type CopySiteOnenoteSectionToSectionGroupOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
	Model        *stable.OnenoteOperation
}

type CopySiteOnenoteSectionToSectionGroupOperationOptions struct {
	Metadata  *odata.Metadata
	RetryFunc client.RequestRetryFunc
}

func DefaultCopySiteOnenoteSectionToSectionGroupOperationOptions() CopySiteOnenoteSectionToSectionGroupOperationOptions {
	return CopySiteOnenoteSectionToSectionGroupOperationOptions{}
}

func (o CopySiteOnenoteSectionToSectionGroupOperationOptions) ToHeaders() *client.Headers {
	out := client.Headers{}

	return &out
}

func (o CopySiteOnenoteSectionToSectionGroupOperationOptions) ToOData() *odata.Query {
	out := odata.Query{}
	if o.Metadata != nil {
		out.Metadata = *o.Metadata
	}
	return &out
}

func (o CopySiteOnenoteSectionToSectionGroupOperationOptions) ToQuery() *client.QueryParams {
	out := client.QueryParams{}

	return &out
}

// CopySiteOnenoteSectionToSectionGroup - Invoke action copyToSectionGroup. For Copy operations, you follow an
// asynchronous calling pattern: First call the Copy action, and then poll the operation endpoint for the result.
func (c SiteOnenoteSectionClient) CopySiteOnenoteSectionToSectionGroup(ctx context.Context, id stable.GroupIdSiteIdOnenoteSectionId, input CopySiteOnenoteSectionToSectionGroupRequest, options CopySiteOnenoteSectionToSectionGroupOperationOptions) (result CopySiteOnenoteSectionToSectionGroupOperationResponse, err error) {
	opts := client.RequestOptions{
		ContentType: "application/json; charset=utf-8",
		ExpectedStatusCodes: []int{
			http.StatusOK,
		},
		HttpMethod:    http.MethodPost,
		OptionsObject: options,
		Path:          fmt.Sprintf("%s/copyToSectionGroup", id.ID()),
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

	var model stable.OnenoteOperation
	result.Model = &model
	if err = resp.Unmarshal(result.Model); err != nil {
		return
	}

	return
}
