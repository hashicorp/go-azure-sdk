package entitlementmanagementassignmentrequest

import (
	"context"
	"net/http"

	"github.com/hashicorp/go-azure-sdk/microsoft-graph/common-types/stable"
	"github.com/hashicorp/go-azure-sdk/sdk/client"
	"github.com/hashicorp/go-azure-sdk/sdk/odata"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type CreateEntitlementManagementAssignmentRequestOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
	Model        *stable.AccessPackageAssignmentRequest
}

type CreateEntitlementManagementAssignmentRequestOperationOptions struct {
	Metadata  *odata.Metadata
	RetryFunc client.RequestRetryFunc
}

func DefaultCreateEntitlementManagementAssignmentRequestOperationOptions() CreateEntitlementManagementAssignmentRequestOperationOptions {
	return CreateEntitlementManagementAssignmentRequestOperationOptions{}
}

func (o CreateEntitlementManagementAssignmentRequestOperationOptions) ToHeaders() *client.Headers {
	out := client.Headers{}

	return &out
}

func (o CreateEntitlementManagementAssignmentRequestOperationOptions) ToOData() *odata.Query {
	out := odata.Query{}
	if o.Metadata != nil {
		out.Metadata = *o.Metadata
	}
	return &out
}

func (o CreateEntitlementManagementAssignmentRequestOperationOptions) ToQuery() *client.QueryParams {
	out := client.QueryParams{}

	return &out
}

// CreateEntitlementManagementAssignmentRequest - Create accessPackageAssignmentRequest. In Microsoft Entra Entitlement
// Management, create a new accessPackageAssignmentRequest object. This operation is used to assign a user to an access
// package, update the assignment, or to remove an access package assignment.
func (c EntitlementManagementAssignmentRequestClient) CreateEntitlementManagementAssignmentRequest(ctx context.Context, input stable.AccessPackageAssignmentRequest, options CreateEntitlementManagementAssignmentRequestOperationOptions) (result CreateEntitlementManagementAssignmentRequestOperationResponse, err error) {
	opts := client.RequestOptions{
		ContentType: "application/json; charset=utf-8",
		ExpectedStatusCodes: []int{
			http.StatusCreated,
		},
		HttpMethod:    http.MethodPost,
		OptionsObject: options,
		Path:          "/identityGovernance/entitlementManagement/assignmentRequests",
		RetryFunc:     options.RetryFunc,
	}

	req, err := c.Client.NewRequest(ctx, opts)
	if err != nil {
		return
	}

	if err = req.Marshal(input); err != nil {
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

	var model stable.AccessPackageAssignmentRequest
	result.Model = &model
	if err = resp.Unmarshal(result.Model); err != nil {
		return
	}

	return
}
