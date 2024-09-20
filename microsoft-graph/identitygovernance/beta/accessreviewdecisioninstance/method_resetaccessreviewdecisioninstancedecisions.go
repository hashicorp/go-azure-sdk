package accessreviewdecisioninstance

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

type ResetAccessReviewDecisionInstanceDecisionsOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
}

type ResetAccessReviewDecisionInstanceDecisionsOperationOptions struct {
	Metadata *odata.Metadata
}

func DefaultResetAccessReviewDecisionInstanceDecisionsOperationOptions() ResetAccessReviewDecisionInstanceDecisionsOperationOptions {
	return ResetAccessReviewDecisionInstanceDecisionsOperationOptions{}
}

func (o ResetAccessReviewDecisionInstanceDecisionsOperationOptions) ToHeaders() *client.Headers {
	out := client.Headers{}

	return &out
}

func (o ResetAccessReviewDecisionInstanceDecisionsOperationOptions) ToOData() *odata.Query {
	out := odata.Query{}
	if o.Metadata != nil {
		out.Metadata = *o.Metadata
	}
	return &out
}

func (o ResetAccessReviewDecisionInstanceDecisionsOperationOptions) ToQuery() *client.QueryParams {
	out := client.QueryParams{}

	return &out
}

// ResetAccessReviewDecisionInstanceDecisions - Invoke action resetDecisions. Resets decisions of all
// accessReviewInstanceDecisionItem objects on an accessReviewInstance to notReviewed.
func (c AccessReviewDecisionInstanceClient) ResetAccessReviewDecisionInstanceDecisions(ctx context.Context, id beta.IdentityGovernanceAccessReviewDecisionId, options ResetAccessReviewDecisionInstanceDecisionsOperationOptions) (result ResetAccessReviewDecisionInstanceDecisionsOperationResponse, err error) {
	opts := client.RequestOptions{
		ContentType: "application/json; charset=utf-8",
		ExpectedStatusCodes: []int{
			http.StatusNoContent,
		},
		HttpMethod:    http.MethodPost,
		OptionsObject: options,
		Path:          fmt.Sprintf("%s/instance/resetDecisions", id.ID()),
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
