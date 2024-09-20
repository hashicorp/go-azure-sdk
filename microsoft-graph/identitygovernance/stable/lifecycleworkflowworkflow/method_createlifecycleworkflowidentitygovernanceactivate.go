package lifecycleworkflowworkflow

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

type CreateLifecycleWorkflowIdentityGovernanceActivateOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
}

type CreateLifecycleWorkflowIdentityGovernanceActivateOperationOptions struct {
	Metadata *odata.Metadata
}

func DefaultCreateLifecycleWorkflowIdentityGovernanceActivateOperationOptions() CreateLifecycleWorkflowIdentityGovernanceActivateOperationOptions {
	return CreateLifecycleWorkflowIdentityGovernanceActivateOperationOptions{}
}

func (o CreateLifecycleWorkflowIdentityGovernanceActivateOperationOptions) ToHeaders() *client.Headers {
	out := client.Headers{}

	return &out
}

func (o CreateLifecycleWorkflowIdentityGovernanceActivateOperationOptions) ToOData() *odata.Query {
	out := odata.Query{}
	if o.Metadata != nil {
		out.Metadata = *o.Metadata
	}
	return &out
}

func (o CreateLifecycleWorkflowIdentityGovernanceActivateOperationOptions) ToQuery() *client.QueryParams {
	out := client.QueryParams{}

	return &out
}

// CreateLifecycleWorkflowIdentityGovernanceActivate - Invoke action activate. Run a workflow object on-demand. You can
// run any workflow on-demand, including scheduled workflows. Workflows created from the 'Real-time employee
// termination' template are run on-demand only. When you run a workflow on demand, the tasks are executed regardless of
// whether the user state matches the scope and trigger execution conditions.
func (c LifecycleWorkflowWorkflowClient) CreateLifecycleWorkflowIdentityGovernanceActivate(ctx context.Context, id stable.IdentityGovernanceLifecycleWorkflowWorkflowId, input CreateLifecycleWorkflowIdentityGovernanceActivateRequest, options CreateLifecycleWorkflowIdentityGovernanceActivateOperationOptions) (result CreateLifecycleWorkflowIdentityGovernanceActivateOperationResponse, err error) {
	opts := client.RequestOptions{
		ContentType: "application/json; charset=utf-8",
		ExpectedStatusCodes: []int{
			http.StatusNoContent,
		},
		HttpMethod:    http.MethodPost,
		OptionsObject: options,
		Path:          fmt.Sprintf("%s/identityGovernance.activate", id.ID()),
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
