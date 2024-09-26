package drivelistitemdriveitemcontentstream

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

type SetDriveListItemDriveItemContentStreamOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
}

type SetDriveListItemDriveItemContentStreamOperationOptions struct {
	ContentType string
	Metadata    *odata.Metadata
	RetryFunc   client.RequestRetryFunc
}

func DefaultSetDriveListItemDriveItemContentStreamOperationOptions() SetDriveListItemDriveItemContentStreamOperationOptions {
	return SetDriveListItemDriveItemContentStreamOperationOptions{}
}

func (o SetDriveListItemDriveItemContentStreamOperationOptions) ToHeaders() *client.Headers {
	out := client.Headers{}

	return &out
}

func (o SetDriveListItemDriveItemContentStreamOperationOptions) ToOData() *odata.Query {
	out := odata.Query{}
	if o.Metadata != nil {
		out.Metadata = *o.Metadata
	}
	return &out
}

func (o SetDriveListItemDriveItemContentStreamOperationOptions) ToQuery() *client.QueryParams {
	out := client.QueryParams{}

	return &out
}

// SetDriveListItemDriveItemContentStream - Update contentStream for the navigation property driveItem in me. The
// content stream, if the item represents a file.
func (c DriveListItemDriveItemContentStreamClient) SetDriveListItemDriveItemContentStream(ctx context.Context, id beta.MeDriveIdListItemId, input []byte, options SetDriveListItemDriveItemContentStreamOperationOptions) (result SetDriveListItemDriveItemContentStreamOperationResponse, err error) {
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
		Path:          fmt.Sprintf("%s/driveItem/contentStream", id.ID()),
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
