package devicehealthscriptrunsummary

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

type DeleteDeviceHealthScriptRunSummaryOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
}

type DeleteDeviceHealthScriptRunSummaryOperationOptions struct {
	IfMatch *string
}

func DefaultDeleteDeviceHealthScriptRunSummaryOperationOptions() DeleteDeviceHealthScriptRunSummaryOperationOptions {
	return DeleteDeviceHealthScriptRunSummaryOperationOptions{}
}

func (o DeleteDeviceHealthScriptRunSummaryOperationOptions) ToHeaders() *client.Headers {
	out := client.Headers{}
	if o.IfMatch != nil {
		out.Append("If-Match", fmt.Sprintf("%v", *o.IfMatch))
	}
	return &out
}

func (o DeleteDeviceHealthScriptRunSummaryOperationOptions) ToOData() *odata.Query {
	out := odata.Query{}

	return &out
}

func (o DeleteDeviceHealthScriptRunSummaryOperationOptions) ToQuery() *client.QueryParams {
	out := client.QueryParams{}

	return &out
}

// DeleteDeviceHealthScriptRunSummary - Delete navigation property runSummary for deviceManagement
func (c DeviceHealthScriptRunSummaryClient) DeleteDeviceHealthScriptRunSummary(ctx context.Context, id beta.DeviceManagementDeviceHealthScriptId, options DeleteDeviceHealthScriptRunSummaryOperationOptions) (result DeleteDeviceHealthScriptRunSummaryOperationResponse, err error) {
	opts := client.RequestOptions{
		ContentType: "application/json; charset=utf-8",
		ExpectedStatusCodes: []int{
			http.StatusNoContent,
		},
		HttpMethod:    http.MethodDelete,
		OptionsObject: options,
		Path:          fmt.Sprintf("%s/runSummary", id.ID()),
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
