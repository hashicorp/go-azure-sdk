package driveitemanalyticsitemactivitystat

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

type CreateDriveItemAnalyticsItemActivityStatOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
	Model        *stable.ItemActivityStat
}

type CreateDriveItemAnalyticsItemActivityStatOperationOptions struct {
	Metadata *odata.Metadata
}

func DefaultCreateDriveItemAnalyticsItemActivityStatOperationOptions() CreateDriveItemAnalyticsItemActivityStatOperationOptions {
	return CreateDriveItemAnalyticsItemActivityStatOperationOptions{}
}

func (o CreateDriveItemAnalyticsItemActivityStatOperationOptions) ToHeaders() *client.Headers {
	out := client.Headers{}

	return &out
}

func (o CreateDriveItemAnalyticsItemActivityStatOperationOptions) ToOData() *odata.Query {
	out := odata.Query{}
	if o.Metadata != nil {
		out.Metadata = *o.Metadata
	}
	return &out
}

func (o CreateDriveItemAnalyticsItemActivityStatOperationOptions) ToQuery() *client.QueryParams {
	out := client.QueryParams{}

	return &out
}

// CreateDriveItemAnalyticsItemActivityStat - Create new navigation property to itemActivityStats for groups
func (c DriveItemAnalyticsItemActivityStatClient) CreateDriveItemAnalyticsItemActivityStat(ctx context.Context, id stable.GroupIdDriveIdItemId, input stable.ItemActivityStat, options CreateDriveItemAnalyticsItemActivityStatOperationOptions) (result CreateDriveItemAnalyticsItemActivityStatOperationResponse, err error) {
	opts := client.RequestOptions{
		ContentType: "application/json; charset=utf-8",
		ExpectedStatusCodes: []int{
			http.StatusCreated,
		},
		HttpMethod:    http.MethodPost,
		OptionsObject: options,
		Path:          fmt.Sprintf("%s/analytics/itemActivityStats", id.ID()),
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

	var model stable.ItemActivityStat
	result.Model = &model
	if err = resp.Unmarshal(result.Model); err != nil {
		return
	}

	return
}
