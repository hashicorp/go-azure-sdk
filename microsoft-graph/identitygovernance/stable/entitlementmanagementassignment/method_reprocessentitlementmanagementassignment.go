package entitlementmanagementassignment

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

type ReprocessEntitlementManagementAssignmentOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
}

type ReprocessEntitlementManagementAssignmentOperationOptions struct {
	Metadata *odata.Metadata
}

func DefaultReprocessEntitlementManagementAssignmentOperationOptions() ReprocessEntitlementManagementAssignmentOperationOptions {
	return ReprocessEntitlementManagementAssignmentOperationOptions{}
}

func (o ReprocessEntitlementManagementAssignmentOperationOptions) ToHeaders() *client.Headers {
	out := client.Headers{}

	return &out
}

func (o ReprocessEntitlementManagementAssignmentOperationOptions) ToOData() *odata.Query {
	out := odata.Query{}
	if o.Metadata != nil {
		out.Metadata = *o.Metadata
	}
	return &out
}

func (o ReprocessEntitlementManagementAssignmentOperationOptions) ToQuery() *client.QueryParams {
	out := client.QueryParams{}

	return &out
}

// ReprocessEntitlementManagementAssignment - Invoke action reprocess. In Microsoft Entra entitlement management,
// callers can automatically reevaluate and enforce an accessPackageAssignment object of a user’s assignments for a
// specific access package. The state of the access package assignment must be Delivered for the administrator to
// reprocess the user's assignment. Only admins with the Access Package Assignment Manager role, or higher, in Microsoft
// Entra entitlement management can perform this action.
func (c EntitlementManagementAssignmentClient) ReprocessEntitlementManagementAssignment(ctx context.Context, id stable.IdentityGovernanceEntitlementManagementAssignmentId, options ReprocessEntitlementManagementAssignmentOperationOptions) (result ReprocessEntitlementManagementAssignmentOperationResponse, err error) {
	opts := client.RequestOptions{
		ContentType: "application/json; charset=utf-8",
		ExpectedStatusCodes: []int{
			http.StatusNoContent,
		},
		HttpMethod:    http.MethodPost,
		OptionsObject: options,
		Path:          fmt.Sprintf("%s/reprocess", id.ID()),
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
