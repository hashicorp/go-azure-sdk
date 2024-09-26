package approleassignment

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

type CreateAppRoleAssignmentOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
	Model        *stable.AppRoleAssignment
}

type CreateAppRoleAssignmentOperationOptions struct {
	Metadata  *odata.Metadata
	RetryFunc client.RequestRetryFunc
}

func DefaultCreateAppRoleAssignmentOperationOptions() CreateAppRoleAssignmentOperationOptions {
	return CreateAppRoleAssignmentOperationOptions{}
}

func (o CreateAppRoleAssignmentOperationOptions) ToHeaders() *client.Headers {
	out := client.Headers{}

	return &out
}

func (o CreateAppRoleAssignmentOperationOptions) ToOData() *odata.Query {
	out := odata.Query{}
	if o.Metadata != nil {
		out.Metadata = *o.Metadata
	}
	return &out
}

func (o CreateAppRoleAssignmentOperationOptions) ToQuery() *client.QueryParams {
	out := client.QueryParams{}

	return &out
}

// CreateAppRoleAssignment - Grant an appRoleAssignment to a group. Use this API to assign an app role to a security
// group. All direct members of the group will be considered assigned. Security groups with dynamic memberships are
// supported. To grant an app role assignment to a group, you need three identifiers: Additional licenses might be
// required to use a group to manage access to applications.
func (c AppRoleAssignmentClient) CreateAppRoleAssignment(ctx context.Context, id stable.GroupId, input stable.AppRoleAssignment, options CreateAppRoleAssignmentOperationOptions) (result CreateAppRoleAssignmentOperationResponse, err error) {
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
		Path:          fmt.Sprintf("%s/appRoleAssignments", id.ID()),
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

	var model stable.AppRoleAssignment
	result.Model = &model
	if err = resp.Unmarshal(result.Model); err != nil {
		return
	}

	return
}
