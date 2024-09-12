package pendingaccessreviewinstancedecisioninstancestage

import (
	"context"
	"net/http"

	"github.com/hashicorp/go-azure-sdk/microsoft-graph/common-types/beta"
	"github.com/hashicorp/go-azure-sdk/sdk/client"
	"github.com/hashicorp/go-azure-sdk/sdk/odata"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type GetPendingAccessReviewInstanceDecisionInstanceStageOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
	Model        *beta.AccessReviewStage
}

type GetPendingAccessReviewInstanceDecisionInstanceStageOperationOptions struct {
	Expand *odata.Expand
	Select *[]string
}

func DefaultGetPendingAccessReviewInstanceDecisionInstanceStageOperationOptions() GetPendingAccessReviewInstanceDecisionInstanceStageOperationOptions {
	return GetPendingAccessReviewInstanceDecisionInstanceStageOperationOptions{}
}

func (o GetPendingAccessReviewInstanceDecisionInstanceStageOperationOptions) ToHeaders() *client.Headers {
	out := client.Headers{}

	return &out
}

func (o GetPendingAccessReviewInstanceDecisionInstanceStageOperationOptions) ToOData() *odata.Query {
	out := odata.Query{}
	if o.Expand != nil {
		out.Expand = *o.Expand
	}
	if o.Select != nil {
		out.Select = *o.Select
	}
	return &out
}

func (o GetPendingAccessReviewInstanceDecisionInstanceStageOperationOptions) ToQuery() *client.QueryParams {
	out := client.QueryParams{}

	return &out
}

// GetPendingAccessReviewInstanceDecisionInstanceStage - Get stages from me. If the instance has multiple stages, this
// returns the collection of stages. A new stage will only be created when the previous stage ends. The existence,
// number, and settings of stages on a review instance are created based on the accessReviewStageSettings on the parent
// accessReviewScheduleDefinition.
func (c PendingAccessReviewInstanceDecisionInstanceStageClient) GetPendingAccessReviewInstanceDecisionInstanceStage(ctx context.Context, id beta.MePendingAccessReviewInstanceIdDecisionIdInstanceStageId, options GetPendingAccessReviewInstanceDecisionInstanceStageOperationOptions) (result GetPendingAccessReviewInstanceDecisionInstanceStageOperationResponse, err error) {
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

	var model beta.AccessReviewStage
	result.Model = &model
	if err = resp.Unmarshal(result.Model); err != nil {
		return
	}

	return
}
