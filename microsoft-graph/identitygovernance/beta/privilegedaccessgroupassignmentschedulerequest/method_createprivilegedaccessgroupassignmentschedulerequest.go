package privilegedaccessgroupassignmentschedulerequest

import (
	"context"
	"net/http"

	"github.com/hashicorp/go-azure-sdk/microsoft-graph/common-types/beta"
	"github.com/hashicorp/go-azure-sdk/sdk/client"
	"github.com/hashicorp/go-azure-sdk/sdk/odata"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type CreatePrivilegedAccessGroupAssignmentScheduleRequestOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
	Model        *beta.PrivilegedAccessGroupAssignmentScheduleRequest
}

type CreatePrivilegedAccessGroupAssignmentScheduleRequestOperationOptions struct {
	Metadata *odata.Metadata
}

func DefaultCreatePrivilegedAccessGroupAssignmentScheduleRequestOperationOptions() CreatePrivilegedAccessGroupAssignmentScheduleRequestOperationOptions {
	return CreatePrivilegedAccessGroupAssignmentScheduleRequestOperationOptions{}
}

func (o CreatePrivilegedAccessGroupAssignmentScheduleRequestOperationOptions) ToHeaders() *client.Headers {
	out := client.Headers{}

	return &out
}

func (o CreatePrivilegedAccessGroupAssignmentScheduleRequestOperationOptions) ToOData() *odata.Query {
	out := odata.Query{}
	if o.Metadata != nil {
		out.Metadata = *o.Metadata
	}
	return &out
}

func (o CreatePrivilegedAccessGroupAssignmentScheduleRequestOperationOptions) ToQuery() *client.QueryParams {
	out := client.QueryParams{}

	return &out
}

// CreatePrivilegedAccessGroupAssignmentScheduleRequest - Create assignmentScheduleRequest. Create a new
// privilegedAccessGroupAssignmentScheduleRequest object.
func (c PrivilegedAccessGroupAssignmentScheduleRequestClient) CreatePrivilegedAccessGroupAssignmentScheduleRequest(ctx context.Context, input beta.PrivilegedAccessGroupAssignmentScheduleRequest, options CreatePrivilegedAccessGroupAssignmentScheduleRequestOperationOptions) (result CreatePrivilegedAccessGroupAssignmentScheduleRequestOperationResponse, err error) {
	opts := client.RequestOptions{
		ContentType: "application/json; charset=utf-8",
		ExpectedStatusCodes: []int{
			http.StatusCreated,
		},
		HttpMethod:    http.MethodPost,
		OptionsObject: options,
		Path:          "/identityGovernance/privilegedAccess/group/assignmentScheduleRequests",
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

	var model beta.PrivilegedAccessGroupAssignmentScheduleRequest
	result.Model = &model
	if err = resp.Unmarshal(result.Model); err != nil {
		return
	}

	return
}
