package accessreviewdefinitioninstancedecisioninstancestage

import (
	"context"
	"net/http"

	"github.com/hashicorp/go-azure-sdk/microsoft-graph/common-types/beta"
	"github.com/hashicorp/go-azure-sdk/sdk/client"
	"github.com/hashicorp/go-azure-sdk/sdk/odata"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type GetAccessReviewDefinitionInstanceDecisionInstanceStageOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
	Model        *beta.AccessReviewStage
}

type GetAccessReviewDefinitionInstanceDecisionInstanceStageOperationOptions struct {
	Expand    *odata.Expand
	Metadata  *odata.Metadata
	RetryFunc client.RequestRetryFunc
	Select    *[]string
}

func DefaultGetAccessReviewDefinitionInstanceDecisionInstanceStageOperationOptions() GetAccessReviewDefinitionInstanceDecisionInstanceStageOperationOptions {
	return GetAccessReviewDefinitionInstanceDecisionInstanceStageOperationOptions{}
}

func (o GetAccessReviewDefinitionInstanceDecisionInstanceStageOperationOptions) ToHeaders() *client.Headers {
	out := client.Headers{}

	return &out
}

func (o GetAccessReviewDefinitionInstanceDecisionInstanceStageOperationOptions) ToOData() *odata.Query {
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

func (o GetAccessReviewDefinitionInstanceDecisionInstanceStageOperationOptions) ToQuery() *client.QueryParams {
	out := client.QueryParams{}

	return &out
}

// GetAccessReviewDefinitionInstanceDecisionInstanceStage - Get stages from identityGovernance. If the instance has
// multiple stages, this returns the collection of stages. A new stage will only be created when the previous stage
// ends. The existence, number, and settings of stages on a review instance are created based on the
// accessReviewStageSettings on the parent accessReviewScheduleDefinition.
func (c AccessReviewDefinitionInstanceDecisionInstanceStageClient) GetAccessReviewDefinitionInstanceDecisionInstanceStage(ctx context.Context, id beta.IdentityGovernanceAccessReviewDefinitionIdInstanceIdDecisionIdInstanceStageId, options GetAccessReviewDefinitionInstanceDecisionInstanceStageOperationOptions) (result GetAccessReviewDefinitionInstanceDecisionInstanceStageOperationResponse, err error) {
	opts := client.RequestOptions{
		ContentType: "application/json; charset=utf-8",
		ExpectedStatusCodes: []int{
			http.StatusOK,
		},
		HttpMethod:    http.MethodGet,
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

	var model beta.AccessReviewStage
	result.Model = &model
	if err = resp.Unmarshal(result.Model); err != nil {
		return
	}

	return
}
