package accessreviewdecisioninstancedecision

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

type CreateAccessReviewDecisionInstanceDecisionRecordAllDecisionOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
}

type CreateAccessReviewDecisionInstanceDecisionRecordAllDecisionOperationOptions struct {
	Metadata *odata.Metadata
}

func DefaultCreateAccessReviewDecisionInstanceDecisionRecordAllDecisionOperationOptions() CreateAccessReviewDecisionInstanceDecisionRecordAllDecisionOperationOptions {
	return CreateAccessReviewDecisionInstanceDecisionRecordAllDecisionOperationOptions{}
}

func (o CreateAccessReviewDecisionInstanceDecisionRecordAllDecisionOperationOptions) ToHeaders() *client.Headers {
	out := client.Headers{}

	return &out
}

func (o CreateAccessReviewDecisionInstanceDecisionRecordAllDecisionOperationOptions) ToOData() *odata.Query {
	out := odata.Query{}
	if o.Metadata != nil {
		out.Metadata = *o.Metadata
	}
	return &out
}

func (o CreateAccessReviewDecisionInstanceDecisionRecordAllDecisionOperationOptions) ToQuery() *client.QueryParams {
	out := client.QueryParams{}

	return &out
}

// CreateAccessReviewDecisionInstanceDecisionRecordAllDecision - Invoke action recordAllDecisions. As a reviewer of an
// access review, record a decision for an accessReviewInstanceDecisionItem that is assigned to you and that matches the
// principal or resource IDs specified. If no IDs are specified, the decisions will apply to every
// accessReviewInstanceDecisionItem for which you are the reviewer.
func (c AccessReviewDecisionInstanceDecisionClient) CreateAccessReviewDecisionInstanceDecisionRecordAllDecision(ctx context.Context, id beta.IdentityGovernanceAccessReviewDecisionId, input CreateAccessReviewDecisionInstanceDecisionRecordAllDecisionRequest, options CreateAccessReviewDecisionInstanceDecisionRecordAllDecisionOperationOptions) (result CreateAccessReviewDecisionInstanceDecisionRecordAllDecisionOperationResponse, err error) {
	opts := client.RequestOptions{
		ContentType: "application/json; charset=utf-8",
		ExpectedStatusCodes: []int{
			http.StatusNoContent,
		},
		HttpMethod:    http.MethodPost,
		OptionsObject: options,
		Path:          fmt.Sprintf("%s/instance/decisions/recordAllDecisions", id.ID()),
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
