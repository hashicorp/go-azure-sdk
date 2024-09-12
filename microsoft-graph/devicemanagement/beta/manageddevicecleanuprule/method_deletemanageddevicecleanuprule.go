package manageddevicecleanuprule

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

type DeleteManagedDeviceCleanupRuleOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
}

type DeleteManagedDeviceCleanupRuleOperationOptions struct {
	IfMatch *string
}

func DefaultDeleteManagedDeviceCleanupRuleOperationOptions() DeleteManagedDeviceCleanupRuleOperationOptions {
	return DeleteManagedDeviceCleanupRuleOperationOptions{}
}

func (o DeleteManagedDeviceCleanupRuleOperationOptions) ToHeaders() *client.Headers {
	out := client.Headers{}
	if o.IfMatch != nil {
		out.Append("If-Match", fmt.Sprintf("%v", *o.IfMatch))
	}
	return &out
}

func (o DeleteManagedDeviceCleanupRuleOperationOptions) ToOData() *odata.Query {
	out := odata.Query{}

	return &out
}

func (o DeleteManagedDeviceCleanupRuleOperationOptions) ToQuery() *client.QueryParams {
	out := client.QueryParams{}

	return &out
}

// DeleteManagedDeviceCleanupRule - Delete navigation property managedDeviceCleanupRules for deviceManagement
func (c ManagedDeviceCleanupRuleClient) DeleteManagedDeviceCleanupRule(ctx context.Context, id beta.DeviceManagementManagedDeviceCleanupRuleId, options DeleteManagedDeviceCleanupRuleOperationOptions) (result DeleteManagedDeviceCleanupRuleOperationResponse, err error) {
	opts := client.RequestOptions{
		ContentType: "application/json; charset=utf-8",
		ExpectedStatusCodes: []int{
			http.StatusNoContent,
		},
		HttpMethod:    http.MethodDelete,
		OptionsObject: options,
		Path:          id.ID(),
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
