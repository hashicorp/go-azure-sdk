package driveitemanalyticsitemactivitystatactivitydriveitem

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

type GetDriveItemAnalyticsItemActivityStatActivityDriveItemOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
	Model        *stable.DriveItem
}

type GetDriveItemAnalyticsItemActivityStatActivityDriveItemOperationOptions struct {
	Expand   *odata.Expand
	Metadata *odata.Metadata
	Select   *[]string
}

func DefaultGetDriveItemAnalyticsItemActivityStatActivityDriveItemOperationOptions() GetDriveItemAnalyticsItemActivityStatActivityDriveItemOperationOptions {
	return GetDriveItemAnalyticsItemActivityStatActivityDriveItemOperationOptions{}
}

func (o GetDriveItemAnalyticsItemActivityStatActivityDriveItemOperationOptions) ToHeaders() *client.Headers {
	out := client.Headers{}

	return &out
}

func (o GetDriveItemAnalyticsItemActivityStatActivityDriveItemOperationOptions) ToOData() *odata.Query {
	out := odata.Query{}
	if o.Expand != nil {
		out.Expand = *o.Expand
	}
	if o.Metadata != nil {
		out.Metadata = *o.Metadata
	}
	if o.Select != nil {
		out.Select = *o.Select
	}
	return &out
}

func (o GetDriveItemAnalyticsItemActivityStatActivityDriveItemOperationOptions) ToQuery() *client.QueryParams {
	out := client.QueryParams{}

	return &out
}

// GetDriveItemAnalyticsItemActivityStatActivityDriveItem - Get driveItem from groups. Exposes the driveItem that was
// the target of this activity.
func (c DriveItemAnalyticsItemActivityStatActivityDriveItemClient) GetDriveItemAnalyticsItemActivityStatActivityDriveItem(ctx context.Context, id stable.GroupIdDriveIdItemIdAnalyticsItemActivityStatIdActivityId, options GetDriveItemAnalyticsItemActivityStatActivityDriveItemOperationOptions) (result GetDriveItemAnalyticsItemActivityStatActivityDriveItemOperationResponse, err error) {
	opts := client.RequestOptions{
		ContentType: "application/json; charset=utf-8",
		ExpectedStatusCodes: []int{
			http.StatusOK,
		},
		HttpMethod:    http.MethodGet,
		OptionsObject: options,
		Path:          fmt.Sprintf("%s/driveItem", id.ID()),
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

	var model stable.DriveItem
	result.Model = &model
	if err = resp.Unmarshal(result.Model); err != nil {
		return
	}

	return
}
