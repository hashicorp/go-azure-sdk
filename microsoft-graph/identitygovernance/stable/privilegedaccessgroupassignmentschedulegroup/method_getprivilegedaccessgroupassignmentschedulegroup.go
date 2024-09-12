package privilegedaccessgroupassignmentschedulegroup

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

type GetPrivilegedAccessGroupAssignmentScheduleGroupOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
	Model        *stable.Group
}

type GetPrivilegedAccessGroupAssignmentScheduleGroupOperationOptions struct {
	Expand *odata.Expand
	Select *[]string
}

func DefaultGetPrivilegedAccessGroupAssignmentScheduleGroupOperationOptions() GetPrivilegedAccessGroupAssignmentScheduleGroupOperationOptions {
	return GetPrivilegedAccessGroupAssignmentScheduleGroupOperationOptions{}
}

func (o GetPrivilegedAccessGroupAssignmentScheduleGroupOperationOptions) ToHeaders() *client.Headers {
	out := client.Headers{}

	return &out
}

func (o GetPrivilegedAccessGroupAssignmentScheduleGroupOperationOptions) ToOData() *odata.Query {
	out := odata.Query{}
	if o.Expand != nil {
		out.Expand = *o.Expand
	}
	if o.Select != nil {
		out.Select = *o.Select
	}
	return &out
}

func (o GetPrivilegedAccessGroupAssignmentScheduleGroupOperationOptions) ToQuery() *client.QueryParams {
	out := client.QueryParams{}

	return &out
}

// GetPrivilegedAccessGroupAssignmentScheduleGroup - Get group from identityGovernance. References the group that is the
// scope of the membership or ownership assignment through PIM for groups. Supports $expand and $select nested in
// $expand for select properties like id, displayName, and mail.
func (c PrivilegedAccessGroupAssignmentScheduleGroupClient) GetPrivilegedAccessGroupAssignmentScheduleGroup(ctx context.Context, id stable.IdentityGovernancePrivilegedAccessGroupAssignmentScheduleId, options GetPrivilegedAccessGroupAssignmentScheduleGroupOperationOptions) (result GetPrivilegedAccessGroupAssignmentScheduleGroupOperationResponse, err error) {
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

	var model stable.Group
	result.Model = &model
	if err = resp.Unmarshal(result.Model); err != nil {
		return
	}

	return
}
