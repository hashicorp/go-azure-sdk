package driverootanalyticsitemactivitystatactivitydriveitemcontent

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

type GetDriveRootAnalyticsItemActivityStatActivityDriveItemContentOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
	Model        *[]byte
}

type GetDriveRootAnalyticsItemActivityStatActivityDriveItemContentOperationOptions struct {
	Format   *odata.Format
	Metadata *odata.Metadata
}

func DefaultGetDriveRootAnalyticsItemActivityStatActivityDriveItemContentOperationOptions() GetDriveRootAnalyticsItemActivityStatActivityDriveItemContentOperationOptions {
	return GetDriveRootAnalyticsItemActivityStatActivityDriveItemContentOperationOptions{}
}

func (o GetDriveRootAnalyticsItemActivityStatActivityDriveItemContentOperationOptions) ToHeaders() *client.Headers {
	out := client.Headers{}

	return &out
}

func (o GetDriveRootAnalyticsItemActivityStatActivityDriveItemContentOperationOptions) ToOData() *odata.Query {
	out := odata.Query{}
	if o.Format != nil {
		out.Format = *o.Format
	}
	if o.Metadata != nil {
		out.Metadata = *o.Metadata
	}
	return &out
}

func (o GetDriveRootAnalyticsItemActivityStatActivityDriveItemContentOperationOptions) ToQuery() *client.QueryParams {
	out := client.QueryParams{}

	return &out
}

// GetDriveRootAnalyticsItemActivityStatActivityDriveItemContent - Get content for the navigation property driveItem
// from me. The content stream, if the item represents a file. The content property will have a potentially breaking
// change in behavior in the future. It will stream content directly instead of redirecting. To proactively opt in to
// the new behavior ahead of time, use the contentStream property instead.
func (c DriveRootAnalyticsItemActivityStatActivityDriveItemContentClient) GetDriveRootAnalyticsItemActivityStatActivityDriveItemContent(ctx context.Context, id beta.MeDriveIdRootAnalyticsItemActivityStatIdActivityId, options GetDriveRootAnalyticsItemActivityStatActivityDriveItemContentOperationOptions) (result GetDriveRootAnalyticsItemActivityStatActivityDriveItemContentOperationResponse, err error) {
	opts := client.RequestOptions{
		ContentType: "application/octet-stream",
		ExpectedStatusCodes: []int{
			http.StatusOK,
		},
		HttpMethod:    http.MethodGet,
		OptionsObject: options,
		Path:          fmt.Sprintf("%s/driveItem/content", id.ID()),
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

	var model []byte
	result.Model = &model
	if err = resp.Unmarshal(result.Model); err != nil {
		return
	}

	return
}
