package plannerplanbucket

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

type ListPlannerPlanBucketsOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
	Model        *[]beta.PlannerBucket
}

type ListPlannerPlanBucketsCompleteResult struct {
	LatestHttpResponse *http.Response
	Items              []beta.PlannerBucket
}

type ListPlannerPlanBucketsCustomPager struct {
	NextLink *odata.Link `json:"@odata.nextLink"`
}

func (p *ListPlannerPlanBucketsCustomPager) NextPageLink() *odata.Link {
	defer func() {
		p.NextLink = nil
	}()

	return p.NextLink
}

// ListPlannerPlanBuckets ...
func (c PlannerPlanBucketClient) ListPlannerPlanBuckets(ctx context.Context, id UserIdPlannerPlanId) (result ListPlannerPlanBucketsOperationResponse, err error) {
	opts := client.RequestOptions{
		ContentType: "application/json; charset=utf-8",
		ExpectedStatusCodes: []int{
			http.StatusOK,
		},
		HttpMethod: http.MethodGet,
		Pager:      &ListPlannerPlanBucketsCustomPager{},
		Path:       fmt.Sprintf("%s/buckets", id.ID()),
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
		Values *[]beta.PlannerBucket `json:"value"`
	}
	if err = resp.Unmarshal(&values); err != nil {
		return
	}

	result.Model = values.Values

	return
}

// ListPlannerPlanBucketsComplete retrieves all the results into a single object
func (c PlannerPlanBucketClient) ListPlannerPlanBucketsComplete(ctx context.Context, id UserIdPlannerPlanId) (ListPlannerPlanBucketsCompleteResult, error) {
	return c.ListPlannerPlanBucketsCompleteMatchingPredicate(ctx, id, PlannerBucketOperationPredicate{})
}

// ListPlannerPlanBucketsCompleteMatchingPredicate retrieves all the results and then applies the predicate
func (c PlannerPlanBucketClient) ListPlannerPlanBucketsCompleteMatchingPredicate(ctx context.Context, id UserIdPlannerPlanId, predicate PlannerBucketOperationPredicate) (result ListPlannerPlanBucketsCompleteResult, err error) {
	items := make([]beta.PlannerBucket, 0)

	resp, err := c.ListPlannerPlanBuckets(ctx, id)
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

	result = ListPlannerPlanBucketsCompleteResult{
		LatestHttpResponse: resp.HttpResponse,
		Items:              items,
	}
	return
}
