package enterpriseapproleeligibilityschedulerequest

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

type CancelEnterpriseAppRoleEligibilityScheduleRequestOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
}

type CancelEnterpriseAppRoleEligibilityScheduleRequestOperationOptions struct {
	Metadata  *odata.Metadata
	RetryFunc client.RequestRetryFunc
}

func DefaultCancelEnterpriseAppRoleEligibilityScheduleRequestOperationOptions() CancelEnterpriseAppRoleEligibilityScheduleRequestOperationOptions {
	return CancelEnterpriseAppRoleEligibilityScheduleRequestOperationOptions{}
}

func (o CancelEnterpriseAppRoleEligibilityScheduleRequestOperationOptions) ToHeaders() *client.Headers {
	out := client.Headers{}

	return &out
}

func (o CancelEnterpriseAppRoleEligibilityScheduleRequestOperationOptions) ToOData() *odata.Query {
	out := odata.Query{}
	if o.Metadata != nil {
		out.Metadata = *o.Metadata
	}
	return &out
}

func (o CancelEnterpriseAppRoleEligibilityScheduleRequestOperationOptions) ToQuery() *client.QueryParams {
	out := client.QueryParams{}

	return &out
}

// CancelEnterpriseAppRoleEligibilityScheduleRequest - Invoke action cancel. Immediately cancel a
// unifiedRoleEligibilityScheduleRequest that is in a Granted status, and have the system automatically delete the
// cancelled request after 30 days. After calling this action, the status of the cancelled
// unifiedRoleEligibilityScheduleRequest changes to Revoked.
func (c EnterpriseAppRoleEligibilityScheduleRequestClient) CancelEnterpriseAppRoleEligibilityScheduleRequest(ctx context.Context, id beta.RoleManagementEnterpriseAppIdRoleEligibilityScheduleRequestId, options CancelEnterpriseAppRoleEligibilityScheduleRequestOperationOptions) (result CancelEnterpriseAppRoleEligibilityScheduleRequestOperationResponse, err error) {
	opts := client.RequestOptions{
		ContentType: "application/json; charset=utf-8",
		ExpectedStatusCodes: []int{
			http.StatusNoContent,
		},
		HttpMethod:    http.MethodPost,
		OptionsObject: options,
		Path:          fmt.Sprintf("%s/cancel", id.ID()),
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
