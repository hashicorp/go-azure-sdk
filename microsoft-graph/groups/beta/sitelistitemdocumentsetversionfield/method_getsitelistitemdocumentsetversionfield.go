package sitelistitemdocumentsetversionfield

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

type GetSiteListItemDocumentSetVersionFieldOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
	Model        *beta.FieldValueSet
}

type GetSiteListItemDocumentSetVersionFieldOperationOptions struct {
	Expand   *odata.Expand
	Metadata *odata.Metadata
	Select   *[]string
}

func DefaultGetSiteListItemDocumentSetVersionFieldOperationOptions() GetSiteListItemDocumentSetVersionFieldOperationOptions {
	return GetSiteListItemDocumentSetVersionFieldOperationOptions{}
}

func (o GetSiteListItemDocumentSetVersionFieldOperationOptions) ToHeaders() *client.Headers {
	out := client.Headers{}

	return &out
}

func (o GetSiteListItemDocumentSetVersionFieldOperationOptions) ToOData() *odata.Query {
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

func (o GetSiteListItemDocumentSetVersionFieldOperationOptions) ToQuery() *client.QueryParams {
	out := client.QueryParams{}

	return &out
}

// GetSiteListItemDocumentSetVersionField - Get fields from groups. A collection of the fields and values for this
// version of the list item.
func (c SiteListItemDocumentSetVersionFieldClient) GetSiteListItemDocumentSetVersionField(ctx context.Context, id beta.GroupIdSiteIdListIdItemIdDocumentSetVersionId, options GetSiteListItemDocumentSetVersionFieldOperationOptions) (result GetSiteListItemDocumentSetVersionFieldOperationResponse, err error) {
	opts := client.RequestOptions{
		ContentType: "application/json; charset=utf-8",
		ExpectedStatusCodes: []int{
			http.StatusOK,
		},
		HttpMethod:    http.MethodGet,
		OptionsObject: options,
		Path:          fmt.Sprintf("%s/fields", id.ID()),
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

	var model beta.FieldValueSet
	result.Model = &model
	if err = resp.Unmarshal(result.Model); err != nil {
		return
	}

	return
}
