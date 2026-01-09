package lifecycleworkflowdeleteditemworkflowtaskreporttaskdefinition

import (
	"context"
	"fmt"
	"net/http"

	"github.com/hashicorp/go-azure-sdk/microsoft-graph/common-types/beta"
	"github.com/hashicorp/go-azure-sdk/sdk/client"
	"github.com/hashicorp/go-azure-sdk/sdk/odata"
)

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type GetLifecycleWorkflowDeletedItemWorkflowTaskReportTaskDefinitionOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
	Model        *beta.IdentityGovernanceTaskDefinition
}

type GetLifecycleWorkflowDeletedItemWorkflowTaskReportTaskDefinitionOperationOptions struct {
	Expand    *odata.Expand
	Metadata  *odata.Metadata
	RetryFunc client.RequestRetryFunc
	Select    *[]string
}

func DefaultGetLifecycleWorkflowDeletedItemWorkflowTaskReportTaskDefinitionOperationOptions() GetLifecycleWorkflowDeletedItemWorkflowTaskReportTaskDefinitionOperationOptions {
	return GetLifecycleWorkflowDeletedItemWorkflowTaskReportTaskDefinitionOperationOptions{}
}

func (o GetLifecycleWorkflowDeletedItemWorkflowTaskReportTaskDefinitionOperationOptions) ToHeaders() *client.Headers {
	out := client.Headers{}

	return &out
}

func (o GetLifecycleWorkflowDeletedItemWorkflowTaskReportTaskDefinitionOperationOptions) ToOData() *odata.Query {
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

func (o GetLifecycleWorkflowDeletedItemWorkflowTaskReportTaskDefinitionOperationOptions) ToQuery() *client.QueryParams {
	out := client.QueryParams{}

	return &out
}

// GetLifecycleWorkflowDeletedItemWorkflowTaskReportTaskDefinition - Get taskDefinition from identityGovernance. The
// taskDefinition associated with the related lifecycle workflow task.Supports $filter(eq, ne) and $expand.
func (c LifecycleWorkflowDeletedItemWorkflowTaskReportTaskDefinitionClient) GetLifecycleWorkflowDeletedItemWorkflowTaskReportTaskDefinition(ctx context.Context, id beta.IdentityGovernanceLifecycleWorkflowDeletedItemWorkflowIdTaskReportId, options GetLifecycleWorkflowDeletedItemWorkflowTaskReportTaskDefinitionOperationOptions) (result GetLifecycleWorkflowDeletedItemWorkflowTaskReportTaskDefinitionOperationResponse, err error) {
	opts := client.RequestOptions{
		ContentType: "application/json; charset=utf-8",
		ExpectedStatusCodes: []int{
			http.StatusOK,
		},
		HttpMethod:    http.MethodGet,
		OptionsObject: options,
		Path:          fmt.Sprintf("%s/taskDefinition", id.ID()),
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

	var model beta.IdentityGovernanceTaskDefinition
	result.Model = &model
	if err = resp.Unmarshal(result.Model); err != nil {
		return
	}

	return
}
