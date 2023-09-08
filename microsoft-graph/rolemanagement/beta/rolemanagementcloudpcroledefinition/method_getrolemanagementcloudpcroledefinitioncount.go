package rolemanagementcloudpcroledefinition

import (
	"context"
	"net/http"

	"github.com/hashicorp/go-azure-sdk/sdk/client"
	"github.com/hashicorp/go-azure-sdk/sdk/odata"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type GetRoleManagementCloudPCRoleDefinitionCountOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
}

// GetRoleManagementCloudPCRoleDefinitionCount ...
func (c RoleManagementCloudPCRoleDefinitionClient) GetRoleManagementCloudPCRoleDefinitionCount(ctx context.Context) (result GetRoleManagementCloudPCRoleDefinitionCountOperationResponse, err error) {
	opts := client.RequestOptions{
		ContentType: "text/plain",
		ExpectedStatusCodes: []int{
			http.StatusOK,
		},
		HttpMethod: http.MethodGet,
		Path:       "/roleManagement/cloudPC/roleDefinitions/$count",
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
