package lifecycleworkflowworkflowversionlastmodifiedby

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

type GetLifecycleWorkflowVersionLastModifiedByOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
	Model        *beta.User
}

type GetLifecycleWorkflowVersionLastModifiedByOperationOptions struct {
	Expand *odata.Expand
	Select *[]string
}

func DefaultGetLifecycleWorkflowVersionLastModifiedByOperationOptions() GetLifecycleWorkflowVersionLastModifiedByOperationOptions {
	return GetLifecycleWorkflowVersionLastModifiedByOperationOptions{}
}

func (o GetLifecycleWorkflowVersionLastModifiedByOperationOptions) ToHeaders() *client.Headers {
	out := client.Headers{}

	return &out
}

func (o GetLifecycleWorkflowVersionLastModifiedByOperationOptions) ToOData() *odata.Query {
	out := odata.Query{}
	if o.Expand != nil {
		out.Expand = *o.Expand
	}
	if o.Select != nil {
		out.Select = *o.Select
	}
	return &out
}

func (o GetLifecycleWorkflowVersionLastModifiedByOperationOptions) ToQuery() *client.QueryParams {
	out := client.QueryParams{}

	return &out
}

// GetLifecycleWorkflowVersionLastModifiedBy - Get lastModifiedBy from identityGovernance. The user who last modified
// the workflow.
func (c LifecycleWorkflowWorkflowVersionLastModifiedByClient) GetLifecycleWorkflowVersionLastModifiedBy(ctx context.Context, id beta.IdentityGovernanceLifecycleWorkflowWorkflowIdVersionId, options GetLifecycleWorkflowVersionLastModifiedByOperationOptions) (result GetLifecycleWorkflowVersionLastModifiedByOperationResponse, err error) {
	opts := client.RequestOptions{
		ContentType: "application/json; charset=utf-8",
		ExpectedStatusCodes: []int{
			http.StatusOK,
		},
		HttpMethod:    http.MethodGet,
		OptionsObject: options,
		Path:          fmt.Sprintf("%s/lastModifiedBy", id.ID()),
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

	var model beta.User
	result.Model = &model
	if err = resp.Unmarshal(result.Model); err != nil {
		return
	}

	return
}
