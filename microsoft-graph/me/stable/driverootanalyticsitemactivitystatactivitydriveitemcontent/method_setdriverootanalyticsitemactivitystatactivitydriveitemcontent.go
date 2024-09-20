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

type SetDriveRootAnalyticsItemActivityStatActivityDriveItemContentOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
}

type SetDriveRootAnalyticsItemActivityStatActivityDriveItemContentOperationOptions struct {
	Metadata *odata.Metadata
}

func DefaultSetDriveRootAnalyticsItemActivityStatActivityDriveItemContentOperationOptions() SetDriveRootAnalyticsItemActivityStatActivityDriveItemContentOperationOptions {
	return SetDriveRootAnalyticsItemActivityStatActivityDriveItemContentOperationOptions{}
}

func (o SetDriveRootAnalyticsItemActivityStatActivityDriveItemContentOperationOptions) ToHeaders() *client.Headers {
	out := client.Headers{}

	return &out
}

func (o SetDriveRootAnalyticsItemActivityStatActivityDriveItemContentOperationOptions) ToOData() *odata.Query {
	out := odata.Query{}
	if o.Metadata != nil {
		out.Metadata = *o.Metadata
	}
	return &out
}

func (o SetDriveRootAnalyticsItemActivityStatActivityDriveItemContentOperationOptions) ToQuery() *client.QueryParams {
	out := client.QueryParams{}

	return &out
}

// SetDriveRootAnalyticsItemActivityStatActivityDriveItemContent - Update content for the navigation property driveItem
// in me. The content stream, if the item represents a file.
func (c DriveRootAnalyticsItemActivityStatActivityDriveItemContentClient) SetDriveRootAnalyticsItemActivityStatActivityDriveItemContent(ctx context.Context, id stable.MeDriveIdRootAnalyticsItemActivityStatIdActivityId, input []byte, options SetDriveRootAnalyticsItemActivityStatActivityDriveItemContentOperationOptions) (result SetDriveRootAnalyticsItemActivityStatActivityDriveItemContentOperationResponse, err error) {
	opts := client.RequestOptions{
		ContentType: "application/json; charset=utf-8",
		ExpectedStatusCodes: []int{
			http.StatusNoContent,
		},
		HttpMethod:    http.MethodPut,
		OptionsObject: options,
		Path:          fmt.Sprintf("%s/driveItem/content", id.ID()),
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
