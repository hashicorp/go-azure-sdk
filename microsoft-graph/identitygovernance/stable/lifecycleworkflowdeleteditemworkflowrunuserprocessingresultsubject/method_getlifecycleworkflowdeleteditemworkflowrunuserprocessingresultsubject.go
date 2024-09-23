package lifecycleworkflowdeleteditemworkflowrunuserprocessingresultsubject

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

type GetLifecycleWorkflowDeletedItemWorkflowRunUserProcessingResultSubjectOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
	Model        *stable.User
}

type GetLifecycleWorkflowDeletedItemWorkflowRunUserProcessingResultSubjectOperationOptions struct {
	Expand    *odata.Expand
	Metadata  *odata.Metadata
	RetryFunc client.RequestRetryFunc
	Select    *[]string
}

func DefaultGetLifecycleWorkflowDeletedItemWorkflowRunUserProcessingResultSubjectOperationOptions() GetLifecycleWorkflowDeletedItemWorkflowRunUserProcessingResultSubjectOperationOptions {
	return GetLifecycleWorkflowDeletedItemWorkflowRunUserProcessingResultSubjectOperationOptions{}
}

func (o GetLifecycleWorkflowDeletedItemWorkflowRunUserProcessingResultSubjectOperationOptions) ToHeaders() *client.Headers {
	out := client.Headers{}

	return &out
}

func (o GetLifecycleWorkflowDeletedItemWorkflowRunUserProcessingResultSubjectOperationOptions) ToOData() *odata.Query {
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

func (o GetLifecycleWorkflowDeletedItemWorkflowRunUserProcessingResultSubjectOperationOptions) ToQuery() *client.QueryParams {
	out := client.QueryParams{}

	return &out
}

// GetLifecycleWorkflowDeletedItemWorkflowRunUserProcessingResultSubject - Get subject from identityGovernance. The
// unique identifier of the AAD user targeted for the taskProcessingResult.Supports $filter(eq, ne) and $expand.
func (c LifecycleWorkflowDeletedItemWorkflowRunUserProcessingResultSubjectClient) GetLifecycleWorkflowDeletedItemWorkflowRunUserProcessingResultSubject(ctx context.Context, id stable.IdentityGovernanceLifecycleWorkflowDeletedItemWorkflowIdRunIdUserProcessingResultId, options GetLifecycleWorkflowDeletedItemWorkflowRunUserProcessingResultSubjectOperationOptions) (result GetLifecycleWorkflowDeletedItemWorkflowRunUserProcessingResultSubjectOperationResponse, err error) {
	opts := client.RequestOptions{
		ContentType: "application/json; charset=utf-8",
		ExpectedStatusCodes: []int{
			http.StatusOK,
		},
		HttpMethod:    http.MethodGet,
		OptionsObject: options,
		Path:          fmt.Sprintf("%s/subject", id.ID()),
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
