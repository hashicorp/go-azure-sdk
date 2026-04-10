package webtestlocationsapis

import (
	"context"
	"fmt"
	"net/http"

	"github.com/hashicorp/go-azure-sdk/sdk/client"
	"github.com/hashicorp/go-azure-sdk/sdk/odata"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type WebTestLocationsListOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
	Model        *[]ApplicationInsightsComponentWebTestLocation
}

type WebTestLocationsListCompleteResult struct {
	LatestHttpResponse *http.Response
	Items              []ApplicationInsightsComponentWebTestLocation
}

type WebTestLocationsListCustomPager struct {
	NextLink *odata.Link `json:"nextLink"`
}

func (p *WebTestLocationsListCustomPager) NextPageLink() *odata.Link {
	defer func() {
		p.NextLink = nil
	}()

	return p.NextLink
}

// WebTestLocationsList ...
func (c WebTestLocationsAPIsClient) WebTestLocationsList(ctx context.Context, id ComponentId) (result WebTestLocationsListOperationResponse, err error) {
	opts := client.RequestOptions{
		ContentType: "application/json; charset=utf-8",
		ExpectedStatusCodes: []int{
			http.StatusOK,
		},
		HttpMethod: http.MethodGet,
		Pager:      &WebTestLocationsListCustomPager{},
		Path:       fmt.Sprintf("%s/syntheticmonitorlocations", id.ID()),
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
		Values *[]ApplicationInsightsComponentWebTestLocation `json:"value"`
	}
	if err = resp.Unmarshal(&values); err != nil {
		return
	}

	result.Model = values.Values

	return
}

// WebTestLocationsListComplete retrieves all the results into a single object
func (c WebTestLocationsAPIsClient) WebTestLocationsListComplete(ctx context.Context, id ComponentId) (WebTestLocationsListCompleteResult, error) {
	return c.WebTestLocationsListCompleteMatchingPredicate(ctx, id, ApplicationInsightsComponentWebTestLocationOperationPredicate{})
}

// WebTestLocationsListCompleteMatchingPredicate retrieves all the results and then applies the predicate
func (c WebTestLocationsAPIsClient) WebTestLocationsListCompleteMatchingPredicate(ctx context.Context, id ComponentId, predicate ApplicationInsightsComponentWebTestLocationOperationPredicate) (result WebTestLocationsListCompleteResult, err error) {
	items := make([]ApplicationInsightsComponentWebTestLocation, 0)

	resp, err := c.WebTestLocationsList(ctx, id)
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

	result = WebTestLocationsListCompleteResult{
		LatestHttpResponse: resp.HttpResponse,
		Items:              items,
	}
	return
}
