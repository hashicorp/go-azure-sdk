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

type GetLifecycleWorkflowDeletedItemWorkflowTaskReportTaskProcessingResultsCountOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
	Model        *[]byte
}

type GetLifecycleWorkflowDeletedItemWorkflowTaskReportTaskProcessingResultsCountOperationOptions struct {
	Filter    *string
	Metadata  *odata.Metadata
	RetryFunc client.RequestRetryFunc
	Search    *string
}

func DefaultGetLifecycleWorkflowDeletedItemWorkflowTaskReportTaskProcessingResultsCountOperationOptions() GetLifecycleWorkflowDeletedItemWorkflowTaskReportTaskProcessingResultsCountOperationOptions {
	return GetLifecycleWorkflowDeletedItemWorkflowTaskReportTaskProcessingResultsCountOperationOptions{}
}

func (o GetLifecycleWorkflowDeletedItemWorkflowTaskReportTaskProcessingResultsCountOperationOptions) ToHeaders() *client.Headers {
	out := client.Headers{}

	return &out
}

func (o GetLifecycleWorkflowDeletedItemWorkflowTaskReportTaskProcessingResultsCountOperationOptions) ToOData() *odata.Query {
	out := odata.Query{}
	if o.Filter != nil {
		out.Filter = *o.Filter
	}
	if o.Metadata != nil {
		out.Metadata = *o.Metadata
	}
	if o.Search != nil {
		out.Search = *o.Search
	}
	return &out
}

func (o GetLifecycleWorkflowDeletedItemWorkflowTaskReportTaskProcessingResultsCountOperationOptions) ToQuery() *client.QueryParams {
	out := client.QueryParams{}

	return &out
}

// GetLifecycleWorkflowDeletedItemWorkflowTaskReportTaskProcessingResultsCount - Get the number of the resource
func (c LifecycleWorkflowDeletedItemWorkflowTaskReportTaskProcessingResultClient) GetLifecycleWorkflowDeletedItemWorkflowTaskReportTaskProcessingResultsCount(ctx context.Context, id beta.IdentityGovernanceLifecycleWorkflowDeletedItemWorkflowIdTaskReportId, options GetLifecycleWorkflowDeletedItemWorkflowTaskReportTaskProcessingResultsCountOperationOptions) (result GetLifecycleWorkflowDeletedItemWorkflowTaskReportTaskProcessingResultsCountOperationResponse, err error) {
	opts := client.RequestOptions{
		ContentType: "text/plain",
		ExpectedStatusCodes: []int{
			http.StatusOK,
		},
		HttpMethod:    http.MethodGet,
		OptionsObject: options,
		Path:          fmt.Sprintf("%s/taskProcessingResults/$count", id.ID()),
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

	var model []byte
	result.Model = &model
	if err = resp.Unmarshal(result.Model); err != nil {
		return
	}

	return
}
