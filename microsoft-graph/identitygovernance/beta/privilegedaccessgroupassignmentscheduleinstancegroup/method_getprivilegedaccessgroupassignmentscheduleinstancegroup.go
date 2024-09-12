package privilegedaccessgroupassignmentscheduleinstancegroup

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

type GetPrivilegedAccessGroupAssignmentScheduleInstanceGroupOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
	Model        *beta.Group
}

type GetPrivilegedAccessGroupAssignmentScheduleInstanceGroupOperationOptions struct {
	Expand *odata.Expand
	Select *[]string
}

func DefaultGetPrivilegedAccessGroupAssignmentScheduleInstanceGroupOperationOptions() GetPrivilegedAccessGroupAssignmentScheduleInstanceGroupOperationOptions {
	return GetPrivilegedAccessGroupAssignmentScheduleInstanceGroupOperationOptions{}
}

func (o GetPrivilegedAccessGroupAssignmentScheduleInstanceGroupOperationOptions) ToHeaders() *client.Headers {
	out := client.Headers{}

	return &out
}

func (o GetPrivilegedAccessGroupAssignmentScheduleInstanceGroupOperationOptions) ToOData() *odata.Query {
	out := odata.Query{}
	if o.Expand != nil {
		out.Expand = *o.Expand
	}
	if o.Select != nil {
		out.Select = *o.Select
	}
	return &out
}

func (o GetPrivilegedAccessGroupAssignmentScheduleInstanceGroupOperationOptions) ToQuery() *client.QueryParams {
	out := client.QueryParams{}

	return &out
}

// GetPrivilegedAccessGroupAssignmentScheduleInstanceGroup - Get group from identityGovernance. References the group
// that is the scope of the membership or ownership assignment through PIM for groups. Supports $expand.
func (c PrivilegedAccessGroupAssignmentScheduleInstanceGroupClient) GetPrivilegedAccessGroupAssignmentScheduleInstanceGroup(ctx context.Context, id beta.IdentityGovernancePrivilegedAccessGroupAssignmentScheduleInstanceId, options GetPrivilegedAccessGroupAssignmentScheduleInstanceGroupOperationOptions) (result GetPrivilegedAccessGroupAssignmentScheduleInstanceGroupOperationResponse, err error) {
	opts := client.RequestOptions{
		ContentType: "application/json; charset=utf-8",
		ExpectedStatusCodes: []int{
			http.StatusOK,
		},
		HttpMethod:    http.MethodGet,
		OptionsObject: options,
		Path:          fmt.Sprintf("%s/group", id.ID()),
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

	var model beta.Group
	result.Model = &model
	if err = resp.Unmarshal(result.Model); err != nil {
		return
	}

	return
}
