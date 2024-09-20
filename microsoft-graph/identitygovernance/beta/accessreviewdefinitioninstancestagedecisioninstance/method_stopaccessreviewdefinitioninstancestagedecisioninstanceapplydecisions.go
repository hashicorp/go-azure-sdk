package accessreviewdefinitioninstancestagedecisioninstance

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

type StopAccessReviewDefinitionInstanceStageDecisionInstanceApplyDecisionsOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
}

type StopAccessReviewDefinitionInstanceStageDecisionInstanceApplyDecisionsOperationOptions struct {
	Metadata *odata.Metadata
}

func DefaultStopAccessReviewDefinitionInstanceStageDecisionInstanceApplyDecisionsOperationOptions() StopAccessReviewDefinitionInstanceStageDecisionInstanceApplyDecisionsOperationOptions {
	return StopAccessReviewDefinitionInstanceStageDecisionInstanceApplyDecisionsOperationOptions{}
}

func (o StopAccessReviewDefinitionInstanceStageDecisionInstanceApplyDecisionsOperationOptions) ToHeaders() *client.Headers {
	out := client.Headers{}

	return &out
}

func (o StopAccessReviewDefinitionInstanceStageDecisionInstanceApplyDecisionsOperationOptions) ToOData() *odata.Query {
	out := odata.Query{}
	if o.Metadata != nil {
		out.Metadata = *o.Metadata
	}
	return &out
}

func (o StopAccessReviewDefinitionInstanceStageDecisionInstanceApplyDecisionsOperationOptions) ToQuery() *client.QueryParams {
	out := client.QueryParams{}

	return &out
}

// StopAccessReviewDefinitionInstanceStageDecisionInstanceApplyDecisions - Invoke action stopApplyDecisions
func (c AccessReviewDefinitionInstanceStageDecisionInstanceClient) StopAccessReviewDefinitionInstanceStageDecisionInstanceApplyDecisions(ctx context.Context, id beta.IdentityGovernanceAccessReviewDefinitionIdInstanceIdStageIdDecisionId, options StopAccessReviewDefinitionInstanceStageDecisionInstanceApplyDecisionsOperationOptions) (result StopAccessReviewDefinitionInstanceStageDecisionInstanceApplyDecisionsOperationResponse, err error) {
	opts := client.RequestOptions{
		ContentType: "application/json; charset=utf-8",
		ExpectedStatusCodes: []int{
			http.StatusNoContent,
		},
		HttpMethod:    http.MethodPost,
		OptionsObject: options,
		Path:          fmt.Sprintf("%s/instance/stopApplyDecisions", id.ID()),
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
