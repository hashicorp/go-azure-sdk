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

type CreateAccessReviewDecisionInstanceApplyDecisionOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
}

type CreateAccessReviewDecisionInstanceApplyDecisionOperationOptions struct {
	Metadata *odata.Metadata
}

func DefaultCreateAccessReviewDecisionInstanceApplyDecisionOperationOptions() CreateAccessReviewDecisionInstanceApplyDecisionOperationOptions {
	return CreateAccessReviewDecisionInstanceApplyDecisionOperationOptions{}
}

func (o CreateAccessReviewDecisionInstanceApplyDecisionOperationOptions) ToHeaders() *client.Headers {
	out := client.Headers{}

	return &out
}

func (o CreateAccessReviewDecisionInstanceApplyDecisionOperationOptions) ToOData() *odata.Query {
	out := odata.Query{}
	if o.Metadata != nil {
		out.Metadata = *o.Metadata
	}
	return &out
}

func (o CreateAccessReviewDecisionInstanceApplyDecisionOperationOptions) ToQuery() *client.QueryParams {
	out := client.QueryParams{}

	return &out
}

// CreateAccessReviewDecisionInstanceApplyDecision - Invoke action applyDecisions. Apply review decisions on an
// accessReviewInstance if the decisions were not applied automatically because the autoApplyDecisionsEnabled property
// is false in the review's accessReviewScheduleSettings. The status of the accessReviewInstance must be Completed to
// call this method.
func (c AccessReviewDecisionInstanceClient) CreateAccessReviewDecisionInstanceApplyDecision(ctx context.Context, id beta.IdentityGovernanceAccessReviewDecisionId, options CreateAccessReviewDecisionInstanceApplyDecisionOperationOptions) (result CreateAccessReviewDecisionInstanceApplyDecisionOperationResponse, err error) {
	opts := client.RequestOptions{
		ContentType: "application/json; charset=utf-8",
		ExpectedStatusCodes: []int{
			http.StatusNoContent,
		},
		HttpMethod:    http.MethodPost,
		OptionsObject: options,
		Path:          fmt.Sprintf("%s/instance/applyDecisions", id.ID()),
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
