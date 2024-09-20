package privilegedaccessgroupassignmentschedulerequesttargetschedule

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

type GetPrivilegedAccessGroupAssignmentScheduleRequestTargetScheduleOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
	Model        *stable.PrivilegedAccessGroupEligibilitySchedule
}

type GetPrivilegedAccessGroupAssignmentScheduleRequestTargetScheduleOperationOptions struct {
	Expand   *odata.Expand
	Metadata *odata.Metadata
	Select   *[]string
}

func DefaultGetPrivilegedAccessGroupAssignmentScheduleRequestTargetScheduleOperationOptions() GetPrivilegedAccessGroupAssignmentScheduleRequestTargetScheduleOperationOptions {
	return GetPrivilegedAccessGroupAssignmentScheduleRequestTargetScheduleOperationOptions{}
}

func (o GetPrivilegedAccessGroupAssignmentScheduleRequestTargetScheduleOperationOptions) ToHeaders() *client.Headers {
	out := client.Headers{}

	return &out
}

func (o GetPrivilegedAccessGroupAssignmentScheduleRequestTargetScheduleOperationOptions) ToOData() *odata.Query {
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

func (o GetPrivilegedAccessGroupAssignmentScheduleRequestTargetScheduleOperationOptions) ToQuery() *client.QueryParams {
	out := client.QueryParams{}

	return &out
}

// GetPrivilegedAccessGroupAssignmentScheduleRequestTargetSchedule - Get targetSchedule from identityGovernance.
// Schedule created by this request. Supports $expand.
func (c PrivilegedAccessGroupAssignmentScheduleRequestTargetScheduleClient) GetPrivilegedAccessGroupAssignmentScheduleRequestTargetSchedule(ctx context.Context, id stable.IdentityGovernancePrivilegedAccessGroupAssignmentScheduleRequestId, options GetPrivilegedAccessGroupAssignmentScheduleRequestTargetScheduleOperationOptions) (result GetPrivilegedAccessGroupAssignmentScheduleRequestTargetScheduleOperationResponse, err error) {
	opts := client.RequestOptions{
		ContentType: "application/json; charset=utf-8",
		ExpectedStatusCodes: []int{
			http.StatusOK,
		},
		HttpMethod:    http.MethodGet,
		OptionsObject: options,
		Path:          fmt.Sprintf("%s/targetSchedule", id.ID()),
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

	var model stable.PrivilegedAccessGroupEligibilitySchedule
	result.Model = &model
	if err = resp.Unmarshal(result.Model); err != nil {
		return
	}

	return
}
