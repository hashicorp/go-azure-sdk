package defenderresourcenamespaceresourceactionresourcescope

import (
	"context"
	"fmt"
	"net/http"

	"github.com/hashicorp/go-azure-sdk/microsoft-graph/common-types/beta"
	"github.com/hashicorp/go-azure-sdk/sdk/client"
	"github.com/hashicorp/go-azure-sdk/sdk/odata"
)

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type UpdateDefenderResourceNamespaceResourceActionResourceScopeOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
}

type UpdateDefenderResourceNamespaceResourceActionResourceScopeOperationOptions struct {
	Metadata  *odata.Metadata
	RetryFunc client.RequestRetryFunc
}

func DefaultUpdateDefenderResourceNamespaceResourceActionResourceScopeOperationOptions() UpdateDefenderResourceNamespaceResourceActionResourceScopeOperationOptions {
	return UpdateDefenderResourceNamespaceResourceActionResourceScopeOperationOptions{}
}

func (o UpdateDefenderResourceNamespaceResourceActionResourceScopeOperationOptions) ToHeaders() *client.Headers {
	out := client.Headers{}

	return &out
}

func (o UpdateDefenderResourceNamespaceResourceActionResourceScopeOperationOptions) ToOData() *odata.Query {
	out := odata.Query{}
	if o.Metadata != nil {
		out.Metadata = *o.Metadata
	}
	return &out
}

func (o UpdateDefenderResourceNamespaceResourceActionResourceScopeOperationOptions) ToQuery() *client.QueryParams {
	out := client.QueryParams{}

	return &out
}

// UpdateDefenderResourceNamespaceResourceActionResourceScope - Update the navigation property resourceScope in
// roleManagement
func (c DefenderResourceNamespaceResourceActionResourceScopeClient) UpdateDefenderResourceNamespaceResourceActionResourceScope(ctx context.Context, id beta.RoleManagementDefenderResourceNamespaceIdResourceActionId, input beta.UnifiedRbacResourceScope, options UpdateDefenderResourceNamespaceResourceActionResourceScopeOperationOptions) (result UpdateDefenderResourceNamespaceResourceActionResourceScopeOperationResponse, err error) {
	opts := client.RequestOptions{
		ContentType: "application/json; charset=utf-8",
		ExpectedStatusCodes: []int{
			http.StatusAccepted,
			http.StatusNoContent,
			http.StatusOK,
		},
		HttpMethod:    http.MethodPatch,
		OptionsObject: options,
		Path:          fmt.Sprintf("%s/resourceScope", id.ID()),
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
