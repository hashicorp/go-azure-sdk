package appconsentrequestsforapprovaluserapproval

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

type UpdateAppConsentRequestsForApprovalUserApprovalOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
}

// UpdateAppConsentRequestsForApprovalUserApproval - Update the navigation property approval in me
func (c AppConsentRequestsForApprovalUserApprovalClient) UpdateAppConsentRequestsForApprovalUserApproval(ctx context.Context, id beta.MeAppConsentRequestsForApprovalIdUserConsentRequestId, input beta.Approval) (result UpdateAppConsentRequestsForApprovalUserApprovalOperationResponse, err error) {
	opts := client.RequestOptions{
		ContentType: "application/json; charset=utf-8",
		ExpectedStatusCodes: []int{
			http.StatusNoContent,
		},
		HttpMethod: http.MethodPatch,
		Path:       fmt.Sprintf("%s/approval", id.ID()),
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
