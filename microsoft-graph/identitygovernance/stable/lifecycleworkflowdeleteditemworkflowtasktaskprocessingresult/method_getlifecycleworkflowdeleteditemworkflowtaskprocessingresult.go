package lifecycleworkflowdeleteditemworkflowtasktaskprocessingresult

import (
	"context"
	"net/http"

	"github.com/hashicorp/go-azure-sdk/microsoft-graph/common-types/stable"
	"github.com/hashicorp/go-azure-sdk/sdk/client"
	"github.com/hashicorp/go-azure-sdk/sdk/odata"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type GetLifecycleWorkflowDeletedItemWorkflowTaskProcessingResultOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
	Model        *stable.IdentityGovernanceTaskProcessingResult
}

type GetLifecycleWorkflowDeletedItemWorkflowTaskProcessingResultOperationOptions struct {
	Expand   *odata.Expand
	Metadata *odata.Metadata
	Select   *[]string
}

func DefaultGetLifecycleWorkflowDeletedItemWorkflowTaskProcessingResultOperationOptions() GetLifecycleWorkflowDeletedItemWorkflowTaskProcessingResultOperationOptions {
	return GetLifecycleWorkflowDeletedItemWorkflowTaskProcessingResultOperationOptions{}
}

func (o GetLifecycleWorkflowDeletedItemWorkflowTaskProcessingResultOperationOptions) ToHeaders() *client.Headers {
	out := client.Headers{}

	return &out
}

func (o GetLifecycleWorkflowDeletedItemWorkflowTaskProcessingResultOperationOptions) ToOData() *odata.Query {
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

func (o GetLifecycleWorkflowDeletedItemWorkflowTaskProcessingResultOperationOptions) ToQuery() *client.QueryParams {
	out := client.QueryParams{}

	return &out
}

// GetLifecycleWorkflowDeletedItemWorkflowTaskProcessingResult - Get taskProcessingResults from identityGovernance. The
// result of processing the task.
func (c LifecycleWorkflowDeletedItemWorkflowTaskTaskProcessingResultClient) GetLifecycleWorkflowDeletedItemWorkflowTaskProcessingResult(ctx context.Context, id stable.IdentityGovernanceLifecycleWorkflowDeletedItemWorkflowIdTaskIdTaskProcessingResultId, options GetLifecycleWorkflowDeletedItemWorkflowTaskProcessingResultOperationOptions) (result GetLifecycleWorkflowDeletedItemWorkflowTaskProcessingResultOperationResponse, err error) {
	opts := client.RequestOptions{
		ContentType: "application/json; charset=utf-8",
		ExpectedStatusCodes: []int{
			http.StatusOK,
		},
		HttpMethod:    http.MethodGet,
		OptionsObject: options,
		Path:          id.ID(),
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

	var model stable.IdentityGovernanceTaskProcessingResult
	result.Model = &model
	if err = resp.Unmarshal(result.Model); err != nil {
		return
	}

	return
}
