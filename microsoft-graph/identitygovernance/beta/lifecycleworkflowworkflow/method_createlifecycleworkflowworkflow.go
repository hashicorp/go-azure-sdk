package lifecycleworkflowworkflow

import (
	"context"
	"net/http"

	"github.com/hashicorp/go-azure-sdk/microsoft-graph/common-types/beta"
	"github.com/hashicorp/go-azure-sdk/sdk/client"
	"github.com/hashicorp/go-azure-sdk/sdk/odata"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type CreateLifecycleWorkflowWorkflowOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
	Model        *beta.IdentityGovernanceWorkflow
}

type CreateLifecycleWorkflowWorkflowOperationOptions struct {
	Metadata *odata.Metadata
}

func DefaultCreateLifecycleWorkflowWorkflowOperationOptions() CreateLifecycleWorkflowWorkflowOperationOptions {
	return CreateLifecycleWorkflowWorkflowOperationOptions{}
}

func (o CreateLifecycleWorkflowWorkflowOperationOptions) ToHeaders() *client.Headers {
	out := client.Headers{}

	return &out
}

func (o CreateLifecycleWorkflowWorkflowOperationOptions) ToOData() *odata.Query {
	out := odata.Query{}
	if o.Metadata != nil {
		out.Metadata = *o.Metadata
	}
	return &out
}

func (o CreateLifecycleWorkflowWorkflowOperationOptions) ToQuery() *client.QueryParams {
	out := client.QueryParams{}

	return &out
}

// CreateLifecycleWorkflowWorkflow - Create workflow. Create a new workflow object. You can create up to 100 workflows
// in a tenant.
func (c LifecycleWorkflowWorkflowClient) CreateLifecycleWorkflowWorkflow(ctx context.Context, input beta.IdentityGovernanceWorkflow, options CreateLifecycleWorkflowWorkflowOperationOptions) (result CreateLifecycleWorkflowWorkflowOperationResponse, err error) {
	opts := client.RequestOptions{
		ContentType: "application/json; charset=utf-8",
		ExpectedStatusCodes: []int{
			http.StatusCreated,
		},
		HttpMethod:    http.MethodPost,
		OptionsObject: options,
		Path:          "/identityGovernance/lifecycleWorkflows/workflows",
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

	var model beta.IdentityGovernanceWorkflow
	result.Model = &model
	if err = resp.Unmarshal(result.Model); err != nil {
		return
	}

	return
}
