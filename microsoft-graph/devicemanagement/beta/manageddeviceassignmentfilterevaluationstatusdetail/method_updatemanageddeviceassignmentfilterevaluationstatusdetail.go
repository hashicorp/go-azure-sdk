package manageddeviceassignmentfilterevaluationstatusdetail

import (
	"context"
	"net/http"

	"github.com/hashicorp/go-azure-sdk/microsoft-graph/common-types/beta"
	"github.com/hashicorp/go-azure-sdk/sdk/client"
	"github.com/hashicorp/go-azure-sdk/sdk/odata"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type UpdateManagedDeviceAssignmentFilterEvaluationStatusDetailOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
}

type UpdateManagedDeviceAssignmentFilterEvaluationStatusDetailOperationOptions struct {
	Metadata  *odata.Metadata
	RetryFunc client.RequestRetryFunc
}

func DefaultUpdateManagedDeviceAssignmentFilterEvaluationStatusDetailOperationOptions() UpdateManagedDeviceAssignmentFilterEvaluationStatusDetailOperationOptions {
	return UpdateManagedDeviceAssignmentFilterEvaluationStatusDetailOperationOptions{}
}

func (o UpdateManagedDeviceAssignmentFilterEvaluationStatusDetailOperationOptions) ToHeaders() *client.Headers {
	out := client.Headers{}

	return &out
}

func (o UpdateManagedDeviceAssignmentFilterEvaluationStatusDetailOperationOptions) ToOData() *odata.Query {
	out := odata.Query{}
	if o.Metadata != nil {
		out.Metadata = *o.Metadata
	}
	return &out
}

func (o UpdateManagedDeviceAssignmentFilterEvaluationStatusDetailOperationOptions) ToQuery() *client.QueryParams {
	out := client.QueryParams{}

	return &out
}

// UpdateManagedDeviceAssignmentFilterEvaluationStatusDetail - Update the navigation property
// assignmentFilterEvaluationStatusDetails in deviceManagement
func (c ManagedDeviceAssignmentFilterEvaluationStatusDetailClient) UpdateManagedDeviceAssignmentFilterEvaluationStatusDetail(ctx context.Context, id beta.DeviceManagementManagedDeviceIdAssignmentFilterEvaluationStatusDetailId, input beta.AssignmentFilterEvaluationStatusDetails, options UpdateManagedDeviceAssignmentFilterEvaluationStatusDetailOperationOptions) (result UpdateManagedDeviceAssignmentFilterEvaluationStatusDetailOperationResponse, err error) {
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
