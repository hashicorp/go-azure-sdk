package directoryroledefinitioninheritspermissionsfrom

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

type CreateDirectoryRoleDefinitionInheritsPermissionsFromOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
	Model        *stable.UnifiedRoleDefinition
}

type CreateDirectoryRoleDefinitionInheritsPermissionsFromOperationOptions struct {
	Metadata  *odata.Metadata
	RetryFunc client.RequestRetryFunc
}

func DefaultCreateDirectoryRoleDefinitionInheritsPermissionsFromOperationOptions() CreateDirectoryRoleDefinitionInheritsPermissionsFromOperationOptions {
	return CreateDirectoryRoleDefinitionInheritsPermissionsFromOperationOptions{}
}

func (o CreateDirectoryRoleDefinitionInheritsPermissionsFromOperationOptions) ToHeaders() *client.Headers {
	out := client.Headers{}

	return &out
}

func (o CreateDirectoryRoleDefinitionInheritsPermissionsFromOperationOptions) ToOData() *odata.Query {
	out := odata.Query{}
	if o.Metadata != nil {
		out.Metadata = *o.Metadata
	}
	return &out
}

func (o CreateDirectoryRoleDefinitionInheritsPermissionsFromOperationOptions) ToQuery() *client.QueryParams {
	out := client.QueryParams{}

	return &out
}

// CreateDirectoryRoleDefinitionInheritsPermissionsFrom - Create new navigation property to inheritsPermissionsFrom for
// roleManagement
func (c DirectoryRoleDefinitionInheritsPermissionsFromClient) CreateDirectoryRoleDefinitionInheritsPermissionsFrom(ctx context.Context, id stable.RoleManagementDirectoryRoleDefinitionId, input stable.UnifiedRoleDefinition, options CreateDirectoryRoleDefinitionInheritsPermissionsFromOperationOptions) (result CreateDirectoryRoleDefinitionInheritsPermissionsFromOperationResponse, err error) {
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
		Path:          fmt.Sprintf("%s/inheritsPermissionsFrom", id.ID()),
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

	var model stable.UnifiedRoleDefinition
	result.Model = &model
	if err = resp.Unmarshal(result.Model); err != nil {
		return
	}

	return
}
