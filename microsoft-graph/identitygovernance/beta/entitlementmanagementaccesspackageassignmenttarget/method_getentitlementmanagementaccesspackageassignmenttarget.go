package entitlementmanagementaccesspackageassignmenttarget

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

type GetEntitlementManagementAccessPackageAssignmentTargetOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
	Model        *beta.AccessPackageSubject
}

type GetEntitlementManagementAccessPackageAssignmentTargetOperationOptions struct {
	Expand    *odata.Expand
	Metadata  *odata.Metadata
	RetryFunc client.RequestRetryFunc
	Select    *[]string
}

func DefaultGetEntitlementManagementAccessPackageAssignmentTargetOperationOptions() GetEntitlementManagementAccessPackageAssignmentTargetOperationOptions {
	return GetEntitlementManagementAccessPackageAssignmentTargetOperationOptions{}
}

func (o GetEntitlementManagementAccessPackageAssignmentTargetOperationOptions) ToHeaders() *client.Headers {
	out := client.Headers{}

	return &out
}

func (o GetEntitlementManagementAccessPackageAssignmentTargetOperationOptions) ToOData() *odata.Query {
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

func (o GetEntitlementManagementAccessPackageAssignmentTargetOperationOptions) ToQuery() *client.QueryParams {
	out := client.QueryParams{}

	return &out
}

// GetEntitlementManagementAccessPackageAssignmentTarget - Get target from identityGovernance. The subject of the access
// package assignment. Read-only. Nullable. Supports $expand. Supports $filter (eq) on objectId.
func (c EntitlementManagementAccessPackageAssignmentTargetClient) GetEntitlementManagementAccessPackageAssignmentTarget(ctx context.Context, id beta.IdentityGovernanceEntitlementManagementAccessPackageAssignmentId, options GetEntitlementManagementAccessPackageAssignmentTargetOperationOptions) (result GetEntitlementManagementAccessPackageAssignmentTargetOperationResponse, err error) {
	opts := client.RequestOptions{
		ContentType: "application/json; charset=utf-8",
		ExpectedStatusCodes: []int{
			http.StatusOK,
		},
		HttpMethod:    http.MethodGet,
		OptionsObject: options,
		Path:          fmt.Sprintf("%s/target", id.ID()),
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

	var model beta.AccessPackageSubject
	result.Model = &model
	if err = resp.Unmarshal(result.Model); err != nil {
		return
	}

	return
}
