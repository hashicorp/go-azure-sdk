package sitecontentmodel

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

type RemoveSiteContentModelFromDriveOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
}

type RemoveSiteContentModelFromDriveOperationOptions struct {
	Metadata *odata.Metadata
}

func DefaultRemoveSiteContentModelFromDriveOperationOptions() RemoveSiteContentModelFromDriveOperationOptions {
	return RemoveSiteContentModelFromDriveOperationOptions{}
}

func (o RemoveSiteContentModelFromDriveOperationOptions) ToHeaders() *client.Headers {
	out := client.Headers{}

	return &out
}

func (o RemoveSiteContentModelFromDriveOperationOptions) ToOData() *odata.Query {
	out := odata.Query{}
	if o.Metadata != nil {
		out.Metadata = *o.Metadata
	}
	return &out
}

func (o RemoveSiteContentModelFromDriveOperationOptions) ToQuery() *client.QueryParams {
	out := client.QueryParams{}

	return &out
}

// RemoveSiteContentModelFromDrive - Invoke action removeFromDrive. Remove a contentModel from a SharePoint document
// library.
func (c SiteContentModelClient) RemoveSiteContentModelFromDrive(ctx context.Context, id beta.GroupIdSiteIdContentModelId, input RemoveSiteContentModelFromDriveRequest, options RemoveSiteContentModelFromDriveOperationOptions) (result RemoveSiteContentModelFromDriveOperationResponse, err error) {
	opts := client.RequestOptions{
		ContentType: "application/json; charset=utf-8",
		ExpectedStatusCodes: []int{
			http.StatusNoContent,
		},
		HttpMethod:    http.MethodPost,
		OptionsObject: options,
		Path:          fmt.Sprintf("%s/removeFromDrive", id.ID()),
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
