package lifecycleworkflowdeleteditemworkflowtask

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

type CreateLifecycleWorkflowDeletedItemWorkflowTaskOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
	Model        *stable.IdentityGovernanceTask
}

type CreateLifecycleWorkflowDeletedItemWorkflowTaskOperationOptions struct {
	Metadata *odata.Metadata
}

func DefaultCreateLifecycleWorkflowDeletedItemWorkflowTaskOperationOptions() CreateLifecycleWorkflowDeletedItemWorkflowTaskOperationOptions {
	return CreateLifecycleWorkflowDeletedItemWorkflowTaskOperationOptions{}
}

func (o CreateLifecycleWorkflowDeletedItemWorkflowTaskOperationOptions) ToHeaders() *client.Headers {
	out := client.Headers{}

	return &out
}

func (o CreateLifecycleWorkflowDeletedItemWorkflowTaskOperationOptions) ToOData() *odata.Query {
	out := odata.Query{}
	if o.Metadata != nil {
		out.Metadata = *o.Metadata
	}
	return &out
}

func (o CreateLifecycleWorkflowDeletedItemWorkflowTaskOperationOptions) ToQuery() *client.QueryParams {
	out := client.QueryParams{}

	return &out
}

// CreateLifecycleWorkflowDeletedItemWorkflowTask - Create new navigation property to tasks for identityGovernance
func (c LifecycleWorkflowDeletedItemWorkflowTaskClient) CreateLifecycleWorkflowDeletedItemWorkflowTask(ctx context.Context, id stable.IdentityGovernanceLifecycleWorkflowDeletedItemWorkflowId, input stable.IdentityGovernanceTask, options CreateLifecycleWorkflowDeletedItemWorkflowTaskOperationOptions) (result CreateLifecycleWorkflowDeletedItemWorkflowTaskOperationResponse, err error) {
	opts := client.RequestOptions{
		ContentType: "application/json; charset=utf-8",
		ExpectedStatusCodes: []int{
			http.StatusCreated,
		},
		HttpMethod:    http.MethodPost,
		OptionsObject: options,
		Path:          fmt.Sprintf("%s/tasks", id.ID()),
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

	var model stable.IdentityGovernanceTask
	result.Model = &model
	if err = resp.Unmarshal(result.Model); err != nil {
		return
	}

	return
}
