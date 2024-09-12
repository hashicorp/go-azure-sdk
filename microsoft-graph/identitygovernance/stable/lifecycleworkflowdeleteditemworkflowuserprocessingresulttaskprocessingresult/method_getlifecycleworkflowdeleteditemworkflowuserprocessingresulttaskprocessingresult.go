package lifecycleworkflowdeleteditemworkflowuserprocessingresulttaskprocessingresult

import (
	"context"
	"net/http"

	"github.com/hashicorp/go-azure-sdk/microsoft-graph/common-types/stable"
	"github.com/hashicorp/go-azure-sdk/sdk/client"
	"github.com/hashicorp/go-azure-sdk/sdk/odata"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type GetLifecycleWorkflowDeletedItemWorkflowUserProcessingResultTaskProcessingResultOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
	Model        *stable.IdentityGovernanceTaskProcessingResult
}

type GetLifecycleWorkflowDeletedItemWorkflowUserProcessingResultTaskProcessingResultOperationOptions struct {
	Expand *odata.Expand
	Select *[]string
}

func DefaultGetLifecycleWorkflowDeletedItemWorkflowUserProcessingResultTaskProcessingResultOperationOptions() GetLifecycleWorkflowDeletedItemWorkflowUserProcessingResultTaskProcessingResultOperationOptions {
	return GetLifecycleWorkflowDeletedItemWorkflowUserProcessingResultTaskProcessingResultOperationOptions{}
}

func (o GetLifecycleWorkflowDeletedItemWorkflowUserProcessingResultTaskProcessingResultOperationOptions) ToHeaders() *client.Headers {
	out := client.Headers{}

	return &out
}

func (o GetLifecycleWorkflowDeletedItemWorkflowUserProcessingResultTaskProcessingResultOperationOptions) ToOData() *odata.Query {
	out := odata.Query{}
	if o.Expand != nil {
		out.Expand = *o.Expand
	}
	if o.Select != nil {
		out.Select = *o.Select
	}
	return &out
}

func (o GetLifecycleWorkflowDeletedItemWorkflowUserProcessingResultTaskProcessingResultOperationOptions) ToQuery() *client.QueryParams {
	out := client.QueryParams{}

	return &out
}

// GetLifecycleWorkflowDeletedItemWorkflowUserProcessingResultTaskProcessingResult - Get taskProcessingResults from
// identityGovernance. The associated individual task execution.
func (c LifecycleWorkflowDeletedItemWorkflowUserProcessingResultTaskProcessingResultClient) GetLifecycleWorkflowDeletedItemWorkflowUserProcessingResultTaskProcessingResult(ctx context.Context, id stable.IdentityGovernanceLifecycleWorkflowDeletedItemWorkflowIdUserProcessingResultIdTaskProcessingResultId, options GetLifecycleWorkflowDeletedItemWorkflowUserProcessingResultTaskProcessingResultOperationOptions) (result GetLifecycleWorkflowDeletedItemWorkflowUserProcessingResultTaskProcessingResultOperationResponse, err error) {
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
