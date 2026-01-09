package pendingaccessreviewinstancestagedecisioninstance

import (
	"context"
	"fmt"
	"net/http"

	"github.com/hashicorp/go-azure-sdk/microsoft-graph/common-types/beta"
	"github.com/hashicorp/go-azure-sdk/sdk/client"
	"github.com/hashicorp/go-azure-sdk/sdk/odata"
)

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type ResetPendingAccessReviewInstanceStageDecisionInstanceDecisionsOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
}

type ResetPendingAccessReviewInstanceStageDecisionInstanceDecisionsOperationOptions struct {
	Metadata  *odata.Metadata
	RetryFunc client.RequestRetryFunc
}

func DefaultResetPendingAccessReviewInstanceStageDecisionInstanceDecisionsOperationOptions() ResetPendingAccessReviewInstanceStageDecisionInstanceDecisionsOperationOptions {
	return ResetPendingAccessReviewInstanceStageDecisionInstanceDecisionsOperationOptions{}
}

func (o ResetPendingAccessReviewInstanceStageDecisionInstanceDecisionsOperationOptions) ToHeaders() *client.Headers {
	out := client.Headers{}

	return &out
}

func (o ResetPendingAccessReviewInstanceStageDecisionInstanceDecisionsOperationOptions) ToOData() *odata.Query {
	out := odata.Query{}
	if o.Metadata != nil {
		out.Metadata = *o.Metadata
	}
	return &out
}

func (o ResetPendingAccessReviewInstanceStageDecisionInstanceDecisionsOperationOptions) ToQuery() *client.QueryParams {
	out := client.QueryParams{}

	return &out
}

// ResetPendingAccessReviewInstanceStageDecisionInstanceDecisions - Invoke action resetDecisions. Resets decisions of
// all accessReviewInstanceDecisionItem objects on an accessReviewInstance to notReviewed.
func (c PendingAccessReviewInstanceStageDecisionInstanceClient) ResetPendingAccessReviewInstanceStageDecisionInstanceDecisions(ctx context.Context, id beta.MePendingAccessReviewInstanceIdStageIdDecisionId, options ResetPendingAccessReviewInstanceStageDecisionInstanceDecisionsOperationOptions) (result ResetPendingAccessReviewInstanceStageDecisionInstanceDecisionsOperationResponse, err error) {
	opts := client.RequestOptions{
		ContentType: "application/json; charset=utf-8",
		ExpectedStatusCodes: []int{
			http.StatusAccepted,
			http.StatusCreated,
			http.StatusNoContent,
			http.StatusOK,
		},
		HttpMethod:    http.MethodPost,
		OptionsObject: options,
		Path:          fmt.Sprintf("%s/instance/resetDecisions", id.ID()),
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

	return
}
