package cloudpc

import (
	"context"
	"fmt"
	"net/http"

	"github.com/hashicorp/go-azure-sdk/microsoft-graph/common-types/beta"
	"github.com/hashicorp/go-azure-sdk/sdk/client"
	"github.com/hashicorp/go-azure-sdk/sdk/odata"
)

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type PowerOffCloudPCOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
}

type PowerOffCloudPCOperationOptions struct {
	Metadata  *odata.Metadata
	RetryFunc client.RequestRetryFunc
}

func DefaultPowerOffCloudPCOperationOptions() PowerOffCloudPCOperationOptions {
	return PowerOffCloudPCOperationOptions{}
}

func (o PowerOffCloudPCOperationOptions) ToHeaders() *client.Headers {
	out := client.Headers{}

	return &out
}

func (o PowerOffCloudPCOperationOptions) ToOData() *odata.Query {
	out := odata.Query{}
	if o.Metadata != nil {
		out.Metadata = *o.Metadata
	}
	return &out
}

func (o PowerOffCloudPCOperationOptions) ToQuery() *client.QueryParams {
	out := client.QueryParams{}

	return &out
}

// PowerOffCloudPC - Invoke action powerOff. Power off a Windows 365 Frontline Cloud PC. This action supports Microsoft
// Endpoint Manager (MEM) admin scenarios. After a Windows 365 Frontline Cloud PC is powered off, it's deallocated, and
// licenses are revoked immediately. Only IT admin users can perform this action.
func (c CloudPCClient) PowerOffCloudPC(ctx context.Context, id beta.MeCloudPCId, options PowerOffCloudPCOperationOptions) (result PowerOffCloudPCOperationResponse, err error) {
	opts := client.RequestOptions{
		ContentType: "application/json; charset=utf-8",
		ExpectedStatusCodes: []int{
			http.StatusAccepted,
			http.StatusCreated,
			http.StatusNoContent,
			http.StatusOK,
		},
		HttpMethod:    http.MethodPost,
		OptionsObject: options,
		Path:          fmt.Sprintf("%s/powerOff", id.ID()),
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
