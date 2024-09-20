package sitecontenttype

import (
	"context"
	"fmt"
	"net/http"

	"github.com/hashicorp/go-azure-sdk/microsoft-graph/common-types/beta"
	"github.com/hashicorp/go-azure-sdk/sdk/client"
	"github.com/hashicorp/go-azure-sdk/sdk/odata"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type AddSiteContentTypesCopyOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
	Model        *beta.ContentType
}

type AddSiteContentTypesCopyOperationOptions struct {
	Metadata *odata.Metadata
}

func DefaultAddSiteContentTypesCopyOperationOptions() AddSiteContentTypesCopyOperationOptions {
	return AddSiteContentTypesCopyOperationOptions{}
}

func (o AddSiteContentTypesCopyOperationOptions) ToHeaders() *client.Headers {
	out := client.Headers{}

	return &out
}

func (o AddSiteContentTypesCopyOperationOptions) ToOData() *odata.Query {
	out := odata.Query{}
	if o.Metadata != nil {
		out.Metadata = *o.Metadata
	}
	return &out
}

func (o AddSiteContentTypesCopyOperationOptions) ToQuery() *client.QueryParams {
	out := client.QueryParams{}

	return &out
}

// AddSiteContentTypesCopy - Invoke action addCopy
func (c SiteContentTypeClient) AddSiteContentTypesCopy(ctx context.Context, id beta.GroupIdSiteId, input AddSiteContentTypesCopyRequest, options AddSiteContentTypesCopyOperationOptions) (result AddSiteContentTypesCopyOperationResponse, err error) {
	opts := client.RequestOptions{
		ContentType: "application/json; charset=utf-8",
		ExpectedStatusCodes: []int{
			http.StatusOK,
		},
		HttpMethod:    http.MethodPost,
		OptionsObject: options,
		Path:          fmt.Sprintf("%s/contentTypes/addCopy", id.ID()),
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

	var model beta.ContentType
	result.Model = &model
	if err = resp.Unmarshal(result.Model); err != nil {
		return
	}

	return
}
