package privilegedaccessgroupeligibilityschedulerequestprincipal

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

type GetPrivilegedAccessGroupEligibilityScheduleRequestPrincipalOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
	Model        stable.DirectoryObject
}

type GetPrivilegedAccessGroupEligibilityScheduleRequestPrincipalOperationOptions struct {
	Expand    *odata.Expand
	Metadata  *odata.Metadata
	RetryFunc client.RequestRetryFunc
	Select    *[]string
}

func DefaultGetPrivilegedAccessGroupEligibilityScheduleRequestPrincipalOperationOptions() GetPrivilegedAccessGroupEligibilityScheduleRequestPrincipalOperationOptions {
	return GetPrivilegedAccessGroupEligibilityScheduleRequestPrincipalOperationOptions{}
}

func (o GetPrivilegedAccessGroupEligibilityScheduleRequestPrincipalOperationOptions) ToHeaders() *client.Headers {
	out := client.Headers{}

	return &out
}

func (o GetPrivilegedAccessGroupEligibilityScheduleRequestPrincipalOperationOptions) ToOData() *odata.Query {
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

func (o GetPrivilegedAccessGroupEligibilityScheduleRequestPrincipalOperationOptions) ToQuery() *client.QueryParams {
	out := client.QueryParams{}

	return &out
}

// GetPrivilegedAccessGroupEligibilityScheduleRequestPrincipal - Get principal from identityGovernance. References the
// principal that's in the scope of the membership or ownership eligibility request through the group that's governed by
// PIM. Supports $expand and $select nested in $expand for id only.
func (c PrivilegedAccessGroupEligibilityScheduleRequestPrincipalClient) GetPrivilegedAccessGroupEligibilityScheduleRequestPrincipal(ctx context.Context, id stable.IdentityGovernancePrivilegedAccessGroupEligibilityScheduleRequestId, options GetPrivilegedAccessGroupEligibilityScheduleRequestPrincipalOperationOptions) (result GetPrivilegedAccessGroupEligibilityScheduleRequestPrincipalOperationResponse, err error) {
	opts := client.RequestOptions{
		ContentType: "application/json; charset=utf-8",
		ExpectedStatusCodes: []int{
			http.StatusOK,
		},
		HttpMethod:    http.MethodGet,
		OptionsObject: options,
		Path:          fmt.Sprintf("%s/principal", id.ID()),
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
