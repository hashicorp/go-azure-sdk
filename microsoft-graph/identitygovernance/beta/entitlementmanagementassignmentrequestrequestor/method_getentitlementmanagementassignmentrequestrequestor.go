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

type GetEntitlementManagementAssignmentRequestRequestorOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
	Model        *beta.AccessPackageSubject
}

type GetEntitlementManagementAssignmentRequestRequestorOperationOptions struct {
	Expand *odata.Expand
	Select *[]string
}

func DefaultGetEntitlementManagementAssignmentRequestRequestorOperationOptions() GetEntitlementManagementAssignmentRequestRequestorOperationOptions {
	return GetEntitlementManagementAssignmentRequestRequestorOperationOptions{}
}

func (o GetEntitlementManagementAssignmentRequestRequestorOperationOptions) ToHeaders() *client.Headers {
	out := client.Headers{}

	return &out
}

func (o GetEntitlementManagementAssignmentRequestRequestorOperationOptions) ToOData() *odata.Query {
	out := odata.Query{}
	if o.Expand != nil {
		out.Expand = *o.Expand
	}
	if o.Select != nil {
		out.Select = *o.Select
	}
	return &out
}

func (o GetEntitlementManagementAssignmentRequestRequestorOperationOptions) ToQuery() *client.QueryParams {
	out := client.QueryParams{}

	return &out
}

// GetEntitlementManagementAssignmentRequestRequestor - Get requestor from identityGovernance. The subject who requested
// or, if a direct assignment, was assigned. Read-only. Nullable. Supports $expand.
func (c EntitlementManagementAssignmentRequestRequestorClient) GetEntitlementManagementAssignmentRequestRequestor(ctx context.Context, id beta.IdentityGovernanceEntitlementManagementAssignmentRequestId, options GetEntitlementManagementAssignmentRequestRequestorOperationOptions) (result GetEntitlementManagementAssignmentRequestRequestorOperationResponse, err error) {
	opts := client.RequestOptions{
		ContentType: "application/json; charset=utf-8",
		ExpectedStatusCodes: []int{
			http.StatusOK,
		},
		HttpMethod:    http.MethodGet,
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

	var model beta.AccessPackageSubject
	result.Model = &model
	if err = resp.Unmarshal(result.Model); err != nil {
		return
	}

	return
}
