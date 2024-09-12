package entitlementmanagementaccesspackageaccesspackagesincompatiblewith

import (
	"context"
	"net/http"

	"github.com/hashicorp/go-azure-sdk/microsoft-graph/common-types/stable"
	"github.com/hashicorp/go-azure-sdk/sdk/client"
	"github.com/hashicorp/go-azure-sdk/sdk/odata"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type GetEntitlementManagementAccessPackageIncompatibleWithOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
	Model        *stable.AccessPackage
}

type GetEntitlementManagementAccessPackageIncompatibleWithOperationOptions struct {
	Expand *odata.Expand
	Select *[]string
}

func DefaultGetEntitlementManagementAccessPackageIncompatibleWithOperationOptions() GetEntitlementManagementAccessPackageIncompatibleWithOperationOptions {
	return GetEntitlementManagementAccessPackageIncompatibleWithOperationOptions{}
}

func (o GetEntitlementManagementAccessPackageIncompatibleWithOperationOptions) ToHeaders() *client.Headers {
	out := client.Headers{}

	return &out
}

func (o GetEntitlementManagementAccessPackageIncompatibleWithOperationOptions) ToOData() *odata.Query {
	out := odata.Query{}
	if o.Expand != nil {
		out.Expand = *o.Expand
	}
	if o.Select != nil {
		out.Select = *o.Select
	}
	return &out
}

func (o GetEntitlementManagementAccessPackageIncompatibleWithOperationOptions) ToQuery() *client.QueryParams {
	out := client.QueryParams{}

	return &out
}

// GetEntitlementManagementAccessPackageIncompatibleWith - Get accessPackagesIncompatibleWith from identityGovernance.
// The access packages that are incompatible with this package. Read-only.
func (c EntitlementManagementAccessPackageAccessPackagesIncompatibleWithClient) GetEntitlementManagementAccessPackageIncompatibleWith(ctx context.Context, id stable.IdentityGovernanceEntitlementManagementAccessPackageIdAccessPackagesIncompatibleWithId, options GetEntitlementManagementAccessPackageIncompatibleWithOperationOptions) (result GetEntitlementManagementAccessPackageIncompatibleWithOperationResponse, err error) {
	opts := client.RequestOptions{
		ContentType: "application/json; charset=utf-8",
		ExpectedStatusCodes: []int{
			http.StatusOK,
		},
		HttpMethod:    http.MethodGet,
		OptionsObject: options,
		Path:          id.ID(),
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

	var model stable.AccessPackage
	result.Model = &model
	if err = resp.Unmarshal(result.Model); err != nil {
		return
	}

	return
}
