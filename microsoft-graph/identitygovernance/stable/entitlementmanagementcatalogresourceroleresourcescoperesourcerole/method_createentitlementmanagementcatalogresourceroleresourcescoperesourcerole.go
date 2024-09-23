package entitlementmanagementcatalogresourceroleresourcescoperesourcerole

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

type CreateEntitlementManagementCatalogResourceRoleResourceScopeResourceRoleOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
	Model        *stable.AccessPackageResourceRole
}

type CreateEntitlementManagementCatalogResourceRoleResourceScopeResourceRoleOperationOptions struct {
	Metadata  *odata.Metadata
	RetryFunc client.RequestRetryFunc
}

func DefaultCreateEntitlementManagementCatalogResourceRoleResourceScopeResourceRoleOperationOptions() CreateEntitlementManagementCatalogResourceRoleResourceScopeResourceRoleOperationOptions {
	return CreateEntitlementManagementCatalogResourceRoleResourceScopeResourceRoleOperationOptions{}
}

func (o CreateEntitlementManagementCatalogResourceRoleResourceScopeResourceRoleOperationOptions) ToHeaders() *client.Headers {
	out := client.Headers{}

	return &out
}

func (o CreateEntitlementManagementCatalogResourceRoleResourceScopeResourceRoleOperationOptions) ToOData() *odata.Query {
	out := odata.Query{}
	if o.Metadata != nil {
		out.Metadata = *o.Metadata
	}
	return &out
}

func (o CreateEntitlementManagementCatalogResourceRoleResourceScopeResourceRoleOperationOptions) ToQuery() *client.QueryParams {
	out := client.QueryParams{}

	return &out
}

// CreateEntitlementManagementCatalogResourceRoleResourceScopeResourceRole - Create new navigation property to roles for
// identityGovernance
func (c EntitlementManagementCatalogResourceRoleResourceScopeResourceRoleClient) CreateEntitlementManagementCatalogResourceRoleResourceScopeResourceRole(ctx context.Context, id stable.IdentityGovernanceEntitlementManagementCatalogIdResourceRoleIdResourceScopeId, input stable.AccessPackageResourceRole, options CreateEntitlementManagementCatalogResourceRoleResourceScopeResourceRoleOperationOptions) (result CreateEntitlementManagementCatalogResourceRoleResourceScopeResourceRoleOperationResponse, err error) {
	opts := client.RequestOptions{
		ContentType: "application/json; charset=utf-8",
		ExpectedStatusCodes: []int{
			http.StatusCreated,
		},
		HttpMethod:    http.MethodPost,
		OptionsObject: options,
		Path:          fmt.Sprintf("%s/resource/roles", id.ID()),
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

	var model stable.AccessPackageResourceRole
	result.Model = &model
	if err = resp.Unmarshal(result.Model); err != nil {
		return
	}

	return
}
