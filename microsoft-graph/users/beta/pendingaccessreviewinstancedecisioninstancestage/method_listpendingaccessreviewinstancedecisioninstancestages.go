package pendingaccessreviewinstancedecisioninstancestage

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

type ListPendingAccessReviewInstanceDecisionInstanceStagesOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
	Model        *[]beta.AccessReviewStage
}

type ListPendingAccessReviewInstanceDecisionInstanceStagesCompleteResult struct {
	LatestHttpResponse *http.Response
	Items              []beta.AccessReviewStage
}

type ListPendingAccessReviewInstanceDecisionInstanceStagesCustomPager struct {
	NextLink *odata.Link `json:"@odata.nextLink"`
}

func (p *ListPendingAccessReviewInstanceDecisionInstanceStagesCustomPager) NextPageLink() *odata.Link {
	defer func() {
		p.NextLink = nil
	}()

	return p.NextLink
}

// ListPendingAccessReviewInstanceDecisionInstanceStages ...
func (c PendingAccessReviewInstanceDecisionInstanceStageClient) ListPendingAccessReviewInstanceDecisionInstanceStages(ctx context.Context, id UserIdPendingAccessReviewInstanceIdDecisionId) (result ListPendingAccessReviewInstanceDecisionInstanceStagesOperationResponse, err error) {
	opts := client.RequestOptions{
		ContentType: "application/json; charset=utf-8",
		ExpectedStatusCodes: []int{
			http.StatusOK,
		},
		HttpMethod: http.MethodGet,
		Pager:      &ListPendingAccessReviewInstanceDecisionInstanceStagesCustomPager{},
		Path:       fmt.Sprintf("%s/instance/stages", id.ID()),
	}

	req, err := c.Client.NewRequest(ctx, opts)
	if err != nil {
		return
	}

	var resp *client.Response
	resp, err = req.ExecutePaged(ctx)
	if resp != nil {
		result.OData = resp.OData
		result.HttpResponse = resp.Response
	}
	if err != nil {
		return
	}

	var values struct {
		Values *[]beta.AccessReviewStage `json:"value"`
	}
	if err = resp.Unmarshal(&values); err != nil {
		return
	}

	result.Model = values.Values

	return
}

// ListPendingAccessReviewInstanceDecisionInstanceStagesComplete retrieves all the results into a single object
func (c PendingAccessReviewInstanceDecisionInstanceStageClient) ListPendingAccessReviewInstanceDecisionInstanceStagesComplete(ctx context.Context, id UserIdPendingAccessReviewInstanceIdDecisionId) (ListPendingAccessReviewInstanceDecisionInstanceStagesCompleteResult, error) {
	return c.ListPendingAccessReviewInstanceDecisionInstanceStagesCompleteMatchingPredicate(ctx, id, AccessReviewStageOperationPredicate{})
}

// ListPendingAccessReviewInstanceDecisionInstanceStagesCompleteMatchingPredicate retrieves all the results and then applies the predicate
func (c PendingAccessReviewInstanceDecisionInstanceStageClient) ListPendingAccessReviewInstanceDecisionInstanceStagesCompleteMatchingPredicate(ctx context.Context, id UserIdPendingAccessReviewInstanceIdDecisionId, predicate AccessReviewStageOperationPredicate) (result ListPendingAccessReviewInstanceDecisionInstanceStagesCompleteResult, err error) {
	items := make([]beta.AccessReviewStage, 0)

	resp, err := c.ListPendingAccessReviewInstanceDecisionInstanceStages(ctx, id)
	if err != nil {
		result.LatestHttpResponse = resp.HttpResponse
		err = fmt.Errorf("loading results: %+v", err)
		return
	}
	if resp.Model != nil {
		for _, v := range *resp.Model {
			if predicate.Matches(v) {
				items = append(items, v)
			}
		}
	}

	result = ListPendingAccessReviewInstanceDecisionInstanceStagesCompleteResult{
		LatestHttpResponse: resp.HttpResponse,
		Items:              items,
	}
	return
}
