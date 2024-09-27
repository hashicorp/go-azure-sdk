package manageddevicelogcollectionrequest

import (
	"context"
	"net/http"

	"github.com/hashicorp/go-azure-sdk/microsoft-graph/common-types/beta"
	"github.com/hashicorp/go-azure-sdk/sdk/client"
	"github.com/hashicorp/go-azure-sdk/sdk/odata"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type UpdateManagedDeviceLogCollectionRequestOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
}

type UpdateManagedDeviceLogCollectionRequestOperationOptions struct {
	Metadata  *odata.Metadata
	RetryFunc client.RequestRetryFunc
}

func DefaultUpdateManagedDeviceLogCollectionRequestOperationOptions() UpdateManagedDeviceLogCollectionRequestOperationOptions {
	return UpdateManagedDeviceLogCollectionRequestOperationOptions{}
}

func (o UpdateManagedDeviceLogCollectionRequestOperationOptions) ToHeaders() *client.Headers {
	out := client.Headers{}

	return &out
}

func (o UpdateManagedDeviceLogCollectionRequestOperationOptions) ToOData() *odata.Query {
	out := odata.Query{}
	if o.Metadata != nil {
		out.Metadata = *o.Metadata
	}
	return &out
}

func (o UpdateManagedDeviceLogCollectionRequestOperationOptions) ToQuery() *client.QueryParams {
	out := client.QueryParams{}

	return &out
}

// UpdateManagedDeviceLogCollectionRequest - Update the navigation property logCollectionRequests in users
func (c ManagedDeviceLogCollectionRequestClient) UpdateManagedDeviceLogCollectionRequest(ctx context.Context, id beta.UserIdManagedDeviceIdLogCollectionRequestId, input beta.DeviceLogCollectionResponse, options UpdateManagedDeviceLogCollectionRequestOperationOptions) (result UpdateManagedDeviceLogCollectionRequestOperationResponse, err error) {
	opts := client.RequestOptions{
		ContentType: "application/json; charset=utf-8",
		ExpectedStatusCodes: []int{
			http.StatusAccepted,
			http.StatusNoContent,
			http.StatusOK,
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
