package sitelistitemactivitydriveitemcontent

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

type SetSiteListItemActivityDriveItemContentOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
}

type SetSiteListItemActivityDriveItemContentOperationOptions struct {
	ContentType string
	Metadata    *odata.Metadata
	RetryFunc   client.RequestRetryFunc
}

func DefaultSetSiteListItemActivityDriveItemContentOperationOptions() SetSiteListItemActivityDriveItemContentOperationOptions {
	return SetSiteListItemActivityDriveItemContentOperationOptions{}
}

func (o SetSiteListItemActivityDriveItemContentOperationOptions) ToHeaders() *client.Headers {
	out := client.Headers{}

	return &out
}

func (o SetSiteListItemActivityDriveItemContentOperationOptions) ToOData() *odata.Query {
	out := odata.Query{}
	if o.Metadata != nil {
		out.Metadata = *o.Metadata
	}
	return &out
}

func (o SetSiteListItemActivityDriveItemContentOperationOptions) ToQuery() *client.QueryParams {
	out := client.QueryParams{}

	return &out
}

// SetSiteListItemActivityDriveItemContent - Update content for the navigation property driveItem in groups. The content
// stream, if the item represents a file. The content property will have a potentially breaking change in behavior in
// the future. It will stream content directly instead of redirecting. To proactively opt in to the new behavior ahead
// of time, use the contentStream property instead.
func (c SiteListItemActivityDriveItemContentClient) SetSiteListItemActivityDriveItemContent(ctx context.Context, id beta.GroupIdSiteIdListIdItemIdActivityId, input []byte, options SetSiteListItemActivityDriveItemContentOperationOptions) (result SetSiteListItemActivityDriveItemContentOperationResponse, err error) {
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
