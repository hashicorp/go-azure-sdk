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

type CreateIosUpdateStatusOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
	Model        *stable.IosUpdateDeviceStatus
}

type CreateIosUpdateStatusOperationOptions struct {
	Metadata *odata.Metadata
}

func DefaultCreateIosUpdateStatusOperationOptions() CreateIosUpdateStatusOperationOptions {
	return CreateIosUpdateStatusOperationOptions{}
}

func (o CreateIosUpdateStatusOperationOptions) ToHeaders() *client.Headers {
	out := client.Headers{}

	return &out
}

func (o CreateIosUpdateStatusOperationOptions) ToOData() *odata.Query {
	out := odata.Query{}
	if o.Metadata != nil {
		out.Metadata = *o.Metadata
	}
	return &out
}

func (o CreateIosUpdateStatusOperationOptions) ToQuery() *client.QueryParams {
	out := client.QueryParams{}

	return &out
}

// CreateIosUpdateStatus - Create iosUpdateDeviceStatus. Create a new iosUpdateDeviceStatus object.
func (c IosUpdateStatusClient) CreateIosUpdateStatus(ctx context.Context, input stable.IosUpdateDeviceStatus, options CreateIosUpdateStatusOperationOptions) (result CreateIosUpdateStatusOperationResponse, err error) {
	opts := client.RequestOptions{
		ContentType: "application/json; charset=utf-8",
		ExpectedStatusCodes: []int{
			http.StatusCreated,
		},
		HttpMethod:    http.MethodPost,
		OptionsObject: options,
		Path:          "/deviceManagement/iosUpdateStatuses",
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

	var model stable.IosUpdateDeviceStatus
	result.Model = &model
	if err = resp.Unmarshal(result.Model); err != nil {
		return
	}

	return
}
