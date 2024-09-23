package lifecycleworkflowdeleteditemworkflowversiontask

import (
	"context"
	"net/http"

	"github.com/hashicorp/go-azure-sdk/microsoft-graph/common-types/stable"
	"github.com/hashicorp/go-azure-sdk/sdk/client"
	"github.com/hashicorp/go-azure-sdk/sdk/odata"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type UpdateLifecycleWorkflowDeletedItemWorkflowVersionTaskOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
}

type UpdateLifecycleWorkflowDeletedItemWorkflowVersionTaskOperationOptions struct {
	Metadata  *odata.Metadata
	RetryFunc client.RequestRetryFunc
}

func DefaultUpdateLifecycleWorkflowDeletedItemWorkflowVersionTaskOperationOptions() UpdateLifecycleWorkflowDeletedItemWorkflowVersionTaskOperationOptions {
	return UpdateLifecycleWorkflowDeletedItemWorkflowVersionTaskOperationOptions{}
}

func (o UpdateLifecycleWorkflowDeletedItemWorkflowVersionTaskOperationOptions) ToHeaders() *client.Headers {
	out := client.Headers{}

	return &out
}

func (o UpdateLifecycleWorkflowDeletedItemWorkflowVersionTaskOperationOptions) ToOData() *odata.Query {
	out := odata.Query{}
	if o.Metadata != nil {
		out.Metadata = *o.Metadata
	}
	return &out
}

func (o UpdateLifecycleWorkflowDeletedItemWorkflowVersionTaskOperationOptions) ToQuery() *client.QueryParams {
	out := client.QueryParams{}

	return &out
}

// UpdateLifecycleWorkflowDeletedItemWorkflowVersionTask - Update the navigation property tasks in identityGovernance
func (c LifecycleWorkflowDeletedItemWorkflowVersionTaskClient) UpdateLifecycleWorkflowDeletedItemWorkflowVersionTask(ctx context.Context, id stable.IdentityGovernanceLifecycleWorkflowDeletedItemWorkflowIdVersionIdTaskId, input stable.IdentityGovernanceTask, options UpdateLifecycleWorkflowDeletedItemWorkflowVersionTaskOperationOptions) (result UpdateLifecycleWorkflowDeletedItemWorkflowVersionTaskOperationResponse, err error) {
	opts := client.RequestOptions{
		ContentType: "application/json; charset=utf-8",
		ExpectedStatusCodes: []int{
			http.StatusNoContent,
		},
		HttpMethod:    http.MethodPatch,
		OptionsObject: options,
		Path:          id.ID(),
		RetryFunc:     options.RetryFunc,
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
