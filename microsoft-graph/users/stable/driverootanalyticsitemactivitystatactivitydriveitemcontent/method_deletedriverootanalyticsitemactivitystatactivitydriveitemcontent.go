package driverootanalyticsitemactivitystatactivitydriveitemcontent

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

type DeleteDriveRootAnalyticsItemActivityStatActivityDriveItemContentOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
}

type DeleteDriveRootAnalyticsItemActivityStatActivityDriveItemContentOperationOptions struct {
	IfMatch   *string
	Metadata  *odata.Metadata
	RetryFunc client.RequestRetryFunc
}

func DefaultDeleteDriveRootAnalyticsItemActivityStatActivityDriveItemContentOperationOptions() DeleteDriveRootAnalyticsItemActivityStatActivityDriveItemContentOperationOptions {
	return DeleteDriveRootAnalyticsItemActivityStatActivityDriveItemContentOperationOptions{}
}

func (o DeleteDriveRootAnalyticsItemActivityStatActivityDriveItemContentOperationOptions) ToHeaders() *client.Headers {
	out := client.Headers{}
	if o.IfMatch != nil {
		out.Append("If-Match", fmt.Sprintf("%v", *o.IfMatch))
	}
	return &out
}

func (o DeleteDriveRootAnalyticsItemActivityStatActivityDriveItemContentOperationOptions) ToOData() *odata.Query {
	out := odata.Query{}
	if o.Metadata != nil {
		out.Metadata = *o.Metadata
	}
	return &out
}

func (o DeleteDriveRootAnalyticsItemActivityStatActivityDriveItemContentOperationOptions) ToQuery() *client.QueryParams {
	out := client.QueryParams{}

	return &out
}

// DeleteDriveRootAnalyticsItemActivityStatActivityDriveItemContent - Delete content for the navigation property
// driveItem in users. The content stream, if the item represents a file.
func (c DriveRootAnalyticsItemActivityStatActivityDriveItemContentClient) DeleteDriveRootAnalyticsItemActivityStatActivityDriveItemContent(ctx context.Context, id stable.UserIdDriveIdRootAnalyticsItemActivityStatIdActivityId, options DeleteDriveRootAnalyticsItemActivityStatActivityDriveItemContentOperationOptions) (result DeleteDriveRootAnalyticsItemActivityStatActivityDriveItemContentOperationResponse, err error) {
	opts := client.RequestOptions{
		ContentType: "application/json; charset=utf-8",
		ExpectedStatusCodes: []int{
			http.StatusNoContent,
			http.StatusOK,
		},
		HttpMethod:    http.MethodDelete,
		OptionsObject: options,
		Path:          fmt.Sprintf("%s/driveItem/content", id.ID()),
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

	return
}
