package entitlementmanagementassignmentrequestaccesspackageassignment

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

type GetEntitlementManagementAssignmentRequestAccessPackageAssignmentOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
	Model        *beta.AccessPackageAssignment
}

type GetEntitlementManagementAssignmentRequestAccessPackageAssignmentOperationOptions struct {
	Expand    *odata.Expand
	Metadata  *odata.Metadata
	RetryFunc client.RequestRetryFunc
	Select    *[]string
}

func DefaultGetEntitlementManagementAssignmentRequestAccessPackageAssignmentOperationOptions() GetEntitlementManagementAssignmentRequestAccessPackageAssignmentOperationOptions {
	return GetEntitlementManagementAssignmentRequestAccessPackageAssignmentOperationOptions{}
}

func (o GetEntitlementManagementAssignmentRequestAccessPackageAssignmentOperationOptions) ToHeaders() *client.Headers {
	out := client.Headers{}

	return &out
}

func (o GetEntitlementManagementAssignmentRequestAccessPackageAssignmentOperationOptions) ToOData() *odata.Query {
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

func (o GetEntitlementManagementAssignmentRequestAccessPackageAssignmentOperationOptions) ToQuery() *client.QueryParams {
	out := client.QueryParams{}

	return &out
}

// GetEntitlementManagementAssignmentRequestAccessPackageAssignment - Get accessPackageAssignment from
// identityGovernance. For a requestType of UserAdd or AdminAdd, an access package assignment requested to be created.
// For a requestType of UserRemove, AdminRemove, or SystemRemove, this property has the id property of an existing
// assignment to be removed. Supports $expand.
func (c EntitlementManagementAssignmentRequestAccessPackageAssignmentClient) GetEntitlementManagementAssignmentRequestAccessPackageAssignment(ctx context.Context, id beta.IdentityGovernanceEntitlementManagementAssignmentRequestId, options GetEntitlementManagementAssignmentRequestAccessPackageAssignmentOperationOptions) (result GetEntitlementManagementAssignmentRequestAccessPackageAssignmentOperationResponse, err error) {
	opts := client.RequestOptions{
		ContentType: "application/json; charset=utf-8",
		ExpectedStatusCodes: []int{
			http.StatusOK,
		},
		HttpMethod:    http.MethodGet,
		OptionsObject: options,
		Path:          fmt.Sprintf("%s/accessPackageAssignment", id.ID()),
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

	var model beta.AccessPackageAssignment
	result.Model = &model
	if err = resp.Unmarshal(result.Model); err != nil {
		return
	}

	return
}
