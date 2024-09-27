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

type ResetPendingAccessReviewInstanceDecisionInstanceDecisionsOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
}

type ResetPendingAccessReviewInstanceDecisionInstanceDecisionsOperationOptions struct {
	Metadata  *odata.Metadata
	RetryFunc client.RequestRetryFunc
}

func DefaultResetPendingAccessReviewInstanceDecisionInstanceDecisionsOperationOptions() ResetPendingAccessReviewInstanceDecisionInstanceDecisionsOperationOptions {
	return ResetPendingAccessReviewInstanceDecisionInstanceDecisionsOperationOptions{}
}

func (o ResetPendingAccessReviewInstanceDecisionInstanceDecisionsOperationOptions) ToHeaders() *client.Headers {
	out := client.Headers{}

	return &out
}

func (o ResetPendingAccessReviewInstanceDecisionInstanceDecisionsOperationOptions) ToOData() *odata.Query {
	out := odata.Query{}
	if o.Metadata != nil {
		out.Metadata = *o.Metadata
	}
	return &out
}

func (o ResetPendingAccessReviewInstanceDecisionInstanceDecisionsOperationOptions) ToQuery() *client.QueryParams {
	out := client.QueryParams{}

	return &out
}

// ResetPendingAccessReviewInstanceDecisionInstanceDecisions - Invoke action resetDecisions. Resets decisions of all
// accessReviewInstanceDecisionItem objects on an accessReviewInstance to notReviewed.
func (c PendingAccessReviewInstanceDecisionInstanceClient) ResetPendingAccessReviewInstanceDecisionInstanceDecisions(ctx context.Context, id beta.UserIdPendingAccessReviewInstanceIdDecisionId, options ResetPendingAccessReviewInstanceDecisionInstanceDecisionsOperationOptions) (result ResetPendingAccessReviewInstanceDecisionInstanceDecisionsOperationResponse, err error) {
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
