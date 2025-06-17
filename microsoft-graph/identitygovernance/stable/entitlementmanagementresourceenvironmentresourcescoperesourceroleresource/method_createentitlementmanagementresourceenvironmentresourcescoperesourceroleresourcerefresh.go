package entitlementmanagementresourceenvironmentresourcescoperesourceroleresource

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

type CreateEntitlementManagementResourceEnvironmentResourceScopeResourceRoleResourceRefreshOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
}

type CreateEntitlementManagementResourceEnvironmentResourceScopeResourceRoleResourceRefreshOperationOptions struct {
	Metadata  *odata.Metadata
	RetryFunc client.RequestRetryFunc
}

func DefaultCreateEntitlementManagementResourceEnvironmentResourceScopeResourceRoleResourceRefreshOperationOptions() CreateEntitlementManagementResourceEnvironmentResourceScopeResourceRoleResourceRefreshOperationOptions {
	return CreateEntitlementManagementResourceEnvironmentResourceScopeResourceRoleResourceRefreshOperationOptions{}
}

func (o CreateEntitlementManagementResourceEnvironmentResourceScopeResourceRoleResourceRefreshOperationOptions) ToHeaders() *client.Headers {
	out := client.Headers{}

	return &out
}

func (o CreateEntitlementManagementResourceEnvironmentResourceScopeResourceRoleResourceRefreshOperationOptions) ToOData() *odata.Query {
	out := odata.Query{}
	if o.Metadata != nil {
		out.Metadata = *o.Metadata
	}
	return &out
}

func (o CreateEntitlementManagementResourceEnvironmentResourceScopeResourceRoleResourceRefreshOperationOptions) ToQuery() *client.QueryParams {
	out := client.QueryParams{}

	return &out
}

// CreateEntitlementManagementResourceEnvironmentResourceScopeResourceRoleResourceRefresh - Invoke action refresh
func (c EntitlementManagementResourceEnvironmentResourceScopeResourceRoleResourceClient) CreateEntitlementManagementResourceEnvironmentResourceScopeResourceRoleResourceRefresh(ctx context.Context, id stable.IdentityGovernanceEntitlementManagementResourceEnvironmentIdResourceIdScopeIdResourceRoleId, options CreateEntitlementManagementResourceEnvironmentResourceScopeResourceRoleResourceRefreshOperationOptions) (result CreateEntitlementManagementResourceEnvironmentResourceScopeResourceRoleResourceRefreshOperationResponse, err error) {
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
