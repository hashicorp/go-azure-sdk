package cloudpcroledefinition

import (
	"context"
	"net/http"

	"github.com/hashicorp/go-azure-sdk/microsoft-graph/common-types/beta"
	"github.com/hashicorp/go-azure-sdk/sdk/client"
	"github.com/hashicorp/go-azure-sdk/sdk/odata"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type UpdateCloudPCRoleDefinitionOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
}

type UpdateCloudPCRoleDefinitionOperationOptions struct {
	Metadata *odata.Metadata
}

func DefaultUpdateCloudPCRoleDefinitionOperationOptions() UpdateCloudPCRoleDefinitionOperationOptions {
	return UpdateCloudPCRoleDefinitionOperationOptions{}
}

func (o UpdateCloudPCRoleDefinitionOperationOptions) ToHeaders() *client.Headers {
	out := client.Headers{}

	return &out
}

func (o UpdateCloudPCRoleDefinitionOperationOptions) ToOData() *odata.Query {
	out := odata.Query{}
	if o.Metadata != nil {
		out.Metadata = *o.Metadata
	}
	return &out
}

func (o UpdateCloudPCRoleDefinitionOperationOptions) ToQuery() *client.QueryParams {
	out := client.QueryParams{}

	return &out
}

// UpdateCloudPCRoleDefinition - Update unifiedRoleDefinition. Update the properties of a unifiedRoleDefinition object
// for an RBAC provider. You cannot update built-in roles. This feature requires a Microsoft Entra ID P1 or P2 license.
// The following RBAC providers are currently supported: - Cloud PC - device management (Intune) - directory (Microsoft
// Entra ID)
func (c CloudPCRoleDefinitionClient) UpdateCloudPCRoleDefinition(ctx context.Context, id beta.RoleManagementCloudPCRoleDefinitionId, input beta.UnifiedRoleDefinition, options UpdateCloudPCRoleDefinitionOperationOptions) (result UpdateCloudPCRoleDefinitionOperationResponse, err error) {
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
