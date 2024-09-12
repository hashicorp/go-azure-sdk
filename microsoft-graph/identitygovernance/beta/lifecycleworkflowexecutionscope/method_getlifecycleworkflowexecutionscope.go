package lifecycleworkflowexecutionscope

import (
	"context"
	"net/http"

	"github.com/hashicorp/go-azure-sdk/microsoft-graph/common-types/beta"
	"github.com/hashicorp/go-azure-sdk/sdk/client"
	"github.com/hashicorp/go-azure-sdk/sdk/odata"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type GetLifecycleWorkflowExecutionScopeOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
	Model        *beta.IdentityGovernanceUserProcessingResult
}

type GetLifecycleWorkflowExecutionScopeOperationOptions struct {
	Expand *odata.Expand
	Select *[]string
}

func DefaultGetLifecycleWorkflowExecutionScopeOperationOptions() GetLifecycleWorkflowExecutionScopeOperationOptions {
	return GetLifecycleWorkflowExecutionScopeOperationOptions{}
}

func (o GetLifecycleWorkflowExecutionScopeOperationOptions) ToHeaders() *client.Headers {
	out := client.Headers{}

	return &out
}

func (o GetLifecycleWorkflowExecutionScopeOperationOptions) ToOData() *odata.Query {
	out := odata.Query{}
	if o.Expand != nil {
		out.Expand = *o.Expand
	}
	if o.Select != nil {
		out.Select = *o.Select
	}
	return &out
}

func (o GetLifecycleWorkflowExecutionScopeOperationOptions) ToQuery() *client.QueryParams {
	out := client.QueryParams{}

	return &out
}

// GetLifecycleWorkflowExecutionScope - Get executionScope from identityGovernance. The unique identifier of the
// Microsoft Entra identity that last modified the workflow object.
func (c LifecycleWorkflowExecutionScopeClient) GetLifecycleWorkflowExecutionScope(ctx context.Context, id beta.IdentityGovernanceLifecycleWorkflowWorkflowIdExecutionScopeId, options GetLifecycleWorkflowExecutionScopeOperationOptions) (result GetLifecycleWorkflowExecutionScopeOperationResponse, err error) {
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

	var model beta.IdentityGovernanceUserProcessingResult
	result.Model = &model
	if err = resp.Unmarshal(result.Model); err != nil {
		return
	}

	return
}
