package driverootlistitemactivity

import (
	"context"
	"net/http"

	"github.com/hashicorp/go-azure-sdk/microsoft-graph/common-types/beta"
	"github.com/hashicorp/go-azure-sdk/sdk/client"
	"github.com/hashicorp/go-azure-sdk/sdk/odata"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type UpdateDriveRootListItemActivityOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
}

type UpdateDriveRootListItemActivityOperationOptions struct {
	Metadata *odata.Metadata
}

func DefaultUpdateDriveRootListItemActivityOperationOptions() UpdateDriveRootListItemActivityOperationOptions {
	return UpdateDriveRootListItemActivityOperationOptions{}
}

func (o UpdateDriveRootListItemActivityOperationOptions) ToHeaders() *client.Headers {
	out := client.Headers{}

	return &out
}

func (o UpdateDriveRootListItemActivityOperationOptions) ToOData() *odata.Query {
	out := odata.Query{}
	if o.Metadata != nil {
		out.Metadata = *o.Metadata
	}
	return &out
}

func (o UpdateDriveRootListItemActivityOperationOptions) ToQuery() *client.QueryParams {
	out := client.QueryParams{}

	return &out
}

// UpdateDriveRootListItemActivity - Update the navigation property activities in groups
func (c DriveRootListItemActivityClient) UpdateDriveRootListItemActivity(ctx context.Context, id beta.GroupIdDriveIdRootListItemActivityId, input beta.ItemActivityOLD, options UpdateDriveRootListItemActivityOperationOptions) (result UpdateDriveRootListItemActivityOperationResponse, err error) {
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
