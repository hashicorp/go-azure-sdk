package insightused

import (
	"context"
	"fmt"
	"net/http"

	"github.com/hashicorp/go-azure-sdk/microsoft-graph/common-types/stable"
	"github.com/hashicorp/go-azure-sdk/sdk/client"
	"github.com/hashicorp/go-azure-sdk/sdk/odata"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type ListInsightUsedsOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
	Model        *[]stable.UsedInsight
}

type ListInsightUsedsCompleteResult struct {
	LatestHttpResponse *http.Response
	Items              []stable.UsedInsight
}

type ListInsightUsedsCustomPager struct {
	NextLink *odata.Link `json:"@odata.nextLink"`
}

func (p *ListInsightUsedsCustomPager) NextPageLink() *odata.Link {
	defer func() {
		p.NextLink = nil
	}()

	return p.NextLink
}

// ListInsightUseds ...
func (c InsightUsedClient) ListInsightUseds(ctx context.Context, id UserId) (result ListInsightUsedsOperationResponse, err error) {
	opts := client.RequestOptions{
		ContentType: "application/json; charset=utf-8",
		ExpectedStatusCodes: []int{
			http.StatusOK,
		},
		HttpMethod: http.MethodGet,
		Pager:      &ListInsightUsedsCustomPager{},
		Path:       fmt.Sprintf("%s/insights/used", id.ID()),
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
		Values *[]stable.UsedInsight `json:"value"`
	}
	if err = resp.Unmarshal(&values); err != nil {
		return
	}

	result.Model = values.Values

	return
}

// ListInsightUsedsComplete retrieves all the results into a single object
func (c InsightUsedClient) ListInsightUsedsComplete(ctx context.Context, id UserId) (ListInsightUsedsCompleteResult, error) {
	return c.ListInsightUsedsCompleteMatchingPredicate(ctx, id, UsedInsightOperationPredicate{})
}

// ListInsightUsedsCompleteMatchingPredicate retrieves all the results and then applies the predicate
func (c InsightUsedClient) ListInsightUsedsCompleteMatchingPredicate(ctx context.Context, id UserId, predicate UsedInsightOperationPredicate) (result ListInsightUsedsCompleteResult, err error) {
	items := make([]stable.UsedInsight, 0)

	resp, err := c.ListInsightUseds(ctx, id)
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

	result = ListInsightUsedsCompleteResult{
		LatestHttpResponse: resp.HttpResponse,
		Items:              items,
	}
	return
}
