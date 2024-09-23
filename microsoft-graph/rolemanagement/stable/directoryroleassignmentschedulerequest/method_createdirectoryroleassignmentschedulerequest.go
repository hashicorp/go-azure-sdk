package directoryroleassignmentschedulerequest

import (
	"context"
	"net/http"

	"github.com/hashicorp/go-azure-sdk/microsoft-graph/common-types/stable"
	"github.com/hashicorp/go-azure-sdk/sdk/client"
	"github.com/hashicorp/go-azure-sdk/sdk/odata"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type CreateDirectoryRoleAssignmentScheduleRequestOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
	Model        *stable.UnifiedRoleAssignmentScheduleRequest
}

type CreateDirectoryRoleAssignmentScheduleRequestOperationOptions struct {
	Metadata  *odata.Metadata
	RetryFunc client.RequestRetryFunc
}

func DefaultCreateDirectoryRoleAssignmentScheduleRequestOperationOptions() CreateDirectoryRoleAssignmentScheduleRequestOperationOptions {
	return CreateDirectoryRoleAssignmentScheduleRequestOperationOptions{}
}

func (o CreateDirectoryRoleAssignmentScheduleRequestOperationOptions) ToHeaders() *client.Headers {
	out := client.Headers{}

	return &out
}

func (o CreateDirectoryRoleAssignmentScheduleRequestOperationOptions) ToOData() *odata.Query {
	out := odata.Query{}
	if o.Metadata != nil {
		out.Metadata = *o.Metadata
	}
	return &out
}

func (o CreateDirectoryRoleAssignmentScheduleRequestOperationOptions) ToQuery() *client.QueryParams {
	out := client.QueryParams{}

	return &out
}

// CreateDirectoryRoleAssignmentScheduleRequest - Create roleAssignmentScheduleRequests. In PIM, carry out the following
// operations through the unifiedRoleAssignmentScheduleRequest object: To call this API to update, renew, and extend
// assignments for yourself, you must have multifactor authentication (MFA) enforced, and running the query in a session
// in which they were challenged for MFA. See Enable per-user Microsoft Entra multifactor authentication to secure
// sign-in events.
func (c DirectoryRoleAssignmentScheduleRequestClient) CreateDirectoryRoleAssignmentScheduleRequest(ctx context.Context, input stable.UnifiedRoleAssignmentScheduleRequest, options CreateDirectoryRoleAssignmentScheduleRequestOperationOptions) (result CreateDirectoryRoleAssignmentScheduleRequestOperationResponse, err error) {
	opts := client.RequestOptions{
		ContentType: "application/json; charset=utf-8",
		ExpectedStatusCodes: []int{
			http.StatusCreated,
		},
		HttpMethod:    http.MethodPost,
		OptionsObject: options,
		Path:          "/roleManagement/directory/roleAssignmentScheduleRequests",
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

	var model stable.UnifiedRoleAssignmentScheduleRequest
	result.Model = &model
	if err = resp.Unmarshal(result.Model); err != nil {
		return
	}

	return
}
