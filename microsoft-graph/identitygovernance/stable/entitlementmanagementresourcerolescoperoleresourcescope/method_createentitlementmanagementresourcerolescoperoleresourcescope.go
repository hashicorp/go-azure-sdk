package entitlementmanagementresourcerolescoperoleresourcescope

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

type CreateEntitlementManagementResourceRoleScopeRoleResourceScopeOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
	Model        *stable.AccessPackageResourceScope
}

type CreateEntitlementManagementResourceRoleScopeRoleResourceScopeOperationOptions struct {
	Metadata  *odata.Metadata
	RetryFunc client.RequestRetryFunc
}

func DefaultCreateEntitlementManagementResourceRoleScopeRoleResourceScopeOperationOptions() CreateEntitlementManagementResourceRoleScopeRoleResourceScopeOperationOptions {
	return CreateEntitlementManagementResourceRoleScopeRoleResourceScopeOperationOptions{}
}

func (o CreateEntitlementManagementResourceRoleScopeRoleResourceScopeOperationOptions) ToHeaders() *client.Headers {
	out := client.Headers{}

	return &out
}

func (o CreateEntitlementManagementResourceRoleScopeRoleResourceScopeOperationOptions) ToOData() *odata.Query {
	out := odata.Query{}
	if o.Metadata != nil {
		out.Metadata = *o.Metadata
	}
	return &out
}

func (o CreateEntitlementManagementResourceRoleScopeRoleResourceScopeOperationOptions) ToQuery() *client.QueryParams {
	out := client.QueryParams{}

	return &out
}

// CreateEntitlementManagementResourceRoleScopeRoleResourceScope - Create new navigation property to scopes for
// identityGovernance
func (c EntitlementManagementResourceRoleScopeRoleResourceScopeClient) CreateEntitlementManagementResourceRoleScopeRoleResourceScope(ctx context.Context, id stable.IdentityGovernanceEntitlementManagementResourceRoleScopeId, input stable.AccessPackageResourceScope, options CreateEntitlementManagementResourceRoleScopeRoleResourceScopeOperationOptions) (result CreateEntitlementManagementResourceRoleScopeRoleResourceScopeOperationResponse, err error) {
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
		Path:          fmt.Sprintf("%s/role/resource/scopes", id.ID()),
		RetryFunc:     options.RetryFunc,
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

	var model stable.AccessPackageResourceScope
	result.Model = &model
	if err = resp.Unmarshal(result.Model); err != nil {
		return
	}

	return
}
