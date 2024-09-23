package entitlementmanagementresourceenvironmentresourcescoperesourceenvironment

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

type GetEntitlementManagementResourceEnvironmentResourceScopeResourceEnvironmentOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
	Model        *stable.AccessPackageResourceEnvironment
}

type GetEntitlementManagementResourceEnvironmentResourceScopeResourceEnvironmentOperationOptions struct {
	Expand    *odata.Expand
	Metadata  *odata.Metadata
	RetryFunc client.RequestRetryFunc
	Select    *[]string
}

func DefaultGetEntitlementManagementResourceEnvironmentResourceScopeResourceEnvironmentOperationOptions() GetEntitlementManagementResourceEnvironmentResourceScopeResourceEnvironmentOperationOptions {
	return GetEntitlementManagementResourceEnvironmentResourceScopeResourceEnvironmentOperationOptions{}
}

func (o GetEntitlementManagementResourceEnvironmentResourceScopeResourceEnvironmentOperationOptions) ToHeaders() *client.Headers {
	out := client.Headers{}

	return &out
}

func (o GetEntitlementManagementResourceEnvironmentResourceScopeResourceEnvironmentOperationOptions) ToOData() *odata.Query {
	out := odata.Query{}
	if o.Expand != nil {
		out.Expand = *o.Expand
	}
	if o.Metadata != nil {
		out.Metadata = *o.Metadata
	}
	if o.Select != nil {
		out.Select = *o.Select
	}
	return &out
}

func (o GetEntitlementManagementResourceEnvironmentResourceScopeResourceEnvironmentOperationOptions) ToQuery() *client.QueryParams {
	out := client.QueryParams{}

	return &out
}

// GetEntitlementManagementResourceEnvironmentResourceScopeResourceEnvironment - Get environment from
// identityGovernance. Contains the environment information for the resource. This can be set using either the
// @odata.bind annotation or the environment's originId.Supports $expand.
func (c EntitlementManagementResourceEnvironmentResourceScopeResourceEnvironmentClient) GetEntitlementManagementResourceEnvironmentResourceScopeResourceEnvironment(ctx context.Context, id stable.IdentityGovernanceEntitlementManagementResourceEnvironmentIdResourceIdScopeId, options GetEntitlementManagementResourceEnvironmentResourceScopeResourceEnvironmentOperationOptions) (result GetEntitlementManagementResourceEnvironmentResourceScopeResourceEnvironmentOperationResponse, err error) {
	opts := client.RequestOptions{
		ContentType: "application/json; charset=utf-8",
		ExpectedStatusCodes: []int{
			http.StatusOK,
		},
		HttpMethod:    http.MethodGet,
		OptionsObject: options,
		Path:          fmt.Sprintf("%s/resource/environment", id.ID()),
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

	var model stable.AccessPackageResourceEnvironment
	result.Model = &model
	if err = resp.Unmarshal(result.Model); err != nil {
		return
	}

	return
}
