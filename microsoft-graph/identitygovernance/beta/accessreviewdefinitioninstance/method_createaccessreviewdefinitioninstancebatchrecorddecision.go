package accessreviewdefinitioninstance

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

type CreateAccessReviewDefinitionInstanceBatchRecordDecisionOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
}

type CreateAccessReviewDefinitionInstanceBatchRecordDecisionOperationOptions struct {
	Metadata *odata.Metadata
}

func DefaultCreateAccessReviewDefinitionInstanceBatchRecordDecisionOperationOptions() CreateAccessReviewDefinitionInstanceBatchRecordDecisionOperationOptions {
	return CreateAccessReviewDefinitionInstanceBatchRecordDecisionOperationOptions{}
}

func (o CreateAccessReviewDefinitionInstanceBatchRecordDecisionOperationOptions) ToHeaders() *client.Headers {
	out := client.Headers{}

	return &out
}

func (o CreateAccessReviewDefinitionInstanceBatchRecordDecisionOperationOptions) ToOData() *odata.Query {
	out := odata.Query{}
	if o.Metadata != nil {
		out.Metadata = *o.Metadata
	}
	return &out
}

func (o CreateAccessReviewDefinitionInstanceBatchRecordDecisionOperationOptions) ToQuery() *client.QueryParams {
	out := client.QueryParams{}

	return &out
}

// CreateAccessReviewDefinitionInstanceBatchRecordDecision - Invoke action batchRecordDecisions. Enables reviewers to
// review all accessReviewInstanceDecisionItem objects in batches by using principalId, resourceId, or neither.
func (c AccessReviewDefinitionInstanceClient) CreateAccessReviewDefinitionInstanceBatchRecordDecision(ctx context.Context, id beta.IdentityGovernanceAccessReviewDefinitionIdInstanceId, input CreateAccessReviewDefinitionInstanceBatchRecordDecisionRequest, options CreateAccessReviewDefinitionInstanceBatchRecordDecisionOperationOptions) (result CreateAccessReviewDefinitionInstanceBatchRecordDecisionOperationResponse, err error) {
	opts := client.RequestOptions{
		ContentType: "application/json; charset=utf-8",
		ExpectedStatusCodes: []int{
			http.StatusNoContent,
		},
		HttpMethod:    http.MethodPost,
		OptionsObject: options,
		Path:          fmt.Sprintf("%s/batchRecordDecisions", id.ID()),
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
