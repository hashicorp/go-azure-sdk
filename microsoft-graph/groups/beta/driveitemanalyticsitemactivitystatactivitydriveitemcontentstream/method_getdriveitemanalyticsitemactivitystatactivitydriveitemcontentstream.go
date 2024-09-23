package driveitemanalyticsitemactivitystatactivitydriveitemcontentstream

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

type GetDriveItemAnalyticsItemActivityStatActivityDriveItemContentStreamOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
	Model        *[]byte
}

type GetDriveItemAnalyticsItemActivityStatActivityDriveItemContentStreamOperationOptions struct {
	Metadata  *odata.Metadata
	RetryFunc client.RequestRetryFunc
}

func DefaultGetDriveItemAnalyticsItemActivityStatActivityDriveItemContentStreamOperationOptions() GetDriveItemAnalyticsItemActivityStatActivityDriveItemContentStreamOperationOptions {
	return GetDriveItemAnalyticsItemActivityStatActivityDriveItemContentStreamOperationOptions{}
}

func (o GetDriveItemAnalyticsItemActivityStatActivityDriveItemContentStreamOperationOptions) ToHeaders() *client.Headers {
	out := client.Headers{}

	return &out
}

func (o GetDriveItemAnalyticsItemActivityStatActivityDriveItemContentStreamOperationOptions) ToOData() *odata.Query {
	out := odata.Query{}
	if o.Metadata != nil {
		out.Metadata = *o.Metadata
	}
	return &out
}

func (o GetDriveItemAnalyticsItemActivityStatActivityDriveItemContentStreamOperationOptions) ToQuery() *client.QueryParams {
	out := client.QueryParams{}

	return &out
}

// GetDriveItemAnalyticsItemActivityStatActivityDriveItemContentStream - Get contentStream for the navigation property
// driveItem from groups. The content stream, if the item represents a file.
func (c DriveItemAnalyticsItemActivityStatActivityDriveItemContentStreamClient) GetDriveItemAnalyticsItemActivityStatActivityDriveItemContentStream(ctx context.Context, id beta.GroupIdDriveIdItemIdAnalyticsItemActivityStatIdActivityId, options GetDriveItemAnalyticsItemActivityStatActivityDriveItemContentStreamOperationOptions) (result GetDriveItemAnalyticsItemActivityStatActivityDriveItemContentStreamOperationResponse, err error) {
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
