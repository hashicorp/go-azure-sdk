package standbycontainergrouppoolruntimeviews

import (
	"context"
	"fmt"
	"net/http"

	"github.com/hashicorp/go-azure-sdk/sdk/client"
	"github.com/hashicorp/go-azure-sdk/sdk/odata"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type ListByStandbyPoolOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
	Model        *[]StandbyContainerGroupPoolRuntimeViewResource
}

type ListByStandbyPoolCompleteResult struct {
	LatestHttpResponse *http.Response
	Items              []StandbyContainerGroupPoolRuntimeViewResource
}

type ListByStandbyPoolCustomPager struct {
	NextLink *odata.Link `json:"nextLink"`
}

func (p *ListByStandbyPoolCustomPager) NextPageLink() *odata.Link {
	defer func() {
		p.NextLink = nil
	}()

	return p.NextLink
}

// ListByStandbyPool ...
func (c StandbyContainerGroupPoolRuntimeViewsClient) ListByStandbyPool(ctx context.Context, id StandbyContainerGroupPoolId) (result ListByStandbyPoolOperationResponse, err error) {
	opts := client.RequestOptions{
		ContentType: "application/json; charset=utf-8",
		ExpectedStatusCodes: []int{
			http.StatusOK,
		},
		HttpMethod: http.MethodGet,
		Pager:      &ListByStandbyPoolCustomPager{},
		Path:       fmt.Sprintf("%s/runtimeViews", id.ID()),
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
		Values *[]StandbyContainerGroupPoolRuntimeViewResource `json:"value"`
	}
	if err = resp.Unmarshal(&values); err != nil {
		return
	}

	result.Model = values.Values

	return
}

// ListByStandbyPoolComplete retrieves all the results into a single object
func (c StandbyContainerGroupPoolRuntimeViewsClient) ListByStandbyPoolComplete(ctx context.Context, id StandbyContainerGroupPoolId) (ListByStandbyPoolCompleteResult, error) {
	return c.ListByStandbyPoolCompleteMatchingPredicate(ctx, id, StandbyContainerGroupPoolRuntimeViewResourceOperationPredicate{})
}

// ListByStandbyPoolCompleteMatchingPredicate retrieves all the results and then applies the predicate
func (c StandbyContainerGroupPoolRuntimeViewsClient) ListByStandbyPoolCompleteMatchingPredicate(ctx context.Context, id StandbyContainerGroupPoolId, predicate StandbyContainerGroupPoolRuntimeViewResourceOperationPredicate) (result ListByStandbyPoolCompleteResult, err error) {
	items := make([]StandbyContainerGroupPoolRuntimeViewResource, 0)

	resp, err := c.ListByStandbyPool(ctx, id)
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

	result = ListByStandbyPoolCompleteResult{
		LatestHttpResponse: resp.HttpResponse,
		Items:              items,
	}
	return
}
