package sitelistitemdriveitemcontent

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

type SetSiteListItemDriveItemContentOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
}

type SetSiteListItemDriveItemContentOperationOptions struct {
	ContentType string
	Metadata    *odata.Metadata
	RetryFunc   client.RequestRetryFunc
}

func DefaultSetSiteListItemDriveItemContentOperationOptions() SetSiteListItemDriveItemContentOperationOptions {
	return SetSiteListItemDriveItemContentOperationOptions{}
}

func (o SetSiteListItemDriveItemContentOperationOptions) ToHeaders() *client.Headers {
	out := client.Headers{}

	return &out
}

func (o SetSiteListItemDriveItemContentOperationOptions) ToOData() *odata.Query {
	out := odata.Query{}
	if o.Metadata != nil {
		out.Metadata = *o.Metadata
	}
	return &out
}

func (o SetSiteListItemDriveItemContentOperationOptions) ToQuery() *client.QueryParams {
	out := client.QueryParams{}

	return &out
}

// SetSiteListItemDriveItemContent - Update content for the navigation property driveItem in groups. The content stream,
// if the item represents a file.
func (c SiteListItemDriveItemContentClient) SetSiteListItemDriveItemContent(ctx context.Context, id stable.GroupIdSiteIdListIdItemId, input []byte, options SetSiteListItemDriveItemContentOperationOptions) (result SetSiteListItemDriveItemContentOperationResponse, err error) {
	opts := client.RequestOptions{
		ContentType: options.ContentType,
		ExpectedStatusCodes: []int{
			http.StatusAccepted,
			http.StatusCreated,
			http.StatusNoContent,
			http.StatusOK,
		},
		HttpMethod:    http.MethodPut,
		OptionsObject: options,
		Path:          fmt.Sprintf("%s/driveItem/content", id.ID()),
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
