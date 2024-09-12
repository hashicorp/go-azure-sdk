package privilegedaccessgroupassignmentapprovalstage

import (
	"context"
	"net/http"

	"github.com/hashicorp/go-azure-sdk/microsoft-graph/common-types/stable"
	"github.com/hashicorp/go-azure-sdk/sdk/client"
	"github.com/hashicorp/go-azure-sdk/sdk/odata"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type UpdatePrivilegedAccessGroupAssignmentApprovalStageOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
}

// UpdatePrivilegedAccessGroupAssignmentApprovalStage - Update the navigation property stages in identityGovernance
func (c PrivilegedAccessGroupAssignmentApprovalStageClient) UpdatePrivilegedAccessGroupAssignmentApprovalStage(ctx context.Context, id stable.IdentityGovernancePrivilegedAccessGroupAssignmentApprovalIdStageId, input stable.ApprovalStage) (result UpdatePrivilegedAccessGroupAssignmentApprovalStageOperationResponse, err error) {
	opts := client.RequestOptions{
		ContentType: "application/json; charset=utf-8",
		ExpectedStatusCodes: []int{
			http.StatusNoContent,
		},
		HttpMethod: http.MethodPatch,
		Path:       id.ID(),
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
