package usermanageddeviceassignmentfilterevaluationstatusdetail

import (
	"context"
	"fmt"
	"net/http"

	"github.com/hashicorp/go-azure-sdk/sdk/client"
	"github.com/hashicorp/go-azure-sdk/sdk/odata"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type GetUserByIdManagedDeviceByIdAssignmentFilterEvaluationStatusDetailCountOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
}

// GetUserByIdManagedDeviceByIdAssignmentFilterEvaluationStatusDetailCount ...
func (c UserManagedDeviceAssignmentFilterEvaluationStatusDetailClient) GetUserByIdManagedDeviceByIdAssignmentFilterEvaluationStatusDetailCount(ctx context.Context, id UserManagedDeviceId) (result GetUserByIdManagedDeviceByIdAssignmentFilterEvaluationStatusDetailCountOperationResponse, err error) {
	opts := client.RequestOptions{
		ContentType: "text/plain",
		ExpectedStatusCodes: []int{
			http.StatusOK,
		},
		HttpMethod: http.MethodGet,
		Path:       fmt.Sprintf("%s/assignmentFilterEvaluationStatusDetails/$count", id.ID()),
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
