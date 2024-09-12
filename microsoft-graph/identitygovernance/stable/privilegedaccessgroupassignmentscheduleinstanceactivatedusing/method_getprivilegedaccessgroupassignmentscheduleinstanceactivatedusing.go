package privilegedaccessgroupassignmentscheduleinstanceactivatedusing

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

type GetPrivilegedAccessGroupAssignmentScheduleInstanceActivatedUsingOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
	Model        *stable.PrivilegedAccessGroupEligibilityScheduleInstance
}

type GetPrivilegedAccessGroupAssignmentScheduleInstanceActivatedUsingOperationOptions struct {
	Expand *odata.Expand
	Select *[]string
}

func DefaultGetPrivilegedAccessGroupAssignmentScheduleInstanceActivatedUsingOperationOptions() GetPrivilegedAccessGroupAssignmentScheduleInstanceActivatedUsingOperationOptions {
	return GetPrivilegedAccessGroupAssignmentScheduleInstanceActivatedUsingOperationOptions{}
}

func (o GetPrivilegedAccessGroupAssignmentScheduleInstanceActivatedUsingOperationOptions) ToHeaders() *client.Headers {
	out := client.Headers{}

	return &out
}

func (o GetPrivilegedAccessGroupAssignmentScheduleInstanceActivatedUsingOperationOptions) ToOData() *odata.Query {
	out := odata.Query{}
	if o.Expand != nil {
		out.Expand = *o.Expand
	}
	if o.Select != nil {
		out.Select = *o.Select
	}
	return &out
}

func (o GetPrivilegedAccessGroupAssignmentScheduleInstanceActivatedUsingOperationOptions) ToQuery() *client.QueryParams {
	out := client.QueryParams{}

	return &out
}

// GetPrivilegedAccessGroupAssignmentScheduleInstanceActivatedUsing - Get activatedUsing from identityGovernance. When
// the request activates a membership or ownership in PIM for groups, this object represents the eligibility request for
// the group. Otherwise, it is null.
func (c PrivilegedAccessGroupAssignmentScheduleInstanceActivatedUsingClient) GetPrivilegedAccessGroupAssignmentScheduleInstanceActivatedUsing(ctx context.Context, id stable.IdentityGovernancePrivilegedAccessGroupAssignmentScheduleInstanceId, options GetPrivilegedAccessGroupAssignmentScheduleInstanceActivatedUsingOperationOptions) (result GetPrivilegedAccessGroupAssignmentScheduleInstanceActivatedUsingOperationResponse, err error) {
	opts := client.RequestOptions{
		ContentType: "application/json; charset=utf-8",
		ExpectedStatusCodes: []int{
			http.StatusOK,
		},
		HttpMethod:    http.MethodGet,
		OptionsObject: options,
		Path:          fmt.Sprintf("%s/activatedUsing", id.ID()),
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

	var model stable.PrivilegedAccessGroupEligibilityScheduleInstance
	result.Model = &model
	if err = resp.Unmarshal(result.Model); err != nil {
		return
	}

	return
}
