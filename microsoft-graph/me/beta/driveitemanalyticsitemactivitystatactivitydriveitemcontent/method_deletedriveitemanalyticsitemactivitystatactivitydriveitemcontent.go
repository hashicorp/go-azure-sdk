package driveitemanalyticsitemactivitystatactivitydriveitemcontent

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

type DeleteDriveItemAnalyticsItemActivityStatActivityDriveItemContentOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
}

type DeleteDriveItemAnalyticsItemActivityStatActivityDriveItemContentOperationOptions struct {
	IfMatch  *string
	Metadata *odata.Metadata
}

func DefaultDeleteDriveItemAnalyticsItemActivityStatActivityDriveItemContentOperationOptions() DeleteDriveItemAnalyticsItemActivityStatActivityDriveItemContentOperationOptions {
	return DeleteDriveItemAnalyticsItemActivityStatActivityDriveItemContentOperationOptions{}
}

func (o DeleteDriveItemAnalyticsItemActivityStatActivityDriveItemContentOperationOptions) ToHeaders() *client.Headers {
	out := client.Headers{}
	if o.IfMatch != nil {
		out.Append("If-Match", fmt.Sprintf("%v", *o.IfMatch))
	}
	return &out
}

func (o DeleteDriveItemAnalyticsItemActivityStatActivityDriveItemContentOperationOptions) ToOData() *odata.Query {
	out := odata.Query{}
	if o.Metadata != nil {
		out.Metadata = *o.Metadata
	}
	return &out
}

func (o DeleteDriveItemAnalyticsItemActivityStatActivityDriveItemContentOperationOptions) ToQuery() *client.QueryParams {
	out := client.QueryParams{}

	return &out
}

// DeleteDriveItemAnalyticsItemActivityStatActivityDriveItemContent - Delete content for the navigation property
// driveItem in me. The content stream, if the item represents a file. The content property will have a potentially
// breaking change in behavior in the future. It will stream content directly instead of redirecting. To proactively opt
// in to the new behavior ahead of time, use the contentStream property instead.
func (c DriveItemAnalyticsItemActivityStatActivityDriveItemContentClient) DeleteDriveItemAnalyticsItemActivityStatActivityDriveItemContent(ctx context.Context, id beta.MeDriveIdItemIdAnalyticsItemActivityStatIdActivityId, options DeleteDriveItemAnalyticsItemActivityStatActivityDriveItemContentOperationOptions) (result DeleteDriveItemAnalyticsItemActivityStatActivityDriveItemContentOperationResponse, err error) {
	opts := client.RequestOptions{
		ContentType: "application/json; charset=utf-8",
		ExpectedStatusCodes: []int{
			http.StatusNoContent,
		},
		HttpMethod:    http.MethodDelete,
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

	return
}
