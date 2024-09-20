package cloudpcroledefinitioninheritspermissionsfrom

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

type CreateCloudPCRoleDefinitionInheritsPermissionsFromOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
	Model        *beta.UnifiedRoleDefinition
}

type CreateCloudPCRoleDefinitionInheritsPermissionsFromOperationOptions struct {
	Metadata *odata.Metadata
}

func DefaultCreateCloudPCRoleDefinitionInheritsPermissionsFromOperationOptions() CreateCloudPCRoleDefinitionInheritsPermissionsFromOperationOptions {
	return CreateCloudPCRoleDefinitionInheritsPermissionsFromOperationOptions{}
}

func (o CreateCloudPCRoleDefinitionInheritsPermissionsFromOperationOptions) ToHeaders() *client.Headers {
	out := client.Headers{}

	return &out
}

func (o CreateCloudPCRoleDefinitionInheritsPermissionsFromOperationOptions) ToOData() *odata.Query {
	out := odata.Query{}
	if o.Metadata != nil {
		out.Metadata = *o.Metadata
	}
	return &out
}

func (o CreateCloudPCRoleDefinitionInheritsPermissionsFromOperationOptions) ToQuery() *client.QueryParams {
	out := client.QueryParams{}

	return &out
}

// CreateCloudPCRoleDefinitionInheritsPermissionsFrom - Create new navigation property to inheritsPermissionsFrom for
// roleManagement
func (c CloudPCRoleDefinitionInheritsPermissionsFromClient) CreateCloudPCRoleDefinitionInheritsPermissionsFrom(ctx context.Context, id beta.RoleManagementCloudPCRoleDefinitionId, input beta.UnifiedRoleDefinition, options CreateCloudPCRoleDefinitionInheritsPermissionsFromOperationOptions) (result CreateCloudPCRoleDefinitionInheritsPermissionsFromOperationResponse, err error) {
	opts := client.RequestOptions{
		ContentType: "application/json; charset=utf-8",
		ExpectedStatusCodes: []int{
			http.StatusCreated,
		},
		HttpMethod:    http.MethodPost,
		OptionsObject: options,
		Path:          fmt.Sprintf("%s/inheritsPermissionsFrom", id.ID()),
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

	var model beta.UnifiedRoleDefinition
	result.Model = &model
	if err = resp.Unmarshal(result.Model); err != nil {
		return
	}

	return
}
