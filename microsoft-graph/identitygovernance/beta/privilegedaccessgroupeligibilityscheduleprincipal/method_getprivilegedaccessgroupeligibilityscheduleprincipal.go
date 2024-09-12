package privilegedaccessgroupeligibilityscheduleprincipal

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/hashicorp/go-azure-sdk/microsoft-graph/common-types/beta"
	"github.com/hashicorp/go-azure-sdk/sdk/client"
	"github.com/hashicorp/go-azure-sdk/sdk/odata"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type GetPrivilegedAccessGroupEligibilitySchedulePrincipalOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
	Model        beta.DirectoryObject
}

type GetPrivilegedAccessGroupEligibilitySchedulePrincipalOperationOptions struct {
	Expand *odata.Expand
	Select *[]string
}

func DefaultGetPrivilegedAccessGroupEligibilitySchedulePrincipalOperationOptions() GetPrivilegedAccessGroupEligibilitySchedulePrincipalOperationOptions {
	return GetPrivilegedAccessGroupEligibilitySchedulePrincipalOperationOptions{}
}

func (o GetPrivilegedAccessGroupEligibilitySchedulePrincipalOperationOptions) ToHeaders() *client.Headers {
	out := client.Headers{}

	return &out
}

func (o GetPrivilegedAccessGroupEligibilitySchedulePrincipalOperationOptions) ToOData() *odata.Query {
	out := odata.Query{}
	if o.Expand != nil {
		out.Expand = *o.Expand
	}
	if o.Select != nil {
		out.Select = *o.Select
	}
	return &out
}

func (o GetPrivilegedAccessGroupEligibilitySchedulePrincipalOperationOptions) ToQuery() *client.QueryParams {
	out := client.QueryParams{}

	return &out
}

// GetPrivilegedAccessGroupEligibilitySchedulePrincipal - Get principal from identityGovernance. References the
// principal that's in the scope of this membership or ownership eligibility request to the group that's governed by
// PIM. Supports $expand.
func (c PrivilegedAccessGroupEligibilitySchedulePrincipalClient) GetPrivilegedAccessGroupEligibilitySchedulePrincipal(ctx context.Context, id beta.IdentityGovernancePrivilegedAccessGroupEligibilityScheduleId, options GetPrivilegedAccessGroupEligibilitySchedulePrincipalOperationOptions) (result GetPrivilegedAccessGroupEligibilitySchedulePrincipalOperationResponse, err error) {
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
	model, err := beta.UnmarshalDirectoryObjectImplementation(respObj)
	if err != nil {
		return
	}
	result.Model = model

	return
}
