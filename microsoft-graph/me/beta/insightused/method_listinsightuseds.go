package insightused

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

type ListInsightUsedsOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
	Model        *[]beta.UsedInsight
}

type ListInsightUsedsCompleteResult struct {
	LatestHttpResponse *http.Response
	Items              []beta.UsedInsight
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
func (c InsightUsedClient) ListInsightUseds(ctx context.Context) (result ListInsightUsedsOperationResponse, err error) {
	opts := client.RequestOptions{
		ContentType: "application/json; charset=utf-8",
		ExpectedStatusCodes: []int{
			http.StatusOK,
		},
		HttpMethod: http.MethodGet,
		Pager:      &ListInsightUsedsCustomPager{},
		Path:       "/me/insights/used",
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
		Values *[]beta.UsedInsight `json:"value"`
	}
	if err = resp.Unmarshal(&values); err != nil {
		return
	}

	result.Model = values.Values

	return
}

// ListInsightUsedsComplete retrieves all the results into a single object
func (c InsightUsedClient) ListInsightUsedsComplete(ctx context.Context) (ListInsightUsedsCompleteResult, error) {
	return c.ListInsightUsedsCompleteMatchingPredicate(ctx, UsedInsightOperationPredicate{})
}

// ListInsightUsedsCompleteMatchingPredicate retrieves all the results and then applies the predicate
func (c InsightUsedClient) ListInsightUsedsCompleteMatchingPredicate(ctx context.Context, predicate UsedInsightOperationPredicate) (result ListInsightUsedsCompleteResult, err error) {
	items := make([]beta.UsedInsight, 0)

	resp, err := c.ListInsightUseds(ctx)
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
