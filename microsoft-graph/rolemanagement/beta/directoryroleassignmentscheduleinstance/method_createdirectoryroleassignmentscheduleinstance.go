package directoryroleassignmentscheduleinstance

import (
	"context"
	"net/http"

	"github.com/hashicorp/go-azure-sdk/microsoft-graph/common-types/beta"
	"github.com/hashicorp/go-azure-sdk/sdk/client"
	"github.com/hashicorp/go-azure-sdk/sdk/odata"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type CreateDirectoryRoleAssignmentScheduleInstanceOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
	Model        *beta.UnifiedRoleAssignmentScheduleInstance
}

type CreateDirectoryRoleAssignmentScheduleInstanceOperationOptions struct {
	Metadata  *odata.Metadata
	RetryFunc client.RequestRetryFunc
}

func DefaultCreateDirectoryRoleAssignmentScheduleInstanceOperationOptions() CreateDirectoryRoleAssignmentScheduleInstanceOperationOptions {
	return CreateDirectoryRoleAssignmentScheduleInstanceOperationOptions{}
}

func (o CreateDirectoryRoleAssignmentScheduleInstanceOperationOptions) ToHeaders() *client.Headers {
	out := client.Headers{}

	return &out
}

func (o CreateDirectoryRoleAssignmentScheduleInstanceOperationOptions) ToOData() *odata.Query {
	out := odata.Query{}
	if o.Metadata != nil {
		out.Metadata = *o.Metadata
	}
	return &out
}

func (o CreateDirectoryRoleAssignmentScheduleInstanceOperationOptions) ToQuery() *client.QueryParams {
	out := client.QueryParams{}

	return &out
}

// CreateDirectoryRoleAssignmentScheduleInstance - Create new navigation property to roleAssignmentScheduleInstances for
// roleManagement
func (c DirectoryRoleAssignmentScheduleInstanceClient) CreateDirectoryRoleAssignmentScheduleInstance(ctx context.Context, input beta.UnifiedRoleAssignmentScheduleInstance, options CreateDirectoryRoleAssignmentScheduleInstanceOperationOptions) (result CreateDirectoryRoleAssignmentScheduleInstanceOperationResponse, err error) {
	opts := client.RequestOptions{
		ContentType: "application/json; charset=utf-8",
		ExpectedStatusCodes: []int{
			http.StatusCreated,
		},
		HttpMethod:    http.MethodPost,
		OptionsObject: options,
		Path:          "/roleManagement/directory/roleAssignmentScheduleInstances",
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

	var model beta.UnifiedRoleAssignmentScheduleInstance
	result.Model = &model
	if err = resp.Unmarshal(result.Model); err != nil {
		return
	}

	return
}
