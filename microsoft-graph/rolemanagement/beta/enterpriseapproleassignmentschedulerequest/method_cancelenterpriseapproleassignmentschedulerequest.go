package enterpriseapproleassignmentschedulerequest

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

type CancelEnterpriseAppRoleAssignmentScheduleRequestOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
}

type CancelEnterpriseAppRoleAssignmentScheduleRequestOperationOptions struct {
	Metadata *odata.Metadata
}

func DefaultCancelEnterpriseAppRoleAssignmentScheduleRequestOperationOptions() CancelEnterpriseAppRoleAssignmentScheduleRequestOperationOptions {
	return CancelEnterpriseAppRoleAssignmentScheduleRequestOperationOptions{}
}

func (o CancelEnterpriseAppRoleAssignmentScheduleRequestOperationOptions) ToHeaders() *client.Headers {
	out := client.Headers{}

	return &out
}

func (o CancelEnterpriseAppRoleAssignmentScheduleRequestOperationOptions) ToOData() *odata.Query {
	out := odata.Query{}
	if o.Metadata != nil {
		out.Metadata = *o.Metadata
	}
	return &out
}

func (o CancelEnterpriseAppRoleAssignmentScheduleRequestOperationOptions) ToQuery() *client.QueryParams {
	out := client.QueryParams{}

	return &out
}

// CancelEnterpriseAppRoleAssignmentScheduleRequest - Invoke action cancel. Immediately cancel a
// unifiedRoleAssignmentScheduleRequest object that is in a Granted status, and have the system automatically delete the
// canceled request after 30 days. After calling this action, the status of the canceled
// unifiedRoleAssignmentScheduleRequest changes to Canceled.
func (c EnterpriseAppRoleAssignmentScheduleRequestClient) CancelEnterpriseAppRoleAssignmentScheduleRequest(ctx context.Context, id beta.RoleManagementEnterpriseAppIdRoleAssignmentScheduleRequestId, options CancelEnterpriseAppRoleAssignmentScheduleRequestOperationOptions) (result CancelEnterpriseAppRoleAssignmentScheduleRequestOperationResponse, err error) {
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
