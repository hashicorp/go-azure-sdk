package pendingaccessreviewinstancestagedecisioninstance

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

type GetPendingAccessReviewInstanceStageDecisionInstanceOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
	Model        *beta.AccessReviewInstance
}

type GetPendingAccessReviewInstanceStageDecisionInstanceOperationOptions struct {
	Expand *odata.Expand
	Select *[]string
}

func DefaultGetPendingAccessReviewInstanceStageDecisionInstanceOperationOptions() GetPendingAccessReviewInstanceStageDecisionInstanceOperationOptions {
	return GetPendingAccessReviewInstanceStageDecisionInstanceOperationOptions{}
}

func (o GetPendingAccessReviewInstanceStageDecisionInstanceOperationOptions) ToHeaders() *client.Headers {
	out := client.Headers{}

	return &out
}

func (o GetPendingAccessReviewInstanceStageDecisionInstanceOperationOptions) ToOData() *odata.Query {
	out := odata.Query{}
	if o.Expand != nil {
		out.Expand = *o.Expand
	}
	if o.Select != nil {
		out.Select = *o.Select
	}
	return &out
}

func (o GetPendingAccessReviewInstanceStageDecisionInstanceOperationOptions) ToQuery() *client.QueryParams {
	out := client.QueryParams{}

	return &out
}

// GetPendingAccessReviewInstanceStageDecisionInstance - Get instance from users. There's exactly one
// accessReviewInstance associated with each decision. The instance is the parent of the decision item, representing the
// recurrence of the access review the decision is made on.
func (c PendingAccessReviewInstanceStageDecisionInstanceClient) GetPendingAccessReviewInstanceStageDecisionInstance(ctx context.Context, id beta.UserIdPendingAccessReviewInstanceIdStageIdDecisionId, options GetPendingAccessReviewInstanceStageDecisionInstanceOperationOptions) (result GetPendingAccessReviewInstanceStageDecisionInstanceOperationResponse, err error) {
	opts := client.RequestOptions{
		ContentType: "application/json; charset=utf-8",
		ExpectedStatusCodes: []int{
			http.StatusOK,
		},
		HttpMethod:    http.MethodGet,
		OptionsObject: options,
		Path:          fmt.Sprintf("%s/instance", id.ID()),
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

	var model beta.AccessReviewInstance
	result.Model = &model
	if err = resp.Unmarshal(result.Model); err != nil {
		return
	}

	return
}
