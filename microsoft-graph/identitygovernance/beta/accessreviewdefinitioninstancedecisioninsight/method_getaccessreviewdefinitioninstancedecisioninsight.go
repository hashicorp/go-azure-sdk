package accessreviewdefinitioninstancedecisioninsight

import (
	"context"
	"encoding/json"
	"net/http"

	"github.com/hashicorp/go-azure-sdk/microsoft-graph/common-types/beta"
	"github.com/hashicorp/go-azure-sdk/sdk/client"
	"github.com/hashicorp/go-azure-sdk/sdk/odata"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type GetAccessReviewDefinitionInstanceDecisionInsightOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
	Model        beta.GovernanceInsight
}

type GetAccessReviewDefinitionInstanceDecisionInsightOperationOptions struct {
	Expand *odata.Expand
	Select *[]string
}

func DefaultGetAccessReviewDefinitionInstanceDecisionInsightOperationOptions() GetAccessReviewDefinitionInstanceDecisionInsightOperationOptions {
	return GetAccessReviewDefinitionInstanceDecisionInsightOperationOptions{}
}

func (o GetAccessReviewDefinitionInstanceDecisionInsightOperationOptions) ToHeaders() *client.Headers {
	out := client.Headers{}

	return &out
}

func (o GetAccessReviewDefinitionInstanceDecisionInsightOperationOptions) ToOData() *odata.Query {
	out := odata.Query{}
	if o.Expand != nil {
		out.Expand = *o.Expand
	}
	if o.Select != nil {
		out.Select = *o.Select
	}
	return &out
}

func (o GetAccessReviewDefinitionInstanceDecisionInsightOperationOptions) ToQuery() *client.QueryParams {
	out := client.QueryParams{}

	return &out
}

// GetAccessReviewDefinitionInstanceDecisionInsight - Get insights from identityGovernance. Insights are recommendations
// to reviewers on whether to approve or deny a decision. There can be multiple insights associated with an
// accessReviewInstanceDecisionItem.
func (c AccessReviewDefinitionInstanceDecisionInsightClient) GetAccessReviewDefinitionInstanceDecisionInsight(ctx context.Context, id beta.IdentityGovernanceAccessReviewDefinitionIdInstanceIdDecisionIdInsightId, options GetAccessReviewDefinitionInstanceDecisionInsightOperationOptions) (result GetAccessReviewDefinitionInstanceDecisionInsightOperationResponse, err error) {
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
	model, err := beta.UnmarshalGovernanceInsightImplementation(respObj)
	if err != nil {
		return
	}
	result.Model = model

	return
}
