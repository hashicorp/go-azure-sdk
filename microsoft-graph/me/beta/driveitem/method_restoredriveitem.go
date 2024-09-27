package driveitem

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

type RestoreDriveItemOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
	Model        *beta.DriveItem
}

type RestoreDriveItemOperationOptions struct {
	Metadata  *odata.Metadata
	RetryFunc client.RequestRetryFunc
}

func DefaultRestoreDriveItemOperationOptions() RestoreDriveItemOperationOptions {
	return RestoreDriveItemOperationOptions{}
}

func (o RestoreDriveItemOperationOptions) ToHeaders() *client.Headers {
	out := client.Headers{}

	return &out
}

func (o RestoreDriveItemOperationOptions) ToOData() *odata.Query {
	out := odata.Query{}
	if o.Metadata != nil {
		out.Metadata = *o.Metadata
	}
	return &out
}

func (o RestoreDriveItemOperationOptions) ToQuery() *client.QueryParams {
	out := client.QueryParams{}

	return &out
}

// RestoreDriveItem - Invoke action restore. Restore a driveItem that has been deleted and is currently in the recycle
// bin.
func (c DriveItemClient) RestoreDriveItem(ctx context.Context, id beta.MeDriveIdItemId, input RestoreDriveItemRequest, options RestoreDriveItemOperationOptions) (result RestoreDriveItemOperationResponse, err error) {
	opts := client.RequestOptions{
		ContentType: "application/json; charset=utf-8",
		ExpectedStatusCodes: []int{
			http.StatusAccepted,
			http.StatusCreated,
			http.StatusNoContent,
			http.StatusOK,
		},
		HttpMethod:    http.MethodPost,
		OptionsObject: options,
		Path:          fmt.Sprintf("%s/restore", id.ID()),
		RetryFunc:     options.RetryFunc,
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

	var model beta.DriveItem
	result.Model = &model
	if err = resp.Unmarshal(result.Model); err != nil {
		return
	}

	return
}
