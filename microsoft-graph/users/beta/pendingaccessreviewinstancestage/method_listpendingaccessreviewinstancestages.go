package pendingaccessreviewinstancestage

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

type ListPendingAccessReviewInstanceStagesOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
	Model        *[]beta.AccessReviewStage
}

type ListPendingAccessReviewInstanceStagesCompleteResult struct {
	LatestHttpResponse *http.Response
	Items              []beta.AccessReviewStage
}

type ListPendingAccessReviewInstanceStagesCustomPager struct {
	NextLink *odata.Link `json:"@odata.nextLink"`
}

func (p *ListPendingAccessReviewInstanceStagesCustomPager) NextPageLink() *odata.Link {
	defer func() {
		p.NextLink = nil
	}()

	return p.NextLink
}

// ListPendingAccessReviewInstanceStages ...
func (c PendingAccessReviewInstanceStageClient) ListPendingAccessReviewInstanceStages(ctx context.Context, id UserIdPendingAccessReviewInstanceId) (result ListPendingAccessReviewInstanceStagesOperationResponse, err error) {
	opts := client.RequestOptions{
		ContentType: "application/json; charset=utf-8",
		ExpectedStatusCodes: []int{
			http.StatusOK,
		},
		HttpMethod: http.MethodGet,
		Pager:      &ListPendingAccessReviewInstanceStagesCustomPager{},
		Path:       fmt.Sprintf("%s/stages", id.ID()),
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

// ListPendingAccessReviewInstanceStagesComplete retrieves all the results into a single object
func (c PendingAccessReviewInstanceStageClient) ListPendingAccessReviewInstanceStagesComplete(ctx context.Context, id UserIdPendingAccessReviewInstanceId) (ListPendingAccessReviewInstanceStagesCompleteResult, error) {
	return c.ListPendingAccessReviewInstanceStagesCompleteMatchingPredicate(ctx, id, AccessReviewStageOperationPredicate{})
}

// ListPendingAccessReviewInstanceStagesCompleteMatchingPredicate retrieves all the results and then applies the predicate
func (c PendingAccessReviewInstanceStageClient) ListPendingAccessReviewInstanceStagesCompleteMatchingPredicate(ctx context.Context, id UserIdPendingAccessReviewInstanceId, predicate AccessReviewStageOperationPredicate) (result ListPendingAccessReviewInstanceStagesCompleteResult, err error) {
	items := make([]beta.AccessReviewStage, 0)

	resp, err := c.ListPendingAccessReviewInstanceStages(ctx, id)
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

	result = ListPendingAccessReviewInstanceStagesCompleteResult{
		LatestHttpResponse: resp.HttpResponse,
		Items:              items,
	}
	return
}
