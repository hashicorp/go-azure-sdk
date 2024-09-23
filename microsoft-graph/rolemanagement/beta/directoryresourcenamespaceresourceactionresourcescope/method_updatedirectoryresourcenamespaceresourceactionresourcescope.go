package directoryresourcenamespaceresourceactionresourcescope

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

type UpdateDirectoryResourceNamespaceResourceActionResourceScopeOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
}

type UpdateDirectoryResourceNamespaceResourceActionResourceScopeOperationOptions struct {
	Metadata  *odata.Metadata
	RetryFunc client.RequestRetryFunc
}

func DefaultUpdateDirectoryResourceNamespaceResourceActionResourceScopeOperationOptions() UpdateDirectoryResourceNamespaceResourceActionResourceScopeOperationOptions {
	return UpdateDirectoryResourceNamespaceResourceActionResourceScopeOperationOptions{}
}

func (o UpdateDirectoryResourceNamespaceResourceActionResourceScopeOperationOptions) ToHeaders() *client.Headers {
	out := client.Headers{}

	return &out
}

func (o UpdateDirectoryResourceNamespaceResourceActionResourceScopeOperationOptions) ToOData() *odata.Query {
	out := odata.Query{}
	if o.Metadata != nil {
		out.Metadata = *o.Metadata
	}
	return &out
}

func (o UpdateDirectoryResourceNamespaceResourceActionResourceScopeOperationOptions) ToQuery() *client.QueryParams {
	out := client.QueryParams{}

	return &out
}

// UpdateDirectoryResourceNamespaceResourceActionResourceScope - Update the navigation property resourceScope in
// roleManagement
func (c DirectoryResourceNamespaceResourceActionResourceScopeClient) UpdateDirectoryResourceNamespaceResourceActionResourceScope(ctx context.Context, id beta.RoleManagementDirectoryResourceNamespaceIdResourceActionId, input beta.UnifiedRbacResourceScope, options UpdateDirectoryResourceNamespaceResourceActionResourceScopeOperationOptions) (result UpdateDirectoryResourceNamespaceResourceActionResourceScopeOperationResponse, err error) {
	opts := client.RequestOptions{
		ContentType: "application/json; charset=utf-8",
		ExpectedStatusCodes: []int{
			http.StatusNoContent,
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
