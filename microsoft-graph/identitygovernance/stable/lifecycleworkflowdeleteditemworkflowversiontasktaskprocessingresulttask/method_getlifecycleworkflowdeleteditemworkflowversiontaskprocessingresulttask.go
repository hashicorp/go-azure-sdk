package lifecycleworkflowdeleteditemworkflowversiontasktaskprocessingresulttask

import (
	"context"
	"fmt"
	"net/http"

	"github.com/hashicorp/go-azure-sdk/microsoft-graph/common-types/stable"
	"github.com/hashicorp/go-azure-sdk/sdk/client"
	"github.com/hashicorp/go-azure-sdk/sdk/odata"
)

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type GetLifecycleWorkflowDeletedItemWorkflowVersionTaskProcessingResultTaskOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
	Model        *stable.IdentityGovernanceTask
}

type GetLifecycleWorkflowDeletedItemWorkflowVersionTaskProcessingResultTaskOperationOptions struct {
	Expand    *odata.Expand
	Metadata  *odata.Metadata
	RetryFunc client.RequestRetryFunc
	Select    *[]string
}

func DefaultGetLifecycleWorkflowDeletedItemWorkflowVersionTaskProcessingResultTaskOperationOptions() GetLifecycleWorkflowDeletedItemWorkflowVersionTaskProcessingResultTaskOperationOptions {
	return GetLifecycleWorkflowDeletedItemWorkflowVersionTaskProcessingResultTaskOperationOptions{}
}

func (o GetLifecycleWorkflowDeletedItemWorkflowVersionTaskProcessingResultTaskOperationOptions) ToHeaders() *client.Headers {
	out := client.Headers{}

	return &out
}

func (o GetLifecycleWorkflowDeletedItemWorkflowVersionTaskProcessingResultTaskOperationOptions) ToOData() *odata.Query {
	out := odata.Query{}
	if o.Expand != nil {
		out.Expand = *o.Expand
	}
	if o.Metadata != nil {
		out.Metadata = *o.Metadata
	}
	if o.Select != nil {
		out.Select = *o.Select
	}
	return &out
}

func (o GetLifecycleWorkflowDeletedItemWorkflowVersionTaskProcessingResultTaskOperationOptions) ToQuery() *client.QueryParams {
	out := client.QueryParams{}

	return &out
}

// GetLifecycleWorkflowDeletedItemWorkflowVersionTaskProcessingResultTask - Get task from identityGovernance. The
// related workflow task
func (c LifecycleWorkflowDeletedItemWorkflowVersionTaskTaskProcessingResultTaskClient) GetLifecycleWorkflowDeletedItemWorkflowVersionTaskProcessingResultTask(ctx context.Context, id stable.IdentityGovernanceLifecycleWorkflowDeletedItemWorkflowIdVersionIdTaskIdTaskProcessingResultId, options GetLifecycleWorkflowDeletedItemWorkflowVersionTaskProcessingResultTaskOperationOptions) (result GetLifecycleWorkflowDeletedItemWorkflowVersionTaskProcessingResultTaskOperationResponse, err error) {
	opts := client.RequestOptions{
		ContentType: "application/json; charset=utf-8",
		ExpectedStatusCodes: []int{
			http.StatusOK,
		},
		HttpMethod:    http.MethodGet,
		OptionsObject: options,
		Path:          fmt.Sprintf("%s/task", id.ID()),
		RetryFunc:     options.RetryFunc,
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

	var model stable.IdentityGovernanceTask
	result.Model = &model
	if err = resp.Unmarshal(result.Model); err != nil {
		return
	}

	return
}
