package devicecompliancepolicydevicestatesummary

import (
	"context"
	"fmt"
	"net/http"

	"github.com/hashicorp/go-azure-sdk/sdk/client"
	"github.com/hashicorp/go-azure-sdk/sdk/odata"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type DeleteDeviceCompliancePolicyDeviceStateSummaryOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
}

type DeleteDeviceCompliancePolicyDeviceStateSummaryOperationOptions struct {
	IfMatch  *string
	Metadata *odata.Metadata
}

func DefaultDeleteDeviceCompliancePolicyDeviceStateSummaryOperationOptions() DeleteDeviceCompliancePolicyDeviceStateSummaryOperationOptions {
	return DeleteDeviceCompliancePolicyDeviceStateSummaryOperationOptions{}
}

func (o DeleteDeviceCompliancePolicyDeviceStateSummaryOperationOptions) ToHeaders() *client.Headers {
	out := client.Headers{}
	if o.IfMatch != nil {
		out.Append("If-Match", fmt.Sprintf("%v", *o.IfMatch))
	}
	return &out
}

func (o DeleteDeviceCompliancePolicyDeviceStateSummaryOperationOptions) ToOData() *odata.Query {
	out := odata.Query{}
	if o.Metadata != nil {
		out.Metadata = *o.Metadata
	}
	return &out
}

func (o DeleteDeviceCompliancePolicyDeviceStateSummaryOperationOptions) ToQuery() *client.QueryParams {
	out := client.QueryParams{}

	return &out
}

// DeleteDeviceCompliancePolicyDeviceStateSummary - Delete navigation property deviceCompliancePolicyDeviceStateSummary
// for deviceManagement
func (c DeviceCompliancePolicyDeviceStateSummaryClient) DeleteDeviceCompliancePolicyDeviceStateSummary(ctx context.Context, options DeleteDeviceCompliancePolicyDeviceStateSummaryOperationOptions) (result DeleteDeviceCompliancePolicyDeviceStateSummaryOperationResponse, err error) {
	opts := client.RequestOptions{
		ContentType: "application/json; charset=utf-8",
		ExpectedStatusCodes: []int{
			http.StatusNoContent,
		},
		HttpMethod:    http.MethodDelete,
		OptionsObject: options,
		Path:          "/deviceManagement/deviceCompliancePolicyDeviceStateSummary",
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
