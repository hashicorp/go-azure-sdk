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

type RevokeManagedDeviceAppleVppLicensesOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
}

type RevokeManagedDeviceAppleVppLicensesOperationOptions struct {
	Metadata  *odata.Metadata
	RetryFunc client.RequestRetryFunc
}

func DefaultRevokeManagedDeviceAppleVppLicensesOperationOptions() RevokeManagedDeviceAppleVppLicensesOperationOptions {
	return RevokeManagedDeviceAppleVppLicensesOperationOptions{}
}

func (o RevokeManagedDeviceAppleVppLicensesOperationOptions) ToHeaders() *client.Headers {
	out := client.Headers{}

	return &out
}

func (o RevokeManagedDeviceAppleVppLicensesOperationOptions) ToOData() *odata.Query {
	out := odata.Query{}
	if o.Metadata != nil {
		out.Metadata = *o.Metadata
	}
	return &out
}

func (o RevokeManagedDeviceAppleVppLicensesOperationOptions) ToQuery() *client.QueryParams {
	out := client.QueryParams{}

	return &out
}

// RevokeManagedDeviceAppleVppLicenses - Invoke action revokeAppleVppLicenses. Revoke all Apple Vpp licenses for a
// device
func (c ManagedDeviceClient) RevokeManagedDeviceAppleVppLicenses(ctx context.Context, id beta.UserIdManagedDeviceId, options RevokeManagedDeviceAppleVppLicensesOperationOptions) (result RevokeManagedDeviceAppleVppLicensesOperationResponse, err error) {
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
		Path:          fmt.Sprintf("%s/revokeAppleVppLicenses", id.ID()),
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
