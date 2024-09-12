package lifecycleworkflowdeleteditemworkflow

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

type CreateLifecycleWorkflowDeletedItemWorkflowIdentityGovernanceRestoreOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
	Model        *stable.IdentityGovernanceWorkflow
}

// CreateLifecycleWorkflowDeletedItemWorkflowIdentityGovernanceRestore - Invoke action restore. Restore a workflow that
// has been deleted. You can only restore a workflow that was deleted within the last 30 days before Microsoft Entra ID
// automatically permanently deletes it.
func (c LifecycleWorkflowDeletedItemWorkflowClient) CreateLifecycleWorkflowDeletedItemWorkflowIdentityGovernanceRestore(ctx context.Context, id stable.IdentityGovernanceLifecycleWorkflowDeletedItemWorkflowId) (result CreateLifecycleWorkflowDeletedItemWorkflowIdentityGovernanceRestoreOperationResponse, err error) {
	opts := client.RequestOptions{
		ContentType: "application/json; charset=utf-8",
		ExpectedStatusCodes: []int{
			http.StatusOK,
		},
		HttpMethod: http.MethodPost,
		Path:       fmt.Sprintf("%s/identityGovernance.restore", id.ID()),
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

	var model stable.IdentityGovernanceWorkflow
	result.Model = &model
	if err = resp.Unmarshal(result.Model); err != nil {
		return
	}

	return
}
