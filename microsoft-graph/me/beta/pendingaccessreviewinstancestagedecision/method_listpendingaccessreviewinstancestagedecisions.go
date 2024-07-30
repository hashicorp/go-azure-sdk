package pendingaccessreviewinstancestagedecision

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

type ListPendingAccessReviewInstanceStageDecisionsOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
	Model        *[]beta.AccessReviewInstanceDecisionItem
}

type ListPendingAccessReviewInstanceStageDecisionsCompleteResult struct {
	LatestHttpResponse *http.Response
	Items              []beta.AccessReviewInstanceDecisionItem
}

type ListPendingAccessReviewInstanceStageDecisionsCustomPager struct {
	NextLink *odata.Link `json:"@odata.nextLink"`
}

func (p *ListPendingAccessReviewInstanceStageDecisionsCustomPager) NextPageLink() *odata.Link {
	defer func() {
		p.NextLink = nil
	}()

	return p.NextLink
}

// ListPendingAccessReviewInstanceStageDecisions ...
func (c PendingAccessReviewInstanceStageDecisionClient) ListPendingAccessReviewInstanceStageDecisions(ctx context.Context, id MePendingAccessReviewInstanceIdStageId) (result ListPendingAccessReviewInstanceStageDecisionsOperationResponse, err error) {
	opts := client.RequestOptions{
		ContentType: "application/json; charset=utf-8",
		ExpectedStatusCodes: []int{
			http.StatusOK,
		},
		HttpMethod: http.MethodGet,
		Pager:      &ListPendingAccessReviewInstanceStageDecisionsCustomPager{},
		Path:       fmt.Sprintf("%s/decisions", id.ID()),
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
		Values *[]beta.AccessReviewInstanceDecisionItem `json:"value"`
	}
	if err = resp.Unmarshal(&values); err != nil {
		return
	}

	result.Model = values.Values

	return
}

// ListPendingAccessReviewInstanceStageDecisionsComplete retrieves all the results into a single object
func (c PendingAccessReviewInstanceStageDecisionClient) ListPendingAccessReviewInstanceStageDecisionsComplete(ctx context.Context, id MePendingAccessReviewInstanceIdStageId) (ListPendingAccessReviewInstanceStageDecisionsCompleteResult, error) {
	return c.ListPendingAccessReviewInstanceStageDecisionsCompleteMatchingPredicate(ctx, id, AccessReviewInstanceDecisionItemOperationPredicate{})
}

// ListPendingAccessReviewInstanceStageDecisionsCompleteMatchingPredicate retrieves all the results and then applies the predicate
func (c PendingAccessReviewInstanceStageDecisionClient) ListPendingAccessReviewInstanceStageDecisionsCompleteMatchingPredicate(ctx context.Context, id MePendingAccessReviewInstanceIdStageId, predicate AccessReviewInstanceDecisionItemOperationPredicate) (result ListPendingAccessReviewInstanceStageDecisionsCompleteResult, err error) {
	items := make([]beta.AccessReviewInstanceDecisionItem, 0)

	resp, err := c.ListPendingAccessReviewInstanceStageDecisions(ctx, id)
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

	result = ListPendingAccessReviewInstanceStageDecisionsCompleteResult{
		LatestHttpResponse: resp.HttpResponse,
		Items:              items,
	}
	return
}
