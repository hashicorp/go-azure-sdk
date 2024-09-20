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

type CreateSiteContentTypeUnpublishOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
}

type CreateSiteContentTypeUnpublishOperationOptions struct {
	Metadata *odata.Metadata
}

func DefaultCreateSiteContentTypeUnpublishOperationOptions() CreateSiteContentTypeUnpublishOperationOptions {
	return CreateSiteContentTypeUnpublishOperationOptions{}
}

func (o CreateSiteContentTypeUnpublishOperationOptions) ToHeaders() *client.Headers {
	out := client.Headers{}

	return &out
}

func (o CreateSiteContentTypeUnpublishOperationOptions) ToOData() *odata.Query {
	out := odata.Query{}
	if o.Metadata != nil {
		out.Metadata = *o.Metadata
	}
	return &out
}

func (o CreateSiteContentTypeUnpublishOperationOptions) ToQuery() *client.QueryParams {
	out := client.QueryParams{}

	return &out
}

// CreateSiteContentTypeUnpublish - Invoke action unpublish. Unpublish a contentType from a content type hub site.
func (c SiteContentTypeClient) CreateSiteContentTypeUnpublish(ctx context.Context, id stable.GroupIdSiteIdContentTypeId, options CreateSiteContentTypeUnpublishOperationOptions) (result CreateSiteContentTypeUnpublishOperationResponse, err error) {
	opts := client.RequestOptions{
		ContentType: "application/json; charset=utf-8",
		ExpectedStatusCodes: []int{
			http.StatusNoContent,
		},
		HttpMethod:    http.MethodPost,
		OptionsObject: options,
		Path:          fmt.Sprintf("%s/unpublish", id.ID()),
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

	return
}
