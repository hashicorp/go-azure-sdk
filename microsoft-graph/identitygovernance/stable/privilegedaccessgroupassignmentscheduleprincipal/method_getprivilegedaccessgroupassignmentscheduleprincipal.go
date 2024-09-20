package privilegedaccessgroupassignmentscheduleprincipal

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/hashicorp/go-azure-sdk/microsoft-graph/common-types/stable"
	"github.com/hashicorp/go-azure-sdk/sdk/client"
	"github.com/hashicorp/go-azure-sdk/sdk/odata"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type GetPrivilegedAccessGroupAssignmentSchedulePrincipalOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
	Model        stable.DirectoryObject
}

type GetPrivilegedAccessGroupAssignmentSchedulePrincipalOperationOptions struct {
	Expand   *odata.Expand
	Metadata *odata.Metadata
	Select   *[]string
}

func DefaultGetPrivilegedAccessGroupAssignmentSchedulePrincipalOperationOptions() GetPrivilegedAccessGroupAssignmentSchedulePrincipalOperationOptions {
	return GetPrivilegedAccessGroupAssignmentSchedulePrincipalOperationOptions{}
}

func (o GetPrivilegedAccessGroupAssignmentSchedulePrincipalOperationOptions) ToHeaders() *client.Headers {
	out := client.Headers{}

	return &out
}

func (o GetPrivilegedAccessGroupAssignmentSchedulePrincipalOperationOptions) ToOData() *odata.Query {
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

func (o GetPrivilegedAccessGroupAssignmentSchedulePrincipalOperationOptions) ToQuery() *client.QueryParams {
	out := client.QueryParams{}

	return &out
}

// GetPrivilegedAccessGroupAssignmentSchedulePrincipal - Get principal from identityGovernance. References the principal
// that's in the scope of this membership or ownership assignment request to the group that's governed through PIM.
// Supports $expand and $select nested in $expand for id only.
func (c PrivilegedAccessGroupAssignmentSchedulePrincipalClient) GetPrivilegedAccessGroupAssignmentSchedulePrincipal(ctx context.Context, id stable.IdentityGovernancePrivilegedAccessGroupAssignmentScheduleId, options GetPrivilegedAccessGroupAssignmentSchedulePrincipalOperationOptions) (result GetPrivilegedAccessGroupAssignmentSchedulePrincipalOperationResponse, err error) {
	opts := client.RequestOptions{
		ContentType: "application/json; charset=utf-8",
		ExpectedStatusCodes: []int{
			http.StatusOK,
		},
		HttpMethod:    http.MethodGet,
		OptionsObject: options,
		Path:          fmt.Sprintf("%s/principal", id.ID()),
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

	var respObj json.RawMessage
	if err = resp.Unmarshal(&respObj); err != nil {
		return
	}
	model, err := stable.UnmarshalDirectoryObjectImplementation(respObj)
	if err != nil {
		return
	}
	result.Model = model

	return
}
