package cloudpcroleassignment

import (
	"context"
	"net/http"

	"github.com/hashicorp/go-azure-sdk/microsoft-graph/common-types/beta"
	"github.com/hashicorp/go-azure-sdk/sdk/client"
	"github.com/hashicorp/go-azure-sdk/sdk/odata"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type CreateCloudPCRoleAssignmentOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
	Model        *beta.UnifiedRoleAssignmentMultiple
}

type CreateCloudPCRoleAssignmentOperationOptions struct {
	Metadata *odata.Metadata
}

func DefaultCreateCloudPCRoleAssignmentOperationOptions() CreateCloudPCRoleAssignmentOperationOptions {
	return CreateCloudPCRoleAssignmentOperationOptions{}
}

func (o CreateCloudPCRoleAssignmentOperationOptions) ToHeaders() *client.Headers {
	out := client.Headers{}

	return &out
}

func (o CreateCloudPCRoleAssignmentOperationOptions) ToOData() *odata.Query {
	out := odata.Query{}
	if o.Metadata != nil {
		out.Metadata = *o.Metadata
	}
	return &out
}

func (o CreateCloudPCRoleAssignmentOperationOptions) ToQuery() *client.QueryParams {
	out := client.QueryParams{}

	return &out
}

// CreateCloudPCRoleAssignment - Create unifiedRoleAssignmentMultiple. Create a new unifiedRoleAssignmentMultiple object
// for an RBAC provider. The following RBAC providers are currently supported: - Cloud PC - device management (Intune)
// For other Microsoft 365 applications (like Microsoft Entra ID), use unifiedRoleAssignment.
func (c CloudPCRoleAssignmentClient) CreateCloudPCRoleAssignment(ctx context.Context, input beta.UnifiedRoleAssignmentMultiple, options CreateCloudPCRoleAssignmentOperationOptions) (result CreateCloudPCRoleAssignmentOperationResponse, err error) {
	opts := client.RequestOptions{
		ContentType: "application/json; charset=utf-8",
		ExpectedStatusCodes: []int{
			http.StatusCreated,
		},
		HttpMethod:    http.MethodPost,
		OptionsObject: options,
		Path:          "/roleManagement/cloudPC/roleAssignments",
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

	var model beta.UnifiedRoleAssignmentMultiple
	result.Model = &model
	if err = resp.Unmarshal(result.Model); err != nil {
		return
	}

	return
}
