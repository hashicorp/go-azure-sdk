package manageddevice

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

type ChangeManagedDeviceAssignmentsOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
}

type ChangeManagedDeviceAssignmentsOperationOptions struct {
	Metadata  *odata.Metadata
	RetryFunc client.RequestRetryFunc
}

func DefaultChangeManagedDeviceAssignmentsOperationOptions() ChangeManagedDeviceAssignmentsOperationOptions {
	return ChangeManagedDeviceAssignmentsOperationOptions{}
}

func (o ChangeManagedDeviceAssignmentsOperationOptions) ToHeaders() *client.Headers {
	out := client.Headers{}

	return &out
}

func (o ChangeManagedDeviceAssignmentsOperationOptions) ToOData() *odata.Query {
	out := odata.Query{}
	if o.Metadata != nil {
		out.Metadata = *o.Metadata
	}
	return &out
}

func (o ChangeManagedDeviceAssignmentsOperationOptions) ToQuery() *client.QueryParams {
	out := client.QueryParams{}

	return &out
}

// ChangeManagedDeviceAssignments - Invoke action changeAssignments
func (c ManagedDeviceClient) ChangeManagedDeviceAssignments(ctx context.Context, id beta.UserIdManagedDeviceId, input ChangeManagedDeviceAssignmentsRequest, options ChangeManagedDeviceAssignmentsOperationOptions) (result ChangeManagedDeviceAssignmentsOperationResponse, err error) {
	opts := client.RequestOptions{
		ContentType: "application/json; charset=utf-8",
		ExpectedStatusCodes: []int{
			http.StatusNoContent,
		},
		HttpMethod:    http.MethodPost,
		OptionsObject: options,
		Path:          fmt.Sprintf("%s/changeAssignments", id.ID()),
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
