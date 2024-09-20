package lifecycleworkflowdeleteditemworkflow

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

type CreateLifecycleWorkflowDeletedItemWorkflowIdentityGovernanceCreateNewVersionOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
	Model        *stable.IdentityGovernanceWorkflow
}

type CreateLifecycleWorkflowDeletedItemWorkflowIdentityGovernanceCreateNewVersionOperationOptions struct {
	Metadata *odata.Metadata
}

func DefaultCreateLifecycleWorkflowDeletedItemWorkflowIdentityGovernanceCreateNewVersionOperationOptions() CreateLifecycleWorkflowDeletedItemWorkflowIdentityGovernanceCreateNewVersionOperationOptions {
	return CreateLifecycleWorkflowDeletedItemWorkflowIdentityGovernanceCreateNewVersionOperationOptions{}
}

func (o CreateLifecycleWorkflowDeletedItemWorkflowIdentityGovernanceCreateNewVersionOperationOptions) ToHeaders() *client.Headers {
	out := client.Headers{}

	return &out
}

func (o CreateLifecycleWorkflowDeletedItemWorkflowIdentityGovernanceCreateNewVersionOperationOptions) ToOData() *odata.Query {
	out := odata.Query{}
	if o.Metadata != nil {
		out.Metadata = *o.Metadata
	}
	return &out
}

func (o CreateLifecycleWorkflowDeletedItemWorkflowIdentityGovernanceCreateNewVersionOperationOptions) ToQuery() *client.QueryParams {
	out := client.QueryParams{}

	return &out
}

// CreateLifecycleWorkflowDeletedItemWorkflowIdentityGovernanceCreateNewVersion - Invoke action createNewVersion. Create
// a new version of the workflow object.
func (c LifecycleWorkflowDeletedItemWorkflowClient) CreateLifecycleWorkflowDeletedItemWorkflowIdentityGovernanceCreateNewVersion(ctx context.Context, id stable.IdentityGovernanceLifecycleWorkflowDeletedItemWorkflowId, input CreateLifecycleWorkflowDeletedItemWorkflowIdentityGovernanceCreateNewVersionRequest, options CreateLifecycleWorkflowDeletedItemWorkflowIdentityGovernanceCreateNewVersionOperationOptions) (result CreateLifecycleWorkflowDeletedItemWorkflowIdentityGovernanceCreateNewVersionOperationResponse, err error) {
	opts := client.RequestOptions{
		ContentType: "application/json; charset=utf-8",
		ExpectedStatusCodes: []int{
			http.StatusOK,
		},
		HttpMethod:    http.MethodPost,
		OptionsObject: options,
		Path:          fmt.Sprintf("%s/identityGovernance.createNewVersion", id.ID()),
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

	var model stable.IdentityGovernanceWorkflow
	result.Model = &model
	if err = resp.Unmarshal(result.Model); err != nil {
		return
	}

	return
}
