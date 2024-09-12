package entitlementmanagementassignmentpolicycatalog

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

type GetEntitlementManagementAssignmentPolicyCatalogOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
	Model        *stable.AccessPackageCatalog
}

type GetEntitlementManagementAssignmentPolicyCatalogOperationOptions struct {
	Expand *odata.Expand
	Select *[]string
}

func DefaultGetEntitlementManagementAssignmentPolicyCatalogOperationOptions() GetEntitlementManagementAssignmentPolicyCatalogOperationOptions {
	return GetEntitlementManagementAssignmentPolicyCatalogOperationOptions{}
}

func (o GetEntitlementManagementAssignmentPolicyCatalogOperationOptions) ToHeaders() *client.Headers {
	out := client.Headers{}

	return &out
}

func (o GetEntitlementManagementAssignmentPolicyCatalogOperationOptions) ToOData() *odata.Query {
	out := odata.Query{}
	if o.Expand != nil {
		out.Expand = *o.Expand
	}
	if o.Select != nil {
		out.Select = *o.Select
	}
	return &out
}

func (o GetEntitlementManagementAssignmentPolicyCatalogOperationOptions) ToQuery() *client.QueryParams {
	out := client.QueryParams{}

	return &out
}

// GetEntitlementManagementAssignmentPolicyCatalog - Get catalog from identityGovernance. Catalog of the access package
// containing this policy. Read-only.
func (c EntitlementManagementAssignmentPolicyCatalogClient) GetEntitlementManagementAssignmentPolicyCatalog(ctx context.Context, id stable.IdentityGovernanceEntitlementManagementAssignmentPolicyId, options GetEntitlementManagementAssignmentPolicyCatalogOperationOptions) (result GetEntitlementManagementAssignmentPolicyCatalogOperationResponse, err error) {
	opts := client.RequestOptions{
		ContentType: "application/json; charset=utf-8",
		ExpectedStatusCodes: []int{
			http.StatusOK,
		},
		HttpMethod:    http.MethodGet,
		OptionsObject: options,
		Path:          fmt.Sprintf("%s/catalog", id.ID()),
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

	var model stable.AccessPackageCatalog
	result.Model = &model
	if err = resp.Unmarshal(result.Model); err != nil {
		return
	}

	return
}
