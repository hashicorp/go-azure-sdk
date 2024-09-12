package lifecycleworkflowruntaskprocessingresult

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

type CreateLifecycleWorkflowRunTaskProcessingResultIdentityGovernanceResumeOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
}

// CreateLifecycleWorkflowRunTaskProcessingResultIdentityGovernanceResume - Invoke action resume. Resume a task
// processing result that's inProgress. In the default case an Azure Logic Apps system-assigned managed identity calls
// this API. For more information, see: Lifecycle Workflows extensibility approach.
func (c LifecycleWorkflowRunTaskProcessingResultClient) CreateLifecycleWorkflowRunTaskProcessingResultIdentityGovernanceResume(ctx context.Context, id stable.IdentityGovernanceLifecycleWorkflowWorkflowIdRunIdTaskProcessingResultId, input CreateLifecycleWorkflowRunTaskProcessingResultIdentityGovernanceResumeRequest) (result CreateLifecycleWorkflowRunTaskProcessingResultIdentityGovernanceResumeOperationResponse, err error) {
	opts := client.RequestOptions{
		ContentType: "application/json; charset=utf-8",
		ExpectedStatusCodes: []int{
			http.StatusNoContent,
		},
		HttpMethod: http.MethodPost,
		Path:       fmt.Sprintf("%s/identityGovernance.resume", id.ID()),
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
