package driverootlistitemactivitydriveitemcontentstream

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

type SetDriveRootListItemActivityDriveItemContentStreamOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
}

type SetDriveRootListItemActivityDriveItemContentStreamOperationOptions struct {
	Metadata *odata.Metadata
}

func DefaultSetDriveRootListItemActivityDriveItemContentStreamOperationOptions() SetDriveRootListItemActivityDriveItemContentStreamOperationOptions {
	return SetDriveRootListItemActivityDriveItemContentStreamOperationOptions{}
}

func (o SetDriveRootListItemActivityDriveItemContentStreamOperationOptions) ToHeaders() *client.Headers {
	out := client.Headers{}

	return &out
}

func (o SetDriveRootListItemActivityDriveItemContentStreamOperationOptions) ToOData() *odata.Query {
	out := odata.Query{}
	if o.Metadata != nil {
		out.Metadata = *o.Metadata
	}
	return &out
}

func (o SetDriveRootListItemActivityDriveItemContentStreamOperationOptions) ToQuery() *client.QueryParams {
	out := client.QueryParams{}

	return &out
}

// SetDriveRootListItemActivityDriveItemContentStream - Update contentStream for the navigation property driveItem in
// groups. The content stream, if the item represents a file.
func (c DriveRootListItemActivityDriveItemContentStreamClient) SetDriveRootListItemActivityDriveItemContentStream(ctx context.Context, id beta.GroupIdDriveIdRootListItemActivityId, input []byte, options SetDriveRootListItemActivityDriveItemContentStreamOperationOptions) (result SetDriveRootListItemActivityDriveItemContentStreamOperationResponse, err error) {
	opts := client.RequestOptions{
		ContentType: "application/json; charset=utf-8",
		ExpectedStatusCodes: []int{
			http.StatusNoContent,
		},
		HttpMethod:    http.MethodPut,
		OptionsObject: options,
		Path:          fmt.Sprintf("%s/driveItem/contentStream", id.ID()),
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
