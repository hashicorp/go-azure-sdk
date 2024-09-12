package entitlementmanagementcatalogresourceroleresource

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

type DeleteEntitlementManagementCatalogResourceRoleResourceOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
}

type DeleteEntitlementManagementCatalogResourceRoleResourceOperationOptions struct {
	IfMatch *string
}

func DefaultDeleteEntitlementManagementCatalogResourceRoleResourceOperationOptions() DeleteEntitlementManagementCatalogResourceRoleResourceOperationOptions {
	return DeleteEntitlementManagementCatalogResourceRoleResourceOperationOptions{}
}

func (o DeleteEntitlementManagementCatalogResourceRoleResourceOperationOptions) ToHeaders() *client.Headers {
	out := client.Headers{}
	if o.IfMatch != nil {
		out.Append("If-Match", fmt.Sprintf("%v", *o.IfMatch))
	}
	return &out
}

func (o DeleteEntitlementManagementCatalogResourceRoleResourceOperationOptions) ToOData() *odata.Query {
	out := odata.Query{}

	return &out
}

func (o DeleteEntitlementManagementCatalogResourceRoleResourceOperationOptions) ToQuery() *client.QueryParams {
	out := client.QueryParams{}

	return &out
}

// DeleteEntitlementManagementCatalogResourceRoleResource - Delete navigation property resource for identityGovernance
func (c EntitlementManagementCatalogResourceRoleResourceClient) DeleteEntitlementManagementCatalogResourceRoleResource(ctx context.Context, id stable.IdentityGovernanceEntitlementManagementCatalogIdResourceIdRoleId, options DeleteEntitlementManagementCatalogResourceRoleResourceOperationOptions) (result DeleteEntitlementManagementCatalogResourceRoleResourceOperationResponse, err error) {
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
