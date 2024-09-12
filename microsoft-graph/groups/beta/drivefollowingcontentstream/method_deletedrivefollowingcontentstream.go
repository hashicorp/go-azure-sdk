package drivefollowingcontentstream

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

type DeleteDriveFollowingContentStreamOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
}

type DeleteDriveFollowingContentStreamOperationOptions struct {
	IfMatch *string
}

func DefaultDeleteDriveFollowingContentStreamOperationOptions() DeleteDriveFollowingContentStreamOperationOptions {
	return DeleteDriveFollowingContentStreamOperationOptions{}
}

func (o DeleteDriveFollowingContentStreamOperationOptions) ToHeaders() *client.Headers {
	out := client.Headers{}
	if o.IfMatch != nil {
		out.Append("If-Match", fmt.Sprintf("%v", *o.IfMatch))
	}
	return &out
}

func (o DeleteDriveFollowingContentStreamOperationOptions) ToOData() *odata.Query {
	out := odata.Query{}

	return &out
}

func (o DeleteDriveFollowingContentStreamOperationOptions) ToQuery() *client.QueryParams {
	out := client.QueryParams{}

	return &out
}

// DeleteDriveFollowingContentStream - Delete contentStream for the navigation property following in groups. The content
// stream, if the item represents a file.
func (c DriveFollowingContentStreamClient) DeleteDriveFollowingContentStream(ctx context.Context, id beta.GroupIdDriveIdFollowingId, options DeleteDriveFollowingContentStreamOperationOptions) (result DeleteDriveFollowingContentStreamOperationResponse, err error) {
	opts := client.RequestOptions{
		ContentType: "application/json; charset=utf-8",
		ExpectedStatusCodes: []int{
			http.StatusNoContent,
		},
		HttpMethod:    http.MethodDelete,
		OptionsObject: options,
		Path:          fmt.Sprintf("%s/contentStream", id.ID()),
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
