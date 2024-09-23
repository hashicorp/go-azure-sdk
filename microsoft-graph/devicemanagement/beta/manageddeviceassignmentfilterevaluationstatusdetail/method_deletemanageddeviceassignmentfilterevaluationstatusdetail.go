package manageddeviceassignmentfilterevaluationstatusdetail

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

type DeleteManagedDeviceAssignmentFilterEvaluationStatusDetailOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
}

type DeleteManagedDeviceAssignmentFilterEvaluationStatusDetailOperationOptions struct {
	IfMatch   *string
	Metadata  *odata.Metadata
	RetryFunc client.RequestRetryFunc
}

func DefaultDeleteManagedDeviceAssignmentFilterEvaluationStatusDetailOperationOptions() DeleteManagedDeviceAssignmentFilterEvaluationStatusDetailOperationOptions {
	return DeleteManagedDeviceAssignmentFilterEvaluationStatusDetailOperationOptions{}
}

func (o DeleteManagedDeviceAssignmentFilterEvaluationStatusDetailOperationOptions) ToHeaders() *client.Headers {
	out := client.Headers{}
	if o.IfMatch != nil {
		out.Append("If-Match", fmt.Sprintf("%v", *o.IfMatch))
	}
	return &out
}

func (o DeleteManagedDeviceAssignmentFilterEvaluationStatusDetailOperationOptions) ToOData() *odata.Query {
	out := odata.Query{}
	if o.Metadata != nil {
		out.Metadata = *o.Metadata
	}
	return &out
}

func (o DeleteManagedDeviceAssignmentFilterEvaluationStatusDetailOperationOptions) ToQuery() *client.QueryParams {
	out := client.QueryParams{}

	return &out
}

// DeleteManagedDeviceAssignmentFilterEvaluationStatusDetail - Delete navigation property
// assignmentFilterEvaluationStatusDetails for deviceManagement
func (c ManagedDeviceAssignmentFilterEvaluationStatusDetailClient) DeleteManagedDeviceAssignmentFilterEvaluationStatusDetail(ctx context.Context, id beta.DeviceManagementManagedDeviceIdAssignmentFilterEvaluationStatusDetailId, options DeleteManagedDeviceAssignmentFilterEvaluationStatusDetailOperationOptions) (result DeleteManagedDeviceAssignmentFilterEvaluationStatusDetailOperationResponse, err error) {
	opts := client.RequestOptions{
		ContentType: "application/json; charset=utf-8",
		ExpectedStatusCodes: []int{
			http.StatusNoContent,
		},
		HttpMethod:    http.MethodDelete,
		OptionsObject: options,
		Path:          id.ID(),
		RetryFunc:     options.RetryFunc,
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
