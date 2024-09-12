package lifecycleworkflowrunuserprocessingresult

import (
	"context"
	"net/http"

	"github.com/hashicorp/go-azure-sdk/microsoft-graph/common-types/stable"
	"github.com/hashicorp/go-azure-sdk/sdk/client"
	"github.com/hashicorp/go-azure-sdk/sdk/odata"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type GetLifecycleWorkflowRunUserProcessingResultOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
	Model        *stable.IdentityGovernanceUserProcessingResult
}

type GetLifecycleWorkflowRunUserProcessingResultOperationOptions struct {
	Expand *odata.Expand
	Select *[]string
}

func DefaultGetLifecycleWorkflowRunUserProcessingResultOperationOptions() GetLifecycleWorkflowRunUserProcessingResultOperationOptions {
	return GetLifecycleWorkflowRunUserProcessingResultOperationOptions{}
}

func (o GetLifecycleWorkflowRunUserProcessingResultOperationOptions) ToHeaders() *client.Headers {
	out := client.Headers{}

	return &out
}

func (o GetLifecycleWorkflowRunUserProcessingResultOperationOptions) ToOData() *odata.Query {
	out := odata.Query{}
	if o.Expand != nil {
		out.Expand = *o.Expand
	}
	if o.Select != nil {
		out.Select = *o.Select
	}
	return &out
}

func (o GetLifecycleWorkflowRunUserProcessingResultOperationOptions) ToQuery() *client.QueryParams {
	out := client.QueryParams{}

	return &out
}

// GetLifecycleWorkflowRunUserProcessingResult - Get userProcessingResult. Get the user processing result of a user
// processing result of a run.
func (c LifecycleWorkflowRunUserProcessingResultClient) GetLifecycleWorkflowRunUserProcessingResult(ctx context.Context, id stable.IdentityGovernanceLifecycleWorkflowWorkflowIdRunIdUserProcessingResultId, options GetLifecycleWorkflowRunUserProcessingResultOperationOptions) (result GetLifecycleWorkflowRunUserProcessingResultOperationResponse, err error) {
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

	var model stable.IdentityGovernanceUserProcessingResult
	result.Model = &model
	if err = resp.Unmarshal(result.Model); err != nil {
		return
	}

	return
}
