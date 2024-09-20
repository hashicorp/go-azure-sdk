package accessreviewdefinitioninstancestagedecisioninsight

import (
	"context"
	"encoding/json"
	"net/http"

	"github.com/hashicorp/go-azure-sdk/microsoft-graph/common-types/stable"
	"github.com/hashicorp/go-azure-sdk/sdk/client"
	"github.com/hashicorp/go-azure-sdk/sdk/odata"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type GetAccessReviewDefinitionInstanceStageDecisionInsightOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
	Model        stable.GovernanceInsight
}

type GetAccessReviewDefinitionInstanceStageDecisionInsightOperationOptions struct {
	Expand   *odata.Expand
	Metadata *odata.Metadata
	Select   *[]string
}

func DefaultGetAccessReviewDefinitionInstanceStageDecisionInsightOperationOptions() GetAccessReviewDefinitionInstanceStageDecisionInsightOperationOptions {
	return GetAccessReviewDefinitionInstanceStageDecisionInsightOperationOptions{}
}

func (o GetAccessReviewDefinitionInstanceStageDecisionInsightOperationOptions) ToHeaders() *client.Headers {
	out := client.Headers{}

	return &out
}

func (o GetAccessReviewDefinitionInstanceStageDecisionInsightOperationOptions) ToOData() *odata.Query {
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

func (o GetAccessReviewDefinitionInstanceStageDecisionInsightOperationOptions) ToQuery() *client.QueryParams {
	out := client.QueryParams{}

	return &out
}

// GetAccessReviewDefinitionInstanceStageDecisionInsight - Get insights from identityGovernance. Insights are
// recommendations to reviewers on whether to approve or deny a decision. There can be multiple insights associated with
// an accessReviewInstanceDecisionItem.
func (c AccessReviewDefinitionInstanceStageDecisionInsightClient) GetAccessReviewDefinitionInstanceStageDecisionInsight(ctx context.Context, id stable.IdentityGovernanceAccessReviewDefinitionIdInstanceIdStageIdDecisionIdInsightId, options GetAccessReviewDefinitionInstanceStageDecisionInsightOperationOptions) (result GetAccessReviewDefinitionInstanceStageDecisionInsightOperationResponse, err error) {
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

	var respObj json.RawMessage
	if err = resp.Unmarshal(&respObj); err != nil {
		return
	}
	model, err := stable.UnmarshalGovernanceInsightImplementation(respObj)
	if err != nil {
		return
	}
	result.Model = model

	return
}
