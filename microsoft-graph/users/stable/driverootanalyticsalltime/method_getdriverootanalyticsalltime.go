package driverootanalyticsalltime

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

type GetDriveRootAnalyticsAllTimeOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
	Model        *stable.ItemActivityStat
}

type GetDriveRootAnalyticsAllTimeOperationOptions struct {
	Expand   *odata.Expand
	Metadata *odata.Metadata
	Select   *[]string
}

func DefaultGetDriveRootAnalyticsAllTimeOperationOptions() GetDriveRootAnalyticsAllTimeOperationOptions {
	return GetDriveRootAnalyticsAllTimeOperationOptions{}
}

func (o GetDriveRootAnalyticsAllTimeOperationOptions) ToHeaders() *client.Headers {
	out := client.Headers{}

	return &out
}

func (o GetDriveRootAnalyticsAllTimeOperationOptions) ToOData() *odata.Query {
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

func (o GetDriveRootAnalyticsAllTimeOperationOptions) ToQuery() *client.QueryParams {
	out := client.QueryParams{}

	return &out
}

// GetDriveRootAnalyticsAllTime - Get allTime from users
func (c DriveRootAnalyticsAllTimeClient) GetDriveRootAnalyticsAllTime(ctx context.Context, id stable.UserIdDriveId, options GetDriveRootAnalyticsAllTimeOperationOptions) (result GetDriveRootAnalyticsAllTimeOperationResponse, err error) {
	opts := client.RequestOptions{
		ContentType: "application/json; charset=utf-8",
		ExpectedStatusCodes: []int{
			http.StatusOK,
		},
		HttpMethod:    http.MethodGet,
		OptionsObject: options,
		Path:          fmt.Sprintf("%s/root/analytics/allTime", id.ID()),
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

	var model stable.ItemActivityStat
	result.Model = &model
	if err = resp.Unmarshal(result.Model); err != nil {
		return
	}

	return
}
