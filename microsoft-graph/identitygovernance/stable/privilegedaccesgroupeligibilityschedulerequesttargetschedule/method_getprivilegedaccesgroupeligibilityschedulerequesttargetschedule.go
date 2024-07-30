package privilegedaccesgroupeligibilityschedulerequesttargetschedule

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

type GetPrivilegedAccesGroupEligibilityScheduleRequestTargetScheduleOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
	Model        *stable.PrivilegedAccessGroupEligibilitySchedule
}

// GetPrivilegedAccesGroupEligibilityScheduleRequestTargetSchedule ...
func (c PrivilegedAccesGroupEligibilityScheduleRequestTargetScheduleClient) GetPrivilegedAccesGroupEligibilityScheduleRequestTargetSchedule(ctx context.Context, id IdentityGovernancePrivilegedAccesGroupEligibilityScheduleRequestId) (result GetPrivilegedAccesGroupEligibilityScheduleRequestTargetScheduleOperationResponse, err error) {
	opts := client.RequestOptions{
		ContentType: "application/json; charset=utf-8",
		ExpectedStatusCodes: []int{
			http.StatusOK,
		},
		HttpMethod: http.MethodGet,
		Path:       fmt.Sprintf("%s/targetSchedule", id.ID()),
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

	var model stable.PrivilegedAccessGroupEligibilitySchedule
	result.Model = &model

	if err = resp.Unmarshal(result.Model); err != nil {
		return
	}

	return
}
