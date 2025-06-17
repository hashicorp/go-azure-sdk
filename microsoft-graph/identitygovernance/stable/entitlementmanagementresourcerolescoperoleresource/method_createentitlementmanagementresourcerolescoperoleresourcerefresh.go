package entitlementmanagementresourcerolescoperoleresource

import (
	"context"
	"fmt"
	"net/http"

	"github.com/hashicorp/go-azure-sdk/microsoft-graph/common-types/stable"
	"github.com/hashicorp/go-azure-sdk/sdk/client"
	"github.com/hashicorp/go-azure-sdk/sdk/odata"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type CreateEntitlementManagementResourceRoleScopeRoleResourceRefreshOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
}

type CreateEntitlementManagementResourceRoleScopeRoleResourceRefreshOperationOptions struct {
	Metadata  *odata.Metadata
	RetryFunc client.RequestRetryFunc
}

func DefaultCreateEntitlementManagementResourceRoleScopeRoleResourceRefreshOperationOptions() CreateEntitlementManagementResourceRoleScopeRoleResourceRefreshOperationOptions {
	return CreateEntitlementManagementResourceRoleScopeRoleResourceRefreshOperationOptions{}
}

func (o CreateEntitlementManagementResourceRoleScopeRoleResourceRefreshOperationOptions) ToHeaders() *client.Headers {
	out := client.Headers{}

	return &out
}

func (o CreateEntitlementManagementResourceRoleScopeRoleResourceRefreshOperationOptions) ToOData() *odata.Query {
	out := odata.Query{}
	if o.Metadata != nil {
		out.Metadata = *o.Metadata
	}
	return &out
}

func (o CreateEntitlementManagementResourceRoleScopeRoleResourceRefreshOperationOptions) ToQuery() *client.QueryParams {
	out := client.QueryParams{}

	return &out
}

// CreateEntitlementManagementResourceRoleScopeRoleResourceRefresh - Invoke action refresh
func (c EntitlementManagementResourceRoleScopeRoleResourceClient) CreateEntitlementManagementResourceRoleScopeRoleResourceRefresh(ctx context.Context, id stable.IdentityGovernanceEntitlementManagementResourceRoleScopeId, options CreateEntitlementManagementResourceRoleScopeRoleResourceRefreshOperationOptions) (result CreateEntitlementManagementResourceRoleScopeRoleResourceRefreshOperationResponse, err error) {
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
		Path:          fmt.Sprintf("%s/role/resource/refresh", id.ID()),
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
