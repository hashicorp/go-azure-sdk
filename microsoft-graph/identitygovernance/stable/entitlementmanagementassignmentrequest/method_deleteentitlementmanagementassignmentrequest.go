package entitlementmanagementassignmentrequest

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

type DeleteEntitlementManagementAssignmentRequestOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
}

type DeleteEntitlementManagementAssignmentRequestOperationOptions struct {
	IfMatch   *string
	Metadata  *odata.Metadata
	RetryFunc client.RequestRetryFunc
}

func DefaultDeleteEntitlementManagementAssignmentRequestOperationOptions() DeleteEntitlementManagementAssignmentRequestOperationOptions {
	return DeleteEntitlementManagementAssignmentRequestOperationOptions{}
}

func (o DeleteEntitlementManagementAssignmentRequestOperationOptions) ToHeaders() *client.Headers {
	out := client.Headers{}
	if o.IfMatch != nil {
		out.Append("If-Match", fmt.Sprintf("%v", *o.IfMatch))
	}
	return &out
}

func (o DeleteEntitlementManagementAssignmentRequestOperationOptions) ToOData() *odata.Query {
	out := odata.Query{}
	if o.Metadata != nil {
		out.Metadata = *o.Metadata
	}
	return &out
}

func (o DeleteEntitlementManagementAssignmentRequestOperationOptions) ToQuery() *client.QueryParams {
	out := client.QueryParams{}

	return &out
}

// DeleteEntitlementManagementAssignmentRequest - Delete accessPackageAssignmentRequest. Delete an
// accessPackageAssignmentRequest object. This request can be made to remove a denied or completed request. You cannot
// delete an access package assignment request if it has any accessPackageAssignment objects.
func (c EntitlementManagementAssignmentRequestClient) DeleteEntitlementManagementAssignmentRequest(ctx context.Context, id stable.IdentityGovernanceEntitlementManagementAssignmentRequestId, options DeleteEntitlementManagementAssignmentRequestOperationOptions) (result DeleteEntitlementManagementAssignmentRequestOperationResponse, err error) {
	opts := client.RequestOptions{
		ContentType: "application/json; charset=utf-8",
		ExpectedStatusCodes: []int{
			http.StatusNoContent,
			http.StatusOK,
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
