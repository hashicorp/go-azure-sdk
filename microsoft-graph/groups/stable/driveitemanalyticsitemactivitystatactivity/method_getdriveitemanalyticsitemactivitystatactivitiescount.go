package driveitemanalyticsitemactivitystatactivity

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

type GetDriveItemAnalyticsItemActivityStatActivitiesCountOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
	Model        *[]byte
}

type GetDriveItemAnalyticsItemActivityStatActivitiesCountOperationOptions struct {
	Filter   *string
	Metadata *odata.Metadata
	Search   *string
}

func DefaultGetDriveItemAnalyticsItemActivityStatActivitiesCountOperationOptions() GetDriveItemAnalyticsItemActivityStatActivitiesCountOperationOptions {
	return GetDriveItemAnalyticsItemActivityStatActivitiesCountOperationOptions{}
}

func (o GetDriveItemAnalyticsItemActivityStatActivitiesCountOperationOptions) ToHeaders() *client.Headers {
	out := client.Headers{}

	return &out
}

func (o GetDriveItemAnalyticsItemActivityStatActivitiesCountOperationOptions) ToOData() *odata.Query {
	out := odata.Query{}
	if o.Filter != nil {
		out.Filter = *o.Filter
	}
	if o.Metadata != nil {
		out.Metadata = *o.Metadata
	}
	if o.Search != nil {
		out.Search = *o.Search
	}
	return &out
}

func (o GetDriveItemAnalyticsItemActivityStatActivitiesCountOperationOptions) ToQuery() *client.QueryParams {
	out := client.QueryParams{}

	return &out
}

// GetDriveItemAnalyticsItemActivityStatActivitiesCount - Get the number of the resource
func (c DriveItemAnalyticsItemActivityStatActivityClient) GetDriveItemAnalyticsItemActivityStatActivitiesCount(ctx context.Context, id stable.GroupIdDriveIdItemIdAnalyticsItemActivityStatId, options GetDriveItemAnalyticsItemActivityStatActivitiesCountOperationOptions) (result GetDriveItemAnalyticsItemActivityStatActivitiesCountOperationResponse, err error) {
	opts := client.RequestOptions{
		ContentType: "text/plain",
		ExpectedStatusCodes: []int{
			http.StatusOK,
		},
		HttpMethod:    http.MethodGet,
		OptionsObject: options,
		Path:          fmt.Sprintf("%s/activities/$count", id.ID()),
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
