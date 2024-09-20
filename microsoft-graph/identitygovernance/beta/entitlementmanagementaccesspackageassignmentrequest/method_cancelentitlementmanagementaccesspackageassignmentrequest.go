package entitlementmanagementaccesspackageassignmentrequest

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

type CancelEntitlementManagementAccessPackageAssignmentRequestOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
}

type CancelEntitlementManagementAccessPackageAssignmentRequestOperationOptions struct {
	Metadata *odata.Metadata
}

func DefaultCancelEntitlementManagementAccessPackageAssignmentRequestOperationOptions() CancelEntitlementManagementAccessPackageAssignmentRequestOperationOptions {
	return CancelEntitlementManagementAccessPackageAssignmentRequestOperationOptions{}
}

func (o CancelEntitlementManagementAccessPackageAssignmentRequestOperationOptions) ToHeaders() *client.Headers {
	out := client.Headers{}

	return &out
}

func (o CancelEntitlementManagementAccessPackageAssignmentRequestOperationOptions) ToOData() *odata.Query {
	out := odata.Query{}
	if o.Metadata != nil {
		out.Metadata = *o.Metadata
	}
	return &out
}

func (o CancelEntitlementManagementAccessPackageAssignmentRequestOperationOptions) ToQuery() *client.QueryParams {
	out := client.QueryParams{}

	return &out
}

// CancelEntitlementManagementAccessPackageAssignmentRequest - Invoke action cancel. In Microsoft Entra Entitlement
// Management, cancel accessPackageAssignmentRequest objects that are in a cancelable state: accepted, pendingApproval,
// pendingNotBefore, pendingApprovalEscalated.
func (c EntitlementManagementAccessPackageAssignmentRequestClient) CancelEntitlementManagementAccessPackageAssignmentRequest(ctx context.Context, id beta.IdentityGovernanceEntitlementManagementAccessPackageAssignmentRequestId, options CancelEntitlementManagementAccessPackageAssignmentRequestOperationOptions) (result CancelEntitlementManagementAccessPackageAssignmentRequestOperationResponse, err error) {
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
