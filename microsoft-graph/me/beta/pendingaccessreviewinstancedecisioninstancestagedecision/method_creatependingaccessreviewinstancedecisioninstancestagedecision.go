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

type CreatePendingAccessReviewInstanceDecisionInstanceStageDecisionOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
	Model        *beta.AccessReviewInstanceDecisionItem
}

type CreatePendingAccessReviewInstanceDecisionInstanceStageDecisionOperationOptions struct {
	Metadata *odata.Metadata
}

func DefaultCreatePendingAccessReviewInstanceDecisionInstanceStageDecisionOperationOptions() CreatePendingAccessReviewInstanceDecisionInstanceStageDecisionOperationOptions {
	return CreatePendingAccessReviewInstanceDecisionInstanceStageDecisionOperationOptions{}
}

func (o CreatePendingAccessReviewInstanceDecisionInstanceStageDecisionOperationOptions) ToHeaders() *client.Headers {
	out := client.Headers{}

	return &out
}

func (o CreatePendingAccessReviewInstanceDecisionInstanceStageDecisionOperationOptions) ToOData() *odata.Query {
	out := odata.Query{}
	if o.Metadata != nil {
		out.Metadata = *o.Metadata
	}
	return &out
}

func (o CreatePendingAccessReviewInstanceDecisionInstanceStageDecisionOperationOptions) ToQuery() *client.QueryParams {
	out := client.QueryParams{}

	return &out
}

// CreatePendingAccessReviewInstanceDecisionInstanceStageDecision - Create new navigation property to decisions for me
func (c PendingAccessReviewInstanceDecisionInstanceStageDecisionClient) CreatePendingAccessReviewInstanceDecisionInstanceStageDecision(ctx context.Context, id beta.MePendingAccessReviewInstanceIdDecisionIdInstanceStageId, input beta.AccessReviewInstanceDecisionItem, options CreatePendingAccessReviewInstanceDecisionInstanceStageDecisionOperationOptions) (result CreatePendingAccessReviewInstanceDecisionInstanceStageDecisionOperationResponse, err error) {
	opts := client.RequestOptions{
		ContentType: "application/json; charset=utf-8",
		ExpectedStatusCodes: []int{
			http.StatusCreated,
		},
		HttpMethod:    http.MethodPost,
		OptionsObject: options,
		Path:          fmt.Sprintf("%s/decisions", id.ID()),
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

	var model beta.AccessReviewInstanceDecisionItem
	result.Model = &model
	if err = resp.Unmarshal(result.Model); err != nil {
		return
	}

	return
}
