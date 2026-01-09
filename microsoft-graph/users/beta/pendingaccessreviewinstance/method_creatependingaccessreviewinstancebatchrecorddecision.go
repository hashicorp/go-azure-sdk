package pendingaccessreviewinstance

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

type CreatePendingAccessReviewInstanceBatchRecordDecisionOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
}

type CreatePendingAccessReviewInstanceBatchRecordDecisionOperationOptions struct {
	Metadata  *odata.Metadata
	RetryFunc client.RequestRetryFunc
}

func DefaultCreatePendingAccessReviewInstanceBatchRecordDecisionOperationOptions() CreatePendingAccessReviewInstanceBatchRecordDecisionOperationOptions {
	return CreatePendingAccessReviewInstanceBatchRecordDecisionOperationOptions{}
}

func (o CreatePendingAccessReviewInstanceBatchRecordDecisionOperationOptions) ToHeaders() *client.Headers {
	out := client.Headers{}

	return &out
}

func (o CreatePendingAccessReviewInstanceBatchRecordDecisionOperationOptions) ToOData() *odata.Query {
	out := odata.Query{}
	if o.Metadata != nil {
		out.Metadata = *o.Metadata
	}
	return &out
}

func (o CreatePendingAccessReviewInstanceBatchRecordDecisionOperationOptions) ToQuery() *client.QueryParams {
	out := client.QueryParams{}

	return &out
}

// CreatePendingAccessReviewInstanceBatchRecordDecision - Invoke action batchRecordDecisions. Enables reviewers to
// review all accessReviewInstanceDecisionItem objects in batches by using principalId, resourceId, or neither.
func (c PendingAccessReviewInstanceClient) CreatePendingAccessReviewInstanceBatchRecordDecision(ctx context.Context, id beta.UserIdPendingAccessReviewInstanceId, input CreatePendingAccessReviewInstanceBatchRecordDecisionRequest, options CreatePendingAccessReviewInstanceBatchRecordDecisionOperationOptions) (result CreatePendingAccessReviewInstanceBatchRecordDecisionOperationResponse, err error) {
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
		Path:          fmt.Sprintf("%s/batchRecordDecisions", id.ID()),
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
