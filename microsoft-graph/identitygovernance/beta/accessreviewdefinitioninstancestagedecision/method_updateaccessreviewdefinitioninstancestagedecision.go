package accessreviewdefinitioninstancestagedecision

import (
	"context"
	"net/http"

	"github.com/hashicorp/go-azure-sdk/microsoft-graph/common-types/beta"
	"github.com/hashicorp/go-azure-sdk/sdk/client"
	"github.com/hashicorp/go-azure-sdk/sdk/odata"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type UpdateAccessReviewDefinitionInstanceStageDecisionOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
}

type UpdateAccessReviewDefinitionInstanceStageDecisionOperationOptions struct {
	Metadata  *odata.Metadata
	RetryFunc client.RequestRetryFunc
}

func DefaultUpdateAccessReviewDefinitionInstanceStageDecisionOperationOptions() UpdateAccessReviewDefinitionInstanceStageDecisionOperationOptions {
	return UpdateAccessReviewDefinitionInstanceStageDecisionOperationOptions{}
}

func (o UpdateAccessReviewDefinitionInstanceStageDecisionOperationOptions) ToHeaders() *client.Headers {
	out := client.Headers{}

	return &out
}

func (o UpdateAccessReviewDefinitionInstanceStageDecisionOperationOptions) ToOData() *odata.Query {
	out := odata.Query{}
	if o.Metadata != nil {
		out.Metadata = *o.Metadata
	}
	return &out
}

func (o UpdateAccessReviewDefinitionInstanceStageDecisionOperationOptions) ToQuery() *client.QueryParams {
	out := client.QueryParams{}

	return &out
}

// UpdateAccessReviewDefinitionInstanceStageDecision - Update accessReviewInstanceDecisionItem. Update access decisions,
// known as accessReviewInstanceDecisionItems, for which the user is the reviewer.
func (c AccessReviewDefinitionInstanceStageDecisionClient) UpdateAccessReviewDefinitionInstanceStageDecision(ctx context.Context, id beta.IdentityGovernanceAccessReviewDefinitionIdInstanceIdStageIdDecisionId, input beta.AccessReviewInstanceDecisionItem, options UpdateAccessReviewDefinitionInstanceStageDecisionOperationOptions) (result UpdateAccessReviewDefinitionInstanceStageDecisionOperationResponse, err error) {
	opts := client.RequestOptions{
		ContentType: "application/json; charset=utf-8",
		ExpectedStatusCodes: []int{
			http.StatusAccepted,
			http.StatusNoContent,
			http.StatusOK,
		},
		HttpMethod:    http.MethodPatch,
		OptionsObject: options,
		Path:          id.ID(),
		RetryFunc:     options.RetryFunc,
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
