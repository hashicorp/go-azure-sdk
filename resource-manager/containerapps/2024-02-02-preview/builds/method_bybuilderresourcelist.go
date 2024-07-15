package builds

import (
	"context"
	"fmt"
	"net/http"

	"github.com/hashicorp/go-azure-sdk/sdk/client"
	"github.com/hashicorp/go-azure-sdk/sdk/odata"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type ByBuilderResourceListOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
	Model        *[]BuildResource
}

type ByBuilderResourceListCompleteResult struct {
	LatestHttpResponse *http.Response
	Items              []BuildResource
}

type ByBuilderResourceListCustomPager struct {
	NextLink *odata.Link `json:"nextLink"`
}

func (p *ByBuilderResourceListCustomPager) NextPageLink() *odata.Link {
	defer func() {
		p.NextLink = nil
	}()

	return p.NextLink
}

// ByBuilderResourceList ...
func (c BuildsClient) ByBuilderResourceList(ctx context.Context, id BuilderId) (result ByBuilderResourceListOperationResponse, err error) {
	opts := client.RequestOptions{
		ContentType: "application/json; charset=utf-8",
		ExpectedStatusCodes: []int{
			http.StatusOK,
		},
		HttpMethod: http.MethodGet,
		Pager:      &ByBuilderResourceListCustomPager{},
		Path:       fmt.Sprintf("%s/builds", id.ID()),
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
		Values *[]BuildResource `json:"value"`
	}
	if err = resp.Unmarshal(&values); err != nil {
		return
	}

	result.Model = values.Values

	return
}

// ByBuilderResourceListComplete retrieves all the results into a single object
func (c BuildsClient) ByBuilderResourceListComplete(ctx context.Context, id BuilderId) (ByBuilderResourceListCompleteResult, error) {
	return c.ByBuilderResourceListCompleteMatchingPredicate(ctx, id, BuildResourceOperationPredicate{})
}

// ByBuilderResourceListCompleteMatchingPredicate retrieves all the results and then applies the predicate
func (c BuildsClient) ByBuilderResourceListCompleteMatchingPredicate(ctx context.Context, id BuilderId, predicate BuildResourceOperationPredicate) (result ByBuilderResourceListCompleteResult, err error) {
	items := make([]BuildResource, 0)

	resp, err := c.ByBuilderResourceList(ctx, id)
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

	result = ByBuilderResourceListCompleteResult{
		LatestHttpResponse: resp.HttpResponse,
		Items:              items,
	}
	return
}
