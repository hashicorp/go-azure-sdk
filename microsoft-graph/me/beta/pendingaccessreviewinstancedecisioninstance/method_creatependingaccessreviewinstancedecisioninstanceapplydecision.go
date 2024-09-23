package pendingaccessreviewinstancedecisioninstance

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

type CreatePendingAccessReviewInstanceDecisionInstanceApplyDecisionOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
}

type CreatePendingAccessReviewInstanceDecisionInstanceApplyDecisionOperationOptions struct {
	Metadata  *odata.Metadata
	RetryFunc client.RequestRetryFunc
}

func DefaultCreatePendingAccessReviewInstanceDecisionInstanceApplyDecisionOperationOptions() CreatePendingAccessReviewInstanceDecisionInstanceApplyDecisionOperationOptions {
	return CreatePendingAccessReviewInstanceDecisionInstanceApplyDecisionOperationOptions{}
}

func (o CreatePendingAccessReviewInstanceDecisionInstanceApplyDecisionOperationOptions) ToHeaders() *client.Headers {
	out := client.Headers{}

	return &out
}

func (o CreatePendingAccessReviewInstanceDecisionInstanceApplyDecisionOperationOptions) ToOData() *odata.Query {
	out := odata.Query{}
	if o.Metadata != nil {
		out.Metadata = *o.Metadata
	}
	return &out
}

func (o CreatePendingAccessReviewInstanceDecisionInstanceApplyDecisionOperationOptions) ToQuery() *client.QueryParams {
	out := client.QueryParams{}

	return &out
}

// CreatePendingAccessReviewInstanceDecisionInstanceApplyDecision - Invoke action applyDecisions. Apply review decisions
// on an accessReviewInstance if the decisions were not applied automatically because the autoApplyDecisionsEnabled
// property is false in the review's accessReviewScheduleSettings. The status of the accessReviewInstance must be
// Completed to call this method.
func (c PendingAccessReviewInstanceDecisionInstanceClient) CreatePendingAccessReviewInstanceDecisionInstanceApplyDecision(ctx context.Context, id beta.MePendingAccessReviewInstanceIdDecisionId, options CreatePendingAccessReviewInstanceDecisionInstanceApplyDecisionOperationOptions) (result CreatePendingAccessReviewInstanceDecisionInstanceApplyDecisionOperationResponse, err error) {
	opts := client.RequestOptions{
		ContentType: "application/json; charset=utf-8",
		ExpectedStatusCodes: []int{
			http.StatusNoContent,
		},
		HttpMethod:    http.MethodPost,
		OptionsObject: options,
		Path:          fmt.Sprintf("%s/instance/applyDecisions", id.ID()),
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
