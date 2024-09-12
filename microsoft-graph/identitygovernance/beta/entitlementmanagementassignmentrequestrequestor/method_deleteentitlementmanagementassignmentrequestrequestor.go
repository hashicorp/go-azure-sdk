package entitlementmanagementassignmentrequestrequestor

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

type DeleteEntitlementManagementAssignmentRequestRequestorOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
}

type DeleteEntitlementManagementAssignmentRequestRequestorOperationOptions struct {
	IfMatch *string
}

func DefaultDeleteEntitlementManagementAssignmentRequestRequestorOperationOptions() DeleteEntitlementManagementAssignmentRequestRequestorOperationOptions {
	return DeleteEntitlementManagementAssignmentRequestRequestorOperationOptions{}
}

func (o DeleteEntitlementManagementAssignmentRequestRequestorOperationOptions) ToHeaders() *client.Headers {
	out := client.Headers{}
	if o.IfMatch != nil {
		out.Append("If-Match", fmt.Sprintf("%v", *o.IfMatch))
	}
	return &out
}

func (o DeleteEntitlementManagementAssignmentRequestRequestorOperationOptions) ToOData() *odata.Query {
	out := odata.Query{}

	return &out
}

func (o DeleteEntitlementManagementAssignmentRequestRequestorOperationOptions) ToQuery() *client.QueryParams {
	out := client.QueryParams{}

	return &out
}

// DeleteEntitlementManagementAssignmentRequestRequestor - Delete navigation property requestor for identityGovernance
func (c EntitlementManagementAssignmentRequestRequestorClient) DeleteEntitlementManagementAssignmentRequestRequestor(ctx context.Context, id beta.IdentityGovernanceEntitlementManagementAssignmentRequestId, options DeleteEntitlementManagementAssignmentRequestRequestorOperationOptions) (result DeleteEntitlementManagementAssignmentRequestRequestorOperationResponse, err error) {
	opts := client.RequestOptions{
		ContentType: "application/json; charset=utf-8",
		ExpectedStatusCodes: []int{
			http.StatusNoContent,
		},
		HttpMethod:    http.MethodDelete,
		OptionsObject: options,
		Path:          fmt.Sprintf("%s/requestor", id.ID()),
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
