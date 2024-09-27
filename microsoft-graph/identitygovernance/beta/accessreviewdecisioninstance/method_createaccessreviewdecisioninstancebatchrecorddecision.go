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

type CreateAccessReviewDecisionInstanceBatchRecordDecisionOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
}

type CreateAccessReviewDecisionInstanceBatchRecordDecisionOperationOptions struct {
	Metadata  *odata.Metadata
	RetryFunc client.RequestRetryFunc
}

func DefaultCreateAccessReviewDecisionInstanceBatchRecordDecisionOperationOptions() CreateAccessReviewDecisionInstanceBatchRecordDecisionOperationOptions {
	return CreateAccessReviewDecisionInstanceBatchRecordDecisionOperationOptions{}
}

func (o CreateAccessReviewDecisionInstanceBatchRecordDecisionOperationOptions) ToHeaders() *client.Headers {
	out := client.Headers{}

	return &out
}

func (o CreateAccessReviewDecisionInstanceBatchRecordDecisionOperationOptions) ToOData() *odata.Query {
	out := odata.Query{}
	if o.Metadata != nil {
		out.Metadata = *o.Metadata
	}
	return &out
}

func (o CreateAccessReviewDecisionInstanceBatchRecordDecisionOperationOptions) ToQuery() *client.QueryParams {
	out := client.QueryParams{}

	return &out
}

// CreateAccessReviewDecisionInstanceBatchRecordDecision - Invoke action batchRecordDecisions. Enables reviewers to
// review all accessReviewInstanceDecisionItem objects in batches by using principalId, resourceId, or neither.
func (c AccessReviewDecisionInstanceClient) CreateAccessReviewDecisionInstanceBatchRecordDecision(ctx context.Context, id beta.IdentityGovernanceAccessReviewDecisionId, input CreateAccessReviewDecisionInstanceBatchRecordDecisionRequest, options CreateAccessReviewDecisionInstanceBatchRecordDecisionOperationOptions) (result CreateAccessReviewDecisionInstanceBatchRecordDecisionOperationResponse, err error) {
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
		Path:          fmt.Sprintf("%s/instance/batchRecordDecisions", id.ID()),
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
