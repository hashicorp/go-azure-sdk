package endpointprivilegemanagementprovisioningstatus

import (
	"context"
	"fmt"
	"net/http"

	"github.com/hashicorp/go-azure-sdk/sdk/client"
	"github.com/hashicorp/go-azure-sdk/sdk/odata"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type DeleteEndpointPrivilegeManagementProvisioningStatusOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
}

type DeleteEndpointPrivilegeManagementProvisioningStatusOperationOptions struct {
	IfMatch *string
}

func DefaultDeleteEndpointPrivilegeManagementProvisioningStatusOperationOptions() DeleteEndpointPrivilegeManagementProvisioningStatusOperationOptions {
	return DeleteEndpointPrivilegeManagementProvisioningStatusOperationOptions{}
}

func (o DeleteEndpointPrivilegeManagementProvisioningStatusOperationOptions) ToHeaders() *client.Headers {
	out := client.Headers{}
	if o.IfMatch != nil {
		out.Append("If-Match", fmt.Sprintf("%v", *o.IfMatch))
	}
	return &out
}

func (o DeleteEndpointPrivilegeManagementProvisioningStatusOperationOptions) ToOData() *odata.Query {
	out := odata.Query{}

	return &out
}

func (o DeleteEndpointPrivilegeManagementProvisioningStatusOperationOptions) ToQuery() *client.QueryParams {
	out := client.QueryParams{}

	return &out
}

// DeleteEndpointPrivilegeManagementProvisioningStatus - Delete navigation property
// endpointPrivilegeManagementProvisioningStatus for deviceManagement
func (c EndpointPrivilegeManagementProvisioningStatusClient) DeleteEndpointPrivilegeManagementProvisioningStatus(ctx context.Context, options DeleteEndpointPrivilegeManagementProvisioningStatusOperationOptions) (result DeleteEndpointPrivilegeManagementProvisioningStatusOperationResponse, err error) {
	opts := client.RequestOptions{
		ContentType: "application/json; charset=utf-8",
		ExpectedStatusCodes: []int{
			http.StatusNoContent,
		},
		HttpMethod:    http.MethodDelete,
		OptionsObject: options,
		Path:          "/deviceManagement/endpointPrivilegeManagementProvisioningStatus",
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
