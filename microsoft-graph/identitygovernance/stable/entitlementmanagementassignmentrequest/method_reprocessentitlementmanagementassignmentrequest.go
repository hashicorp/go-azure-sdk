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

type ReprocessEntitlementManagementAssignmentRequestOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
}

type ReprocessEntitlementManagementAssignmentRequestOperationOptions struct {
	Metadata *odata.Metadata
}

func DefaultReprocessEntitlementManagementAssignmentRequestOperationOptions() ReprocessEntitlementManagementAssignmentRequestOperationOptions {
	return ReprocessEntitlementManagementAssignmentRequestOperationOptions{}
}

func (o ReprocessEntitlementManagementAssignmentRequestOperationOptions) ToHeaders() *client.Headers {
	out := client.Headers{}

	return &out
}

func (o ReprocessEntitlementManagementAssignmentRequestOperationOptions) ToOData() *odata.Query {
	out := odata.Query{}
	if o.Metadata != nil {
		out.Metadata = *o.Metadata
	}
	return &out
}

func (o ReprocessEntitlementManagementAssignmentRequestOperationOptions) ToQuery() *client.QueryParams {
	out := client.QueryParams{}

	return &out
}

// ReprocessEntitlementManagementAssignmentRequest - Invoke action reprocess. In Microsoft Entra entitlement management,
// callers can automatically retry a user's request for access to an access package. It's performed on an
// accessPackageAssignmentRequest object whose requestState is in a DeliveryFailed or PartiallyDelivered state. You can
// only reprocess a request within 14 days from the time the original request was completed. For requests completed more
// than 14 days, you will need to ask the users to cancel the request(s) and make a new request in the MyAccess portal.
func (c EntitlementManagementAssignmentRequestClient) ReprocessEntitlementManagementAssignmentRequest(ctx context.Context, id stable.IdentityGovernanceEntitlementManagementAssignmentRequestId, options ReprocessEntitlementManagementAssignmentRequestOperationOptions) (result ReprocessEntitlementManagementAssignmentRequestOperationResponse, err error) {
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
