package siterecyclebin

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

type UpdateSiteRecycleBinOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
}

type UpdateSiteRecycleBinOperationOptions struct {
	Metadata *odata.Metadata
}

func DefaultUpdateSiteRecycleBinOperationOptions() UpdateSiteRecycleBinOperationOptions {
	return UpdateSiteRecycleBinOperationOptions{}
}

func (o UpdateSiteRecycleBinOperationOptions) ToHeaders() *client.Headers {
	out := client.Headers{}

	return &out
}

func (o UpdateSiteRecycleBinOperationOptions) ToOData() *odata.Query {
	out := odata.Query{}
	if o.Metadata != nil {
		out.Metadata = *o.Metadata
	}
	return &out
}

func (o UpdateSiteRecycleBinOperationOptions) ToQuery() *client.QueryParams {
	out := client.QueryParams{}

	return &out
}

// UpdateSiteRecycleBin - Update the navigation property recycleBin in groups
func (c SiteRecycleBinClient) UpdateSiteRecycleBin(ctx context.Context, id beta.GroupIdSiteId, input beta.RecycleBin, options UpdateSiteRecycleBinOperationOptions) (result UpdateSiteRecycleBinOperationResponse, err error) {
	opts := client.RequestOptions{
		ContentType: "application/json; charset=utf-8",
		ExpectedStatusCodes: []int{
			http.StatusNoContent,
		},
		HttpMethod:    http.MethodPatch,
		OptionsObject: options,
		Path:          fmt.Sprintf("%s/recycleBin", id.ID()),
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
