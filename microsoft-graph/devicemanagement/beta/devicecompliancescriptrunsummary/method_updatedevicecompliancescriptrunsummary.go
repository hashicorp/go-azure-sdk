package devicecompliancescriptrunsummary

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

type UpdateDeviceComplianceScriptRunSummaryOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
}

type UpdateDeviceComplianceScriptRunSummaryOperationOptions struct {
	Metadata *odata.Metadata
}

func DefaultUpdateDeviceComplianceScriptRunSummaryOperationOptions() UpdateDeviceComplianceScriptRunSummaryOperationOptions {
	return UpdateDeviceComplianceScriptRunSummaryOperationOptions{}
}

func (o UpdateDeviceComplianceScriptRunSummaryOperationOptions) ToHeaders() *client.Headers {
	out := client.Headers{}

	return &out
}

func (o UpdateDeviceComplianceScriptRunSummaryOperationOptions) ToOData() *odata.Query {
	out := odata.Query{}
	if o.Metadata != nil {
		out.Metadata = *o.Metadata
	}
	return &out
}

func (o UpdateDeviceComplianceScriptRunSummaryOperationOptions) ToQuery() *client.QueryParams {
	out := client.QueryParams{}

	return &out
}

// UpdateDeviceComplianceScriptRunSummary - Update the navigation property runSummary in deviceManagement
func (c DeviceComplianceScriptRunSummaryClient) UpdateDeviceComplianceScriptRunSummary(ctx context.Context, id beta.DeviceManagementDeviceComplianceScriptId, input beta.DeviceComplianceScriptRunSummary, options UpdateDeviceComplianceScriptRunSummaryOperationOptions) (result UpdateDeviceComplianceScriptRunSummaryOperationResponse, err error) {
	opts := client.RequestOptions{
		ContentType: "application/json; charset=utf-8",
		ExpectedStatusCodes: []int{
			http.StatusNoContent,
		},
		HttpMethod:    http.MethodPatch,
		OptionsObject: options,
		Path:          fmt.Sprintf("%s/runSummary", id.ID()),
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
