package directoryroleassignmentschedulerequest

import (
	"context"
	"net/http"

	"github.com/hashicorp/go-azure-sdk/microsoft-graph/common-types/beta"
	"github.com/hashicorp/go-azure-sdk/sdk/client"
	"github.com/hashicorp/go-azure-sdk/sdk/odata"
)

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type CreateDirectoryRoleAssignmentScheduleRequestOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
	Model        *beta.UnifiedRoleAssignmentScheduleRequest
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

// CreateDirectoryRoleAssignmentScheduleRequest - Create roleAssignmentScheduleRequests. Create a new
// unifiedRoleAssignmentScheduleRequest object. This operation allows both admins and users to add, remove, extend, or
// renew assignments. To run this request, the calling user must have multifactor authentication (MFA) enforced, and
// running the query in a session in which they were challenged for MFA. See Enable per-user Microsoft Entra multifactor
// authentication to secure sign-in events.
func (c DirectoryRoleAssignmentScheduleRequestClient) CreateDirectoryRoleAssignmentScheduleRequest(ctx context.Context, input beta.UnifiedRoleAssignmentScheduleRequest, options CreateDirectoryRoleAssignmentScheduleRequestOperationOptions) (result CreateDirectoryRoleAssignmentScheduleRequestOperationResponse, err error) {
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

	var model beta.UnifiedRoleAssignmentScheduleRequest
	result.Model = &model
	if err = resp.Unmarshal(result.Model); err != nil {
		return
	}

	return
}
