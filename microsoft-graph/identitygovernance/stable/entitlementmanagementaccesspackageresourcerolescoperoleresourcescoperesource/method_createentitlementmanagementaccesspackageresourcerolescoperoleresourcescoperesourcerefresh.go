package entitlementmanagementaccesspackageresourcerolescoperoleresourcescoperesource

import (
	"context"
	"fmt"
	"net/http"

	"github.com/hashicorp/go-azure-sdk/microsoft-graph/common-types/stable"
	"github.com/hashicorp/go-azure-sdk/sdk/client"
	"github.com/hashicorp/go-azure-sdk/sdk/odata"
)

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type CreateEntitlementManagementAccessPackageResourceRoleScopeRoleResourceScopeResourceRefreshOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
}

type CreateEntitlementManagementAccessPackageResourceRoleScopeRoleResourceScopeResourceRefreshOperationOptions struct {
	Metadata  *odata.Metadata
	RetryFunc client.RequestRetryFunc
}

func DefaultCreateEntitlementManagementAccessPackageResourceRoleScopeRoleResourceScopeResourceRefreshOperationOptions() CreateEntitlementManagementAccessPackageResourceRoleScopeRoleResourceScopeResourceRefreshOperationOptions {
	return CreateEntitlementManagementAccessPackageResourceRoleScopeRoleResourceScopeResourceRefreshOperationOptions{}
}

func (o CreateEntitlementManagementAccessPackageResourceRoleScopeRoleResourceScopeResourceRefreshOperationOptions) ToHeaders() *client.Headers {
	out := client.Headers{}

	return &out
}

func (o CreateEntitlementManagementAccessPackageResourceRoleScopeRoleResourceScopeResourceRefreshOperationOptions) ToOData() *odata.Query {
	out := odata.Query{}
	if o.Metadata != nil {
		out.Metadata = *o.Metadata
	}
	return &out
}

func (o CreateEntitlementManagementAccessPackageResourceRoleScopeRoleResourceScopeResourceRefreshOperationOptions) ToQuery() *client.QueryParams {
	out := client.QueryParams{}

	return &out
}

// CreateEntitlementManagementAccessPackageResourceRoleScopeRoleResourceScopeResourceRefresh - Invoke action refresh
func (c EntitlementManagementAccessPackageResourceRoleScopeRoleResourceScopeResourceClient) CreateEntitlementManagementAccessPackageResourceRoleScopeRoleResourceScopeResourceRefresh(ctx context.Context, id stable.IdentityGovernanceEntitlementManagementAccessPackageIdResourceRoleScopeIdRoleResourceScopeId, options CreateEntitlementManagementAccessPackageResourceRoleScopeRoleResourceScopeResourceRefreshOperationOptions) (result CreateEntitlementManagementAccessPackageResourceRoleScopeRoleResourceScopeResourceRefreshOperationResponse, err error) {
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
		Path:          fmt.Sprintf("%s/resource/refresh", id.ID()),
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
