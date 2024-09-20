package enterpriseappresourcenamespaceresourceaction

import (
	"context"
	"net/http"

	"github.com/hashicorp/go-azure-sdk/microsoft-graph/common-types/beta"
	"github.com/hashicorp/go-azure-sdk/sdk/client"
	"github.com/hashicorp/go-azure-sdk/sdk/odata"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type UpdateEnterpriseAppResourceNamespaceResourceActionOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
}

type UpdateEnterpriseAppResourceNamespaceResourceActionOperationOptions struct {
	Metadata *odata.Metadata
}

func DefaultUpdateEnterpriseAppResourceNamespaceResourceActionOperationOptions() UpdateEnterpriseAppResourceNamespaceResourceActionOperationOptions {
	return UpdateEnterpriseAppResourceNamespaceResourceActionOperationOptions{}
}

func (o UpdateEnterpriseAppResourceNamespaceResourceActionOperationOptions) ToHeaders() *client.Headers {
	out := client.Headers{}

	return &out
}

func (o UpdateEnterpriseAppResourceNamespaceResourceActionOperationOptions) ToOData() *odata.Query {
	out := odata.Query{}
	if o.Metadata != nil {
		out.Metadata = *o.Metadata
	}
	return &out
}

func (o UpdateEnterpriseAppResourceNamespaceResourceActionOperationOptions) ToQuery() *client.QueryParams {
	out := client.QueryParams{}

	return &out
}

// UpdateEnterpriseAppResourceNamespaceResourceAction - Update the navigation property resourceActions in roleManagement
func (c EnterpriseAppResourceNamespaceResourceActionClient) UpdateEnterpriseAppResourceNamespaceResourceAction(ctx context.Context, id beta.RoleManagementEnterpriseAppIdResourceNamespaceIdResourceActionId, input beta.UnifiedRbacResourceAction, options UpdateEnterpriseAppResourceNamespaceResourceActionOperationOptions) (result UpdateEnterpriseAppResourceNamespaceResourceActionOperationResponse, err error) {
	opts := client.RequestOptions{
		ContentType: "application/json; charset=utf-8",
		ExpectedStatusCodes: []int{
			http.StatusNoContent,
		},
		HttpMethod:    http.MethodPatch,
		OptionsObject: options,
		Path:          id.ID(),
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
