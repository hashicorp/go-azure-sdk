package lifecycleworkflowdeleteditemworkflowtaskreporttaskprocessingresult

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

type CreateLifecycleWorkflowDeletedItemWorkflowTaskReportTaskProcessingResultIdentityGovernanceResumeOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
}

type CreateLifecycleWorkflowDeletedItemWorkflowTaskReportTaskProcessingResultIdentityGovernanceResumeOperationOptions struct {
	Metadata *odata.Metadata
}

func DefaultCreateLifecycleWorkflowDeletedItemWorkflowTaskReportTaskProcessingResultIdentityGovernanceResumeOperationOptions() CreateLifecycleWorkflowDeletedItemWorkflowTaskReportTaskProcessingResultIdentityGovernanceResumeOperationOptions {
	return CreateLifecycleWorkflowDeletedItemWorkflowTaskReportTaskProcessingResultIdentityGovernanceResumeOperationOptions{}
}

func (o CreateLifecycleWorkflowDeletedItemWorkflowTaskReportTaskProcessingResultIdentityGovernanceResumeOperationOptions) ToHeaders() *client.Headers {
	out := client.Headers{}

	return &out
}

func (o CreateLifecycleWorkflowDeletedItemWorkflowTaskReportTaskProcessingResultIdentityGovernanceResumeOperationOptions) ToOData() *odata.Query {
	out := odata.Query{}
	if o.Metadata != nil {
		out.Metadata = *o.Metadata
	}
	return &out
}

func (o CreateLifecycleWorkflowDeletedItemWorkflowTaskReportTaskProcessingResultIdentityGovernanceResumeOperationOptions) ToQuery() *client.QueryParams {
	out := client.QueryParams{}

	return &out
}

// CreateLifecycleWorkflowDeletedItemWorkflowTaskReportTaskProcessingResultIdentityGovernanceResume - Invoke action
// resume. Resume a task processing result that's inProgress. In the default case an Azure Logic Apps system-assigned
// managed identity calls this API. For more information, see: Lifecycle Workflows extensibility approach.
func (c LifecycleWorkflowDeletedItemWorkflowTaskReportTaskProcessingResultClient) CreateLifecycleWorkflowDeletedItemWorkflowTaskReportTaskProcessingResultIdentityGovernanceResume(ctx context.Context, id beta.IdentityGovernanceLifecycleWorkflowDeletedItemWorkflowIdTaskReportIdTaskProcessingResultId, input CreateLifecycleWorkflowDeletedItemWorkflowTaskReportTaskProcessingResultIdentityGovernanceResumeRequest, options CreateLifecycleWorkflowDeletedItemWorkflowTaskReportTaskProcessingResultIdentityGovernanceResumeOperationOptions) (result CreateLifecycleWorkflowDeletedItemWorkflowTaskReportTaskProcessingResultIdentityGovernanceResumeOperationResponse, err error) {
	opts := client.RequestOptions{
		ContentType: "application/json; charset=utf-8",
		ExpectedStatusCodes: []int{
			http.StatusNoContent,
		},
		HttpMethod:    http.MethodPost,
		OptionsObject: options,
		Path:          fmt.Sprintf("%s/identityGovernance.resume", id.ID()),
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
