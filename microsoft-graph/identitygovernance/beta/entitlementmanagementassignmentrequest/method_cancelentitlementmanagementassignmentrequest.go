package entitlementmanagementassignmentrequest

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

type CancelEntitlementManagementAssignmentRequestOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
}

type CancelEntitlementManagementAssignmentRequestOperationOptions struct {
	Metadata *odata.Metadata
}

func DefaultCancelEntitlementManagementAssignmentRequestOperationOptions() CancelEntitlementManagementAssignmentRequestOperationOptions {
	return CancelEntitlementManagementAssignmentRequestOperationOptions{}
}

func (o CancelEntitlementManagementAssignmentRequestOperationOptions) ToHeaders() *client.Headers {
	out := client.Headers{}

	return &out
}

func (o CancelEntitlementManagementAssignmentRequestOperationOptions) ToOData() *odata.Query {
	out := odata.Query{}
	if o.Metadata != nil {
		out.Metadata = *o.Metadata
	}
	return &out
}

func (o CancelEntitlementManagementAssignmentRequestOperationOptions) ToQuery() *client.QueryParams {
	out := client.QueryParams{}

	return &out
}

// CancelEntitlementManagementAssignmentRequest - Invoke action cancel. In Microsoft Entra Entitlement Management,
// cancel accessPackageAssignmentRequest objects that are in a cancelable state: accepted, pendingApproval,
// pendingNotBefore, pendingApprovalEscalated.
func (c EntitlementManagementAssignmentRequestClient) CancelEntitlementManagementAssignmentRequest(ctx context.Context, id beta.IdentityGovernanceEntitlementManagementAssignmentRequestId, options CancelEntitlementManagementAssignmentRequestOperationOptions) (result CancelEntitlementManagementAssignmentRequestOperationResponse, err error) {
	opts := client.RequestOptions{
		ContentType: "application/json; charset=utf-8",
		ExpectedStatusCodes: []int{
			http.StatusNoContent,
		},
		HttpMethod:    http.MethodPost,
		OptionsObject: options,
		Path:          fmt.Sprintf("%s/cancel", id.ID()),
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
