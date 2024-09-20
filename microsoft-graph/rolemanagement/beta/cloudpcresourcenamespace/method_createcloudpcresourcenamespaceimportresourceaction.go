package cloudpcresourcenamespace

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

type CreateCloudPCResourceNamespaceImportResourceActionOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
	Model        *beta.UnifiedRbacResourceNamespace
}

type CreateCloudPCResourceNamespaceImportResourceActionOperationOptions struct {
	Metadata *odata.Metadata
}

func DefaultCreateCloudPCResourceNamespaceImportResourceActionOperationOptions() CreateCloudPCResourceNamespaceImportResourceActionOperationOptions {
	return CreateCloudPCResourceNamespaceImportResourceActionOperationOptions{}
}

func (o CreateCloudPCResourceNamespaceImportResourceActionOperationOptions) ToHeaders() *client.Headers {
	out := client.Headers{}

	return &out
}

func (o CreateCloudPCResourceNamespaceImportResourceActionOperationOptions) ToOData() *odata.Query {
	out := odata.Query{}
	if o.Metadata != nil {
		out.Metadata = *o.Metadata
	}
	return &out
}

func (o CreateCloudPCResourceNamespaceImportResourceActionOperationOptions) ToQuery() *client.QueryParams {
	out := client.QueryParams{}

	return &out
}

// CreateCloudPCResourceNamespaceImportResourceAction - Invoke action importResourceActions
func (c CloudPCResourceNamespaceClient) CreateCloudPCResourceNamespaceImportResourceAction(ctx context.Context, id beta.RoleManagementCloudPCResourceNamespaceId, input CreateCloudPCResourceNamespaceImportResourceActionRequest, options CreateCloudPCResourceNamespaceImportResourceActionOperationOptions) (result CreateCloudPCResourceNamespaceImportResourceActionOperationResponse, err error) {
	opts := client.RequestOptions{
		ContentType: "application/json; charset=utf-8",
		ExpectedStatusCodes: []int{
			http.StatusOK,
		},
		HttpMethod:    http.MethodPost,
		OptionsObject: options,
		Path:          fmt.Sprintf("%s/importResourceActions", id.ID()),
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

	var model beta.UnifiedRbacResourceNamespace
	result.Model = &model
	if err = resp.Unmarshal(result.Model); err != nil {
		return
	}

	return
}
