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

type DeleteEntitlementManagementAccessPackageAssignmentRequestOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
}

type DeleteEntitlementManagementAccessPackageAssignmentRequestOperationOptions struct {
	IfMatch *string
}

func DefaultDeleteEntitlementManagementAccessPackageAssignmentRequestOperationOptions() DeleteEntitlementManagementAccessPackageAssignmentRequestOperationOptions {
	return DeleteEntitlementManagementAccessPackageAssignmentRequestOperationOptions{}
}

func (o DeleteEntitlementManagementAccessPackageAssignmentRequestOperationOptions) ToHeaders() *client.Headers {
	out := client.Headers{}
	if o.IfMatch != nil {
		out.Append("If-Match", fmt.Sprintf("%v", *o.IfMatch))
	}
	return &out
}

func (o DeleteEntitlementManagementAccessPackageAssignmentRequestOperationOptions) ToOData() *odata.Query {
	out := odata.Query{}

	return &out
}

func (o DeleteEntitlementManagementAccessPackageAssignmentRequestOperationOptions) ToQuery() *client.QueryParams {
	out := client.QueryParams{}

	return &out
}

// DeleteEntitlementManagementAccessPackageAssignmentRequest - Delete accessPackageAssignmentRequest. Delete an
// accessPackageAssignmentRequest object. This request can be made to remove a denied or completed request. You can't
// delete an access package assignment request if it has any accessPackageAssignment objects.
func (c EntitlementManagementAccessPackageAssignmentRequestClient) DeleteEntitlementManagementAccessPackageAssignmentRequest(ctx context.Context, id beta.IdentityGovernanceEntitlementManagementAccessPackageAssignmentRequestId, options DeleteEntitlementManagementAccessPackageAssignmentRequestOperationOptions) (result DeleteEntitlementManagementAccessPackageAssignmentRequestOperationResponse, err error) {
	opts := client.RequestOptions{
		ContentType: "application/json; charset=utf-8",
		ExpectedStatusCodes: []int{
			http.StatusNoContent,
		},
		HttpMethod:    http.MethodDelete,
		OptionsObject: options,
		Path:          id.ID(),
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
