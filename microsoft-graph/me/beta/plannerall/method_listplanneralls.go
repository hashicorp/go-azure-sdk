package plannerall

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

type ListPlannerAllsOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
	Model        *[]beta.PlannerDelta
}

type ListPlannerAllsCompleteResult struct {
	LatestHttpResponse *http.Response
	Items              []beta.PlannerDelta
}

type ListPlannerAllsCustomPager struct {
	NextLink *odata.Link `json:"@odata.nextLink"`
}

func (p *ListPlannerAllsCustomPager) NextPageLink() *odata.Link {
	defer func() {
		p.NextLink = nil
	}()

	return p.NextLink
}

// ListPlannerAlls ...
func (c PlannerAllClient) ListPlannerAlls(ctx context.Context) (result ListPlannerAllsOperationResponse, err error) {
	opts := client.RequestOptions{
		ContentType: "application/json; charset=utf-8",
		ExpectedStatusCodes: []int{
			http.StatusOK,
		},
		HttpMethod: http.MethodGet,
		Pager:      &ListPlannerAllsCustomPager{},
		Path:       "/me/planner/all",
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
		Values *[]beta.PlannerDelta `json:"value"`
	}
	if err = resp.Unmarshal(&values); err != nil {
		return
	}

	result.Model = values.Values

	return
}

// ListPlannerAllsComplete retrieves all the results into a single object
func (c PlannerAllClient) ListPlannerAllsComplete(ctx context.Context) (ListPlannerAllsCompleteResult, error) {
	return c.ListPlannerAllsCompleteMatchingPredicate(ctx, PlannerDeltaOperationPredicate{})
}

// ListPlannerAllsCompleteMatchingPredicate retrieves all the results and then applies the predicate
func (c PlannerAllClient) ListPlannerAllsCompleteMatchingPredicate(ctx context.Context, predicate PlannerDeltaOperationPredicate) (result ListPlannerAllsCompleteResult, err error) {
	items := make([]beta.PlannerDelta, 0)

	resp, err := c.ListPlannerAlls(ctx)
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

	result = ListPlannerAllsCompleteResult{
		LatestHttpResponse: resp.HttpResponse,
		Items:              items,
	}
	return
}
