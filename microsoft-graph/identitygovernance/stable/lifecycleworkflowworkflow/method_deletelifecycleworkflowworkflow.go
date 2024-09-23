package lifecycleworkflowworkflow

import (
	"context"
	"fmt"
	"net/http"

	"github.com/hashicorp/go-azure-sdk/microsoft-graph/common-types/stable"
	"github.com/hashicorp/go-azure-sdk/sdk/client"
	"github.com/hashicorp/go-azure-sdk/sdk/odata"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type DeleteLifecycleWorkflowWorkflowOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
}

type DeleteLifecycleWorkflowWorkflowOperationOptions struct {
	IfMatch   *string
	Metadata  *odata.Metadata
	RetryFunc client.RequestRetryFunc
}

func DefaultDeleteLifecycleWorkflowWorkflowOperationOptions() DeleteLifecycleWorkflowWorkflowOperationOptions {
	return DeleteLifecycleWorkflowWorkflowOperationOptions{}
}

func (o DeleteLifecycleWorkflowWorkflowOperationOptions) ToHeaders() *client.Headers {
	out := client.Headers{}
	if o.IfMatch != nil {
		out.Append("If-Match", fmt.Sprintf("%v", *o.IfMatch))
	}
	return &out
}

func (o DeleteLifecycleWorkflowWorkflowOperationOptions) ToOData() *odata.Query {
	out := odata.Query{}
	if o.Metadata != nil {
		out.Metadata = *o.Metadata
	}
	return &out
}

func (o DeleteLifecycleWorkflowWorkflowOperationOptions) ToQuery() *client.QueryParams {
	out := client.QueryParams{}

	return &out
}

// DeleteLifecycleWorkflowWorkflow - Delete workflow. Delete a workflow object and its associated tasks,
// taskProcessingResults and versions. You can restore a deleted workflow and its associated objects within 30 days of
// deletion.
func (c LifecycleWorkflowWorkflowClient) DeleteLifecycleWorkflowWorkflow(ctx context.Context, id stable.IdentityGovernanceLifecycleWorkflowWorkflowId, options DeleteLifecycleWorkflowWorkflowOperationOptions) (result DeleteLifecycleWorkflowWorkflowOperationResponse, err error) {
	opts := client.RequestOptions{
		ContentType: "application/json; charset=utf-8",
		ExpectedStatusCodes: []int{
			http.StatusNoContent,
		},
		HttpMethod:    http.MethodDelete,
		OptionsObject: options,
		Path:          id.ID(),
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

	return
}
