package driverootanalyticsitemactivitystatactivitydriveitemcontentstream

import (
	"context"
	"fmt"
	"net/http"

	"github.com/hashicorp/go-azure-sdk/microsoft-graph/common-types/beta"
	"github.com/hashicorp/go-azure-sdk/sdk/client"
	"github.com/hashicorp/go-azure-sdk/sdk/odata"
)

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type GetDriveRootAnalyticsItemActivityStatActivityDriveItemContentStreamOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
	Model        *[]byte
}

type GetDriveRootAnalyticsItemActivityStatActivityDriveItemContentStreamOperationOptions struct {
	Metadata  *odata.Metadata
	RetryFunc client.RequestRetryFunc
}

func DefaultGetDriveRootAnalyticsItemActivityStatActivityDriveItemContentStreamOperationOptions() GetDriveRootAnalyticsItemActivityStatActivityDriveItemContentStreamOperationOptions {
	return GetDriveRootAnalyticsItemActivityStatActivityDriveItemContentStreamOperationOptions{}
}

func (o GetDriveRootAnalyticsItemActivityStatActivityDriveItemContentStreamOperationOptions) ToHeaders() *client.Headers {
	out := client.Headers{}

	return &out
}

func (o GetDriveRootAnalyticsItemActivityStatActivityDriveItemContentStreamOperationOptions) ToOData() *odata.Query {
	out := odata.Query{}
	if o.Metadata != nil {
		out.Metadata = *o.Metadata
	}
	return &out
}

func (o GetDriveRootAnalyticsItemActivityStatActivityDriveItemContentStreamOperationOptions) ToQuery() *client.QueryParams {
	out := client.QueryParams{}

	return &out
}

// GetDriveRootAnalyticsItemActivityStatActivityDriveItemContentStream - Get contentStream for the navigation property
// driveItem from users. The content stream, if the item represents a file.
func (c DriveRootAnalyticsItemActivityStatActivityDriveItemContentStreamClient) GetDriveRootAnalyticsItemActivityStatActivityDriveItemContentStream(ctx context.Context, id beta.UserIdDriveIdRootAnalyticsItemActivityStatIdActivityId, options GetDriveRootAnalyticsItemActivityStatActivityDriveItemContentStreamOperationOptions) (result GetDriveRootAnalyticsItemActivityStatActivityDriveItemContentStreamOperationResponse, err error) {
	opts := client.RequestOptions{
		ContentType: "application/octet-stream",
		ExpectedStatusCodes: []int{
			http.StatusOK,
		},
		HttpMethod:    http.MethodGet,
		OptionsObject: options,
		Path:          fmt.Sprintf("%s/driveItem/contentStream", id.ID()),
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

	var model []byte
	result.Model = &model
	if err = resp.Unmarshal(result.Model); err != nil {
		return
	}

	return
}
