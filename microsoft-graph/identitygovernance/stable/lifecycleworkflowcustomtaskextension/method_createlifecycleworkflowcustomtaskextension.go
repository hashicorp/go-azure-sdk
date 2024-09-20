package lifecycleworkflowcustomtaskextension

import (
	"context"
	"net/http"

	"github.com/hashicorp/go-azure-sdk/microsoft-graph/common-types/stable"
	"github.com/hashicorp/go-azure-sdk/sdk/client"
	"github.com/hashicorp/go-azure-sdk/sdk/odata"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type CreateLifecycleWorkflowCustomTaskExtensionOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
	Model        *stable.IdentityGovernanceCustomTaskExtension
}

type CreateLifecycleWorkflowCustomTaskExtensionOperationOptions struct {
	Metadata *odata.Metadata
}

func DefaultCreateLifecycleWorkflowCustomTaskExtensionOperationOptions() CreateLifecycleWorkflowCustomTaskExtensionOperationOptions {
	return CreateLifecycleWorkflowCustomTaskExtensionOperationOptions{}
}

func (o CreateLifecycleWorkflowCustomTaskExtensionOperationOptions) ToHeaders() *client.Headers {
	out := client.Headers{}

	return &out
}

func (o CreateLifecycleWorkflowCustomTaskExtensionOperationOptions) ToOData() *odata.Query {
	out := odata.Query{}
	if o.Metadata != nil {
		out.Metadata = *o.Metadata
	}
	return &out
}

func (o CreateLifecycleWorkflowCustomTaskExtensionOperationOptions) ToQuery() *client.QueryParams {
	out := client.QueryParams{}

	return &out
}

// CreateLifecycleWorkflowCustomTaskExtension - Create customTaskExtensions. Create a new customTaskExtension object.
func (c LifecycleWorkflowCustomTaskExtensionClient) CreateLifecycleWorkflowCustomTaskExtension(ctx context.Context, input stable.IdentityGovernanceCustomTaskExtension, options CreateLifecycleWorkflowCustomTaskExtensionOperationOptions) (result CreateLifecycleWorkflowCustomTaskExtensionOperationResponse, err error) {
	opts := client.RequestOptions{
		ContentType: "application/json; charset=utf-8",
		ExpectedStatusCodes: []int{
			http.StatusCreated,
		},
		HttpMethod:    http.MethodPost,
		OptionsObject: options,
		Path:          "/identityGovernance/lifecycleWorkflows/customTaskExtensions",
	}

	req, err := c.Client.NewRequest(ctx, opts)
	if err != nil {
		return
	}

	if err = req.Marshal(input); err != nil {
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

	var model stable.IdentityGovernanceCustomTaskExtension
	result.Model = &model
	if err = resp.Unmarshal(result.Model); err != nil {
		return
	}

	return
}
