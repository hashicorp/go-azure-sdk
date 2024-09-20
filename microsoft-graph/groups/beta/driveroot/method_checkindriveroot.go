package driveroot

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

type CheckinDriveRootOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
}

type CheckinDriveRootOperationOptions struct {
	Metadata *odata.Metadata
}

func DefaultCheckinDriveRootOperationOptions() CheckinDriveRootOperationOptions {
	return CheckinDriveRootOperationOptions{}
}

func (o CheckinDriveRootOperationOptions) ToHeaders() *client.Headers {
	out := client.Headers{}

	return &out
}

func (o CheckinDriveRootOperationOptions) ToOData() *odata.Query {
	out := odata.Query{}
	if o.Metadata != nil {
		out.Metadata = *o.Metadata
	}
	return &out
}

func (o CheckinDriveRootOperationOptions) ToQuery() *client.QueryParams {
	out := client.QueryParams{}

	return &out
}

// CheckinDriveRoot - Invoke action checkin. Check in a checked out driveItem resource, which makes the version of the
// document available to others.
func (c DriveRootClient) CheckinDriveRoot(ctx context.Context, id beta.GroupIdDriveId, input CheckinDriveRootRequest, options CheckinDriveRootOperationOptions) (result CheckinDriveRootOperationResponse, err error) {
	opts := client.RequestOptions{
		ContentType: "application/json; charset=utf-8",
		ExpectedStatusCodes: []int{
			http.StatusNoContent,
		},
		HttpMethod:    http.MethodPost,
		OptionsObject: options,
		Path:          fmt.Sprintf("%s/root/checkin", id.ID()),
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
