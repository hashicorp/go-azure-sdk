package iosupdatestatus

import (
	"context"
	"net/http"

	"github.com/hashicorp/go-azure-sdk/microsoft-graph/common-types/stable"
	"github.com/hashicorp/go-azure-sdk/sdk/client"
	"github.com/hashicorp/go-azure-sdk/sdk/odata"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type UpdateIosUpdateStatusOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
}

type UpdateIosUpdateStatusOperationOptions struct {
	Metadata  *odata.Metadata
	RetryFunc client.RequestRetryFunc
}

func DefaultUpdateIosUpdateStatusOperationOptions() UpdateIosUpdateStatusOperationOptions {
	return UpdateIosUpdateStatusOperationOptions{}
}

func (o UpdateIosUpdateStatusOperationOptions) ToHeaders() *client.Headers {
	out := client.Headers{}

	return &out
}

func (o UpdateIosUpdateStatusOperationOptions) ToOData() *odata.Query {
	out := odata.Query{}
	if o.Metadata != nil {
		out.Metadata = *o.Metadata
	}
	return &out
}

func (o UpdateIosUpdateStatusOperationOptions) ToQuery() *client.QueryParams {
	out := client.QueryParams{}

	return &out
}

// UpdateIosUpdateStatus - Update iosUpdateDeviceStatus. Update the properties of a iosUpdateDeviceStatus object.
func (c IosUpdateStatusClient) UpdateIosUpdateStatus(ctx context.Context, id stable.DeviceManagementIosUpdateStatusId, input stable.IosUpdateDeviceStatus, options UpdateIosUpdateStatusOperationOptions) (result UpdateIosUpdateStatusOperationResponse, err error) {
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
