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

type DeleteEntitlementManagementResourceEnvironmentResourceScopeResourceRoleResourceOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
}

type DeleteEntitlementManagementResourceEnvironmentResourceScopeResourceRoleResourceOperationOptions struct {
	IfMatch *string
}

func DefaultDeleteEntitlementManagementResourceEnvironmentResourceScopeResourceRoleResourceOperationOptions() DeleteEntitlementManagementResourceEnvironmentResourceScopeResourceRoleResourceOperationOptions {
	return DeleteEntitlementManagementResourceEnvironmentResourceScopeResourceRoleResourceOperationOptions{}
}

func (o DeleteEntitlementManagementResourceEnvironmentResourceScopeResourceRoleResourceOperationOptions) ToHeaders() *client.Headers {
	out := client.Headers{}
	if o.IfMatch != nil {
		out.Append("If-Match", fmt.Sprintf("%v", *o.IfMatch))
	}
	return &out
}

func (o DeleteEntitlementManagementResourceEnvironmentResourceScopeResourceRoleResourceOperationOptions) ToOData() *odata.Query {
	out := odata.Query{}

	return &out
}

func (o DeleteEntitlementManagementResourceEnvironmentResourceScopeResourceRoleResourceOperationOptions) ToQuery() *client.QueryParams {
	out := client.QueryParams{}

	return &out
}

// DeleteEntitlementManagementResourceEnvironmentResourceScopeResourceRoleResource - Delete navigation property resource
// for identityGovernance
func (c EntitlementManagementResourceEnvironmentResourceScopeResourceRoleResourceClient) DeleteEntitlementManagementResourceEnvironmentResourceScopeResourceRoleResource(ctx context.Context, id stable.IdentityGovernanceEntitlementManagementResourceEnvironmentIdResourceIdScopeIdResourceRoleId, options DeleteEntitlementManagementResourceEnvironmentResourceScopeResourceRoleResourceOperationOptions) (result DeleteEntitlementManagementResourceEnvironmentResourceScopeResourceRoleResourceOperationResponse, err error) {
	opts := client.RequestOptions{
		ContentType: "application/json; charset=utf-8",
		ExpectedStatusCodes: []int{
			http.StatusNoContent,
		},
		HttpMethod:    http.MethodDelete,
		OptionsObject: options,
		Path:          fmt.Sprintf("%s/resource", id.ID()),
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
