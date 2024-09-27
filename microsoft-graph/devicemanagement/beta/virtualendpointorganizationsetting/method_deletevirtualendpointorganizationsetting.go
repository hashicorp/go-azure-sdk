package virtualendpointorganizationsetting

import (
	"context"
	"fmt"
	"net/http"

	"github.com/hashicorp/go-azure-sdk/sdk/client"
	"github.com/hashicorp/go-azure-sdk/sdk/odata"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type DeleteVirtualEndpointOrganizationSettingOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
}

type DeleteVirtualEndpointOrganizationSettingOperationOptions struct {
	IfMatch   *string
	Metadata  *odata.Metadata
	RetryFunc client.RequestRetryFunc
}

func DefaultDeleteVirtualEndpointOrganizationSettingOperationOptions() DeleteVirtualEndpointOrganizationSettingOperationOptions {
	return DeleteVirtualEndpointOrganizationSettingOperationOptions{}
}

func (o DeleteVirtualEndpointOrganizationSettingOperationOptions) ToHeaders() *client.Headers {
	out := client.Headers{}
	if o.IfMatch != nil {
		out.Append("If-Match", fmt.Sprintf("%v", *o.IfMatch))
	}
	return &out
}

func (o DeleteVirtualEndpointOrganizationSettingOperationOptions) ToOData() *odata.Query {
	out := odata.Query{}
	if o.Metadata != nil {
		out.Metadata = *o.Metadata
	}
	return &out
}

func (o DeleteVirtualEndpointOrganizationSettingOperationOptions) ToQuery() *client.QueryParams {
	out := client.QueryParams{}

	return &out
}

// DeleteVirtualEndpointOrganizationSetting - Delete navigation property organizationSettings for deviceManagement
func (c VirtualEndpointOrganizationSettingClient) DeleteVirtualEndpointOrganizationSetting(ctx context.Context, options DeleteVirtualEndpointOrganizationSettingOperationOptions) (result DeleteVirtualEndpointOrganizationSettingOperationResponse, err error) {
	opts := client.RequestOptions{
		ContentType: "application/json; charset=utf-8",
		ExpectedStatusCodes: []int{
			http.StatusNoContent,
			http.StatusOK,
		},
		HttpMethod:    http.MethodDelete,
		OptionsObject: options,
		Path:          "/deviceManagement/virtualEndpoint/organizationSettings",
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
