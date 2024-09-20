package cloudpcresourcenamespace

import (
	"context"
	"net/http"

	"github.com/hashicorp/go-azure-sdk/microsoft-graph/common-types/beta"
	"github.com/hashicorp/go-azure-sdk/sdk/client"
	"github.com/hashicorp/go-azure-sdk/sdk/odata"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type UpdateCloudPCResourceNamespaceOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
}

type UpdateCloudPCResourceNamespaceOperationOptions struct {
	Metadata *odata.Metadata
}

func DefaultUpdateCloudPCResourceNamespaceOperationOptions() UpdateCloudPCResourceNamespaceOperationOptions {
	return UpdateCloudPCResourceNamespaceOperationOptions{}
}

func (o UpdateCloudPCResourceNamespaceOperationOptions) ToHeaders() *client.Headers {
	out := client.Headers{}

	return &out
}

func (o UpdateCloudPCResourceNamespaceOperationOptions) ToOData() *odata.Query {
	out := odata.Query{}
	if o.Metadata != nil {
		out.Metadata = *o.Metadata
	}
	return &out
}

func (o UpdateCloudPCResourceNamespaceOperationOptions) ToQuery() *client.QueryParams {
	out := client.QueryParams{}

	return &out
}

// UpdateCloudPCResourceNamespace - Update the navigation property resourceNamespaces in roleManagement
func (c CloudPCResourceNamespaceClient) UpdateCloudPCResourceNamespace(ctx context.Context, id beta.RoleManagementCloudPCResourceNamespaceId, input beta.UnifiedRbacResourceNamespace, options UpdateCloudPCResourceNamespaceOperationOptions) (result UpdateCloudPCResourceNamespaceOperationResponse, err error) {
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
