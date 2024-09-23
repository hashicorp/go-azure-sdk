package entitlementmanagementresourcenamespaceresourceaction

import (
	"context"
	"net/http"

	"github.com/hashicorp/go-azure-sdk/microsoft-graph/common-types/beta"
	"github.com/hashicorp/go-azure-sdk/sdk/client"
	"github.com/hashicorp/go-azure-sdk/sdk/odata"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type UpdateEntitlementManagementResourceNamespaceResourceActionOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
}

type UpdateEntitlementManagementResourceNamespaceResourceActionOperationOptions struct {
	Metadata  *odata.Metadata
	RetryFunc client.RequestRetryFunc
}

func DefaultUpdateEntitlementManagementResourceNamespaceResourceActionOperationOptions() UpdateEntitlementManagementResourceNamespaceResourceActionOperationOptions {
	return UpdateEntitlementManagementResourceNamespaceResourceActionOperationOptions{}
}

func (o UpdateEntitlementManagementResourceNamespaceResourceActionOperationOptions) ToHeaders() *client.Headers {
	out := client.Headers{}

	return &out
}

func (o UpdateEntitlementManagementResourceNamespaceResourceActionOperationOptions) ToOData() *odata.Query {
	out := odata.Query{}
	if o.Metadata != nil {
		out.Metadata = *o.Metadata
	}
	return &out
}

func (o UpdateEntitlementManagementResourceNamespaceResourceActionOperationOptions) ToQuery() *client.QueryParams {
	out := client.QueryParams{}

	return &out
}

// UpdateEntitlementManagementResourceNamespaceResourceAction - Update the navigation property resourceActions in
// roleManagement
func (c EntitlementManagementResourceNamespaceResourceActionClient) UpdateEntitlementManagementResourceNamespaceResourceAction(ctx context.Context, id beta.RoleManagementEntitlementManagementResourceNamespaceIdResourceActionId, input beta.UnifiedRbacResourceAction, options UpdateEntitlementManagementResourceNamespaceResourceActionOperationOptions) (result UpdateEntitlementManagementResourceNamespaceResourceActionOperationResponse, err error) {
	opts := client.RequestOptions{
		ContentType: "application/json; charset=utf-8",
		ExpectedStatusCodes: []int{
			http.StatusNoContent,
		},
		HttpMethod:    http.MethodPatch,
		OptionsObject: options,
		Path:          id.ID(),
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

	return
}
