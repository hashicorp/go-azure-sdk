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

type SetDriveFollowingContentStreamOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
}

type SetDriveFollowingContentStreamOperationOptions struct {
	ContentType string
	Metadata    *odata.Metadata
	RetryFunc   client.RequestRetryFunc
}

func DefaultSetDriveFollowingContentStreamOperationOptions() SetDriveFollowingContentStreamOperationOptions {
	return SetDriveFollowingContentStreamOperationOptions{}
}

func (o SetDriveFollowingContentStreamOperationOptions) ToHeaders() *client.Headers {
	out := client.Headers{}

	return &out
}

func (o SetDriveFollowingContentStreamOperationOptions) ToOData() *odata.Query {
	out := odata.Query{}
	if o.Metadata != nil {
		out.Metadata = *o.Metadata
	}
	return &out
}

func (o SetDriveFollowingContentStreamOperationOptions) ToQuery() *client.QueryParams {
	out := client.QueryParams{}

	return &out
}

// SetDriveFollowingContentStream - Update contentStream for the navigation property following in users. The content
// stream, if the item represents a file.
func (c DriveFollowingContentStreamClient) SetDriveFollowingContentStream(ctx context.Context, id beta.UserIdDriveIdFollowingId, input []byte, options SetDriveFollowingContentStreamOperationOptions) (result SetDriveFollowingContentStreamOperationResponse, err error) {
	opts := client.RequestOptions{
		ContentType: options.ContentType,
		ExpectedStatusCodes: []int{
			http.StatusAccepted,
			http.StatusCreated,
			http.StatusNoContent,
			http.StatusOK,
		},
		HttpMethod:    http.MethodPut,
		OptionsObject: options,
		Path:          fmt.Sprintf("%s/contentStream", id.ID()),
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

	return
}
