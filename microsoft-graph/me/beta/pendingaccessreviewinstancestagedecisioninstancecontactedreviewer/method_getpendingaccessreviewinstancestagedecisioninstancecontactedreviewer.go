package pendingaccessreviewinstancestagedecisioninstancecontactedreviewer

import (
	"context"
	"net/http"

	"github.com/hashicorp/go-azure-sdk/microsoft-graph/common-types/beta"
	"github.com/hashicorp/go-azure-sdk/sdk/client"
	"github.com/hashicorp/go-azure-sdk/sdk/odata"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type GetPendingAccessReviewInstanceStageDecisionInstanceContactedReviewerOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
	Model        *beta.AccessReviewReviewer
}

type GetPendingAccessReviewInstanceStageDecisionInstanceContactedReviewerOperationOptions struct {
	Expand   *odata.Expand
	Metadata *odata.Metadata
	Select   *[]string
}

func DefaultGetPendingAccessReviewInstanceStageDecisionInstanceContactedReviewerOperationOptions() GetPendingAccessReviewInstanceStageDecisionInstanceContactedReviewerOperationOptions {
	return GetPendingAccessReviewInstanceStageDecisionInstanceContactedReviewerOperationOptions{}
}

func (o GetPendingAccessReviewInstanceStageDecisionInstanceContactedReviewerOperationOptions) ToHeaders() *client.Headers {
	out := client.Headers{}

	return &out
}

func (o GetPendingAccessReviewInstanceStageDecisionInstanceContactedReviewerOperationOptions) ToOData() *odata.Query {
	out := odata.Query{}
	if o.Expand != nil {
		out.Expand = *o.Expand
	}
	if o.Metadata != nil {
		out.Metadata = *o.Metadata
	}
	if o.Select != nil {
		out.Select = *o.Select
	}
	return &out
}

func (o GetPendingAccessReviewInstanceStageDecisionInstanceContactedReviewerOperationOptions) ToQuery() *client.QueryParams {
	out := client.QueryParams{}

	return &out
}

// GetPendingAccessReviewInstanceStageDecisionInstanceContactedReviewer - Get contactedReviewers from me. Returns the
// collection of reviewers who were contacted to complete this review. While the reviewers and fallbackReviewers
// properties of the accessReviewScheduleDefinition might specify group owners or managers as reviewers,
// contactedReviewers returns their individual identities. Supports $select. Read-only.
func (c PendingAccessReviewInstanceStageDecisionInstanceContactedReviewerClient) GetPendingAccessReviewInstanceStageDecisionInstanceContactedReviewer(ctx context.Context, id beta.MePendingAccessReviewInstanceIdStageIdDecisionIdInstanceContactedReviewerId, options GetPendingAccessReviewInstanceStageDecisionInstanceContactedReviewerOperationOptions) (result GetPendingAccessReviewInstanceStageDecisionInstanceContactedReviewerOperationResponse, err error) {
	opts := client.RequestOptions{
		ContentType: "application/json; charset=utf-8",
		ExpectedStatusCodes: []int{
			http.StatusOK,
		},
		HttpMethod:    http.MethodGet,
		OptionsObject: options,
		Path:          id.ID(),
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

	var model beta.AccessReviewReviewer
	result.Model = &model
	if err = resp.Unmarshal(result.Model); err != nil {
		return
	}

	return
}
