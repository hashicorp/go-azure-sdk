package pendingaccessreviewinstancestagedecisioninstance

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

type CreatePendingAccessReviewInstanceStageDecisionInstanceBatchRecordDecisionOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
}

type CreatePendingAccessReviewInstanceStageDecisionInstanceBatchRecordDecisionOperationOptions struct {
	Metadata *odata.Metadata
}

func DefaultCreatePendingAccessReviewInstanceStageDecisionInstanceBatchRecordDecisionOperationOptions() CreatePendingAccessReviewInstanceStageDecisionInstanceBatchRecordDecisionOperationOptions {
	return CreatePendingAccessReviewInstanceStageDecisionInstanceBatchRecordDecisionOperationOptions{}
}

func (o CreatePendingAccessReviewInstanceStageDecisionInstanceBatchRecordDecisionOperationOptions) ToHeaders() *client.Headers {
	out := client.Headers{}

	return &out
}

func (o CreatePendingAccessReviewInstanceStageDecisionInstanceBatchRecordDecisionOperationOptions) ToOData() *odata.Query {
	out := odata.Query{}
	if o.Metadata != nil {
		out.Metadata = *o.Metadata
	}
	return &out
}

func (o CreatePendingAccessReviewInstanceStageDecisionInstanceBatchRecordDecisionOperationOptions) ToQuery() *client.QueryParams {
	out := client.QueryParams{}

	return &out
}

// CreatePendingAccessReviewInstanceStageDecisionInstanceBatchRecordDecision - Invoke action batchRecordDecisions.
// Enables reviewers to review all accessReviewInstanceDecisionItem objects in batches by using principalId, resourceId,
// or neither.
func (c PendingAccessReviewInstanceStageDecisionInstanceClient) CreatePendingAccessReviewInstanceStageDecisionInstanceBatchRecordDecision(ctx context.Context, id beta.UserIdPendingAccessReviewInstanceIdStageIdDecisionId, input CreatePendingAccessReviewInstanceStageDecisionInstanceBatchRecordDecisionRequest, options CreatePendingAccessReviewInstanceStageDecisionInstanceBatchRecordDecisionOperationOptions) (result CreatePendingAccessReviewInstanceStageDecisionInstanceBatchRecordDecisionOperationResponse, err error) {
	opts := client.RequestOptions{
		ContentType: "application/json; charset=utf-8",
		ExpectedStatusCodes: []int{
			http.StatusNoContent,
		},
		HttpMethod:    http.MethodPost,
		OptionsObject: options,
		Path:          fmt.Sprintf("%s/instance/batchRecordDecisions", id.ID()),
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
