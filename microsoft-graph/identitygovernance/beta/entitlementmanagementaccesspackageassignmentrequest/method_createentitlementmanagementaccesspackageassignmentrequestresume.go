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

type CreateEntitlementManagementAccessPackageAssignmentRequestResumeOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
}

type CreateEntitlementManagementAccessPackageAssignmentRequestResumeOperationOptions struct {
	Metadata *odata.Metadata
}

func DefaultCreateEntitlementManagementAccessPackageAssignmentRequestResumeOperationOptions() CreateEntitlementManagementAccessPackageAssignmentRequestResumeOperationOptions {
	return CreateEntitlementManagementAccessPackageAssignmentRequestResumeOperationOptions{}
}

func (o CreateEntitlementManagementAccessPackageAssignmentRequestResumeOperationOptions) ToHeaders() *client.Headers {
	out := client.Headers{}

	return &out
}

func (o CreateEntitlementManagementAccessPackageAssignmentRequestResumeOperationOptions) ToOData() *odata.Query {
	out := odata.Query{}
	if o.Metadata != nil {
		out.Metadata = *o.Metadata
	}
	return &out
}

func (o CreateEntitlementManagementAccessPackageAssignmentRequestResumeOperationOptions) ToQuery() *client.QueryParams {
	out := client.QueryParams{}

	return &out
}

// CreateEntitlementManagementAccessPackageAssignmentRequestResume - Invoke action resume. In Microsoft Entra
// entitlement management, when an access package policy has been enabled to call out a custom extension and the request
// processing is waiting for the callback from the customer, the customer can initiate a resume action. It is performed
// on an accessPackageAssignmentRequest object whose requestStatus is in a WaitingForCallback state.
func (c EntitlementManagementAccessPackageAssignmentRequestClient) CreateEntitlementManagementAccessPackageAssignmentRequestResume(ctx context.Context, id beta.IdentityGovernanceEntitlementManagementAccessPackageAssignmentRequestId, input CreateEntitlementManagementAccessPackageAssignmentRequestResumeRequest, options CreateEntitlementManagementAccessPackageAssignmentRequestResumeOperationOptions) (result CreateEntitlementManagementAccessPackageAssignmentRequestResumeOperationResponse, err error) {
	opts := client.RequestOptions{
		ContentType: "application/json; charset=utf-8",
		ExpectedStatusCodes: []int{
			http.StatusNoContent,
		},
		HttpMethod:    http.MethodPost,
		OptionsObject: options,
		Path:          fmt.Sprintf("%s/resume", id.ID()),
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

	return
}
