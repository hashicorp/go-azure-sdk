package entitlementmanagementroleeligibilityschedulerequest

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

type CancelEntitlementManagementRoleEligibilityScheduleRequestOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
}

type CancelEntitlementManagementRoleEligibilityScheduleRequestOperationOptions struct {
	Metadata *odata.Metadata
}

func DefaultCancelEntitlementManagementRoleEligibilityScheduleRequestOperationOptions() CancelEntitlementManagementRoleEligibilityScheduleRequestOperationOptions {
	return CancelEntitlementManagementRoleEligibilityScheduleRequestOperationOptions{}
}

func (o CancelEntitlementManagementRoleEligibilityScheduleRequestOperationOptions) ToHeaders() *client.Headers {
	out := client.Headers{}

	return &out
}

func (o CancelEntitlementManagementRoleEligibilityScheduleRequestOperationOptions) ToOData() *odata.Query {
	out := odata.Query{}
	if o.Metadata != nil {
		out.Metadata = *o.Metadata
	}
	return &out
}

func (o CancelEntitlementManagementRoleEligibilityScheduleRequestOperationOptions) ToQuery() *client.QueryParams {
	out := client.QueryParams{}

	return &out
}

// CancelEntitlementManagementRoleEligibilityScheduleRequest - Invoke action cancel. Immediately cancel a
// unifiedRoleEligibilityScheduleRequest object whose status is Granted and have the system automatically delete the
// cancelled request after 30 days. After calling this action, the status of the cancelled
// unifiedRoleEligibilityScheduleRequest changes to Revoked.
func (c EntitlementManagementRoleEligibilityScheduleRequestClient) CancelEntitlementManagementRoleEligibilityScheduleRequest(ctx context.Context, id stable.RoleManagementEntitlementManagementRoleEligibilityScheduleRequestId, options CancelEntitlementManagementRoleEligibilityScheduleRequestOperationOptions) (result CancelEntitlementManagementRoleEligibilityScheduleRequestOperationResponse, err error) {
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
