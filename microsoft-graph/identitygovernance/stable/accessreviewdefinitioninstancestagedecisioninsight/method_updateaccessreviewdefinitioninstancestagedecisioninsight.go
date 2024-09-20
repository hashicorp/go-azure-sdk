package accessreviewdefinitioninstancestagedecisioninsight

import (
	"context"
	"net/http"

	"github.com/hashicorp/go-azure-sdk/microsoft-graph/common-types/stable"
	"github.com/hashicorp/go-azure-sdk/sdk/client"
	"github.com/hashicorp/go-azure-sdk/sdk/odata"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type UpdateAccessReviewDefinitionInstanceStageDecisionInsightOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
}

type UpdateAccessReviewDefinitionInstanceStageDecisionInsightOperationOptions struct {
	Metadata *odata.Metadata
}

func DefaultUpdateAccessReviewDefinitionInstanceStageDecisionInsightOperationOptions() UpdateAccessReviewDefinitionInstanceStageDecisionInsightOperationOptions {
	return UpdateAccessReviewDefinitionInstanceStageDecisionInsightOperationOptions{}
}

func (o UpdateAccessReviewDefinitionInstanceStageDecisionInsightOperationOptions) ToHeaders() *client.Headers {
	out := client.Headers{}

	return &out
}

func (o UpdateAccessReviewDefinitionInstanceStageDecisionInsightOperationOptions) ToOData() *odata.Query {
	out := odata.Query{}
	if o.Metadata != nil {
		out.Metadata = *o.Metadata
	}
	return &out
}

func (o UpdateAccessReviewDefinitionInstanceStageDecisionInsightOperationOptions) ToQuery() *client.QueryParams {
	out := client.QueryParams{}

	return &out
}

// UpdateAccessReviewDefinitionInstanceStageDecisionInsight - Update the navigation property insights in
// identityGovernance
func (c AccessReviewDefinitionInstanceStageDecisionInsightClient) UpdateAccessReviewDefinitionInstanceStageDecisionInsight(ctx context.Context, id stable.IdentityGovernanceAccessReviewDefinitionIdInstanceIdStageIdDecisionIdInsightId, input stable.GovernanceInsight, options UpdateAccessReviewDefinitionInstanceStageDecisionInsightOperationOptions) (result UpdateAccessReviewDefinitionInstanceStageDecisionInsightOperationResponse, err error) {
	opts := client.RequestOptions{
		ContentType: "application/json; charset=utf-8",
		ExpectedStatusCodes: []int{
			http.StatusNoContent,
		},
		HttpMethod:    http.MethodPatch,
		OptionsObject: options,
		Path:          id.ID(),
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

	return
}
