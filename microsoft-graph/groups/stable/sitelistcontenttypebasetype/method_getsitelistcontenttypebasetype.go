package sitelistcontenttypebasetype

import (
	"context"
	"net/http"

	"github.com/hashicorp/go-azure-sdk/microsoft-graph/common-types/stable"
	"github.com/hashicorp/go-azure-sdk/sdk/client"
	"github.com/hashicorp/go-azure-sdk/sdk/odata"
)

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type GetSiteListContentTypeBaseTypeOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
	Model        *stable.ContentType
}

type GetSiteListContentTypeBaseTypeOperationOptions struct {
	Expand    *odata.Expand
	Metadata  *odata.Metadata
	RetryFunc client.RequestRetryFunc
	Select    *[]string
}

func DefaultGetSiteListContentTypeBaseTypeOperationOptions() GetSiteListContentTypeBaseTypeOperationOptions {
	return GetSiteListContentTypeBaseTypeOperationOptions{}
}

func (o GetSiteListContentTypeBaseTypeOperationOptions) ToHeaders() *client.Headers {
	out := client.Headers{}

	return &out
}

func (o GetSiteListContentTypeBaseTypeOperationOptions) ToOData() *odata.Query {
	out := odata.Query{}
	if o.Expand != nil {
		out.Expand = *o.Expand
	}
	if o.Metadata != nil {
		out.Metadata = *o.Metadata
	}
	if o.Select != nil {
		out.Select = *o.Select
	}
	return &out
}

func (o GetSiteListContentTypeBaseTypeOperationOptions) ToQuery() *client.QueryParams {
	out := client.QueryParams{}

	return &out
}

// GetSiteListContentTypeBaseType - Get baseTypes from groups. The collection of content types that are ancestors of
// this content type.
func (c SiteListContentTypeBaseTypeClient) GetSiteListContentTypeBaseType(ctx context.Context, id stable.GroupIdSiteIdListIdContentTypeIdBaseTypeId, options GetSiteListContentTypeBaseTypeOperationOptions) (result GetSiteListContentTypeBaseTypeOperationResponse, err error) {
	opts := client.RequestOptions{
		ContentType: "application/json; charset=utf-8",
		ExpectedStatusCodes: []int{
			http.StatusOK,
		},
		HttpMethod:    http.MethodGet,
		OptionsObject: options,
		Path:          id.ID(),
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

	var model stable.ContentType
	result.Model = &model
	if err = resp.Unmarshal(result.Model); err != nil {
		return
	}

	return
}
