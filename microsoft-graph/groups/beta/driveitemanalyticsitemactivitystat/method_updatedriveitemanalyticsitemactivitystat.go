package driveitemanalyticsitemactivitystat

import (
	"context"
	"net/http"

	"github.com/hashicorp/go-azure-sdk/microsoft-graph/common-types/beta"
	"github.com/hashicorp/go-azure-sdk/sdk/client"
	"github.com/hashicorp/go-azure-sdk/sdk/odata"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type UpdateDriveItemAnalyticsItemActivityStatOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
}

type UpdateDriveItemAnalyticsItemActivityStatOperationOptions struct {
	Metadata *odata.Metadata
}

func DefaultUpdateDriveItemAnalyticsItemActivityStatOperationOptions() UpdateDriveItemAnalyticsItemActivityStatOperationOptions {
	return UpdateDriveItemAnalyticsItemActivityStatOperationOptions{}
}

func (o UpdateDriveItemAnalyticsItemActivityStatOperationOptions) ToHeaders() *client.Headers {
	out := client.Headers{}

	return &out
}

func (o UpdateDriveItemAnalyticsItemActivityStatOperationOptions) ToOData() *odata.Query {
	out := odata.Query{}
	if o.Metadata != nil {
		out.Metadata = *o.Metadata
	}
	return &out
}

func (o UpdateDriveItemAnalyticsItemActivityStatOperationOptions) ToQuery() *client.QueryParams {
	out := client.QueryParams{}

	return &out
}

// UpdateDriveItemAnalyticsItemActivityStat - Update the navigation property itemActivityStats in groups
func (c DriveItemAnalyticsItemActivityStatClient) UpdateDriveItemAnalyticsItemActivityStat(ctx context.Context, id beta.GroupIdDriveIdItemIdAnalyticsItemActivityStatId, input beta.ItemActivityStat, options UpdateDriveItemAnalyticsItemActivityStatOperationOptions) (result UpdateDriveItemAnalyticsItemActivityStatOperationResponse, err error) {
	opts := client.RequestOptions{
		ContentType: "application/json; charset=utf-8",
		ExpectedStatusCodes: []int{
			http.StatusNoContent,
		},
		HttpMethod:    http.MethodPatch,
		OptionsObject: options,
		Path:          id.ID(),
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
