package devicemanagement

import (
	"context"
	"net/http"

	"github.com/hashicorp/go-azure-sdk/sdk/client"
	"github.com/hashicorp/go-azure-sdk/sdk/odata"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type EnableUnlicensedAdminstratorsOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
}

type EnableUnlicensedAdminstratorsOperationOptions struct {
	Metadata *odata.Metadata
}

func DefaultEnableUnlicensedAdminstratorsOperationOptions() EnableUnlicensedAdminstratorsOperationOptions {
	return EnableUnlicensedAdminstratorsOperationOptions{}
}

func (o EnableUnlicensedAdminstratorsOperationOptions) ToHeaders() *client.Headers {
	out := client.Headers{}

	return &out
}

func (o EnableUnlicensedAdminstratorsOperationOptions) ToOData() *odata.Query {
	out := odata.Query{}
	if o.Metadata != nil {
		out.Metadata = *o.Metadata
	}
	return &out
}

func (o EnableUnlicensedAdminstratorsOperationOptions) ToQuery() *client.QueryParams {
	out := client.QueryParams{}

	return &out
}

// EnableUnlicensedAdminstrators - Invoke action enableUnlicensedAdminstrators. Upon enabling, users assigned as
// administrators via Role Assignment Memberships will no longer require an assigned Intune license. You are limited to
// 350 unlicensed direct members for each AAD security group in a role assignment, but you can assign multiple AAD
// security groups to a role if you need to support more than 350 unlicensed administrators. Licensed administrators
// will continue to function as-is in that transitive memberships apply and are not subject to the 350 member limit.
func (c DeviceManagementClient) EnableUnlicensedAdminstrators(ctx context.Context, options EnableUnlicensedAdminstratorsOperationOptions) (result EnableUnlicensedAdminstratorsOperationResponse, err error) {
	opts := client.RequestOptions{
		ContentType: "application/json; charset=utf-8",
		ExpectedStatusCodes: []int{
			http.StatusNoContent,
		},
		HttpMethod:    http.MethodPost,
		OptionsObject: options,
		Path:          "/deviceManagement/enableUnlicensedAdminstrators",
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
