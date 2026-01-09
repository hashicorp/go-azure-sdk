package privilegedaccessgroupassignmentschedulerequestactivatedusing

import (
	"context"
	"fmt"
	"net/http"

	"github.com/hashicorp/go-azure-sdk/microsoft-graph/common-types/beta"
	"github.com/hashicorp/go-azure-sdk/sdk/client"
	"github.com/hashicorp/go-azure-sdk/sdk/odata"
)

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type GetPrivilegedAccessGroupAssignmentScheduleRequestActivatedUsingOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
	Model        *beta.PrivilegedAccessGroupEligibilitySchedule
}

type GetPrivilegedAccessGroupAssignmentScheduleRequestActivatedUsingOperationOptions struct {
	Expand    *odata.Expand
	Metadata  *odata.Metadata
	RetryFunc client.RequestRetryFunc
	Select    *[]string
}

func DefaultGetPrivilegedAccessGroupAssignmentScheduleRequestActivatedUsingOperationOptions() GetPrivilegedAccessGroupAssignmentScheduleRequestActivatedUsingOperationOptions {
	return GetPrivilegedAccessGroupAssignmentScheduleRequestActivatedUsingOperationOptions{}
}

func (o GetPrivilegedAccessGroupAssignmentScheduleRequestActivatedUsingOperationOptions) ToHeaders() *client.Headers {
	out := client.Headers{}

	return &out
}

func (o GetPrivilegedAccessGroupAssignmentScheduleRequestActivatedUsingOperationOptions) ToOData() *odata.Query {
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

func (o GetPrivilegedAccessGroupAssignmentScheduleRequestActivatedUsingOperationOptions) ToQuery() *client.QueryParams {
	out := client.QueryParams{}

	return &out
}

// GetPrivilegedAccessGroupAssignmentScheduleRequestActivatedUsing - Get activatedUsing from identityGovernance. When
// the request activates a membership or ownership assignment in PIM for groups, this object represents the eligibility
// policy for the group. Otherwise, it is null. Supports $expand.
func (c PrivilegedAccessGroupAssignmentScheduleRequestActivatedUsingClient) GetPrivilegedAccessGroupAssignmentScheduleRequestActivatedUsing(ctx context.Context, id beta.IdentityGovernancePrivilegedAccessGroupAssignmentScheduleRequestId, options GetPrivilegedAccessGroupAssignmentScheduleRequestActivatedUsingOperationOptions) (result GetPrivilegedAccessGroupAssignmentScheduleRequestActivatedUsingOperationResponse, err error) {
	opts := client.RequestOptions{
		ContentType: "application/json; charset=utf-8",
		ExpectedStatusCodes: []int{
			http.StatusOK,
		},
		HttpMethod:    http.MethodGet,
		OptionsObject: options,
		Path:          fmt.Sprintf("%s/activatedUsing", id.ID()),
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

	var model beta.PrivilegedAccessGroupEligibilitySchedule
	result.Model = &model
	if err = resp.Unmarshal(result.Model); err != nil {
		return
	}

	return
}
