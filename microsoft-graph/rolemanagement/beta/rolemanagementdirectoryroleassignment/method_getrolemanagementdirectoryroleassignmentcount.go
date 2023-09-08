package rolemanagementdirectoryroleassignment

import (
	"context"
	"net/http"

	"github.com/hashicorp/go-azure-sdk/sdk/client"
	"github.com/hashicorp/go-azure-sdk/sdk/odata"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type GetRoleManagementDirectoryRoleAssignmentCountOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
}

// GetRoleManagementDirectoryRoleAssignmentCount ...
func (c RoleManagementDirectoryRoleAssignmentClient) GetRoleManagementDirectoryRoleAssignmentCount(ctx context.Context) (result GetRoleManagementDirectoryRoleAssignmentCountOperationResponse, err error) {
	opts := client.RequestOptions{
		ContentType: "text/plain",
		ExpectedStatusCodes: []int{
			http.StatusOK,
		},
		HttpMethod: http.MethodGet,
		Path:       "/roleManagement/directory/roleAssignments/$count",
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
