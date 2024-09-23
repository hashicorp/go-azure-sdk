package drivelistitemactivitydriveitemcontent

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

type SetDriveListItemActivityDriveItemContentOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
}

type SetDriveListItemActivityDriveItemContentOperationOptions struct {
	ContentType string
	Metadata    *odata.Metadata
	RetryFunc   client.RequestRetryFunc
}

func DefaultSetDriveListItemActivityDriveItemContentOperationOptions() SetDriveListItemActivityDriveItemContentOperationOptions {
	return SetDriveListItemActivityDriveItemContentOperationOptions{}
}

func (o SetDriveListItemActivityDriveItemContentOperationOptions) ToHeaders() *client.Headers {
	out := client.Headers{}

	return &out
}

func (o SetDriveListItemActivityDriveItemContentOperationOptions) ToOData() *odata.Query {
	out := odata.Query{}
	if o.Metadata != nil {
		out.Metadata = *o.Metadata
	}
	return &out
}

func (o SetDriveListItemActivityDriveItemContentOperationOptions) ToQuery() *client.QueryParams {
	out := client.QueryParams{}

	return &out
}

// SetDriveListItemActivityDriveItemContent - Update content for the navigation property driveItem in users. The content
// stream, if the item represents a file. The content property will have a potentially breaking change in behavior in
// the future. It will stream content directly instead of redirecting. To proactively opt in to the new behavior ahead
// of time, use the contentStream property instead.
func (c DriveListItemActivityDriveItemContentClient) SetDriveListItemActivityDriveItemContent(ctx context.Context, id beta.UserIdDriveIdListItemIdActivityId, input []byte, options SetDriveListItemActivityDriveItemContentOperationOptions) (result SetDriveListItemActivityDriveItemContentOperationResponse, err error) {
	opts := client.RequestOptions{
		ContentType: options.ContentType,
		ExpectedStatusCodes: []int{
			http.StatusNoContent,
		},
		HttpMethod:    http.MethodPut,
		OptionsObject: options,
		Path:          fmt.Sprintf("%s/driveItem/content", id.ID()),
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
