package approleassignedto

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

type CreateAppRoleAssignedToOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
	Model        *stable.AppRoleAssignment
}

// CreateAppRoleAssignedTo - Grant an appRoleAssignment for a service principal. Assign an app role for a resource
// service principal, to a user, group, or client service principal. App roles that are assigned to service principals
// are also known as application permissions. Application permissions can be granted directly with app role assignments,
// or through a consent experience. To grant an app role assignment, you need three identifiers:
func (c AppRoleAssignedToClient) CreateAppRoleAssignedTo(ctx context.Context, id stable.ServicePrincipalId, input stable.AppRoleAssignment) (result CreateAppRoleAssignedToOperationResponse, err error) {
	opts := client.RequestOptions{
		ContentType: "application/json; charset=utf-8",
		ExpectedStatusCodes: []int{
			http.StatusCreated,
		},
		HttpMethod: http.MethodPost,
		Path:       fmt.Sprintf("%s/appRoleAssignedTo", id.ID()),
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

	var model stable.AppRoleAssignment
	result.Model = &model
	if err = resp.Unmarshal(result.Model); err != nil {
		return
	}

	return
}
