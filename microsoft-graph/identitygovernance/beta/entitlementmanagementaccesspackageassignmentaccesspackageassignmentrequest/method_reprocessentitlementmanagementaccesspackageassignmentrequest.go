package entitlementmanagementaccesspackageassignmentaccesspackageassignmentrequest

import (
	"context"
	"fmt"
	"net/http"

	"github.com/hashicorp/go-azure-sdk/microsoft-graph/common-types/beta"
	"github.com/hashicorp/go-azure-sdk/sdk/client"
	"github.com/hashicorp/go-azure-sdk/sdk/odata"
)

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type ReprocessEntitlementManagementAccessPackageAssignmentRequestOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
}

type ReprocessEntitlementManagementAccessPackageAssignmentRequestOperationOptions struct {
	Metadata  *odata.Metadata
	RetryFunc client.RequestRetryFunc
}

func DefaultReprocessEntitlementManagementAccessPackageAssignmentRequestOperationOptions() ReprocessEntitlementManagementAccessPackageAssignmentRequestOperationOptions {
	return ReprocessEntitlementManagementAccessPackageAssignmentRequestOperationOptions{}
}

func (o ReprocessEntitlementManagementAccessPackageAssignmentRequestOperationOptions) ToHeaders() *client.Headers {
	out := client.Headers{}

	return &out
}

func (o ReprocessEntitlementManagementAccessPackageAssignmentRequestOperationOptions) ToOData() *odata.Query {
	out := odata.Query{}
	if o.Metadata != nil {
		out.Metadata = *o.Metadata
	}
	return &out
}

func (o ReprocessEntitlementManagementAccessPackageAssignmentRequestOperationOptions) ToQuery() *client.QueryParams {
	out := client.QueryParams{}

	return &out
}

// ReprocessEntitlementManagementAccessPackageAssignmentRequest - Invoke action reprocess
func (c EntitlementManagementAccessPackageAssignmentAccessPackageAssignmentRequestClient) ReprocessEntitlementManagementAccessPackageAssignmentRequest(ctx context.Context, id beta.IdentityGovernanceEntitlementManagementAccessPackageAssignmentIdAccessPackageAssignmentRequestId, options ReprocessEntitlementManagementAccessPackageAssignmentRequestOperationOptions) (result ReprocessEntitlementManagementAccessPackageAssignmentRequestOperationResponse, err error) {
	opts := client.RequestOptions{
		ContentType: "application/json; charset=utf-8",
		ExpectedStatusCodes: []int{
			http.StatusAccepted,
			http.StatusCreated,
			http.StatusNoContent,
			http.StatusOK,
		},
		HttpMethod:    http.MethodPost,
		OptionsObject: options,
		Path:          fmt.Sprintf("%s/reprocess", id.ID()),
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
