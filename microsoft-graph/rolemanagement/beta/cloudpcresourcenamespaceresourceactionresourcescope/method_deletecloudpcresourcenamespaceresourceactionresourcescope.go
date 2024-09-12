package cloudpcresourcenamespaceresourceactionresourcescope

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

type DeleteCloudPCResourceNamespaceResourceActionResourceScopeOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
}

type DeleteCloudPCResourceNamespaceResourceActionResourceScopeOperationOptions struct {
	IfMatch *string
}

func DefaultDeleteCloudPCResourceNamespaceResourceActionResourceScopeOperationOptions() DeleteCloudPCResourceNamespaceResourceActionResourceScopeOperationOptions {
	return DeleteCloudPCResourceNamespaceResourceActionResourceScopeOperationOptions{}
}

func (o DeleteCloudPCResourceNamespaceResourceActionResourceScopeOperationOptions) ToHeaders() *client.Headers {
	out := client.Headers{}
	if o.IfMatch != nil {
		out.Append("If-Match", fmt.Sprintf("%v", *o.IfMatch))
	}
	return &out
}

func (o DeleteCloudPCResourceNamespaceResourceActionResourceScopeOperationOptions) ToOData() *odata.Query {
	out := odata.Query{}

	return &out
}

func (o DeleteCloudPCResourceNamespaceResourceActionResourceScopeOperationOptions) ToQuery() *client.QueryParams {
	out := client.QueryParams{}

	return &out
}

// DeleteCloudPCResourceNamespaceResourceActionResourceScope - Delete navigation property resourceScope for
// roleManagement
func (c CloudPCResourceNamespaceResourceActionResourceScopeClient) DeleteCloudPCResourceNamespaceResourceActionResourceScope(ctx context.Context, id beta.RoleManagementCloudPCResourceNamespaceIdResourceActionId, options DeleteCloudPCResourceNamespaceResourceActionResourceScopeOperationOptions) (result DeleteCloudPCResourceNamespaceResourceActionResourceScopeOperationResponse, err error) {
	opts := client.RequestOptions{
		ContentType: "application/json; charset=utf-8",
		ExpectedStatusCodes: []int{
			http.StatusNoContent,
		},
		HttpMethod:    http.MethodDelete,
		OptionsObject: options,
		Path:          fmt.Sprintf("%s/resourceScope", id.ID()),
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
