package lifecycleworkflowcustomtaskextension

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

type DeleteLifecycleWorkflowCustomTaskExtensionOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
}

type DeleteLifecycleWorkflowCustomTaskExtensionOperationOptions struct {
	IfMatch   *string
	Metadata  *odata.Metadata
	RetryFunc client.RequestRetryFunc
}

func DefaultDeleteLifecycleWorkflowCustomTaskExtensionOperationOptions() DeleteLifecycleWorkflowCustomTaskExtensionOperationOptions {
	return DeleteLifecycleWorkflowCustomTaskExtensionOperationOptions{}
}

func (o DeleteLifecycleWorkflowCustomTaskExtensionOperationOptions) ToHeaders() *client.Headers {
	out := client.Headers{}
	if o.IfMatch != nil {
		out.Append("If-Match", fmt.Sprintf("%v", *o.IfMatch))
	}
	return &out
}

func (o DeleteLifecycleWorkflowCustomTaskExtensionOperationOptions) ToOData() *odata.Query {
	out := odata.Query{}
	if o.Metadata != nil {
		out.Metadata = *o.Metadata
	}
	return &out
}

func (o DeleteLifecycleWorkflowCustomTaskExtensionOperationOptions) ToQuery() *client.QueryParams {
	out := client.QueryParams{}

	return &out
}

// DeleteLifecycleWorkflowCustomTaskExtension - Delete customTaskExtension. Delete a customTaskExtension object. A
// custom task extension can only be deleted if it is not referenced in any task objects in a lifecycle workflow.
func (c LifecycleWorkflowCustomTaskExtensionClient) DeleteLifecycleWorkflowCustomTaskExtension(ctx context.Context, id stable.IdentityGovernanceLifecycleWorkflowCustomTaskExtensionId, options DeleteLifecycleWorkflowCustomTaskExtensionOperationOptions) (result DeleteLifecycleWorkflowCustomTaskExtensionOperationResponse, err error) {
	opts := client.RequestOptions{
		ContentType: "application/json; charset=utf-8",
		ExpectedStatusCodes: []int{
			http.StatusNoContent,
			http.StatusOK,
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
