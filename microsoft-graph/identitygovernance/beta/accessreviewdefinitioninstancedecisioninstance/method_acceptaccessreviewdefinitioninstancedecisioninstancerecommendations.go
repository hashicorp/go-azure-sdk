package accessreviewdefinitioninstancedecisioninstance

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

type AcceptAccessReviewDefinitionInstanceDecisionInstanceRecommendationsOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
}

type AcceptAccessReviewDefinitionInstanceDecisionInstanceRecommendationsOperationOptions struct {
	Metadata  *odata.Metadata
	RetryFunc client.RequestRetryFunc
}

func DefaultAcceptAccessReviewDefinitionInstanceDecisionInstanceRecommendationsOperationOptions() AcceptAccessReviewDefinitionInstanceDecisionInstanceRecommendationsOperationOptions {
	return AcceptAccessReviewDefinitionInstanceDecisionInstanceRecommendationsOperationOptions{}
}

func (o AcceptAccessReviewDefinitionInstanceDecisionInstanceRecommendationsOperationOptions) ToHeaders() *client.Headers {
	out := client.Headers{}

	return &out
}

func (o AcceptAccessReviewDefinitionInstanceDecisionInstanceRecommendationsOperationOptions) ToOData() *odata.Query {
	out := odata.Query{}
	if o.Metadata != nil {
		out.Metadata = *o.Metadata
	}
	return &out
}

func (o AcceptAccessReviewDefinitionInstanceDecisionInstanceRecommendationsOperationOptions) ToQuery() *client.QueryParams {
	out := client.QueryParams{}

	return &out
}

// AcceptAccessReviewDefinitionInstanceDecisionInstanceRecommendations - Invoke action acceptRecommendations. Allows the
// acceptance of recommendations on all accessReviewInstanceDecisionItem objects that haven't been reviewed for an
// accessReviewInstance object for which the calling user is a reviewer. Recommendations are generated if
// recommendationsEnabled is true on the accessReviewScheduleDefinition object. If there isn't a recommendation on an
// accessReviewInstanceDecisionItem object no decision will be recorded.
func (c AccessReviewDefinitionInstanceDecisionInstanceClient) AcceptAccessReviewDefinitionInstanceDecisionInstanceRecommendations(ctx context.Context, id beta.IdentityGovernanceAccessReviewDefinitionIdInstanceIdDecisionId, options AcceptAccessReviewDefinitionInstanceDecisionInstanceRecommendationsOperationOptions) (result AcceptAccessReviewDefinitionInstanceDecisionInstanceRecommendationsOperationResponse, err error) {
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
