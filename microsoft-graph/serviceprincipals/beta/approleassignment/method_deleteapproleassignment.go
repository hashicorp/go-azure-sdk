package approleassignment

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

type DeleteAppRoleAssignmentOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
}

type DeleteAppRoleAssignmentOperationOptions struct {
	IfMatch   *string
	Metadata  *odata.Metadata
	RetryFunc client.RequestRetryFunc
}

func DefaultDeleteAppRoleAssignmentOperationOptions() DeleteAppRoleAssignmentOperationOptions {
	return DeleteAppRoleAssignmentOperationOptions{}
}

func (o DeleteAppRoleAssignmentOperationOptions) ToHeaders() *client.Headers {
	out := client.Headers{}
	if o.IfMatch != nil {
		out.Append("If-Match", fmt.Sprintf("%v", *o.IfMatch))
	}
	return &out
}

func (o DeleteAppRoleAssignmentOperationOptions) ToOData() *odata.Query {
	out := odata.Query{}
	if o.Metadata != nil {
		out.Metadata = *o.Metadata
	}
	return &out
}

func (o DeleteAppRoleAssignmentOperationOptions) ToQuery() *client.QueryParams {
	out := client.QueryParams{}

	return &out
}

// DeleteAppRoleAssignment - Delete appRoleAssignment. Deletes an appRoleAssignment that a service principal has been
// granted. App roles that are assigned to service principals are also known as application permissions. Deleting an app
// role assignment for a service principal is equivalent to revoking the app-only permission grant.
func (c AppRoleAssignmentClient) DeleteAppRoleAssignment(ctx context.Context, id beta.ServicePrincipalIdAppRoleAssignmentId, options DeleteAppRoleAssignmentOperationOptions) (result DeleteAppRoleAssignmentOperationResponse, err error) {
	opts := client.RequestOptions{
		ContentType: "application/json; charset=utf-8",
		ExpectedStatusCodes: []int{
			http.StatusNoContent,
			http.StatusOK,
		},
		HttpMethod:    http.MethodDelete,
		OptionsObject: options,
		Path:          id.ID(),
		RetryFunc:     options.RetryFunc,
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
