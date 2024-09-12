package driveitemanalyticslastsevenday

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

type GetDriveItemAnalyticsLastSevenDayOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
	Model        *beta.ItemActivityStat
}

type GetDriveItemAnalyticsLastSevenDayOperationOptions struct {
	Expand *odata.Expand
	Select *[]string
}

func DefaultGetDriveItemAnalyticsLastSevenDayOperationOptions() GetDriveItemAnalyticsLastSevenDayOperationOptions {
	return GetDriveItemAnalyticsLastSevenDayOperationOptions{}
}

func (o GetDriveItemAnalyticsLastSevenDayOperationOptions) ToHeaders() *client.Headers {
	out := client.Headers{}

	return &out
}

func (o GetDriveItemAnalyticsLastSevenDayOperationOptions) ToOData() *odata.Query {
	out := odata.Query{}
	if o.Expand != nil {
		out.Expand = *o.Expand
	}
	if o.Select != nil {
		out.Select = *o.Select
	}
	return &out
}

func (o GetDriveItemAnalyticsLastSevenDayOperationOptions) ToQuery() *client.QueryParams {
	out := client.QueryParams{}

	return &out
}

// GetDriveItemAnalyticsLastSevenDay - Get lastSevenDays from groups
func (c DriveItemAnalyticsLastSevenDayClient) GetDriveItemAnalyticsLastSevenDay(ctx context.Context, id beta.GroupIdDriveIdItemId, options GetDriveItemAnalyticsLastSevenDayOperationOptions) (result GetDriveItemAnalyticsLastSevenDayOperationResponse, err error) {
	opts := client.RequestOptions{
		ContentType: "application/json; charset=utf-8",
		ExpectedStatusCodes: []int{
			http.StatusOK,
		},
		HttpMethod:    http.MethodGet,
		OptionsObject: options,
		Path:          fmt.Sprintf("%s/analytics/lastSevenDays", id.ID()),
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

	var model beta.ItemActivityStat
	result.Model = &model
	if err = resp.Unmarshal(result.Model); err != nil {
		return
	}

	return
}
