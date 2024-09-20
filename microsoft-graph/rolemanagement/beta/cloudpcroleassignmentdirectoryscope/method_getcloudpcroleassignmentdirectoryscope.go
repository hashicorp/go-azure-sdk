package cloudpcroleassignmentdirectoryscope

import (
	"context"
	"encoding/json"
	"net/http"

	"github.com/hashicorp/go-azure-sdk/microsoft-graph/common-types/beta"
	"github.com/hashicorp/go-azure-sdk/sdk/client"
	"github.com/hashicorp/go-azure-sdk/sdk/odata"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type GetCloudPCRoleAssignmentDirectoryScopeOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
	Model        beta.DirectoryObject
}

type GetCloudPCRoleAssignmentDirectoryScopeOperationOptions struct {
	Expand   *odata.Expand
	Metadata *odata.Metadata
	Select   *[]string
}

func DefaultGetCloudPCRoleAssignmentDirectoryScopeOperationOptions() GetCloudPCRoleAssignmentDirectoryScopeOperationOptions {
	return GetCloudPCRoleAssignmentDirectoryScopeOperationOptions{}
}

func (o GetCloudPCRoleAssignmentDirectoryScopeOperationOptions) ToHeaders() *client.Headers {
	out := client.Headers{}

	return &out
}

func (o GetCloudPCRoleAssignmentDirectoryScopeOperationOptions) ToOData() *odata.Query {
	out := odata.Query{}
	if o.Expand != nil {
		out.Expand = *o.Expand
	}
	if o.Metadata != nil {
		out.Metadata = *o.Metadata
	}
	if o.Select != nil {
		out.Select = *o.Select
	}
	return &out
}

func (o GetCloudPCRoleAssignmentDirectoryScopeOperationOptions) ToQuery() *client.QueryParams {
	out := client.QueryParams{}

	return &out
}

// GetCloudPCRoleAssignmentDirectoryScope - Get directoryScopes from roleManagement. Read-only collection that
// references the directory objects that are scope of the assignment. Provided so that callers can get the directory
// objects using $expand at the same time as getting the role assignment. Read-only. Supports $expand.
func (c CloudPCRoleAssignmentDirectoryScopeClient) GetCloudPCRoleAssignmentDirectoryScope(ctx context.Context, id beta.RoleManagementCloudPCRoleAssignmentIdDirectoryScopeId, options GetCloudPCRoleAssignmentDirectoryScopeOperationOptions) (result GetCloudPCRoleAssignmentDirectoryScopeOperationResponse, err error) {
	opts := client.RequestOptions{
		ContentType: "application/json; charset=utf-8",
		ExpectedStatusCodes: []int{
			http.StatusOK,
		},
		HttpMethod:    http.MethodGet,
		OptionsObject: options,
		Path:          id.ID(),
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

	var respObj json.RawMessage
	if err = resp.Unmarshal(&respObj); err != nil {
		return
	}
	model, err := beta.UnmarshalDirectoryObjectImplementation(respObj)
	if err != nil {
		return
	}
	result.Model = model

	return
}
