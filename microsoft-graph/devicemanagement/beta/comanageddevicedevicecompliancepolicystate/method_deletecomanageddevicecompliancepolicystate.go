package comanageddevicedevicecompliancepolicystate

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

type DeleteComanagedDeviceCompliancePolicyStateOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
}

type DeleteComanagedDeviceCompliancePolicyStateOperationOptions struct {
	IfMatch  *string
	Metadata *odata.Metadata
}

func DefaultDeleteComanagedDeviceCompliancePolicyStateOperationOptions() DeleteComanagedDeviceCompliancePolicyStateOperationOptions {
	return DeleteComanagedDeviceCompliancePolicyStateOperationOptions{}
}

func (o DeleteComanagedDeviceCompliancePolicyStateOperationOptions) ToHeaders() *client.Headers {
	out := client.Headers{}
	if o.IfMatch != nil {
		out.Append("If-Match", fmt.Sprintf("%v", *o.IfMatch))
	}
	return &out
}

func (o DeleteComanagedDeviceCompliancePolicyStateOperationOptions) ToOData() *odata.Query {
	out := odata.Query{}
	if o.Metadata != nil {
		out.Metadata = *o.Metadata
	}
	return &out
}

func (o DeleteComanagedDeviceCompliancePolicyStateOperationOptions) ToQuery() *client.QueryParams {
	out := client.QueryParams{}

	return &out
}

// DeleteComanagedDeviceCompliancePolicyState - Delete navigation property deviceCompliancePolicyStates for
// deviceManagement
func (c ComanagedDeviceDeviceCompliancePolicyStateClient) DeleteComanagedDeviceCompliancePolicyState(ctx context.Context, id beta.DeviceManagementComanagedDeviceIdDeviceCompliancePolicyStateId, options DeleteComanagedDeviceCompliancePolicyStateOperationOptions) (result DeleteComanagedDeviceCompliancePolicyStateOperationResponse, err error) {
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
