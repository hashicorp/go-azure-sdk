package driveitemanalyticsitemactivitystatactivitydriveitemcontent

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

type SetDriveItemAnalyticsItemActivityStatActivityDriveItemContentOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
}

type SetDriveItemAnalyticsItemActivityStatActivityDriveItemContentOperationOptions struct {
	Metadata *odata.Metadata
}

func DefaultSetDriveItemAnalyticsItemActivityStatActivityDriveItemContentOperationOptions() SetDriveItemAnalyticsItemActivityStatActivityDriveItemContentOperationOptions {
	return SetDriveItemAnalyticsItemActivityStatActivityDriveItemContentOperationOptions{}
}

func (o SetDriveItemAnalyticsItemActivityStatActivityDriveItemContentOperationOptions) ToHeaders() *client.Headers {
	out := client.Headers{}

	return &out
}

func (o SetDriveItemAnalyticsItemActivityStatActivityDriveItemContentOperationOptions) ToOData() *odata.Query {
	out := odata.Query{}
	if o.Metadata != nil {
		out.Metadata = *o.Metadata
	}
	return &out
}

func (o SetDriveItemAnalyticsItemActivityStatActivityDriveItemContentOperationOptions) ToQuery() *client.QueryParams {
	out := client.QueryParams{}

	return &out
}

// SetDriveItemAnalyticsItemActivityStatActivityDriveItemContent - Update content for the navigation property driveItem
// in me. The content stream, if the item represents a file.
func (c DriveItemAnalyticsItemActivityStatActivityDriveItemContentClient) SetDriveItemAnalyticsItemActivityStatActivityDriveItemContent(ctx context.Context, id stable.MeDriveIdItemIdAnalyticsItemActivityStatIdActivityId, input []byte, options SetDriveItemAnalyticsItemActivityStatActivityDriveItemContentOperationOptions) (result SetDriveItemAnalyticsItemActivityStatActivityDriveItemContentOperationResponse, err error) {
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
