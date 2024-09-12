package devicecustomattributeshellscriptgroupassignment

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

type GetDeviceCustomAttributeShellScriptGroupAssignmentsCountOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
	Model        *[]byte
}

type GetDeviceCustomAttributeShellScriptGroupAssignmentsCountOperationOptions struct {
	Filter *string
	Search *string
}

func DefaultGetDeviceCustomAttributeShellScriptGroupAssignmentsCountOperationOptions() GetDeviceCustomAttributeShellScriptGroupAssignmentsCountOperationOptions {
	return GetDeviceCustomAttributeShellScriptGroupAssignmentsCountOperationOptions{}
}

func (o GetDeviceCustomAttributeShellScriptGroupAssignmentsCountOperationOptions) ToHeaders() *client.Headers {
	out := client.Headers{}

	return &out
}

func (o GetDeviceCustomAttributeShellScriptGroupAssignmentsCountOperationOptions) ToOData() *odata.Query {
	out := odata.Query{}
	if o.Filter != nil {
		out.Filter = *o.Filter
	}
	if o.Search != nil {
		out.Search = *o.Search
	}
	return &out
}

func (o GetDeviceCustomAttributeShellScriptGroupAssignmentsCountOperationOptions) ToQuery() *client.QueryParams {
	out := client.QueryParams{}

	return &out
}

// GetDeviceCustomAttributeShellScriptGroupAssignmentsCount - Get the number of the resource
func (c DeviceCustomAttributeShellScriptGroupAssignmentClient) GetDeviceCustomAttributeShellScriptGroupAssignmentsCount(ctx context.Context, id beta.DeviceManagementDeviceCustomAttributeShellScriptId, options GetDeviceCustomAttributeShellScriptGroupAssignmentsCountOperationOptions) (result GetDeviceCustomAttributeShellScriptGroupAssignmentsCountOperationResponse, err error) {
	opts := client.RequestOptions{
		ContentType: "text/plain",
		ExpectedStatusCodes: []int{
			http.StatusOK,
		},
		HttpMethod:    http.MethodGet,
		OptionsObject: options,
		Path:          fmt.Sprintf("%s/groupAssignments/$count", id.ID()),
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

	var model []byte
	result.Model = &model
	if err = resp.Unmarshal(result.Model); err != nil {
		return
	}

	return
}
