package entitlementmanagementresourcerolescopescoperesourcescope

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

type CreateEntitlementManagementResourceRoleScopeResourceScopeOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
	Model        *stable.AccessPackageResourceScope
}

type CreateEntitlementManagementResourceRoleScopeResourceScopeOperationOptions struct {
	Metadata *odata.Metadata
}

func DefaultCreateEntitlementManagementResourceRoleScopeResourceScopeOperationOptions() CreateEntitlementManagementResourceRoleScopeResourceScopeOperationOptions {
	return CreateEntitlementManagementResourceRoleScopeResourceScopeOperationOptions{}
}

func (o CreateEntitlementManagementResourceRoleScopeResourceScopeOperationOptions) ToHeaders() *client.Headers {
	out := client.Headers{}

	return &out
}

func (o CreateEntitlementManagementResourceRoleScopeResourceScopeOperationOptions) ToOData() *odata.Query {
	out := odata.Query{}
	if o.Metadata != nil {
		out.Metadata = *o.Metadata
	}
	return &out
}

func (o CreateEntitlementManagementResourceRoleScopeResourceScopeOperationOptions) ToQuery() *client.QueryParams {
	out := client.QueryParams{}

	return &out
}

// CreateEntitlementManagementResourceRoleScopeResourceScope - Create new navigation property to scopes for
// identityGovernance
func (c EntitlementManagementResourceRoleScopeScopeResourceScopeClient) CreateEntitlementManagementResourceRoleScopeResourceScope(ctx context.Context, id stable.IdentityGovernanceEntitlementManagementResourceRoleScopeId, input stable.AccessPackageResourceScope, options CreateEntitlementManagementResourceRoleScopeResourceScopeOperationOptions) (result CreateEntitlementManagementResourceRoleScopeResourceScopeOperationResponse, err error) {
	opts := client.RequestOptions{
		ContentType: "application/json; charset=utf-8",
		ExpectedStatusCodes: []int{
			http.StatusCreated,
		},
		HttpMethod:    http.MethodPost,
		OptionsObject: options,
		Path:          fmt.Sprintf("%s/scope/resource/scopes", id.ID()),
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
