package manageddevicedevicecategory

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

type SetManagedDeviceCategoryRefOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
}

type SetManagedDeviceCategoryRefOperationOptions struct {
	Metadata  *odata.Metadata
	RetryFunc client.RequestRetryFunc
}

func DefaultSetManagedDeviceCategoryRefOperationOptions() SetManagedDeviceCategoryRefOperationOptions {
	return SetManagedDeviceCategoryRefOperationOptions{}
}

func (o SetManagedDeviceCategoryRefOperationOptions) ToHeaders() *client.Headers {
	out := client.Headers{}

	return &out
}

func (o SetManagedDeviceCategoryRefOperationOptions) ToOData() *odata.Query {
	out := odata.Query{}
	if o.Metadata != nil {
		out.Metadata = *o.Metadata
	}
	return &out
}

func (o SetManagedDeviceCategoryRefOperationOptions) ToQuery() *client.QueryParams {
	out := client.QueryParams{}

	return &out
}

// SetManagedDeviceCategoryRef - Update the ref of navigation property deviceCategory in me
func (c ManagedDeviceDeviceCategoryClient) SetManagedDeviceCategoryRef(ctx context.Context, id beta.MeManagedDeviceId, input beta.ReferenceUpdate, options SetManagedDeviceCategoryRefOperationOptions) (result SetManagedDeviceCategoryRefOperationResponse, err error) {
	opts := client.RequestOptions{
		ContentType: "application/json; charset=utf-8",
		ExpectedStatusCodes: []int{
			http.StatusAccepted,
			http.StatusCreated,
			http.StatusNoContent,
			http.StatusOK,
		},
		HttpMethod:    http.MethodPut,
		OptionsObject: options,
		Path:          fmt.Sprintf("%s/deviceCategory/$ref", id.ID()),
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
