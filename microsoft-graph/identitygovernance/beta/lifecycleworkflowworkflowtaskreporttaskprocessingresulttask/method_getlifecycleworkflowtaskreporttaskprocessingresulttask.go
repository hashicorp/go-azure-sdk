package lifecycleworkflowworkflowtaskreporttaskprocessingresulttask

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

type GetLifecycleWorkflowTaskReportTaskProcessingResultTaskOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
	Model        *beta.IdentityGovernanceTask
}

type GetLifecycleWorkflowTaskReportTaskProcessingResultTaskOperationOptions struct {
	Expand    *odata.Expand
	Metadata  *odata.Metadata
	RetryFunc client.RequestRetryFunc
	Select    *[]string
}

func DefaultGetLifecycleWorkflowTaskReportTaskProcessingResultTaskOperationOptions() GetLifecycleWorkflowTaskReportTaskProcessingResultTaskOperationOptions {
	return GetLifecycleWorkflowTaskReportTaskProcessingResultTaskOperationOptions{}
}

func (o GetLifecycleWorkflowTaskReportTaskProcessingResultTaskOperationOptions) ToHeaders() *client.Headers {
	out := client.Headers{}

	return &out
}

func (o GetLifecycleWorkflowTaskReportTaskProcessingResultTaskOperationOptions) ToOData() *odata.Query {
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

func (o GetLifecycleWorkflowTaskReportTaskProcessingResultTaskOperationOptions) ToQuery() *client.QueryParams {
	out := client.QueryParams{}

	return &out
}

// GetLifecycleWorkflowTaskReportTaskProcessingResultTask - Get task from identityGovernance. The related workflow task
func (c LifecycleWorkflowWorkflowTaskReportTaskProcessingResultTaskClient) GetLifecycleWorkflowTaskReportTaskProcessingResultTask(ctx context.Context, id beta.IdentityGovernanceLifecycleWorkflowWorkflowIdTaskReportIdTaskProcessingResultId, options GetLifecycleWorkflowTaskReportTaskProcessingResultTaskOperationOptions) (result GetLifecycleWorkflowTaskReportTaskProcessingResultTaskOperationResponse, err error) {
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

	var model beta.IdentityGovernanceTask
	result.Model = &model
	if err = resp.Unmarshal(result.Model); err != nil {
		return
	}

	return
}
