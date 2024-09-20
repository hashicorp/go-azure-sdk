package pendingaccessreviewinstancedecisioninstancestagedecision

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

type CreatePendingAccessReviewInstanceDecisionInstanceStageDecisionRecordAllDecisionOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
}

type CreatePendingAccessReviewInstanceDecisionInstanceStageDecisionRecordAllDecisionOperationOptions struct {
	Metadata *odata.Metadata
}

func DefaultCreatePendingAccessReviewInstanceDecisionInstanceStageDecisionRecordAllDecisionOperationOptions() CreatePendingAccessReviewInstanceDecisionInstanceStageDecisionRecordAllDecisionOperationOptions {
	return CreatePendingAccessReviewInstanceDecisionInstanceStageDecisionRecordAllDecisionOperationOptions{}
}

func (o CreatePendingAccessReviewInstanceDecisionInstanceStageDecisionRecordAllDecisionOperationOptions) ToHeaders() *client.Headers {
	out := client.Headers{}

	return &out
}

func (o CreatePendingAccessReviewInstanceDecisionInstanceStageDecisionRecordAllDecisionOperationOptions) ToOData() *odata.Query {
	out := odata.Query{}
	if o.Metadata != nil {
		out.Metadata = *o.Metadata
	}
	return &out
}

func (o CreatePendingAccessReviewInstanceDecisionInstanceStageDecisionRecordAllDecisionOperationOptions) ToQuery() *client.QueryParams {
	out := client.QueryParams{}

	return &out
}

// CreatePendingAccessReviewInstanceDecisionInstanceStageDecisionRecordAllDecision - Invoke action recordAllDecisions.
// As a reviewer of an access review, record a decision for an accessReviewInstanceDecisionItem that is assigned to you
// and that matches the principal or resource IDs specified. If no IDs are specified, the decisions will apply to every
// accessReviewInstanceDecisionItem for which you are the reviewer.
func (c PendingAccessReviewInstanceDecisionInstanceStageDecisionClient) CreatePendingAccessReviewInstanceDecisionInstanceStageDecisionRecordAllDecision(ctx context.Context, id beta.MePendingAccessReviewInstanceIdDecisionIdInstanceStageId, input CreatePendingAccessReviewInstanceDecisionInstanceStageDecisionRecordAllDecisionRequest, options CreatePendingAccessReviewInstanceDecisionInstanceStageDecisionRecordAllDecisionOperationOptions) (result CreatePendingAccessReviewInstanceDecisionInstanceStageDecisionRecordAllDecisionOperationResponse, err error) {
	opts := client.RequestOptions{
		ContentType: "application/json; charset=utf-8",
		ExpectedStatusCodes: []int{
			http.StatusNoContent,
		},
		HttpMethod:    http.MethodPost,
		OptionsObject: options,
		Path:          fmt.Sprintf("%s/decisions/recordAllDecisions", id.ID()),
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
