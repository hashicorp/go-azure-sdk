package privilegedaccessgroupeligibilityschedulerequest

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

type CancelPrivilegedAccessGroupEligibilityScheduleRequestOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
}

type CancelPrivilegedAccessGroupEligibilityScheduleRequestOperationOptions struct {
	Metadata *odata.Metadata
}

func DefaultCancelPrivilegedAccessGroupEligibilityScheduleRequestOperationOptions() CancelPrivilegedAccessGroupEligibilityScheduleRequestOperationOptions {
	return CancelPrivilegedAccessGroupEligibilityScheduleRequestOperationOptions{}
}

func (o CancelPrivilegedAccessGroupEligibilityScheduleRequestOperationOptions) ToHeaders() *client.Headers {
	out := client.Headers{}

	return &out
}

func (o CancelPrivilegedAccessGroupEligibilityScheduleRequestOperationOptions) ToOData() *odata.Query {
	out := odata.Query{}
	if o.Metadata != nil {
		out.Metadata = *o.Metadata
	}
	return &out
}

func (o CancelPrivilegedAccessGroupEligibilityScheduleRequestOperationOptions) ToQuery() *client.QueryParams {
	out := client.QueryParams{}

	return &out
}

// CancelPrivilegedAccessGroupEligibilityScheduleRequest - Invoke action cancel. Cancel an eligibility assignment
// request to a group whose membership and ownership are governed by PIM.
func (c PrivilegedAccessGroupEligibilityScheduleRequestClient) CancelPrivilegedAccessGroupEligibilityScheduleRequest(ctx context.Context, id beta.IdentityGovernancePrivilegedAccessGroupEligibilityScheduleRequestId, options CancelPrivilegedAccessGroupEligibilityScheduleRequestOperationOptions) (result CancelPrivilegedAccessGroupEligibilityScheduleRequestOperationResponse, err error) {
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
