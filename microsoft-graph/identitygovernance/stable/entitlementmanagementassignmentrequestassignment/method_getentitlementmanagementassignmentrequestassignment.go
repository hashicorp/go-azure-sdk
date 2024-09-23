package entitlementmanagementassignmentrequestassignment

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

type GetEntitlementManagementAssignmentRequestAssignmentOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
	Model        *stable.AccessPackageAssignment
}

type GetEntitlementManagementAssignmentRequestAssignmentOperationOptions struct {
	Expand    *odata.Expand
	Metadata  *odata.Metadata
	RetryFunc client.RequestRetryFunc
	Select    *[]string
}

func DefaultGetEntitlementManagementAssignmentRequestAssignmentOperationOptions() GetEntitlementManagementAssignmentRequestAssignmentOperationOptions {
	return GetEntitlementManagementAssignmentRequestAssignmentOperationOptions{}
}

func (o GetEntitlementManagementAssignmentRequestAssignmentOperationOptions) ToHeaders() *client.Headers {
	out := client.Headers{}

	return &out
}

func (o GetEntitlementManagementAssignmentRequestAssignmentOperationOptions) ToOData() *odata.Query {
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

func (o GetEntitlementManagementAssignmentRequestAssignmentOperationOptions) ToQuery() *client.QueryParams {
	out := client.QueryParams{}

	return &out
}

// GetEntitlementManagementAssignmentRequestAssignment - Get assignment from identityGovernance. For a requestType of
// userAdd or adminAdd, this is an access package assignment requested to be created. For a requestType of userRemove,
// adminRemove or systemRemove, this has the id property of an existing assignment to be removed. Supports $expand.
func (c EntitlementManagementAssignmentRequestAssignmentClient) GetEntitlementManagementAssignmentRequestAssignment(ctx context.Context, id stable.IdentityGovernanceEntitlementManagementAssignmentRequestId, options GetEntitlementManagementAssignmentRequestAssignmentOperationOptions) (result GetEntitlementManagementAssignmentRequestAssignmentOperationResponse, err error) {
	opts := client.RequestOptions{
		ContentType: "application/json; charset=utf-8",
		ExpectedStatusCodes: []int{
			http.StatusOK,
		},
		HttpMethod:    http.MethodGet,
		OptionsObject: options,
		Path:          fmt.Sprintf("%s/assignment", id.ID()),
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

	var model stable.AccessPackageAssignment
	result.Model = &model
	if err = resp.Unmarshal(result.Model); err != nil {
		return
	}

	return
}
