package entitlementmanagementaccesspackageresourceaccesspackageresourcescope

import (
	"context"
	"fmt"
	"net/http"

	"github.com/hashicorp/go-azure-sdk/microsoft-graph/common-types/beta"
	"github.com/hashicorp/go-azure-sdk/sdk/client"
	"github.com/hashicorp/go-azure-sdk/sdk/odata"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type CreateEntitlementManagementAccessPackageResourceScopeOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
	Model        *beta.AccessPackageResourceScope
}

type CreateEntitlementManagementAccessPackageResourceScopeOperationOptions struct {
	Metadata  *odata.Metadata
	RetryFunc client.RequestRetryFunc
}

func DefaultCreateEntitlementManagementAccessPackageResourceScopeOperationOptions() CreateEntitlementManagementAccessPackageResourceScopeOperationOptions {
	return CreateEntitlementManagementAccessPackageResourceScopeOperationOptions{}
}

func (o CreateEntitlementManagementAccessPackageResourceScopeOperationOptions) ToHeaders() *client.Headers {
	out := client.Headers{}

	return &out
}

func (o CreateEntitlementManagementAccessPackageResourceScopeOperationOptions) ToOData() *odata.Query {
	out := odata.Query{}
	if o.Metadata != nil {
		out.Metadata = *o.Metadata
	}
	return &out
}

func (o CreateEntitlementManagementAccessPackageResourceScopeOperationOptions) ToQuery() *client.QueryParams {
	out := client.QueryParams{}

	return &out
}

// CreateEntitlementManagementAccessPackageResourceScope - Create new navigation property to accessPackageResourceScopes
// for identityGovernance
func (c EntitlementManagementAccessPackageResourceAccessPackageResourceScopeClient) CreateEntitlementManagementAccessPackageResourceScope(ctx context.Context, id beta.IdentityGovernanceEntitlementManagementAccessPackageResourceId, input beta.AccessPackageResourceScope, options CreateEntitlementManagementAccessPackageResourceScopeOperationOptions) (result CreateEntitlementManagementAccessPackageResourceScopeOperationResponse, err error) {
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
		Path:          fmt.Sprintf("%s/accessPackageResourceScopes", id.ID()),
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

	var model beta.AccessPackageResourceScope
	result.Model = &model
	if err = resp.Unmarshal(result.Model); err != nil {
		return
	}

	return
}
