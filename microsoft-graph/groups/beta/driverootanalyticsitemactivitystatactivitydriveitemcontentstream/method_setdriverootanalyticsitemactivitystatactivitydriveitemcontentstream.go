package driverootanalyticsitemactivitystatactivitydriveitemcontentstream

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

type SetDriveRootAnalyticsItemActivityStatActivityDriveItemContentStreamOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
}

type SetDriveRootAnalyticsItemActivityStatActivityDriveItemContentStreamOperationOptions struct {
	Metadata *odata.Metadata
}

func DefaultSetDriveRootAnalyticsItemActivityStatActivityDriveItemContentStreamOperationOptions() SetDriveRootAnalyticsItemActivityStatActivityDriveItemContentStreamOperationOptions {
	return SetDriveRootAnalyticsItemActivityStatActivityDriveItemContentStreamOperationOptions{}
}

func (o SetDriveRootAnalyticsItemActivityStatActivityDriveItemContentStreamOperationOptions) ToHeaders() *client.Headers {
	out := client.Headers{}

	return &out
}

func (o SetDriveRootAnalyticsItemActivityStatActivityDriveItemContentStreamOperationOptions) ToOData() *odata.Query {
	out := odata.Query{}
	if o.Metadata != nil {
		out.Metadata = *o.Metadata
	}
	return &out
}

func (o SetDriveRootAnalyticsItemActivityStatActivityDriveItemContentStreamOperationOptions) ToQuery() *client.QueryParams {
	out := client.QueryParams{}

	return &out
}

// SetDriveRootAnalyticsItemActivityStatActivityDriveItemContentStream - Update contentStream for the navigation
// property driveItem in groups. The content stream, if the item represents a file.
func (c DriveRootAnalyticsItemActivityStatActivityDriveItemContentStreamClient) SetDriveRootAnalyticsItemActivityStatActivityDriveItemContentStream(ctx context.Context, id beta.GroupIdDriveIdRootAnalyticsItemActivityStatIdActivityId, input []byte, options SetDriveRootAnalyticsItemActivityStatActivityDriveItemContentStreamOperationOptions) (result SetDriveRootAnalyticsItemActivityStatActivityDriveItemContentStreamOperationResponse, err error) {
	opts := client.RequestOptions{
		ContentType: "application/json; charset=utf-8",
		ExpectedStatusCodes: []int{
			http.StatusNoContent,
		},
		HttpMethod:    http.MethodPut,
		OptionsObject: options,
		Path:          fmt.Sprintf("%s/driveItem/contentStream", id.ID()),
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
