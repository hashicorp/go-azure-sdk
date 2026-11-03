package instancepooloperations

import (
	"context"
	"fmt"
	"net/http"

	"github.com/hashicorp/go-azure-sdk/sdk/client"
	"github.com/hashicorp/go-azure-sdk/sdk/odata"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type ListByInstancePoolOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
	Model        *[]InstancePoolOperation
}

type ListByInstancePoolCompleteResult struct {
	LatestHttpResponse *http.Response
	Items              []InstancePoolOperation
}

type ListByInstancePoolCustomPager struct {
	NextLink *odata.Link `json:"nextLink"`
}

func (p *ListByInstancePoolCustomPager) NextPageLink() *odata.Link {
	defer func() {
		p.NextLink = nil
	}()

	return p.NextLink
}

// ListByInstancePool ...
func (c InstancePoolOperationsClient) ListByInstancePool(ctx context.Context, id InstancePoolId) (result ListByInstancePoolOperationResponse, err error) {
	opts := client.RequestOptions{
		ContentType: "application/json; charset=utf-8",
		ExpectedStatusCodes: []int{
			http.StatusOK,
		},
		HttpMethod: http.MethodGet,
		Pager:      &ListByInstancePoolCustomPager{},
		Path:       fmt.Sprintf("%s/operations", id.ID()),
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
		Values *[]InstancePoolOperation `json:"value"`
	}
	if err = resp.Unmarshal(&values); err != nil {
		return
	}

	result.Model = values.Values

	return
}

// ListByInstancePoolComplete retrieves all the results into a single object
func (c InstancePoolOperationsClient) ListByInstancePoolComplete(ctx context.Context, id InstancePoolId) (ListByInstancePoolCompleteResult, error) {
	return c.ListByInstancePoolCompleteMatchingPredicate(ctx, id, InstancePoolOperationOperationPredicate{})
}

// ListByInstancePoolCompleteMatchingPredicate retrieves all the results and then applies the predicate
func (c InstancePoolOperationsClient) ListByInstancePoolCompleteMatchingPredicate(ctx context.Context, id InstancePoolId, predicate InstancePoolOperationOperationPredicate) (result ListByInstancePoolCompleteResult, err error) {
	items := make([]InstancePoolOperation, 0)

	resp, err := c.ListByInstancePool(ctx, id)
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

	result = ListByInstancePoolCompleteResult{
		LatestHttpResponse: resp.HttpResponse,
		Items:              items,
	}
	return
}
