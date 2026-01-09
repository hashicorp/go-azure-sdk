package lifecycleworkflowcustomtaskextensionlastmodifiedby

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

type GetLifecycleWorkflowCustomTaskExtensionLastModifiedByOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
	Model        *stable.User
}

type GetLifecycleWorkflowCustomTaskExtensionLastModifiedByOperationOptions struct {
	Expand    *odata.Expand
	Metadata  *odata.Metadata
	RetryFunc client.RequestRetryFunc
	Select    *[]string
}

func DefaultGetLifecycleWorkflowCustomTaskExtensionLastModifiedByOperationOptions() GetLifecycleWorkflowCustomTaskExtensionLastModifiedByOperationOptions {
	return GetLifecycleWorkflowCustomTaskExtensionLastModifiedByOperationOptions{}
}

func (o GetLifecycleWorkflowCustomTaskExtensionLastModifiedByOperationOptions) ToHeaders() *client.Headers {
	out := client.Headers{}

	return &out
}

func (o GetLifecycleWorkflowCustomTaskExtensionLastModifiedByOperationOptions) ToOData() *odata.Query {
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

func (o GetLifecycleWorkflowCustomTaskExtensionLastModifiedByOperationOptions) ToQuery() *client.QueryParams {
	out := client.QueryParams{}

	return &out
}

// GetLifecycleWorkflowCustomTaskExtensionLastModifiedBy - Get lastModifiedBy from identityGovernance. The unique
// identifier of the Microsoft Entra user that modified the custom task extension last.Supports $filter(eq, ne) and
// $expand.
func (c LifecycleWorkflowCustomTaskExtensionLastModifiedByClient) GetLifecycleWorkflowCustomTaskExtensionLastModifiedBy(ctx context.Context, id stable.IdentityGovernanceLifecycleWorkflowCustomTaskExtensionId, options GetLifecycleWorkflowCustomTaskExtensionLastModifiedByOperationOptions) (result GetLifecycleWorkflowCustomTaskExtensionLastModifiedByOperationResponse, err error) {
	opts := client.RequestOptions{
		ContentType: "application/json; charset=utf-8",
		ExpectedStatusCodes: []int{
			http.StatusOK,
		},
		HttpMethod:    http.MethodGet,
		OptionsObject: options,
		Path:          fmt.Sprintf("%s/lastModifiedBy", id.ID()),
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

	var model stable.User
	result.Model = &model
	if err = resp.Unmarshal(result.Model); err != nil {
		return
	}

	return
}
