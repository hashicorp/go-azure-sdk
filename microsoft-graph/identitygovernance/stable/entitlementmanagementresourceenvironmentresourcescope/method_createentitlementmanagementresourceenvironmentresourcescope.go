package entitlementmanagementresourceenvironmentresourcescope

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

type CreateEntitlementManagementResourceEnvironmentResourceScopeOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
	Model        *stable.AccessPackageResourceScope
}

type CreateEntitlementManagementResourceEnvironmentResourceScopeOperationOptions struct {
	Metadata  *odata.Metadata
	RetryFunc client.RequestRetryFunc
}

func DefaultCreateEntitlementManagementResourceEnvironmentResourceScopeOperationOptions() CreateEntitlementManagementResourceEnvironmentResourceScopeOperationOptions {
	return CreateEntitlementManagementResourceEnvironmentResourceScopeOperationOptions{}
}

func (o CreateEntitlementManagementResourceEnvironmentResourceScopeOperationOptions) ToHeaders() *client.Headers {
	out := client.Headers{}

	return &out
}

func (o CreateEntitlementManagementResourceEnvironmentResourceScopeOperationOptions) ToOData() *odata.Query {
	out := odata.Query{}
	if o.Metadata != nil {
		out.Metadata = *o.Metadata
	}
	return &out
}

func (o CreateEntitlementManagementResourceEnvironmentResourceScopeOperationOptions) ToQuery() *client.QueryParams {
	out := client.QueryParams{}

	return &out
}

// CreateEntitlementManagementResourceEnvironmentResourceScope - Create new navigation property to scopes for
// identityGovernance
func (c EntitlementManagementResourceEnvironmentResourceScopeClient) CreateEntitlementManagementResourceEnvironmentResourceScope(ctx context.Context, id stable.IdentityGovernanceEntitlementManagementResourceEnvironmentIdResourceId, input stable.AccessPackageResourceScope, options CreateEntitlementManagementResourceEnvironmentResourceScopeOperationOptions) (result CreateEntitlementManagementResourceEnvironmentResourceScopeOperationResponse, err error) {
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
		Path:          fmt.Sprintf("%s/scopes", id.ID()),
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
