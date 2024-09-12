package privilegedaccessgroupassignmentschedulerequest

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

type CancelPrivilegedAccessGroupAssignmentScheduleRequestOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
}

// CancelPrivilegedAccessGroupAssignmentScheduleRequest - Invoke action cancel. Cancel an access assignment request to a
// group whose membership and ownership are governed by PIM.
func (c PrivilegedAccessGroupAssignmentScheduleRequestClient) CancelPrivilegedAccessGroupAssignmentScheduleRequest(ctx context.Context, id beta.IdentityGovernancePrivilegedAccessGroupAssignmentScheduleRequestId) (result CancelPrivilegedAccessGroupAssignmentScheduleRequestOperationResponse, err error) {
	opts := client.RequestOptions{
		ContentType: "application/json; charset=utf-8",
		ExpectedStatusCodes: []int{
			http.StatusNoContent,
		},
		HttpMethod: http.MethodPost,
		Path:       fmt.Sprintf("%s/cancel", id.ID()),
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
