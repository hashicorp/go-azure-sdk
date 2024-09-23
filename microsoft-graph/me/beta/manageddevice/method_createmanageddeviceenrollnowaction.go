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

type CreateManagedDeviceEnrollNowActionOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
}

type CreateManagedDeviceEnrollNowActionOperationOptions struct {
	Metadata  *odata.Metadata
	RetryFunc client.RequestRetryFunc
}

func DefaultCreateManagedDeviceEnrollNowActionOperationOptions() CreateManagedDeviceEnrollNowActionOperationOptions {
	return CreateManagedDeviceEnrollNowActionOperationOptions{}
}

func (o CreateManagedDeviceEnrollNowActionOperationOptions) ToHeaders() *client.Headers {
	out := client.Headers{}

	return &out
}

func (o CreateManagedDeviceEnrollNowActionOperationOptions) ToOData() *odata.Query {
	out := odata.Query{}
	if o.Metadata != nil {
		out.Metadata = *o.Metadata
	}
	return &out
}

func (o CreateManagedDeviceEnrollNowActionOperationOptions) ToQuery() *client.QueryParams {
	out := client.QueryParams{}

	return &out
}

// CreateManagedDeviceEnrollNowAction - Invoke action enrollNowAction. Trigger comanagement enrollment action on
// ConfigurationManager client
func (c ManagedDeviceClient) CreateManagedDeviceEnrollNowAction(ctx context.Context, id beta.MeManagedDeviceId, options CreateManagedDeviceEnrollNowActionOperationOptions) (result CreateManagedDeviceEnrollNowActionOperationResponse, err error) {
	opts := client.RequestOptions{
		ContentType: "application/json; charset=utf-8",
		ExpectedStatusCodes: []int{
			http.StatusNoContent,
		},
		HttpMethod:    http.MethodPost,
		OptionsObject: options,
		Path:          fmt.Sprintf("%s/enrollNowAction", id.ID()),
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
