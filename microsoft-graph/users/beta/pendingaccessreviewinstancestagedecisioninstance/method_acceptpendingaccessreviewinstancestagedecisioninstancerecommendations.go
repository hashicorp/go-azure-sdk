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

type AcceptPendingAccessReviewInstanceStageDecisionInstanceRecommendationsOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
}

type AcceptPendingAccessReviewInstanceStageDecisionInstanceRecommendationsOperationOptions struct {
	Metadata  *odata.Metadata
	RetryFunc client.RequestRetryFunc
}

func DefaultAcceptPendingAccessReviewInstanceStageDecisionInstanceRecommendationsOperationOptions() AcceptPendingAccessReviewInstanceStageDecisionInstanceRecommendationsOperationOptions {
	return AcceptPendingAccessReviewInstanceStageDecisionInstanceRecommendationsOperationOptions{}
}

func (o AcceptPendingAccessReviewInstanceStageDecisionInstanceRecommendationsOperationOptions) ToHeaders() *client.Headers {
	out := client.Headers{}

	return &out
}

func (o AcceptPendingAccessReviewInstanceStageDecisionInstanceRecommendationsOperationOptions) ToOData() *odata.Query {
	out := odata.Query{}
	if o.Metadata != nil {
		out.Metadata = *o.Metadata
	}
	return &out
}

func (o AcceptPendingAccessReviewInstanceStageDecisionInstanceRecommendationsOperationOptions) ToQuery() *client.QueryParams {
	out := client.QueryParams{}

	return &out
}

// AcceptPendingAccessReviewInstanceStageDecisionInstanceRecommendations - Invoke action acceptRecommendations. Allows
// the acceptance of recommendations on all accessReviewInstanceDecisionItem objects that haven't been reviewed for an
// accessReviewInstance object for which the calling user is a reviewer. Recommendations are generated if
// recommendationsEnabled is true on the accessReviewScheduleDefinition object. If there isn't a recommendation on an
// accessReviewInstanceDecisionItem object no decision will be recorded.
func (c PendingAccessReviewInstanceStageDecisionInstanceClient) AcceptPendingAccessReviewInstanceStageDecisionInstanceRecommendations(ctx context.Context, id beta.UserIdPendingAccessReviewInstanceIdStageIdDecisionId, options AcceptPendingAccessReviewInstanceStageDecisionInstanceRecommendationsOperationOptions) (result AcceptPendingAccessReviewInstanceStageDecisionInstanceRecommendationsOperationResponse, err error) {
	opts := client.RequestOptions{
		ContentType: "application/json; charset=utf-8",
		ExpectedStatusCodes: []int{
			http.StatusNoContent,
		},
		HttpMethod:    http.MethodPost,
		OptionsObject: options,
		Path:          fmt.Sprintf("%s/instance/acceptRecommendations", id.ID()),
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
