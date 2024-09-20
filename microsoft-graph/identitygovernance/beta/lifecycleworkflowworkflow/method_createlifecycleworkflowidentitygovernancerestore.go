package lifecycleworkflowworkflow

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

type CreateLifecycleWorkflowIdentityGovernanceRestoreOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
	Model        *beta.IdentityGovernanceWorkflow
}

type CreateLifecycleWorkflowIdentityGovernanceRestoreOperationOptions struct {
	Metadata *odata.Metadata
}

func DefaultCreateLifecycleWorkflowIdentityGovernanceRestoreOperationOptions() CreateLifecycleWorkflowIdentityGovernanceRestoreOperationOptions {
	return CreateLifecycleWorkflowIdentityGovernanceRestoreOperationOptions{}
}

func (o CreateLifecycleWorkflowIdentityGovernanceRestoreOperationOptions) ToHeaders() *client.Headers {
	out := client.Headers{}

	return &out
}

func (o CreateLifecycleWorkflowIdentityGovernanceRestoreOperationOptions) ToOData() *odata.Query {
	out := odata.Query{}
	if o.Metadata != nil {
		out.Metadata = *o.Metadata
	}
	return &out
}

func (o CreateLifecycleWorkflowIdentityGovernanceRestoreOperationOptions) ToQuery() *client.QueryParams {
	out := client.QueryParams{}

	return &out
}

// CreateLifecycleWorkflowIdentityGovernanceRestore - Invoke action restore. Restore a workflow that has been deleted.
// You can only restore a workflow that was deleted within the last 30 days before Microsoft Entra ID automatically
// permanently deletes it.
func (c LifecycleWorkflowWorkflowClient) CreateLifecycleWorkflowIdentityGovernanceRestore(ctx context.Context, id beta.IdentityGovernanceLifecycleWorkflowWorkflowId, options CreateLifecycleWorkflowIdentityGovernanceRestoreOperationOptions) (result CreateLifecycleWorkflowIdentityGovernanceRestoreOperationResponse, err error) {
	opts := client.RequestOptions{
		ContentType: "application/json; charset=utf-8",
		ExpectedStatusCodes: []int{
			http.StatusOK,
		},
		HttpMethod:    http.MethodPost,
		OptionsObject: options,
		Path:          fmt.Sprintf("%s/identityGovernance.restore", id.ID()),
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

	var model beta.IdentityGovernanceWorkflow
	result.Model = &model
	if err = resp.Unmarshal(result.Model); err != nil {
		return
	}

	return
}
