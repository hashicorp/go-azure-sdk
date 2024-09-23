package sitecontenttype

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

type CopySiteContentTypeToDefaultContentLocationOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
}

type CopySiteContentTypeToDefaultContentLocationOperationOptions struct {
	Metadata  *odata.Metadata
	RetryFunc client.RequestRetryFunc
}

func DefaultCopySiteContentTypeToDefaultContentLocationOperationOptions() CopySiteContentTypeToDefaultContentLocationOperationOptions {
	return CopySiteContentTypeToDefaultContentLocationOperationOptions{}
}

func (o CopySiteContentTypeToDefaultContentLocationOperationOptions) ToHeaders() *client.Headers {
	out := client.Headers{}

	return &out
}

func (o CopySiteContentTypeToDefaultContentLocationOperationOptions) ToOData() *odata.Query {
	out := odata.Query{}
	if o.Metadata != nil {
		out.Metadata = *o.Metadata
	}
	return &out
}

func (o CopySiteContentTypeToDefaultContentLocationOperationOptions) ToQuery() *client.QueryParams {
	out := client.QueryParams{}

	return &out
}

// CopySiteContentTypeToDefaultContentLocation - Invoke action copyToDefaultContentLocation. Copy a file to a default
// content location in a content type. The file can then be added as a default file or template via a POST operation.
func (c SiteContentTypeClient) CopySiteContentTypeToDefaultContentLocation(ctx context.Context, id stable.GroupIdSiteIdContentTypeId, input CopySiteContentTypeToDefaultContentLocationRequest, options CopySiteContentTypeToDefaultContentLocationOperationOptions) (result CopySiteContentTypeToDefaultContentLocationOperationResponse, err error) {
	opts := client.RequestOptions{
		ContentType: "application/json; charset=utf-8",
		ExpectedStatusCodes: []int{
			http.StatusNoContent,
		},
		HttpMethod:    http.MethodPost,
		OptionsObject: options,
		Path:          fmt.Sprintf("%s/copyToDefaultContentLocation", id.ID()),
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
