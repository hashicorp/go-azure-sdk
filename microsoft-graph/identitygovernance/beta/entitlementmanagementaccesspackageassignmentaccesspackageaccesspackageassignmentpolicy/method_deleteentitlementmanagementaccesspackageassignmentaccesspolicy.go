package entitlementmanagementaccesspackageassignmentaccesspackageaccesspackageassignmentpolicy

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

type DeleteEntitlementManagementAccessPackageAssignmentAccessPolicyOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
}

type DeleteEntitlementManagementAccessPackageAssignmentAccessPolicyOperationOptions struct {
	IfMatch   *string
	Metadata  *odata.Metadata
	RetryFunc client.RequestRetryFunc
}

func DefaultDeleteEntitlementManagementAccessPackageAssignmentAccessPolicyOperationOptions() DeleteEntitlementManagementAccessPackageAssignmentAccessPolicyOperationOptions {
	return DeleteEntitlementManagementAccessPackageAssignmentAccessPolicyOperationOptions{}
}

func (o DeleteEntitlementManagementAccessPackageAssignmentAccessPolicyOperationOptions) ToHeaders() *client.Headers {
	out := client.Headers{}
	if o.IfMatch != nil {
		out.Append("If-Match", fmt.Sprintf("%v", *o.IfMatch))
	}
	return &out
}

func (o DeleteEntitlementManagementAccessPackageAssignmentAccessPolicyOperationOptions) ToOData() *odata.Query {
	out := odata.Query{}
	if o.Metadata != nil {
		out.Metadata = *o.Metadata
	}
	return &out
}

func (o DeleteEntitlementManagementAccessPackageAssignmentAccessPolicyOperationOptions) ToQuery() *client.QueryParams {
	out := client.QueryParams{}

	return &out
}

// DeleteEntitlementManagementAccessPackageAssignmentAccessPolicy - Delete navigation property
// accessPackageAssignmentPolicies for identityGovernance
func (c EntitlementManagementAccessPackageAssignmentAccessPackageAccessPackageAssignmentPolicyClient) DeleteEntitlementManagementAccessPackageAssignmentAccessPolicy(ctx context.Context, id beta.IdentityGovernanceEntitlementManagementAccessPackageAssignmentIdAccessPackageAccessPackageAssignmentPolicyId, options DeleteEntitlementManagementAccessPackageAssignmentAccessPolicyOperationOptions) (result DeleteEntitlementManagementAccessPackageAssignmentAccessPolicyOperationResponse, err error) {
	opts := client.RequestOptions{
		ContentType: "application/json; charset=utf-8",
		ExpectedStatusCodes: []int{
			http.StatusNoContent,
		},
		HttpMethod:    http.MethodDelete,
		OptionsObject: options,
		Path:          id.ID(),
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

	return
}
