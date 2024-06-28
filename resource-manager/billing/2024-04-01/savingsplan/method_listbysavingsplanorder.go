package savingsplan

import (
	"context"
	"fmt"
	"net/http"

	"github.com/hashicorp/go-azure-sdk/sdk/client"
	"github.com/hashicorp/go-azure-sdk/sdk/odata"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type ListBySavingsPlanOrderOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
	Model        *[]SavingsPlanModel
}

type ListBySavingsPlanOrderCompleteResult struct {
	LatestHttpResponse *http.Response
	Items              []SavingsPlanModel
}

type ListBySavingsPlanOrderCustomPager struct {
	NextLink *odata.Link `json:"nextLink"`
}

func (p *ListBySavingsPlanOrderCustomPager) NextPageLink() *odata.Link {
	defer func() {
		p.NextLink = nil
	}()

	return p.NextLink
}

// ListBySavingsPlanOrder ...
func (c SavingsPlanClient) ListBySavingsPlanOrder(ctx context.Context, id SavingsPlanOrderId) (result ListBySavingsPlanOrderOperationResponse, err error) {
	opts := client.RequestOptions{
		ContentType: "application/json; charset=utf-8",
		ExpectedStatusCodes: []int{
			http.StatusOK,
		},
		HttpMethod: http.MethodGet,
		Pager:      &ListBySavingsPlanOrderCustomPager{},
		Path:       fmt.Sprintf("%s/savingsPlans", id.ID()),
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
		Values *[]SavingsPlanModel `json:"value"`
	}
	if err = resp.Unmarshal(&values); err != nil {
		return
	}

	result.Model = values.Values

	return
}

// ListBySavingsPlanOrderComplete retrieves all the results into a single object
func (c SavingsPlanClient) ListBySavingsPlanOrderComplete(ctx context.Context, id SavingsPlanOrderId) (ListBySavingsPlanOrderCompleteResult, error) {
	return c.ListBySavingsPlanOrderCompleteMatchingPredicate(ctx, id, SavingsPlanModelOperationPredicate{})
}

// ListBySavingsPlanOrderCompleteMatchingPredicate retrieves all the results and then applies the predicate
func (c SavingsPlanClient) ListBySavingsPlanOrderCompleteMatchingPredicate(ctx context.Context, id SavingsPlanOrderId, predicate SavingsPlanModelOperationPredicate) (result ListBySavingsPlanOrderCompleteResult, err error) {
	items := make([]SavingsPlanModel, 0)

	resp, err := c.ListBySavingsPlanOrder(ctx, id)
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

	result = ListBySavingsPlanOrderCompleteResult{
		LatestHttpResponse: resp.HttpResponse,
		Items:              items,
	}
	return
}
