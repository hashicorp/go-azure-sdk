package driveroot

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

type UnfollowDriveRootOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
}

type UnfollowDriveRootOperationOptions struct {
	Metadata  *odata.Metadata
	RetryFunc client.RequestRetryFunc
}

func DefaultUnfollowDriveRootOperationOptions() UnfollowDriveRootOperationOptions {
	return UnfollowDriveRootOperationOptions{}
}

func (o UnfollowDriveRootOperationOptions) ToHeaders() *client.Headers {
	out := client.Headers{}

	return &out
}

func (o UnfollowDriveRootOperationOptions) ToOData() *odata.Query {
	out := odata.Query{}
	if o.Metadata != nil {
		out.Metadata = *o.Metadata
	}
	return &out
}

func (o UnfollowDriveRootOperationOptions) ToQuery() *client.QueryParams {
	out := client.QueryParams{}

	return &out
}

// UnfollowDriveRoot - Invoke action unfollow. Unfollow a driveItem.
func (c DriveRootClient) UnfollowDriveRoot(ctx context.Context, id stable.MeDriveId, options UnfollowDriveRootOperationOptions) (result UnfollowDriveRootOperationResponse, err error) {
	opts := client.RequestOptions{
		ContentType: "application/json; charset=utf-8",
		ExpectedStatusCodes: []int{
			http.StatusNoContent,
		},
		HttpMethod:    http.MethodPost,
		OptionsObject: options,
		Path:          fmt.Sprintf("%s/root/unfollow", id.ID()),
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

	return
}
