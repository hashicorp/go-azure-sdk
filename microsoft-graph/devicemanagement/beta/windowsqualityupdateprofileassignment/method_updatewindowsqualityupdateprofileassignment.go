package windowsqualityupdateprofileassignment

import (
	"context"
	"net/http"

	"github.com/hashicorp/go-azure-sdk/microsoft-graph/common-types/beta"
	"github.com/hashicorp/go-azure-sdk/sdk/client"
	"github.com/hashicorp/go-azure-sdk/sdk/odata"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type UpdateWindowsQualityUpdateProfileAssignmentOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
}

type UpdateWindowsQualityUpdateProfileAssignmentOperationOptions struct {
	Metadata  *odata.Metadata
	RetryFunc client.RequestRetryFunc
}

func DefaultUpdateWindowsQualityUpdateProfileAssignmentOperationOptions() UpdateWindowsQualityUpdateProfileAssignmentOperationOptions {
	return UpdateWindowsQualityUpdateProfileAssignmentOperationOptions{}
}

func (o UpdateWindowsQualityUpdateProfileAssignmentOperationOptions) ToHeaders() *client.Headers {
	out := client.Headers{}

	return &out
}

func (o UpdateWindowsQualityUpdateProfileAssignmentOperationOptions) ToOData() *odata.Query {
	out := odata.Query{}
	if o.Metadata != nil {
		out.Metadata = *o.Metadata
	}
	return &out
}

func (o UpdateWindowsQualityUpdateProfileAssignmentOperationOptions) ToQuery() *client.QueryParams {
	out := client.QueryParams{}

	return &out
}

// UpdateWindowsQualityUpdateProfileAssignment - Update the navigation property assignments in deviceManagement
func (c WindowsQualityUpdateProfileAssignmentClient) UpdateWindowsQualityUpdateProfileAssignment(ctx context.Context, id beta.DeviceManagementWindowsQualityUpdateProfileIdAssignmentId, input beta.WindowsQualityUpdateProfileAssignment, options UpdateWindowsQualityUpdateProfileAssignmentOperationOptions) (result UpdateWindowsQualityUpdateProfileAssignmentOperationResponse, err error) {
	opts := client.RequestOptions{
		ContentType: "application/json; charset=utf-8",
		ExpectedStatusCodes: []int{
			http.StatusNoContent,
		},
		HttpMethod:    http.MethodPatch,
		OptionsObject: options,
		Path:          id.ID(),
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
